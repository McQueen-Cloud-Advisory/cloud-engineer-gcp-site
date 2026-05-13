# Cloud Provider Differences Overview

## Purpose

This page explains the differences between AWS, Azure, and Google Cloud in a way that is useful for architecture and operations, not just for exam-style service mapping.

Most cloud platforms offer the same broad capability categories: identity, networking, compute, storage, databases, observability, infrastructure automation, analytics, and AI. The hard part is that they do not package those categories in the same operating model.

That is where the real differences show up.

## What Stays Similar Across Providers

The underlying engineering questions are still familiar across all three major clouds:

- How are environments separated?
- How are humans and workloads authenticated?
- How does traffic enter the system?
- How do services reach each other privately?
- How is data stored, protected, and queried?
- How are deployments automated?
- How are failures observed?
- Where do the main cost risks come from?

If you learn those questions well, it becomes much easier to translate a design from one cloud to another.

## What Actually Differs

The biggest cloud-provider differences are not usually the product names. They are the default boundaries, identity models, network assumptions, and operational workflows that shape how systems are built.

### Governance Hierarchy

The first meaningful difference is where each provider expects you to create boundaries.

| Area | AWS | Azure | Google Cloud |
| --- | --- | --- | --- |
| Top governance layer | AWS Organizations | Entra tenant + Management Groups | Organization + Folders |
| Main billing and isolation unit | Account | Subscription | Project |
| Common grouping layer inside that unit | Resource tags and service-specific grouping | Resource Groups | Resource hierarchy inside Projects plus labels |
| Practical design effect | Accounts are a strong isolation boundary and are used heavily for environment separation | Subscriptions and resource groups are central to operational organization | Projects are the main everyday isolation and permission boundary |

This matters because a design that spans three AWS accounts is not a direct match for a design that spans three Azure resource groups or three Google Cloud projects. The names may sound similar, but the blast radius, permission model, and billing behavior are different.

### Identity and Workload Identity

Identity design is one of the clearest places where the providers feel different in practice.

- AWS relies heavily on IAM roles, policies, and trust relationships. Cross-account access through role assumption is normal. Human access is often separated from workload access through federation and role-based workflows.
- Azure centers identity around Microsoft Entra ID. Human identity, enterprise integration, application registrations, managed identities, and Azure RBAC all fit into a more directory-centric operating model.
- Google Cloud uses IAM bindings and service accounts in a strongly project-centric model. Workload identity patterns are usually clean, but the team has to think carefully about how permissions are granted across projects.

In practical terms, AWS often feels role-first, Azure feels directory-first, and Google Cloud feels project-and-service-account-first.

### Networking Defaults

All three providers support private networks, routing, load balancing, and private service access. The difference is how those pieces are presented and what defaults engineers have to understand.

- AWS networking often feels explicit and component-driven. VPCs, subnets, route tables, security groups, gateways, and endpoints are all highly visible parts of the architecture.
- Azure networking often feels closely tied to platform governance and resource organization. VNets, subnets, private endpoints, NSGs, and service integration patterns are common, but many teams also need to think about how the network design fits subscriptions and resource groups.
- Google Cloud networking has a different feel because VPCs are global while subnets are regional. That changes how teams think about multi-region design, project boundaries, and network reuse.

This means network diagrams that look similar at a high level can behave quite differently when you actually implement private access, peering, egress control, or service exposure.

### Compute and Application Hosting

Each provider offers VMs, containers, and serverless, but the operational feel is not identical.

- AWS provides a wide range of compute choices and often expects teams to make more explicit service selections across EC2, ECS, EKS, Lambda, App Runner, and other runtimes.
- Azure often feels strong in integrated application hosting patterns through Functions, Container Apps, App Service, and broader identity and management integration.
- Google Cloud is often praised for a simpler managed application story around Cloud Run, Cloud Functions, and GKE, with fewer overlapping choices in some areas.

The difference is not that one provider "has serverless" and another does not. The difference is how many runtime choices exist, how integrated they are with identity and operations, and how much platform complexity the team has to own.

### Data and Analytics

The data stacks are comparable in capability, but not in shape.

- AWS commonly combines S3, EventBridge, Lambda, Glue, Athena, Redshift, and related services into modular data pipelines.
- Azure often combines Blob Storage, Data Factory, Fabric, Synapse-style patterns, and Azure-native governance and identity integration.
- Google Cloud often feels more opinionated and streamlined in analytics workflows, especially with Cloud Storage feeding BigQuery and related services.

As a result, the same analytics intent can feel more composable on AWS, more enterprise-platform-centric on Azure, and more warehouse-centric on Google Cloud.

### Observability and Operations

All three providers offer logs, metrics, tracing, dashboards, and alerting, but the surrounding operational model differs.

- AWS CloudWatch often works well as a base operational layer, with many teams supplementing it for richer cross-service observability.
- Azure Monitor and Application Insights often feel closely integrated with application and platform operations together.
- Google Cloud Monitoring is strong as a project- and service-oriented monitoring layer, especially in environments built around Google Cloud-native managed services.

Those differences change how quickly a team gets useful default telemetry and how much extra work is needed to build a complete incident-response view.

### AI Platform Differences

The provider differences are especially visible in AI workloads.

- AWS Bedrock emphasizes managed model access and surrounding services for guardrails, retrieval, and agents.
- Azure AI often emphasizes enterprise governance and integration through Azure OpenAI, Azure AI Search, Content Safety, and Azure AI Foundry.
- Google Cloud often emphasizes a broader AI platform story through Vertex AI, Agent Builder, Model Garden, and related tooling.

All three can support prompt-based apps, RAG systems, and agentic workflows. The difference is how the surrounding platform organizes model choice, retrieval, safety, and runtime integration.

### Cost and Procurement Shape

Every provider charges for compute, storage, networking, and managed services, but the billing posture feels different.

- AWS often exposes more granular service-level pricing decisions and account-based spend patterns.
- Azure cost discussions often interact closely with enterprise licensing, Microsoft estate decisions, and procurement relationships.
- Google Cloud often feels simpler in some project-level billing flows, though large analytics and AI workloads still need careful cost design.

The important lesson is that cloud cost comparison is never just a SKU comparison. It is a comparison of architecture choices, data movement, identity boundaries, and operational habits.

## How To Translate a Design Between Providers

Do not translate cloud architectures by matching service names one-to-one.

Translate them in this order instead:

1. Define the actual system requirements.
2. Define the boundary model: accounts, subscriptions, or projects.
3. Rebuild the identity model using the target provider's native approach.
4. Re-evaluate networking assumptions.
5. Choose the compute and data services that fit the target provider's strengths.
6. Re-check observability, cost, and operational ownership.

That is the difference between migration thinking and engineering thinking.

## Common Mistakes

- Comparing services only by category labels instead of by operating model.
- Assuming a familiar capability behaves the same way on every provider.
- Ignoring governance structure until resources already exist.
- Treating account, subscription, and project boundaries as interchangeable.
- Underestimating networking and identity differences during migration or multi-cloud design.
- Comparing provider AI offerings only by model names instead of by the surrounding platform controls.

## How This Fits Into Cloud Engineering

Cloud engineers often have to evaluate providers, translate designs, migrate workloads, or explain why similar systems are implemented differently across platforms. Strong provider comparison skill means understanding where the platforms are conceptually similar and where their real operational differences start to matter.

## Official References

- [Google Cloud global infrastructure](https://cloud.google.com/about/locations)
- [AWS global infrastructure](https://aws.amazon.com/about-aws/global-infrastructure/)
- [Azure geographies](https://learn.microsoft.com/en-us/azure/reliability/regions-list)
