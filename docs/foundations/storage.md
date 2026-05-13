# Storage

## Purpose

This page explains how cloud engineers choose storage types and why access pattern, durability, lifecycle, and cost matter more than raw capacity alone.

## Core Concepts

Storage is not one thing. Cloud platforms expose multiple storage models because applications do not all interact with data in the same way.

- Object storage is suited to files, assets, backups, logs, and large unstructured data collections.
- Block storage is used when an operating system or application needs attached volumes with low-latency access.
- File storage supports shared file semantics for workloads that expect a filesystem across instances.
- Archive and lifecycle-managed tiers reduce cost for data that is rarely accessed.

Choosing between them depends on how the data is read, written, shared, backed up, and retained.

## What To Evaluate

- Does the workload need filesystem semantics or is object access enough?
- Is low latency more important than very high durability and scale?
- How often is the data accessed?
- How long must the data be retained?
- What replication, backup, or regional recovery model is required?

Storage is often a long-lived decision. Moving application state or large datasets later can be expensive and operationally disruptive.

## Common Mistakes

- Choosing storage by habit instead of by access pattern.
- Making buckets or blobs public without deliberate review.
- Keeping data forever because lifecycle rules were never defined.
- Ignoring egress and cross-region transfer costs.
- Treating backup, restore, and retention as someone else's problem.

## How This Fits Into Cloud Engineering

Cloud engineers need to know where data lives, how it is protected, who can access it, and what it costs to keep. Storage decisions influence application design, analytics architecture, disaster recovery, and compliance. Good storage design is practical, durable, and intentional from the start.

## Official References

- [Google Cloud Storage overview](https://cloud.google.com/storage/docs/introduction)
- [Amazon S3 documentation](https://docs.aws.amazon.com/s3/)
- [Azure Blob Storage introduction](https://learn.microsoft.com/en-us/azure/storage/blobs/storage-blobs-introduction)
