# Project 03: Scheduled API Ingestion on Azure

## Purpose

Build a scheduled ingestion workflow so you practice recurring data movement, secret handling, and operational visibility on Azure.

## Scenario

Assume a team needs to collect data from an external API or source system on a schedule and land it in Azure for later analysis. The challenge is not only moving data. It is doing so repeatedly with clear secrets management, visible failures, and a storage layout that supports later use.

This project is valuable because it exposes the operational side of recurring data movement, not just the happy path of a one-time copy.

## Architecture

```text
Scheduled trigger
-> Azure Data Factory
-> External API or source system
-> Azure Blob Storage
-> Azure Monitor and Application Insights
```

## What You Will Build

- A scheduled pipeline that pulls data from an external source.
- A storage landing zone for raw data.
- Monitoring for pipeline runs and failure visibility.
- An explanation of how freshness and failure handling work.

## Why This Architecture Works

Azure Data Factory is a natural orchestration layer for recurring data movement. Blob Storage is a simple landing zone for raw data. Key Vault, Monitor, and Application Insights complete the design by making secret handling and operational visibility explicit instead of implicit.

## Services Used

- [Azure Data Factory](../services/data-factory.md)
- [Azure Blob Storage](../services/blob-storage.md)
- [Azure Key Vault](../services/key-vault.md)
- [Azure Monitor](../services/monitor.md)
- [Application Insights](../services/application-insights.md)

## Skills Practiced

- Pipeline orchestration
- External data integration
- Data landing zone design
- Operational alerting
- Thinking about freshness and retry behavior

## Implementation Steps

1. Define the ingestion schedule, source contract, and the raw storage layout you want in Blob Storage.
2. Build the Data Factory pipeline and parameterize it cleanly enough that you can explain and reuse it.
3. Store external credentials in Key Vault and scope the pipeline identity to only the resources it needs.
4. Land the raw data in a structure that supports replay, review, and later transformation.
5. Add monitoring for pipeline failures, latency, and freshness so quiet data loss is visible.
6. Document how the schedule operates, how secrets are handled, and what operators should check when runs fail.

## Security and Operations Considerations

Store external credentials in Key Vault, scope pipeline identities carefully, and review who can access the landing storage. Think about malformed source data, partial ingestion, duplicate runs, and what counts as a stale dataset.

Operationally, recurring ingestion work is only credible when you can show how missing or bad data would be detected.

## Cost Considerations

Pipeline activity, storage growth, monitoring retention, and any added downstream processing are the main cost areas to watch.

## How to Extend This Project

- Add data validation checks.
- Add a transformation step after landing.
- Add freshness alerts for delayed or failed runs.
- Add a curated zone after the raw landing layer.

## Portfolio Value

This project demonstrates recurring integration work, operational monitoring, and the discipline needed for scheduled data movement in production-style Azure environments.

## Official References

- [Azure Data Factory documentation](https://learn.microsoft.com/en-us/azure/data-factory/)
- [Azure Blob Storage documentation](https://learn.microsoft.com/en-us/azure/storage/blobs/)
- [Azure Key Vault documentation](https://learn.microsoft.com/en-us/azure/key-vault/)
