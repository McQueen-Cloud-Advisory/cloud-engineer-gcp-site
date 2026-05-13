# Project 04: Analytics Platform on Azure

## Purpose

Build a small analytics platform so you practice data landing, orchestration, transformation, and reporting concepts on Azure.

## Scenario

Assume a team needs to collect data from one or more sources, land it in Azure, clean it into a more useful form, and expose it for reporting or analysis. The goal is not to reproduce a full enterprise data estate. The goal is to build a small but explainable analytics flow with clear raw, curated, and reporting layers.

This project is valuable because it introduces platform thinking around data movement, storage, and reporting rather than only application delivery.

## Architecture

```text
Data sources
-> Azure Data Factory
-> Azure Blob Storage raw and curated layers
-> Microsoft Fabric
-> Azure Monitor
```

## What You Will Build

- A small raw and curated data layout.
- A pipeline that moves or reshapes data.
- A query or reporting surface over curated data.
- Monitoring or visibility around pipeline and reporting behavior.

## Why This Architecture Works

Data Factory provides a managed orchestration layer, Blob Storage gives the platform a raw and curated landing pattern, and Microsoft Fabric gives the project a reporting or analytical surface without requiring a hand-built warehouse stack. Azure Monitor keeps the flow observable.

## Services Used

- [Azure Data Factory](../services/data-factory.md)
- [Azure Blob Storage](../services/blob-storage.md)
- [Microsoft Fabric](../services/microsoft-fabric.md)
- [Azure Monitor](../services/monitor.md)

## Skills Practiced

- Analytics pipeline design
- Object storage data layout
- Platform monitoring
- Data governance thinking
- Explaining raw versus curated data boundaries

## Implementation Steps

1. Choose a small dataset and define the raw, curated, and reporting outcomes you want.
2. Create the Blob Storage structure and decide how the raw and curated layers will be separated.
3. Build the Data Factory orchestration flow that moves or reshapes the data.
4. Expose the curated data through a simple analytical or reporting layer in Fabric.
5. Add monitoring and freshness visibility so bad or missing data does not stay silent.
6. Document how data moves through the system and how identities and access boundaries apply.

## Security and Operations Considerations

Review who can access raw versus curated data, how identities reach storage and analytics services, and how any sensitive data is protected. A small analytics platform still needs clear answers about data ownership, freshness expectations, and who can change pipeline logic.

## Cost Considerations

Storage, orchestration activity, query patterns, and Fabric capacity or usage can become significant if the scope grows too quickly.

## How to Extend This Project

- Add data quality checks.
- Add partitioning and lifecycle controls.
- Add a small dashboard or scheduled report refresh.
- Add a governed semantic layer or curated dataset boundary.

## Portfolio Value

This project shows that you can think in terms of data movement, storage design, transformation, reporting, and operations rather than only application hosting.

## Official References

- [Azure Data Factory documentation](https://learn.microsoft.com/en-us/azure/data-factory/)
- [Azure Blob Storage documentation](https://learn.microsoft.com/en-us/azure/storage/blobs/)
- [Microsoft Fabric documentation](https://learn.microsoft.com/en-us/fabric/)
