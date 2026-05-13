# Secret Manager

## Definition

Secret Manager is Google Cloud's managed service for storing and controlling access to secrets such as API keys, tokens, and credentials.

It matters because secret handling is not only about hiding values. Teams also need versioning, access control, rotation, and auditability so application credentials do not become scattered across code, pipelines, and runtime configuration.

## How It Is Commonly Used

Secret Manager is commonly used by Cloud Run services, Cloud Functions, scheduled jobs, and other Google Cloud workloads that need sensitive configuration at runtime. A service account is granted access to a narrow set of secrets, and the workload retrieves those values only when needed.

This service is most effective when it is combined with careful IAM design and a clear decision about which secrets should exist at all.

## What To Pay Attention To

### Identity and Access

Use IAM and service accounts so each workload can only read the secrets it actually needs. Secret access should be as narrow as the application's real dependency set.

### Networking

Review how workloads retrieve secrets and whether network restrictions or private service access affect that path. Secret retrieval is part of runtime behavior, not a separate administrative action.

### Security

Treat secret lifecycle, rotation, and access logging as part of the security design. The question is not only where a secret is stored but also how it is changed and who can observe that change.

### Operations and Observability

Track secret access failures and unusual access patterns so credential issues are visible quickly. Authentication problems often appear first as application errors rather than explicit secret-manager incidents.

### Cost

Stored versions and access operations add cost, so clean up unused secrets and avoid noisy access loops.

## Common Mistakes

- Giving one service account access to too many secrets.
- Treating secret storage as a substitute for rotation discipline.
- Leaving unused secret versions and old integration credentials in place indefinitely.
- Forgetting that secret retrieval failures can become application outages.

## How This Fits Into Cloud Engineering

Secret Manager sits at the center of identity, runtime configuration, and operational safety. Cloud engineers need to treat secret handling as part of the architecture because it directly affects deployment quality, incident response, and blast radius.

## Related Projects

- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)
- [Project 03: Scheduled API Ingestion](../projects/project-03-scheduled-api-ingestion.md)
- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)

## Related Patterns

- [Serverless API](../patterns/serverless-api.md)
- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

## Official References

- [Secret Manager documentation](https://cloud.google.com/secret-manager/docs)
- [Secret Manager overview](https://cloud.google.com/secret-manager/docs/overview)
