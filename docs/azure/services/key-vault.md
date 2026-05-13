# Azure Key Vault

## Definition

Azure Key Vault is Azure's managed service for storing and controlling access to secrets, keys, and certificates.

It matters because secret handling is not only a storage problem. Teams need to decide who can read sensitive values, how workloads retrieve them, how rotation works, and how access is audited over time.

## How It Is Commonly Used

Azure Key Vault is commonly used to hold application secrets, certificates, and integration credentials for Functions, Container Apps, automation pipelines, and other Azure services. In a well-designed environment, workloads retrieve only the values they need at runtime, usually through managed identities rather than embedded credentials.

Key Vault is most useful when it is part of a broader identity and deployment design instead of an isolated security checkbox.

## What To Pay Attention To

### Identity and Access

Use RBAC or access policies deliberately and prefer managed identities for runtime secret access. Vault-wide read access is usually broader than the workload requires.

### Networking

Review private endpoints and how workloads reach the vault across environments. Secret access failures often show up as application startup problems or intermittent runtime errors.

### Security

Treat secret naming, rotation, and access logging as part of the security design, not just storage details. A secret that exists but is never rotated or reviewed is still a risk.

### Operations and Observability

Track access failures and unusual access patterns so credential problems surface quickly. Secret platforms are operational dependencies, not background utilities.

### Cost

Operation volume and premium features can increase cost, so remove unused secrets and avoid noisy access patterns.

## Common Mistakes

- Granting broad vault access when a workload needs only one or two values.
- Using the vault without a clear rotation plan.
- Treating Key Vault as a reason to ignore authorization design elsewhere.
- Forgetting that networking restrictions can break secret retrieval in production.

## How This Fits Into Cloud Engineering

Key Vault sits at the center of identity, secret management, deployment safety, and runtime access. Cloud engineers need to treat it as part of the system design because credential handling shapes both security posture and operational reliability.

## Related Projects

- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)
- [Project 03: Scheduled API Ingestion](../projects/project-03-scheduled-api-ingestion.md)
- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)

## Related Patterns

- [Serverless API](../patterns/serverless-api.md)
- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

## Official References

- [Azure Key Vault documentation](https://learn.microsoft.com/en-us/azure/key-vault/)
- [Azure Key Vault overview](https://learn.microsoft.com/en-us/azure/key-vault/general/overview)
