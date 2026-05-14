# Azure Role-Based Access Control

## Purpose

Azure RBAC is the authorization system that determines which identities can perform which actions on Azure resources and at what scope.

## Definition

It matters because identity without authorization is not enough. Cloud engineers need a way to decide who can deploy, who can operate, which workloads can access which resources, and how those permissions inherit across management groups, subscriptions, resource groups, and resources.

RBAC is not the same thing as authentication. It takes an identity that already exists and decides what that identity is allowed to do. That distinction matters because many Azure access mistakes happen when teams understand who can sign in but not how scope and inheritance change what they can actually change.

In simple terms:

> Azure RBAC is the permission system that answers what a user or workload can do in Azure and where those permissions apply.

## What Problem It Solves

Azure RBAC solves the problem of needing controlled, least-privilege authorization across a hierarchy of Azure resources.

It gives teams a structured way to separate deployment access, operator access, and runtime access rather than relying on broad standing privileges.

That does not remove engineering responsibility. Engineers still need to understand scope, inheritance, role choice, and how privilege boundaries align with real service ownership.

## How It Is Commonly Used

Azure RBAC is commonly used to assign least-privilege access to people, groups, service principals, and managed identities. It is the main way teams separate deployment responsibilities from runtime access and keep different environments or resource groups from collapsing into one broad permission set.

RBAC is especially important in Azure because scope and inheritance shape the real blast radius of every access decision.

## Foundational Concepts Connected to Azure RBAC

Azure RBAC connects directly to several cloud engineering foundations.

### Authorization

RBAC is the main Azure authorization model in this learning path. It turns identity into actual resource permissions.

### Scope and Inheritance

Management groups, subscriptions, resource groups, and individual resources form the hierarchy where permissions expand or narrow. Understanding that hierarchy is part of safe Azure design.

### Separation of Duties

Deployment paths, operational access, and workload identities should not all share the same role assignments. RBAC is one of the main ways teams enforce that separation.

### Governance

Role assignments influence how safely environments are managed over time. Poorly structured authorization creates long-term governance and audit problems.

### Observability and Audit

Permission changes and authorization failures need to be visible because access problems often surface first as broken deployments, failed applications, or risky admin actions.

## When to Use It

Use Azure RBAC any time a human user, service principal, managed identity, or automation path needs Azure resource access.

Good use cases include:

- separating deployment permissions from runtime permissions,
- scoping access by subscription, resource group, or resource,
- limiting operator access to specific environments,
- granting workloads only the Azure actions they actually need.

RBAC is strongest when the permission model mirrors real ownership and operational responsibility.

## When Not to Use It

Do not use broad built-in roles as a shortcut when a narrower role or narrower scope will work.

Do not confuse RBAC with network protection or application-level authorization. It is one part of the control model, not the whole security posture.

Do not treat a single successful deployment as proof that the permission model is well designed.

## Compare To

### Azure RBAC vs. Entra ID

Entra ID authenticates the user or workload.

Azure RBAC authorizes what that identity can do on Azure resources.

### Azure RBAC vs. Application Authorization

RBAC controls access to Azure resources and Azure control-plane or data-plane actions.

Application authorization controls what your users can do inside your application. These are related but separate decisions.

## Tradeoffs

Azure RBAC's biggest advantage is precision across Azure's resource hierarchy. Teams can grant specific permissions at scopes that align with ownership and operational needs.

The tradeoff is complexity. Scope inheritance and built-in role breadth can create access outcomes that are less obvious than they first appear.

RBAC also makes least-privilege design practical. That is valuable, but only if teams review assignments over time instead of letting broad roles accumulate.

Another tradeoff is that authorization models age badly when teams change structure, subscriptions, or workload boundaries without rethinking permissions.

## Common Mistakes

- Granting Owner when a narrower role would work.
- Forgetting how inheritance changes the effective permission set.
- Mixing human, deployment, and runtime access in the same identity path.
- Treating role assignment as finished work instead of something that needs review.
- Assuming a broad resource-group assignment is harmless because it feels smaller than subscription scope.

## Cloud Engineering Considerations

### Identity and Access

Review scope, inheritance, and role definition carefully before assigning access. A seemingly small assignment at a broad scope can affect far more resources than intended.

### Networking

RBAC is not a network boundary, so it should work alongside network controls rather than replace them.

### Security

Least privilege, periodic access review, and separation of duties matter as much as the role definitions themselves. Good authorization design is continuous, not one-time setup.

### Operations and Observability

Track role changes and authorization failures so access problems are visible quickly. Many production issues come from recent permission changes rather than from service outages.

### Reliability

Reliable Azure operations depend on predictable permissions. Deployments fail, services break, and incident response slows down when role assignments are unclear, oversized, or inconsistent across environments.

### Cost

RBAC does not add meaningful direct cost, but poor access control can contribute to accidental spend and operational risk.

## Project and Pattern Connections

Azure RBAC is foundational across the Azure learning path and is most directly connected to:

- [Project 01: Static Site](../projects/project-01-static-site.md)
- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)
- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)
- [Static Site](../patterns/static-site.md)
- [Serverless API](../patterns/serverless-api.md)
- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

It matters in every Azure project because every deployment path and every runtime permission set depends on it.

## Official References

- [Azure RBAC documentation](https://learn.microsoft.com/en-us/azure/role-based-access-control/)
- [Azure RBAC overview](https://learn.microsoft.com/en-us/azure/role-based-access-control/overview)
- [Understand scope for Azure RBAC](https://learn.microsoft.com/en-us/azure/role-based-access-control/scope-overview)
