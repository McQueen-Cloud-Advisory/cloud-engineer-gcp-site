# Azure Cosmos DB

## Purpose

Azure Cosmos DB is a managed database service for globally distributed and application-focused data workloads.

## Definition

Azure Cosmos DB is a managed database platform designed for application data models that need low-latency access, flexible scaling, and a managed operational footprint. It is commonly used for document-oriented and other non-relational application workloads where the access pattern and data model fit its strengths.

The important point is that Cosmos DB is not just "a database in Azure." It makes different tradeoffs from a relational store, especially around partitioning, throughput, and how data models are shaped.

In simple terms:

> Cosmos DB is an application database for workloads that need scalable, managed access patterns more than they need traditional relational behavior.

## What Problem It Solves

Cosmos DB gives teams a way to store application data at scale without running database servers directly. It is useful when predictable low-latency reads and writes matter and when the data model fits a document or NoSQL-oriented design.

## How It Is Commonly Used

Cosmos DB is commonly used for:

- application records and state for APIs,
- user profiles, preferences, and session-like data,
- global or distributed application workloads,
- serverless backends that need a managed operational database,
- cases where the application model fits partitioned, non-relational access patterns.

## When to Use It

- Use it for application data that fits a managed document or NoSQL model.
- Use it when predictable scale and low-latency reads or writes matter.
- Use it when the data model and access paths are better suited to Cosmos DB than to a relational store.
- Use it when the team wants managed scale without owning database infrastructure.

## When Not to Use It

- Do not choose it before understanding the data model and access patterns.
- Do not assume it behaves like a relational database with the same design tradeoffs.
- Do not select it only because it is managed if the workload really needs relational joins, transactions, or reporting patterns better served elsewhere.

## Common Mistakes

- Choosing a poor partition key and creating uneven traffic or hot partitions.
- Treating request units as a billing detail instead of a design constraint.
- Building a schema and query model that really belongs in a relational database.
- Granting broad access through keys or wide roles when narrower identity-based access would work.
- Ignoring the cost effect of inefficient queries and throughput settings.

## Cloud Engineering Considerations

### Identity and Access

Use RBAC, managed identities, and scoped keys or tokens carefully so application access stays narrow. Data access should line up with service boundaries, not broad environment-level convenience.

### Networking

Review private endpoints and how application runtimes access the database across environments. Data access paths are part of system architecture, not just application wiring.

### Security

Protect sensitive data, review backup behavior, and make sure only the right services can access the account. Managed databases still need deliberate data-protection choices.

### Observability

Track request units, latency, failure behavior, and partition stress against expected usage patterns. Many runtime symptoms reflect modeling issues, not platform instability.

### Cost

Throughput configuration, storage, backups, and inefficient access patterns all affect cost. Good data modeling is one of the main cost controls.

## How This Fits Into Cloud Engineering

Database selection changes how an application scales, what it costs to run, and how easy it is to evolve. Cosmos DB fits well when the cloud engineering choices around partitioning, access patterns, and service boundaries are made intentionally.

## Related Projects

- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)

## Related Patterns

- [Serverless API](../patterns/serverless-api.md)

## Official References

- [Azure Cosmos DB documentation](https://learn.microsoft.com/en-us/azure/cosmos-db/)
- [Azure Cosmos DB overview](https://learn.microsoft.com/en-us/azure/cosmos-db/introduction)
