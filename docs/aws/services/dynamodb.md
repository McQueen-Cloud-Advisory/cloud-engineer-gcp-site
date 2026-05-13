# Amazon DynamoDB

## Purpose

Amazon DynamoDB is a managed NoSQL database for high-scale, low-latency key-value and document workloads.

## Definition

DynamoDB is a managed database built around key-based access patterns rather than around relational joins and general-purpose SQL modeling. It is designed for applications that need predictable latency, flexible scale, and a managed operational model.

That does not make it a universal database choice. DynamoDB works best when the access patterns are understood and the table design is built around them deliberately.

In simple terms:

> DynamoDB is a fast application database for workloads that know how they need to read and write data.

## What Problem It Solves

DynamoDB gives teams an application data store that can scale well without running database servers, patching clusters, or manually handling replication and high availability. It is useful when relational features are not the main requirement and low-latency access matters.

## How It Is Commonly Used

DynamoDB is commonly used for:

- application state and metadata,
- request-driven backend data for APIs,
- user profiles, sessions, preferences, or workflow state,
- event-driven systems that need durable key-value access,
- serverless architectures that want a managed data layer.

## When to Use It

- Use it for request-driven application state and metadata.
- Use it when access patterns are known and fit a key-value or document model.
- Use it when you need managed scaling and high availability without operating database servers.
- Use it when a serverless or event-driven architecture benefits from a fully managed low-latency store.

## When Not to Use It

- Do not use it before you understand the access patterns and key design.
- Do not assume it is a drop-in replacement for relational querying and joins.
- Do not choose it only because it is managed if the workload really needs relational behavior.

## Common Mistakes

- Designing the table before identifying the read and write patterns.
- Using scans where targeted key access should exist.
- Choosing poor partition keys that create hot partitions or uneven traffic.
- Granting broad table permissions when narrow item or table access would be better.
- Ignoring the cost impact of inefficient access patterns or unused capacity settings.

## Cloud Engineering Considerations

### Identity and Access

Limit read and write permissions to the specific tables and actions each workload needs. Data access should be scoped to the service boundary, not to the convenience of the development team.

### Networking

Review how application runtimes reach DynamoDB and whether private access patterns, regional architecture, or cross-account access paths matter.

### Security

Protect sensitive attributes, enable encryption, and review access paths for exports, streams, and administrative actions.

### Observability

Monitor throttling, latency, failed requests, and capacity behavior against expected access patterns. Many DynamoDB problems are actually data-model problems surfacing as runtime symptoms.

### Cost

Read and write volume, storage, backups, streams, and inefficient access design all affect cost. Good data modeling is also cost modeling.

## How This Fits Into Cloud Engineering

Database choice shapes how an application behaves under load, how it evolves, and what it costs to operate. DynamoDB is a strong fit when the cloud engineering decisions around data model, traffic shape, and service boundaries are clear.

## Related Projects

- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)

## Related Patterns

- [Serverless API](../patterns/serverless-api.md)

## Official References

- [Amazon DynamoDB documentation](https://docs.aws.amazon.com/dynamodb/)
- [DynamoDB developer guide](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Introduction.html)
