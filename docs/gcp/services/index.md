# Google Cloud Services

## Purpose

This section explains Google Cloud services from a practical cloud engineering perspective and focuses on the services that power the current learning path.

## How To Read This Section

These pages are not intended to cover the full Google Cloud catalog. They are intended to make the current provider path easier to understand by answering a focused set of questions.

- What problem does this service solve in Google Cloud?
- How does it fit into a project or pattern?
- What identity, networking, and observability assumptions come with it?
- How does it compare to the surrounding runtime and data choices?

That makes the section most useful when you read it as part of a project, not as a disconnected glossary.

## What This Service Set Is Designed To Teach

The Google Cloud path emphasizes the services most useful for learning project-based application delivery, eventing, secret handling, monitoring, analytics, and later AI extension.

## Service Categories

### Identity and Access

- [IAM and Service Accounts](iam-and-service-accounts.md)

### Storage and Data

- [Cloud Storage](cloud-storage.md)
- [BigQuery](bigquery.md)

### Compute and Application Hosting

- [Cloud Run](cloud-run.md)
- [Cloud Functions](cloud-functions.md)

### Integration and Scheduling

- [Pub/Sub](pubsub.md)

### Observability and Operations

- [Cloud Monitoring](cloud-monitoring.md)

### AI and Agentic Platforms

- [AI and Agentic Workloads](ai-and-agents.md)
- [Vertex AI](vertex-ai.md)
- [Vertex AI Agent Builder](vertex-ai-agent-builder.md)
- [Agent Development Kit](agent-development-kit.md)
- [Model Garden](model-garden.md)
- [Model Armor](model-armor.md)

### Security and Secrets

- [Secret Manager](secret-manager.md)

## How The Categories Connect

Typical Google Cloud workloads in this site start with IAM and service accounts, then connect Cloud Run or Cloud Functions to Cloud Storage, Pub/Sub, and Secret Manager. Cloud Monitoring gives you the operational view, while BigQuery extends the path into analytics. Vertex AI and related AI services build on top of that same application foundation rather than replacing it.

## How This Fits Into Cloud Engineering

Cloud engineers need service knowledge that supports design, deployment, and operations. This section helps you explain how Google Cloud services work together in a system and what tradeoffs they introduce.

## Official References

- [Google Cloud documentation](https://cloud.google.com/docs)
- [Google Cloud architecture framework](https://cloud.google.com/architecture/framework)
