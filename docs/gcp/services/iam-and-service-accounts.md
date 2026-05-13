# IAM and Service Accounts

## Purpose

Google Cloud IAM and service accounts control who or what can access Google Cloud resources and what actions they can perform.

## Definition

Google Cloud IAM is the policy model for granting permissions to users, groups, and workloads. Service accounts are identities used by applications and automation instead of by human operators.

These two ideas have to be understood together because they answer different parts of the same problem. IAM decides what permissions exist and where they are granted. Service accounts provide a runtime identity that can receive those permissions without sharing human credentials.

In simple terms:

> IAM decides what is allowed in Google Cloud, and service accounts give workloads a safe way to act inside those rules.

## What Problem It Solves

Without clear identity design, projects drift toward broad editor permissions, shared credentials, and unclear boundaries between human administrators and application workloads. IAM and service accounts give Google Cloud environments a way to separate those concerns cleanly.

## How It Is Commonly Used

This model is commonly used for:

- granting user and group permissions through IAM roles,
- giving Cloud Run, Cloud Functions, and other services a workload identity,
- separating deployment identities from runtime identities,
- handling cross-project access between applications and shared services,
- reducing the need for long-lived keys.

## When to Use It

- Use IAM roles to scope user and group permissions.
- Use service accounts for workload identity and automation.
- Use them to enforce least privilege across projects and services.
- Use them to separate human administration from application runtime access.

## When Not to Use It

- Do not use broad project-wide permissions when narrower roles are enough.
- Do not rely on unmanaged service account keys when platform-managed identity paths are available.
- Do not confuse cloud-resource identity with your application's end-user identity model.

## Common Mistakes

- Leaving default service accounts with more access than the workload actually needs.
- Granting powerful roles at the project level because it is faster than designing a narrower policy.
- Reusing one service account across multiple unrelated services.
- Creating and forgetting service account keys that effectively become long-lived secrets.
- Mixing deployment permissions and runtime permissions into the same identity.

## Cloud Engineering Considerations

### Identity and Access

Review IAM bindings, service account usage, and the difference between acting as a workload and administering the platform. Clear identity separation makes systems easier to secure and easier to explain.

### Networking

Identity is separate from networking, but project structure, ingress decisions, and cross-project access patterns still shape the design. Secure networking often depends on correct identity assumptions.

### Security

Avoid unnecessary service account keys, review permission inheritance, and keep powerful roles tightly limited. Workload identity should reduce secret sprawl, not recreate it.

### Observability

Monitor access changes, failed authorizations, and workload identity behavior so permission problems surface before they become outages or security issues.

### Cost

IAM is not a major direct cost driver, but weak access controls can contribute to operational waste, overspending, or risky experimentation elsewhere.

## How This Fits Into Cloud Engineering

Cloud systems are not only networks and runtimes. They are also trust relationships. In Google Cloud, IAM and service accounts are the layer that decides whether a deployment pipeline, application service, or engineer can safely do its job.

## Related Projects

- [Project 01: Static Site](../projects/project-01-static-site.md)
- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)
- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)

## Related Patterns

- [Static Site](../patterns/static-site.md)
- [Serverless API](../patterns/serverless-api.md)
- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

## Official References

- [Google Cloud IAM overview](https://cloud.google.com/iam/docs/overview)
- [Service accounts overview](https://cloud.google.com/iam/docs/service-account-overview)
