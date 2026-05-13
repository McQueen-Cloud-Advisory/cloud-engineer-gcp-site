# Managed Identities

## Definition

Managed identities give Azure workloads a platform-managed identity for accessing other Azure resources.

They matter because service-to-service access is one of the easiest places for credential sprawl to grow. Managed identities reduce the need to store secrets, but they only help if role assignments and resource access are designed carefully.

## How It Is Commonly Used

Managed identities are commonly used by Azure Functions, Container Apps, automation workflows, and other Azure services that need to reach resources such as Key Vault, Storage, or databases. Instead of embedding a credential, the workload authenticates as its managed identity and authorization is determined through Azure RBAC or another supported access model.

This is one of the most important Azure-native patterns to understand because it connects identity design directly to application architecture.

## What To Pay Attention To

### Identity and Access

Treat managed identities as first-class workload identities and review every role assignment they receive. System-assigned and user-assigned identities both reduce secret handling, but neither removes the need for least privilege.

### Networking

Managed identities are not a network control, so review how the workload still reaches the target service. Identity success does not guarantee network success.

### Security

They reduce secret exposure, but authorization scope still determines blast radius. Over-privileged managed identities can become silent platform risk.

### Operations and Observability

Track access failures and permission changes as part of normal platform operations. Identity problems often surface as application or pipeline issues first.

### Cost

Managed identities are primarily a security and operational benefit rather than a direct cost driver.

## Common Mistakes

- Treating managed identity as equivalent to least privilege.
- Reusing one identity for many unrelated workloads.
- Forgetting to review inherited access through broad RBAC scope.
- Assuming an access problem must be networking or code when the role assignment is wrong.

## How This Fits Into Cloud Engineering

Managed identities are central to Azure cloud engineering because they connect platform identity, workload design, authorization scope, and runtime behavior. Understanding them well leads to safer systems and simpler secret management.

## Related Projects

- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)
- [Project 03: Scheduled API Ingestion](../projects/project-03-scheduled-api-ingestion.md)
- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)

## Related Patterns

- [Serverless API](../patterns/serverless-api.md)
- [Scheduled Job](../patterns/scheduled-job.md)
- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

## Official References

- [Managed identities documentation](https://learn.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/overview)
- [Managed identities overview](https://learn.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/overview)
