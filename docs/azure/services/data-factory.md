# Azure Data Factory

## Purpose

Azure Data Factory orchestrates data movement and transformation workflows across systems.

## Definition

Azure Data Factory is a managed data integration and orchestration service. It coordinates how data is copied, transformed, and moved between systems on a schedule or as part of repeatable pipeline flows.

Its value is not only that it can move data. It gives teams a control plane for recurring ingestion and transformation work so they do not have to hand-build every orchestration step inside custom code.

Data Factory is not the whole data platform. It is the orchestration and movement layer. Teams still need to define source ownership, target models, data quality expectations, and how pipeline failures affect the downstream system.

In simple terms:

> Data Factory is the workflow layer that helps Azure data pipelines run on purpose instead of by ad hoc scripts and manual effort.

## What Problem It Solves

Without an orchestration layer, recurring data movement often turns into fragile scripts, hidden credentials, unclear schedules, and poor operational visibility. Data Factory gives teams a managed way to define, schedule, and monitor repeatable ingestion and movement workflows.

It solves the problem of needing repeatable data movement and orchestration without turning every pipeline into a custom application.

That does not remove engineering responsibility. Engineers still need to define contracts, retries, backfills, permissions, and how freshness is measured.

## How It Is Commonly Used

Data Factory is commonly used for:

- scheduled ingestion from APIs, files, or databases,
- movement of data between storage systems and analytics layers,
- orchestration of transformation or staging workflows,
- repeatable pipeline runs that need centralized visibility,
- integration patterns where a managed control plane is easier to operate than fully custom code.

In many Azure analytics designs, Data Factory becomes the line between where data is sourced and where it is staged, transformed, or handed off to an analytical platform. That makes its scheduling, identity model, and failure behavior part of the architecture.

## Foundational Concepts Connected to Data Factory

Data Factory connects directly to several cloud engineering foundations.

### Data Orchestration

Data pipelines need a control layer that decides when data should move, how steps depend on each other, and what happens when something fails.

### Integration and Scheduling

Data Factory is often the scheduled workflow layer for recurring ingestion and movement tasks. Schedule design and dependency design both matter.

### Identity and Access

Pipelines often touch many systems. Managed identities and scoped source and destination access are essential to keep convenience from turning into excessive privilege.

### Observability and Freshness

Pipeline health is not just about whether a run completed. It is also about whether the right data arrived on time and in usable form.

### Cost Management

Movement volume, integration runtimes, and orchestration frequency all influence cost. Pipeline design is also cost design.

## When to Use It

- Use it for scheduled ingestion and movement of data between systems.
- Use it when analytics or integration workflows need pipeline orchestration.
- Use it when the team needs a managed control plane for recurring data movement tasks.
- Use it when visibility into pipeline status and failure handling matters as much as the movement itself.

Data Factory is strongest when the system needs repeatable, observable data movement rather than a one-off script or a narrow runtime job.

## When Not to Use It

- Do not use it for simple tasks that are better handled by a lighter runtime.
- Do not build pipelines without clear ownership of data quality and failure handling.
- Do not assume a visual pipeline makes the underlying data contract simple.

## Compare To

### Data Factory vs. Functions

Functions is better for narrow code-driven tasks or event handlers.

Data Factory is better when the problem is orchestration of recurring data movement and pipeline steps across systems.

### Data Factory vs. Microsoft Fabric

Data Factory is the orchestration and movement layer.

Microsoft Fabric is the broader analytics platform surface where storage, transformation, and reporting may come together.

## Tradeoffs

Data Factory's biggest advantage is that it gives teams a managed control plane for recurring ingestion and movement workflows.

The tradeoff is platform weight. A visual or managed pipeline tool can be powerful, but it can also be heavier than necessary for very small tasks.

Data Factory also makes recurring movement easier to operationalize. That is useful, but it can hide weak data contracts if teams treat successful movement as proof of useful data.

Another tradeoff is that orchestration centralization can create broad dependency and permission surfaces unless pipeline ownership is deliberate.

## Common Mistakes

- Using a heavy orchestration tool for tasks that could be simpler code.
- Moving data without clear contracts around freshness, quality, or ownership.
- Treating pipeline success as proof that the data is correct and usable.
- Granting broad source and destination permissions to convenience identities.
- Ignoring retry, idempotency, and backfill behavior in recurring workflows.
- Assuming the first green pipeline run means the data platform is trustworthy.

## Cloud Engineering Considerations

### Identity and Access

Use managed identities and scoped data-source permissions so pipelines only access approved systems. Pipeline convenience should not override least-privilege design.

### Networking

Review self-hosted versus managed connectivity, private endpoints, and cross-system data paths. Data movement architecture is also network architecture.

### Security

Protect source credentials, review sensitive data movement, and make sure pipeline permissions are tightly scoped. Data pipelines often touch more systems than any single application service does.

### Observability

Track pipeline runs, failures, latency, and data freshness so ingestion problems are visible quickly. A successful scheduler is not the same thing as a healthy data platform.

### Reliability

Reliable pipeline design depends on safe retries, clear backfill behavior, predictable schedules, and knowing how downstream systems behave when data arrives late or partially.

### Cost

Pipeline activity, movement volume, integration runtimes, and frequency all influence cost. Orchestration should be deliberate, not constant by default.

## Project and Pattern Connections

Azure Data Factory is most directly connected to:

- [Project 03: Scheduled API Ingestion](../projects/project-03-scheduled-api-ingestion.md)
- [Project 04: Analytics Platform](../projects/project-04-analytics-platform.md)
- [Scheduled Job](../patterns/scheduled-job.md)
- [Analytics Platform](../patterns/analytics-platform.md)

It matters wherever the Azure learning path needs repeatable data movement to become an operable platform concern rather than an ad hoc script.

## Official References

- [Azure Data Factory documentation](https://learn.microsoft.com/en-us/azure/data-factory/)
- [Introduction to Azure Data Factory](https://learn.microsoft.com/en-us/azure/data-factory/introduction)
- [Pipelines and activities in Azure Data Factory](https://learn.microsoft.com/en-us/azure/data-factory/concepts-pipelines-activities)
