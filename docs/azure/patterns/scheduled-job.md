# Scheduled Job on Azure

## Purpose

This pattern explains how to run recurring workloads on Azure using managed scheduling and integration services.

## Pattern Summary

A scheduled job pattern on Azure often uses Data Factory or another managed trigger path to run a recurring ingestion or maintenance workload. The design centers on repeatability, controlled access to source systems, and clear failure visibility.

This pattern matters because recurring jobs often become operationally important even when they look simple at first.

## When This Pattern Fits

Use this pattern when:

- work needs to happen on a predictable schedule,
- the team wants a managed orchestration layer,
- the main concern is moving data or running controlled batch logic,
- and strong visibility into failures or freshness is required.

## When Not to Use It

Do not use this pattern when a lighter runtime is clearly enough or when the workflow needs a more complex orchestration model than a straightforward scheduled job.

## Common Use Cases

- Scheduled ingestion
- Batch maintenance tasks
- Recurring synchronization jobs

## Reference Architecture

```text
Schedule
-> Azure Data Factory
-> External source or internal system
-> Blob Storage or downstream processing
-> Monitoring
```

## Why This Pattern Works

It works because Azure can separate the orchestration layer, source access, storage, and monitoring model cleanly. That makes recurring automation easier to explain and easier to operate than ad hoc scripts scattered across environments.

## Provider Services

- Azure Data Factory
- Azure Blob Storage
- Azure Key Vault
- Azure Monitor
- Application Insights

## Design Considerations

### Security

Protect source credentials, scope pipeline permissions carefully, and review who can operate the schedule and downstream data access.

### Reliability

Design for retry, idempotency, and visible failure states so recurring jobs do not silently drift out of compliance.

### Observability

Track pipeline success, latency, and data freshness rather than only trigger execution.

### Cost

Pipeline activity, storage growth, and retained telemetry are the main cost areas to monitor.

### Deployment

Treat schedule definitions and pipeline configuration as part of the deployable system rather than manual platform setup.

## Common Mistakes

- Using a heavy orchestration path for work that could be simpler.
- Ignoring freshness and only checking whether the trigger fired.
- Leaving pipeline identities too broad.
- Treating data landing as successful just because the pipeline completed.
- Failing to define what safe retry means for the downstream system.

## Related Projects

- [Project 03: Scheduled API Ingestion](../projects/project-03-scheduled-api-ingestion.md)

## How This Fits Into Cloud Engineering

This pattern matters because recurring jobs are operational systems, not just scheduled tasks. Cloud engineering here means designing them so they are secure, visible, and recoverable over time.

## Official References

- [Azure Data Factory documentation](https://learn.microsoft.com/en-us/azure/data-factory/)
- [Azure Blob Storage documentation](https://learn.microsoft.com/en-us/azure/storage/blobs/)
