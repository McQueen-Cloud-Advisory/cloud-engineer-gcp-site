# Amazon DynamoDB

## Purpose

Amazon DynamoDB is a managed NoSQL database for high-scale, low-latency key-value and document workloads.

It is used when an application needs predictable request-time data access without operating database servers, but the access patterns can be designed around keys and known query paths.

## Definition

DynamoDB is a managed database built around key-based access patterns rather than around relational joins and general-purpose SQL modeling. It is designed for applications that need predictable latency, flexible scale, and a managed operational model.

That does not make it a universal database choice. DynamoDB works best when the access patterns are understood and the table design is built around them deliberately.

DynamoDB is not a drop-in replacement for relational database design. That boundary matters because many DynamoDB problems begin when teams choose it for its scale and management benefits before deciding whether the workload actually fits a key-value or document model.

In simple terms:

> DynamoDB is a fast application database for workloads that know how they need to read and write data.

## What Problem It Solves

DynamoDB gives teams an application data store that can scale well without running database servers, patching clusters, or manually handling replication and high availability. It is useful when relational features are not the main requirement and low-latency access matters.

That does not remove engineering responsibility. Engineers still need to model access patterns, partition keys, write paths, secondary indexes, and cost behavior intentionally.

## How It Is Commonly Used

DynamoDB is commonly used for:

- application state and metadata,
- request-driven backend data for APIs,
- user profiles, sessions, preferences, or workflow state,
- event-driven systems that need durable key-value access,
- serverless architectures that want a managed data layer.

In many AWS workloads, DynamoDB is the request-time data store behind Lambda functions and API Gateway, where the table design is shaped by the API contract rather than by traditional normalized schemas.

## Foundational Concepts Connected to DynamoDB

DynamoDB connects directly to several cloud engineering foundations.

### Application Data Modeling

DynamoDB design starts from access patterns. The way the application reads and writes data should shape the table, not the other way around.

### Performance and Scale

Partition design, key choice, and traffic distribution determine whether the database behaves predictably under load.

### Identity and Access

Application runtimes should get only the table and action permissions they need. Data access should be scoped to the service boundary.

### Observability

Throttling, latency, and failed requests reveal whether the data model and traffic shape actually match the design assumptions.

### Cost Management

Read patterns, write patterns, streams, backups, and capacity choices all affect cost. Inefficient data models become operational cost problems.

## When to Use It

- Use it for request-driven application state and metadata.
- Use it when access patterns are known and fit a key-value or document model.
- Use it when you need managed scaling and high availability without operating database servers.
- Use it when a serverless or event-driven architecture benefits from a fully managed low-latency store.

DynamoDB is strongest when the team can clearly explain how the application will access its data.

## When Not to Use It

- Do not use it before you understand the access patterns and key design.
- Do not assume it is a drop-in replacement for relational querying and joins.
- Do not choose it only because it is managed if the workload really needs relational behavior.

## Compare To

### DynamoDB vs. S3

S3 is object storage for files, assets, and raw data.

DynamoDB is an operational application database for low-latency reads and writes. They solve very different storage problems.

### DynamoDB vs. Relational Databases

Relational databases are better when the workload depends on joins, relational constraints, and ad hoc query flexibility.

DynamoDB is better when the workload can be modeled around known key-based access patterns and needs a fully managed, low-latency operational store.

## Tradeoffs

DynamoDB's biggest advantage is predictable managed scale for request-driven application data. Teams can serve low-latency workloads without taking on database host operations.

The tradeoff is modeling discipline. DynamoDB works well when the access patterns are clear, but it becomes painful when teams try to treat it like a general relational database.

DynamoDB also makes it easy to scale application storage quickly. That is useful, but it can hide poor table design until throttling or cost makes the mistakes visible.

Another tradeoff is that schema flexibility does not remove the need for structure. Weak modeling still creates operational complexity.

## Common Mistakes

- Designing the table before identifying the read and write patterns.
- Using scans where targeted key access should exist.
- Choosing poor partition keys that create hot partitions or uneven traffic.
- Granting broad table permissions when narrow item or table access would be better.
- Ignoring the cost impact of inefficient access patterns or unused capacity settings.
- Treating the first working table design as if it is automatically production-ready.

## Cloud Engineering Considerations

### Identity and Access

Limit read and write permissions to the specific tables and actions each workload needs. Data access should be scoped to the service boundary, not to the convenience of the development team.

### Networking

Review how application runtimes reach DynamoDB and whether private access patterns, regional architecture, or cross-account access paths matter.

### Security

Protect sensitive attributes, enable encryption, and review access paths for exports, streams, and administrative actions.

### Observability

Monitor throttling, latency, failed requests, and capacity behavior against expected access patterns. Many DynamoDB problems are actually data-model problems surfacing as runtime symptoms.

### Reliability

Reliable DynamoDB design depends on stable key patterns, safe retry behavior, and understanding how the application reacts when reads are stale, writes fail, or traffic concentrates unevenly.

### Cost

Read and write volume, storage, backups, streams, and inefficient access design all affect cost. Good data modeling is also cost modeling.

## Project and Pattern Connections

Amazon DynamoDB is most directly connected to:

- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)
- [Project 03: Scheduled API Ingestion](../projects/project-03-scheduled-api-ingestion.md)
- [Serverless API](../patterns/serverless-api.md)

It matters when an AWS workload needs a real application database, not just durable file storage.

## Official References

- [Amazon DynamoDB documentation](https://docs.aws.amazon.com/dynamodb/)
- [DynamoDB developer guide](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Introduction.html)
- [Best practices for designing and using partition keys effectively](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-partition-key-design.html)
