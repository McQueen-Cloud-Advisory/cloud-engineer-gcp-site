# Azure Role-Based Access Control

## Definition

Azure RBAC is the authorization system that determines which identities can perform which actions on Azure resources and at what scope.

It matters because identity without authorization is not enough. Cloud engineers need a way to decide who can deploy, who can operate, which workloads can access which resources, and how those permissions inherit across management groups, subscriptions, resource groups, and resources.

## How It Is Commonly Used

Azure RBAC is commonly used to assign least-privilege access to people, groups, service principals, and managed identities. It is the main way teams separate deployment responsibilities from runtime access and keep different environments or resource groups from collapsing into one broad permission set.

RBAC is especially important in Azure because scope and inheritance shape the real blast radius of every access decision.

## What To Pay Attention To

### Identity and Access

Review scope, inheritance, and role definition carefully before assigning access. A seemingly small assignment at a broad scope can affect far more resources than intended.

### Networking

RBAC is not a network boundary, so it should work alongside network controls rather than replace them.

### Security

Least privilege, periodic access review, and separation of duties matter as much as the role definitions themselves. Good authorization design is continuous, not one-time setup.

### Operations and Observability

Track role changes and authorization failures so access problems are visible quickly. Many production issues come from recent permission changes rather than from service outages.

### Cost

RBAC does not add meaningful direct cost, but poor access control can contribute to accidental spend and operational risk.

## Common Mistakes

- Granting Owner when a narrower role would work.
- Forgetting how inheritance changes the effective permission set.
- Mixing human, deployment, and runtime access in the same identity path.
- Treating role assignment as finished work instead of something that needs review.

## How This Fits Into Cloud Engineering

Azure RBAC is foundational cloud engineering because it defines who can change the platform and what workloads are allowed to do inside it. Without clear authorization design, the rest of the architecture becomes hard to secure and hard to operate.

## Related Projects

- [Project 01: Static Site](../projects/project-01-static-site.md)
- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)
- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)

## Related Patterns

- [Static Site](../patterns/static-site.md)
- [Serverless API](../patterns/serverless-api.md)
- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

## Official References

- [Azure RBAC documentation](https://learn.microsoft.com/en-us/azure/role-based-access-control/)
- [Azure RBAC overview](https://learn.microsoft.com/en-us/azure/role-based-access-control/overview)
