# Azure Services

## Purpose

This section explains Azure services from a practical cloud engineering perspective and focuses on the services that power the current learning path.

It is meant to help you understand how services fit into real systems, not just memorize short definitions in isolation.

## Definition

This section is a curated guide to the Azure services used most directly across the site's projects, patterns, and provider-specific learning path.

It is not a full Azure catalog. It is a working set of services chosen because they help explain the main architectural decisions a cloud engineer has to make: identity, authorization, runtime, storage, data movement, observability, secrets, and AI platform design.

In simple terms:

> This section is the Azure service map for the systems this site teaches you to build.

## What Problem This Section Solves

Cloud service catalogs are broad, and that breadth can make early learning messy. People often read isolated service descriptions without understanding how those services work together in a real workload.

This section narrows the scope. It helps you answer practical questions such as:

- which service solves this problem,
- where that service belongs in an architecture,
- what tradeoffs come with choosing it,
- and what identity, networking, observability, and cost implications follow from that choice.

## How To Read This Section

These pages are intended to make the current provider path easier to understand by answering a focused set of questions.

- What problem does this service solve in Azure?
- How does it fit into a project or pattern?
- What identity, networking, and observability assumptions come with it?
- How does it compare to the surrounding runtime, data, and governance choices?

This section is most useful when you read it alongside a project or pattern rather than as a disconnected glossary.

A practical reading order is:

1. Start with a project or pattern.
2. Use the relevant service pages to understand the main building blocks.
3. Compare nearby services when a design choice is not obvious.
4. Return to the foundations pages when a service keeps depending on the same core concept.

## What This Service Set Is Designed To Teach

The Azure path emphasizes the services most useful for learning:

- identity-centered governance and authorization,
- application hosting and runtime choice,
- object storage and application data boundaries,
- data movement and analytics-platform design,
- monitoring and operational visibility,
- secret handling and safer runtime configuration,
- and how AI features fit into normal cloud engineering rather than sitting outside it.

## Service Categories

### Identity and Access

Azure systems start with trust boundaries. These pages explain how user identity, workload identity, and Azure resource authorization should be separated.

- [Entra ID](entra-id.md)
- [Role-Based Access Control](role-based-access-control.md)
- [Managed Identities](managed-identities.md)

### Application Hosting and Delivery

These services define how application code or containers are delivered. The main question here is whether the workload is a function, a managed containerized service, or a container delivery path.

- [Functions](functions.md)
- [Container Apps](container-apps.md)
- [Container Registry](container-registry.md)

### Storage and Data

These services hold files, application records, or broader analytics-platform components. The key boundary is whether the system needs durable objects, an operational database, a pipeline control plane, or a larger analytical platform.

- [Blob Storage](blob-storage.md)
- [Cosmos DB](cosmos-db.md)
- [Data Factory](data-factory.md)
- [Microsoft Fabric](microsoft-fabric.md)

### Integration and Scheduling

These services help systems receive API traffic or sit behind a governed API edge rather than exposing every backend directly.

- [API Management](api-management.md)

### Observability and Operations

These services help teams detect failures, watch health, and operate workloads like systems instead of demos.

- [Monitor](monitor.md)
- [Application Insights](application-insights.md)

### AI and Agentic Platforms

These pages explain how AI features become part of a real Azure application architecture, including model access, retrieval, safety controls, and managed agent behavior.

- [Microsoft Foundry](microsoft-foundry.md)
- [Foundry Agent Service](foundry-agent-service.md)
- [Azure OpenAI](azure-openai.md)
- [Azure AI Search](azure-ai-search.md)
- [Content Safety](content-safety.md)

### Security and Secrets

These services focus on sensitive runtime configuration and the operational discipline around secret access.

- [Key Vault](key-vault.md)

## How The Categories Connect

Typical Azure workloads in this site follow a pattern like this:

1. Entra ID, RBAC, and managed identities define who or what can act.
2. API Management, Functions, or Container Apps provide the entry and runtime boundary.
3. Blob Storage, Cosmos DB, and Key Vault support data, state, and sensitive configuration.
4. Data Factory and Microsoft Fabric extend the architecture into movement and analytics when the workload needs broader data-platform capabilities.
5. Azure Monitor and Application Insights provide the operational view.
6. Foundry, Azure OpenAI, Azure AI Search, and Content Safety build on top of that same application foundation rather than replacing it.

The important lesson is that Azure service choice is not a list-building exercise. It is a system-design exercise.

## What This Section Does Not Try To Do

This section does not try to summarize every Azure product or give certification-style one-liners for the whole platform.

It is intentionally narrower than that. The goal is to build judgment about which services belong in a practical architecture and why.

## Project and Pattern Connections

Use this section together with the Azure project and pattern pages:

- [Projects Overview](../projects/index.md)
- [Patterns Overview](../patterns/index.md)

## How This Fits Into Cloud Engineering

Cloud engineers need service knowledge that supports design, deployment, and operations. This section helps you explain how Azure services work together in a system and what tradeoffs they introduce.

## Official References

- [Microsoft Learn for Azure](https://learn.microsoft.com/en-us/azure/)
- [Azure Well-Architected Framework](https://learn.microsoft.com/en-us/azure/well-architected/)
