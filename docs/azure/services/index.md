# Azure Services

## Purpose

This section explains Azure services from a practical cloud engineering perspective and focuses on the services that matter most in the current learning path.

## How To Read This Section

These pages are here to help you understand how Azure services participate in real systems, not to summarize the full Microsoft catalog. Use them to answer a smaller set of questions.

- What role does this service play in an Azure workload?
- How does it interact with identity, subscriptions, resource groups, and observability?
- What tradeoffs come with choosing it?
- When should it appear in a project or pattern?

This is why the section works best when you read it alongside a project or pattern instead of as a long isolated reading list.

## What This Service Set Is Designed To Teach

The Azure path emphasizes the services most useful for learning identity-centered governance, application hosting, secrets, data, monitoring, and AI-oriented extension inside Azure.

## Service Categories

### Identity and Access

- [Entra ID](entra-id.md)
- [Role-Based Access Control](role-based-access-control.md)
- [Managed Identities](managed-identities.md)

### Storage and Data

- [Blob Storage](blob-storage.md)
- [Cosmos DB](cosmos-db.md)
- [Data Factory](data-factory.md)
- [Microsoft Fabric](microsoft-fabric.md)

### Compute and Application Hosting

- [Functions](functions.md)
- [Container Apps](container-apps.md)
- [Container Registry](container-registry.md)

### Integration and Scheduling

- [API Management](api-management.md)

### Observability and Operations

- [Monitor](monitor.md)
- [Application Insights](application-insights.md)

### AI and Agentic Platforms

- [Microsoft Foundry](microsoft-foundry.md)
- [Foundry Agent Service](foundry-agent-service.md)
- [Azure OpenAI](azure-openai.md)
- [Azure AI Search](azure-ai-search.md)
- [Content Safety](content-safety.md)

### Security and Secrets

- [Key Vault](key-vault.md)

## How The Categories Connect

Azure systems often begin with identity and governance decisions. Entra ID, RBAC, and managed identities define how both people and workloads get access. Storage, compute, APIs, and data services then sit inside that control model. Monitor and Application Insights give you the operational view. Later, Foundry, Azure OpenAI, and AI Search extend the same application path into AI workloads.

## How This Fits Into Cloud Engineering

Cloud engineers need service knowledge that helps them design and operate systems, not just browse product names. This section is designed to help you explain why a service belongs in an Azure architecture and what operating responsibility comes with it.

## Official References

- [Microsoft Learn for Azure](https://learn.microsoft.com/en-us/azure/)
- [Azure Well-Architected Framework](https://learn.microsoft.com/en-us/azure/well-architected/)
