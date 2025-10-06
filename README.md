# Lineage Lattice Design

**Lineage Lattice** is lineage.cloud's **Key-Value Store (KV Store)** will provide a lightweight, internal data layer for configuration, metadata, and future caching.

# System Design Diagram

Will draw later

## System Goals and Strategies

| Goal/Problem        | Strategy           |
| ------------------- |:------------------:|
| Data Partition            | Consistent Hashing      |
| CAP            | AP + Eventual Consistency           |
| Inconsistency       | Versioning with vector clock          |
| Failure Detection       | Gossip Protocol          |
| Temporary Failure       | Sloppy Quorum          |
| Permanent Failure       | Merkle Tree          |
| Failure Detection       | Gossip Protocol          |
| Memory       | Memory Cache & SSTables          |

## Write Path
Will write the ones based on Cassandra

## Read Path
Will write the ones based on Cassandra