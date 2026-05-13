# Microsoft Entra ID

## Purpose

Microsoft Entra ID provides identity management for users, groups, applications, and federated access in Azure.

## Definition

Microsoft Entra ID is the identity platform behind workforce sign-in, application registration, service principals, and much of Azure's access model. It is where organizations manage who can authenticate, how applications establish identity, and how those identities connect to access decisions.

The important distinction is that identity in Azure is not only about logging in to a portal. It also shapes how applications authenticate, how automation runs, and how services receive permissions through related mechanisms such as managed identities and Azure RBAC.

In simple terms:

> Entra ID is the identity system that lets Azure know who a person or workload is before any access decision can happen.

## What Problem It Solves

Without a central identity platform, teams end up with fragmented authentication, inconsistent application access patterns, and weak control over privileged administration. Entra ID gives Azure environments a common identity layer for users, groups, applications, federation, and security controls such as MFA and conditional access.

## How It Is Commonly Used

Entra ID commonly appears in work such as:

- workforce identity and single sign-on,
- application registrations for custom services and APIs,
- service principals and managed identity-backed workload access,
- conditional access and MFA policy enforcement,
- B2B or cross-organization access scenarios.

## When to Use It

- Use it for human identity, application registration, and authentication flows.
- Use it when Azure services and applications need centralized identity control.
- Use it as the foundation for workload and user access design across Azure.
- Use it when you need a durable model for privileged administration and sign-in governance.

## When Not to Use It

- Do not treat identity as only a sign-in problem.
- Do not confuse Entra ID authentication with Azure resource authorization. Authentication and RBAC are connected, but they are not identical.
- Do not rely on overly broad tenant-wide privileges when scoped roles or application-specific access paths are possible.

## Common Mistakes

- Confusing tenant roles, application roles, and Azure resource roles.
- Leaving too many standing administrators with broad directory privileges.
- Creating app registrations and client secrets without ownership, rotation, or cleanup discipline.
- Ignoring managed identities and falling back to static credentials too quickly.
- Treating MFA or conditional access as optional add-ons instead of basic control points.

## Cloud Engineering Considerations

### Identity and Access

Design users, groups, applications, and workload identities deliberately. Good cloud engineering depends on being able to explain which identities are interactive, which are automated, and which privileges they actually need.

### Networking

Identity is a control-plane layer, but network restrictions and application exposure still affect how identities are used. Internal applications, public APIs, and admin access paths rarely need identical treatment.

### Security

Use conditional access, MFA, strong application registration practices, and managed identities wherever they fit. Identity is one of the highest-leverage security layers in Azure.

### Observability

Review sign-in logs, audit behavior, application authentication failures, and privileged role usage as part of normal security and platform operations.

### Cost

Licensing tiers can matter depending on governance, conditional access, identity protection, and access review needs. Identity design can therefore have both security and licensing implications.

## How This Fits Into Cloud Engineering

Cloud systems are built out of services, but they are governed through identity. In Azure, Entra ID often sits at the center of secure access design, application authentication, and privileged administration. If identity is weak, the rest of the platform design becomes fragile.

## Related Projects

- [Project 01: Static Site](../projects/project-01-static-site.md)
- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)
- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)

## Related Patterns

- [Static Site](../patterns/static-site.md)
- [Serverless API](../patterns/serverless-api.md)

## Official References

- [Microsoft Entra ID documentation](https://learn.microsoft.com/en-us/entra/fundamentals/)
- [What is Microsoft Entra ID](https://learn.microsoft.com/en-us/entra/fundamentals/whatis)
