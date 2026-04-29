# delta

Leaderless eventually consistent replicated data store.

Single binary, no external dependencies. A consistent hashing ring determines which nodes own which keys. CRDT merge resolves conflicts without coordination. Anti-entropy detects and repairs divergence between replicas in the background. Each read or write specifies how many replicas must respond (R and W) out of the total (N). The caller picks these per request, so a low-stakes preference update can tolerate stale reads while a high-stakes record can require stronger agreement.

## Data Flow

```mermaid
graph LR
    subgraph "Write Path"
        C1[Client] -->|Put key, value| COORD[Coordinator]
        COORD -->|fan out| R1[Replica 1]
        COORD -->|fan out| R2[Replica 2]
        COORD -->|fan out| R3[Replica 3]
        R1 -->|ack| COORD
        R2 -->|ack| COORD
        COORD -->|W=2 satisfied| C1
    end

    subgraph "Read Path"
        C2[Client] -->|Get key| COORD2[Coordinator]
        COORD2 -->|fan out| R4[Replica 1]
        COORD2 -->|fan out| R5[Replica 2]
        R4 -->|value + version| COORD2
        R5 -->|value + version| COORD2
        COORD2 -->|R=2 satisfied, merge if divergent| C2
    end
```

## Architecture

```mermaid
graph TD
    subgraph "cmd/delta"
        MAIN[main.go<br/>composition root]
    end

    subgraph "internal/domain/"
        KEY[key.go<br/>hash to ring position]
        RING[ring.go<br/>consistent hashing<br/>preference list]
        REP[replica.go<br/>value + version vector<br/>CRDT merge]
        QRM[quorum.go<br/>R/W/N evaluation]
        MRK[merkle.go<br/>hash tree build + diff]
    end

    subgraph "internal/service/"
        COORD[coordinator.go<br/>route to preference list<br/>gather quorum]
        AE[antientropy.go<br/>periodic Merkle comparison<br/>repair divergent keys]
    end

    subgraph "internal/handler/"
        GRPC[grpc/<br/>client operations]
        SYNC[sync/<br/>edge client sync<br/>+ node-to-node repair]
    end

    subgraph "internal/client/ (ports)"
        PER[persister/<br/>per-node SQLite]
        TEL[telemetry/<br/>wide events]
    end

    subgraph "meld (library dependency)"
        CRDT[crdt/<br/>version vectors, OR-Set, LWW]
        MEMB[membership/swim]
        GOSSIP[gossip/]
    end

    MAIN --> COORD
    MAIN --> AE
    COORD --> RING
    COORD --> REP
    COORD --> QRM
    COORD --> PER
    COORD --> TEL
    AE --> MRK
    AE --> PER
    AE --> TEL
    GRPC --> COORD
    SYNC --> COORD
    SYNC --> AE
    REP --> CRDT
    AE --> MEMB
    AE --> GOSSIP
```

## Consistent Hashing Ring

Every node is hashed to one or more positions on a circular ring of integers (0 to 2^64). To find which nodes own a key, hash the key to a position on the ring and walk clockwise. The first N distinct nodes you encounter are the preference list: the nodes responsible for storing that key's replicas.

This is deterministic. Every node in the cluster can independently compute the same preference list for the same key without any coordination. When a node joins or leaves, only the keys adjacent to its ring positions are affected. Everything else stays where it is.

## Anti-Entropy

```mermaid
sequenceDiagram
    participant A as Node A
    participant B as Node B

    Note over A,B: Periodic background comparison
    A->>B: Send Merkle root hash
    B->>A: Roots differ, send subtree hashes
    A->>B: Identify divergent key ranges
    B->>A: Send replicas for divergent keys
    A->>A: CRDT merge (commutative, idempotent)
    Note over A,B: Replicas converged
```

Not on the hot path. Background process that guarantees convergence given sufficient time.

## Sync Relay Use Case

```mermaid
graph TD
    subgraph "Devices (edge instances)"
        L[Laptop<br/>edge + SQLite]
        D[Desktop<br/>edge + SQLite]
        P[Phone<br/>edge + SQLite]
    end

    subgraph "Homelab (delta cluster)"
        N1[Node 1]
        N2[Node 2]
        N3[Node 3]
    end

    L -->|sync when available| N1
    D -->|sync when available| N2
    P -->|sync when available| N3
    N1 -->|anti-entropy| N2
    N2 -->|anti-entropy| N1
    N2 -->|anti-entropy| N3
    N3 -->|anti-entropy| N2

    style L fill:#0f3460,stroke:#e94560,color:#eee
    style D fill:#0f3460,stroke:#e94560,color:#eee
    style P fill:#0f3460,stroke:#e94560,color:#eee
```

delta is never the source of truth. If homelab is down, edge replicas keep working. When it comes back, sync catches up. Partitions are expected, not errors.

## Dependencies

- **meld**: CRDT types, version vectors, gossip transport, SWIM membership. Must be complete.

Observability via Telemetry port (OTel).
