# IAM and Service Accounts

## Purpose

Google Cloud IAM and service accounts control who or what can access Google Cloud resources and what actions they can perform.

They are used when teams need to separate human access from workload access, define permission boundaries, and make cloud systems safe enough to operate at scale.

## Definition

Google Cloud IAM is the policy model for granting permissions to users, groups, and workloads. Service accounts are identities used by applications and automation instead of by human operators.

These two ideas have to be understood together because they answer different parts of the same problem. IAM decides what permissions exist and where they are granted. Service accounts provide a runtime identity that can receive those permissions without sharing human credentials.

IAM and service accounts are not the same thing as end-user authentication inside an application. They control access to Google Cloud resources. That boundary matters because many security and architecture mistakes begin when platform identity, deployment identity, and application user identity are treated as if they were interchangeable.

In simple terms:

> IAM decides what is allowed in Google Cloud, and service accounts give workloads a safe way to act inside those rules.

## What Problem It Solves

IAM and service accounts solve the problem of uncontrolled trust. Without clear identity design, projects drift toward broad editor permissions, shared credentials, and unclear boundaries between human administrators and application workloads.

They give Google Cloud environments a way to separate human responsibilities from runtime behavior, narrow access to only what is needed, and reduce the need for long-lived keys embedded in scripts or applications.

That does not make identity simple. Engineers still need to choose the right scope for permissions, define which identities should exist at all, and understand how cross-project access, inheritance, and impersonation work.

## How It Is Commonly Used

This model is commonly used for:

- granting user and group permissions through IAM roles,
- giving Cloud Run, Cloud Functions, and other services a workload identity,
- separating deployment identities from runtime identities,
- handling cross-project access between applications and shared services,
- reducing the need for long-lived keys.

In a practical Google Cloud architecture, IAM and service accounts are the trust boundary behind almost every other service decision. They determine what a runtime can read, write, invoke, or administer.

## Foundational Concepts Connected to IAM and Service Accounts

IAM and service accounts connect directly to several cloud engineering foundations.

### Identity and Access

This is the main access-control layer for Google Cloud resources. Good identity design starts with least privilege, clear ownership, and clean separation between humans and workloads.

### Governance

IAM is where organizational policy becomes enforceable. Roles, project boundaries, inherited permissions, and administrative responsibilities all affect the control model.

### Workload Identity

Service accounts let applications act as managed identities instead of carrying static credentials. That changes how runtimes, pipelines, and automation should be designed.

### Security

Many cloud security failures are really identity failures: excessive roles, forgotten service account keys, or weak separation between deployment and runtime access.

### Observability

Permission changes, denied access, impersonation, and privileged actions all need to be visible enough for troubleshooting and audit.

## When to Use It

Use IAM and service accounts whenever users, applications, or automation need access to Google Cloud resources.

Good use cases include:

- scoping human access with roles and groups,
- giving workloads a runtime identity,
- separating deployment permissions from runtime permissions,
- enabling cross-project access without sharing user credentials,
- reducing or eliminating the need for static keys.

They should be treated as a default design concern, not as cleanup work added after the first permissions incident.

## When Not to Use It

Do not use broad project-wide permissions when narrower roles are enough. Convenience is not a design principle.

Do not rely on unmanaged service account keys when platform-managed identity paths are available. Long-lived keys recreate the same secret-management problem that service accounts are supposed to reduce.

Do not confuse Google Cloud resource identity with your application's end-user identity model. A user who can sign in to your app is not automatically someone who should administer the cloud resources behind it.

## Compare To

### IAM and Service Accounts vs. Secret Manager

IAM and service accounts control identity and permissions. They decide who or what can access a resource.

Secret Manager stores sensitive values such as tokens, passwords, or API keys. Teams often use both together, but they solve different problems. Secret Manager does not replace least privilege, and IAM does not replace secret storage.

### IAM and Service Accounts vs. Application User Authentication

Application user authentication answers who the end user is inside your product.

IAM and service accounts answer who can access Google Cloud resources and which workloads are allowed to act on the platform. Mixing those layers leads to weak architecture and confusing trust boundaries.

## Tradeoffs

IAM's biggest advantage is control. It gives teams a structured way to express least privilege and separation of responsibility across a cloud environment.

The tradeoff is complexity. Fine-grained access control requires teams to think clearly about roles, resource scope, inheritance, and operational ownership.

Service accounts reduce the need for static credentials, which is a major advantage. The tradeoff is that poorly designed service-account usage can still create broad blast radius if multiple workloads share the same identity or if powerful roles are granted lazily.

Another tradeoff is that identity decisions are rarely visible to end users. That makes them easy to postpone, even though they often determine whether the platform is secure and explainable.

## Common Mistakes

- Leaving default service accounts with more access than the workload actually needs.
- Granting powerful roles at the project level because it is faster than designing a narrower policy.
- Reusing one service account across multiple unrelated services.
- Creating and forgetting service account keys that effectively become long-lived secrets.
- Mixing deployment permissions and runtime permissions into the same identity.
- Assuming permission problems can be debugged later without documenting who owns what access path.

## Cloud Engineering Considerations

### Identity and Access

Review IAM bindings, service account usage, and the difference between acting as a workload and administering the platform. Clear identity separation makes systems easier to secure and easier to explain.

### Networking and Project Boundaries

Identity is separate from networking, but project structure, ingress decisions, and cross-project access patterns still shape the design. Secure networking often depends on correct identity assumptions.

### Security and Governance

Avoid unnecessary service account keys, review permission inheritance, and keep powerful roles tightly limited. Workload identity should reduce secret sprawl, not recreate it.

### Observability

Monitor access changes, failed authorizations, impersonation behavior, and workload identity usage so permission problems surface before they become outages or security issues.

### Reliability

Identity design affects reliability because systems fail when they cannot authenticate or authorize correctly. Clear role assignment, documented ownership, and predictable service-account usage reduce permission-related outages.

### Cost

IAM is not a major direct cost driver, but weak access controls can contribute to operational waste, overspending, or risky experimentation elsewhere.

## Project and Pattern Connections

IAM and service accounts are most directly connected to:

- [Project 01: Static Site](../projects/project-01-static-site.md)
- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)
- [Project 03: Scheduled API Ingestion](../projects/project-03-scheduled-api-ingestion.md)
- [Project 04: Analytics Platform](../projects/project-04-analytics-platform.md)
- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)
- [Static Site](../patterns/static-site.md)
- [Serverless API](../patterns/serverless-api.md)
- [Scheduled Job](../patterns/scheduled-job.md)
- [Analytics Platform](../patterns/analytics-platform.md)
- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

They appear across almost every Google Cloud design because identity is the control plane behind deployment, runtime behavior, and platform trust.

## Official References

- [Google Cloud IAM overview](https://cloud.google.com/iam/docs/overview)
- [Service accounts overview](https://cloud.google.com/iam/docs/service-account-overview)
- [Understanding roles](https://cloud.google.com/iam/docs/understanding-roles)
