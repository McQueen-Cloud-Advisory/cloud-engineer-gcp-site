# Azure Data Factory

## Purpose

Azure Data Factory orchestrates data movement and transformation workflows across systems.

## Definition

Azure Data Factory is a managed data integration and orchestration service. It coordinates how data is copied, transformed, and moved between systems on a schedule or as part of repeatable pipeline flows.

Its value is not only that it can move data. It gives teams a control plane for recurring ingestion and transformation work so they do not have to hand-build every orchestration step inside custom code.

In simple terms:

> Data Factory is the workflow layer that helps Azure data pipelines run on purpose instead of by ad hoc scripts and manual effort.

## What Problem It Solves

Without an orchestration layer, recurring data movement often turns into fragile scripts, hidden credentials, unclear schedules, and poor operational visibility. Data Factory gives teams a managed way to define, schedule, and monitor repeatable ingestion and movement workflows.

## How It Is Commonly Used

Data Factory is commonly used for:

- scheduled ingestion from APIs, files, or databases,
- movement of data between storage systems and analytics layers,
- orchestration of transformation or staging workflows,
- repeatable pipeline runs that need centralized visibility,
- integration patterns where a managed control plane is easier to operate than fully custom code.

## When to Use It

- Use it for scheduled ingestion and movement of data between systems.
- Use it when analytics or integration workflows need pipeline orchestration.
- Use it when the team needs a managed control plane for recurring data movement tasks.
- Use it when visibility into pipeline status and failure handling matters as much as the movement itself.

## When Not to Use It

- Do not use it for simple tasks that are better handled by a lighter runtime.
- Do not build pipelines without clear ownership of data quality and failure handling.
- Do not assume a visual pipeline makes the underlying data contract simple.

## Common Mistakes

- Using a heavy orchestration tool for tasks that could be simpler code.
- Moving data without clear contracts around freshness, quality, or ownership.
- Treating pipeline success as proof that the data is correct and usable.
- Granting broad source and destination permissions to convenience identities.
- Ignoring retry, idempotency, and backfill behavior in recurring workflows.

## Cloud Engineering Considerations

### Identity and Access

Use managed identities and scoped data-source permissions so pipelines only access approved systems. Pipeline convenience should not override least-privilege design.

### Networking

Review self-hosted versus managed connectivity, private endpoints, and cross-system data paths. Data movement architecture is also network architecture.

### Security

Protect source credentials, review sensitive data movement, and make sure pipeline permissions are tightly scoped. Data pipelines often touch more systems than any single application service does.

### Observability

Track pipeline runs, failures, latency, and data freshness so ingestion problems are visible quickly. A successful scheduler is not the same thing as a healthy data platform.

### Cost

Pipeline activity, movement volume, integration runtimes, and frequency all influence cost. Orchestration should be deliberate, not constant by default.

## How This Fits Into Cloud Engineering

Cloud data systems need more than storage and analytics engines. They need a reliable way to move data between stages and systems. Data Factory matters because it turns repeatable movement and orchestration into an operable part of the platform.

## Related Projects

- [Project 03: Scheduled API Ingestion](../projects/project-03-scheduled-api-ingestion.md)
- [Project 04: Analytics Platform](../projects/project-04-analytics-platform.md)

## Related Patterns

- [Scheduled Job](../patterns/scheduled-job.md)
- [Analytics Platform](../patterns/analytics-platform.md)

## Official References

- [Azure Data Factory documentation](https://learn.microsoft.com/en-us/azure/data-factory/)
- [Introduction to Azure Data Factory](https://learn.microsoft.com/en-us/azure/data-factory/introduction)
