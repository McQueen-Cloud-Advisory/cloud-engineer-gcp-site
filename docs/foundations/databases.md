# Databases

## Purpose

This page explains how cloud engineers choose between database models and why those choices affect application behavior, scaling, reliability, and analytics.

## Core Concepts

Not all data platforms solve the same problem. Transactional systems, analytical platforms, caches, document stores, and object storage-backed data lakes all serve different access patterns. The right database choice depends on how the application reads, writes, relates, and protects data.

- Relational databases are strong when transactions, constraints, and joins matter.
- Document or key-value stores are useful when access patterns are simple and scale requirements are different.
- Analytical platforms are optimized for large scans, aggregation, and reporting rather than transactional writes.
- Caches improve read performance but are not a substitute for a source of truth.

The biggest question is usually not "Which database is best?" It is "What kind of workload are we actually serving?"

## What To Evaluate

- Transaction requirements and acceptable consistency behavior.
- Expected read and write patterns.
- Data model flexibility and how much schema discipline is needed.
- Backup, recovery, replication, and regional resilience requirements.
- Access control, network placement, and how the application reaches the data platform.

Managed database services reduce some administrative effort, but they do not remove the need for indexing, backup planning, access review, cost awareness, or capacity thinking.

## Common Mistakes

- Choosing a database based on trend or familiarity instead of workload shape.
- Using an analytical service for transactional application behavior.
- Ignoring backup and restore until after data becomes critical.
- Treating scaling as automatic without understanding hot partitions, query cost, or index design.
- Storing every kind of data in the same platform for convenience.

## How This Fits Into Cloud Engineering

Cloud engineers need to understand where application state lives and how it behaves under failure, growth, and change. Database decisions influence latency, availability, security boundaries, deployment patterns, and long-term operating cost. Good database choices are foundational architecture choices.

## Official References

- [Google Cloud databases overview](https://cloud.google.com/products/databases)
- [AWS databases overview](https://aws.amazon.com/products/databases/)
- [Azure data storage technology choices](https://learn.microsoft.com/en-us/azure/architecture/data-guide/technology-choices/data-store-overview)
