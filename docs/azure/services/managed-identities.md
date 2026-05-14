# Managed Identities

## Purpose

Managed identities give Azure workloads a platform-managed identity for accessing other Azure resources.

## Definition

They matter because service-to-service access is one of the easiest places for credential sprawl to grow. Managed identities reduce the need to store secrets, but they only help if role assignments and resource access are designed carefully.

Managed identities are not permission systems by themselves. They provide an identity for the workload, and authorization is usually granted later through Azure RBAC or a supported data-plane access model. That boundary matters because teams sometimes hear “managed identity” and assume the access design is finished.

In simple terms:

> Managed identities let Azure workloads authenticate as themselves without storing their own credentials.

## What Problem It Solves

Managed identities solve the problem of workloads needing credentials to reach Azure services without forcing teams to create, store, rotate, and secure those credentials manually.

They give Azure applications and automation a safer default identity path for service-to-service access.

That does not remove engineering responsibility. Engineers still need to scope permissions narrowly, choose between system-assigned and user-assigned identities carefully, and understand how the workload reaches the target service.

## How It Is Commonly Used

Managed identities are commonly used by Azure Functions, Container Apps, automation workflows, and other Azure services that need to reach resources such as Key Vault, Storage, or databases. Instead of embedding a credential, the workload authenticates as its managed identity and authorization is determined through Azure RBAC or another supported access model.

This is one of the most important Azure-native patterns to understand because it connects identity design directly to application architecture.

## Foundational Concepts Connected to Managed Identities

Managed identities connect directly to several cloud engineering foundations.

### Workload Identity

Managed identities are the main Azure-native workload identity pattern in this learning path. They let applications authenticate without embedding secrets in code or configuration.

### Least Privilege

The identity is only useful when its permissions are narrow. A managed identity with broad access is still a security and operational risk.

### Application Architecture

The runtime boundary and service dependencies determine where the identity should exist and what it should be allowed to do.

### Secret Reduction

Managed identities reduce credential sprawl, but they do not remove the need for Key Vault, RBAC, or careful dependency design.

### Observability and Troubleshooting

Identity-based failures often look like application or networking failures. Teams need clear ways to tell which part of the access path is broken.

## When to Use It

Use managed identities when an Azure-hosted workload needs to access Azure resources or supported downstream services without storing long-lived credentials.

Good use cases include:

- Functions reading from Blob Storage or Key Vault,
- Container Apps accessing secrets or databases,
- automation jobs interacting with Azure resources,
- AI workloads reaching search, storage, or secret-management services.

Managed identities are strongest when each workload has a clear service boundary and a narrowly scoped permission set.

## When Not to Use It

Do not assume managed identities replace authorization design. They only solve the credential part of the problem.

Do not reuse one identity for many unrelated applications just because it seems convenient.

Do not treat a working access path as proof that the permissions are least privilege.

## Compare To

### Managed Identities vs. Service Principals with Secrets

Service principals with client secrets require explicit secret storage and rotation.

Managed identities are the more Azure-native path when the workload is running on Azure and can use a platform-managed identity directly.

### Managed Identities vs. Azure RBAC

Managed identities provide the workload identity.

Azure RBAC decides what that identity can do.

## Tradeoffs

Managed identities' biggest advantage is secret reduction. They let teams remove one of the most common sources of credential sprawl in Azure applications.

The tradeoff is that identity simplicity can hide authorization complexity. A workload may authenticate cleanly while still having the wrong permissions, the wrong scope, or the wrong dependency design.

Managed identities also make Azure-native service integration safer. That is valuable, but it can encourage overuse of one shared identity unless teams enforce ownership boundaries.

Another tradeoff is that identity success does not guarantee network reachability or correct runtime behavior. The workload still has to reach the target service safely.

## Common Mistakes

- Treating managed identity as equivalent to least privilege.
- Reusing one identity for many unrelated workloads.
- Forgetting to review inherited access through broad RBAC scope.
- Assuming an access problem must be networking or code when the role assignment is wrong.
- Keeping fallback secrets in the application even after moving to managed identity.

## Cloud Engineering Considerations

### Identity and Access

Treat managed identities as first-class workload identities and review every role assignment they receive. System-assigned and user-assigned identities both reduce secret handling, but neither removes the need for least privilege.

### Networking

Managed identities are not a network control, so review how the workload still reaches the target service. Identity success does not guarantee network success.

### Security

They reduce secret exposure, but authorization scope still determines blast radius. Over-privileged managed identities can become silent platform risk.

### Operations and Observability

Track access failures and permission changes as part of normal platform operations. Identity problems often surface as application or pipeline issues first.

### Reliability

Reliable workload identity design depends on predictable role assignments, clear ownership of identities, and applications that fail safely when access is missing or misconfigured.

### Cost

Managed identities are primarily a security and operational benefit rather than a direct cost driver.

## Project and Pattern Connections

Managed identities are most directly connected to:

- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)
- [Project 03: Scheduled API Ingestion](../projects/project-03-scheduled-api-ingestion.md)
- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)
- [Serverless API](../patterns/serverless-api.md)
- [Scheduled Job](../patterns/scheduled-job.md)
- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

It matters whenever an Azure workload needs a real runtime identity instead of embedded credentials.

## Official References

- [Managed identities documentation](https://learn.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/overview)
- [Managed identities overview](https://learn.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/overview)
- [Managed identity types and how to choose them](https://learn.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/managed-identities-overview)
