# Azure Projects

## Purpose

This section contains hands-on Azure projects designed to produce real artifacts while reinforcing the Azure operating model.

## What The Sequence Is Meant To Build

The Azure project path is cumulative. Each project layers more delivery and operations responsibility on top of the earlier work.

- [Project 01: Static Site](project-01-static-site.md) introduces object storage delivery, public access choices, and deployment discipline.
- [Project 02: Serverless Contact Form](project-02-serverless-contact-form.md) adds application hosting, secrets, data, and monitoring.
- [Project 03: Scheduled API Ingestion](project-03-scheduled-api-ingestion.md) introduces recurring automation, integration, and operational visibility.
- [Project 04: Analytics Platform](project-04-analytics-platform.md) extends the path into data movement, curation, and analytics workloads.
- [Project 05: Agentic RAG Assistant](project-05-agentic-rag-assistant.md) adds AI services, retrieval, safety, and higher-complexity runtime concerns.

## How To Use The Project Path

Start in order unless you already have strong experience with the earlier Azure concepts. The point of the sequence is to let identity, deployment, observability, and cost concerns compound gradually instead of all at once.

As you work through a project, try to capture five things every time:

- the main request or data flow,
- the Azure identity and RBAC model,
- where secrets and data live,
- how failures become visible,
- and what the next iteration would improve.

## What Good Output Looks Like

Good Azure project output is not just "the resources deployed successfully." It is a small architecture story that shows judgment.

The strongest output usually explains:

- why the chosen Azure services fit,
- how identities are separated between deployment and runtime,
- how traffic or data moves through the system,
- what the main cost and security considerations are,
- and how the workload would be operated after launch.

## Portfolio Goal

Each project should produce material you can explain in a resume, interview, or portfolio. The ideal output is not just a working Azure deployment. It is a project story that shows architecture judgment and operational awareness.

## How This Fits Into Cloud Engineering

Projects are where provider knowledge becomes engineering skill. They force you to connect governance, identity, runtime, data, and monitoring into one system that can be explained and improved over time.

## Official References

- [Microsoft Learn for Azure](https://learn.microsoft.com/en-us/azure/)
- [Azure Well-Architected Framework](https://learn.microsoft.com/en-us/azure/well-architected/)
