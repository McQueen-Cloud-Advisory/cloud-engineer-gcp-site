# AWS Secrets Manager

## Definition

AWS Secrets Manager is AWS's managed service for storing, accessing, and rotating sensitive application values such as API keys, tokens, and connection details.

It matters because secret handling is not only about hiding values from source control. Teams also need controlled runtime retrieval, auditability, rotation, and a clear boundary between operator access and application access.

## How It Is Commonly Used

AWS Secrets Manager is commonly used by Lambda functions, App Runner services, scheduled jobs, and other workloads that need credentials at runtime. In a good design, the workload role can read only the specific secret versions it needs, while operators and pipelines have separate responsibilities.

This service is especially useful when the team wants secret lifecycle management to be part of the platform rather than a set of ad hoc scripts and environment variables.

## What To Pay Attention To

### Identity and Access

Grant workloads access to the smallest set of secrets possible and separate operator access from runtime access. Secrets should map to application needs, not to broad environment-wide permissions.

### Networking

Review how workloads retrieve secrets and whether private connectivity or network controls are required. Secret retrieval is part of runtime dependency behavior.

### Security

Treat secret naming, rotation, and access logging as part of the security design, not just storage details. A secret that never rotates or is shared too widely remains a risk even in a managed platform.

### Operations and Observability

Monitor secret retrieval failures and unusual access patterns so credential issues are visible quickly. Many application outages that look like code bugs are really secret-access problems.

### Cost

Stored secrets and API usage add cost, so remove unused secrets and avoid unnecessary retrieval loops.

## Common Mistakes

- Giving workloads access to many secrets they do not use.
- Treating secret storage as a replacement for rotation discipline.
- Hard-coding fallback credentials because runtime retrieval was not designed well.
- Ignoring secret-access failures until they become production incidents.

## How This Fits Into Cloud Engineering

Secrets Manager matters because credential handling sits between application architecture, security, and operations. Cloud engineers need to make secret access reliable and narrow enough that the platform stays both usable and defensible.

## Related Projects

- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)
- [Project 03: Scheduled API Ingestion](../projects/project-03-scheduled-api-ingestion.md)
- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)

## Related Patterns

- [Serverless API](../patterns/serverless-api.md)
- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

## Official References

- [AWS Secrets Manager documentation](https://docs.aws.amazon.com/secretsmanager/)
- [Secrets Manager user guide](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html)
