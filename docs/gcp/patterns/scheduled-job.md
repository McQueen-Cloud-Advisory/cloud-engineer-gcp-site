# Scheduled Job on Google Cloud

## Purpose

This pattern explains how to run recurring workloads on Google Cloud using managed scheduling and event delivery services.

## Pattern Summary

A scheduled job pattern on Google Cloud commonly uses Cloud Scheduler to publish to Pub/Sub and a downstream runtime such as Cloud Functions to do the work. That creates a managed recurring pipeline without operating a dedicated scheduler host.

This pattern matters because recurring workloads need more than a cron expression.

## When This Pattern Fits

Use this pattern when:

- work needs to happen on a predictable schedule,
- asynchronous triggering is useful,
- the job can run as a small managed runtime,
- and the team wants clear separation between the scheduler, queue, and worker.

## When Not to Use It

Do not use this pattern when the workflow needs deeper orchestration, long-lived stateful execution, or a much simpler path than scheduler-plus-messaging.

## Common Use Cases

- Scheduled ingestion
- Maintenance tasks
- Recurring synchronization jobs

## Reference Architecture

```text
Schedule
-> Cloud Scheduler
-> Pub/Sub
-> Cloud Functions
-> Storage or downstream API
-> Monitoring
```

## Why This Pattern Works

It works because Google Cloud separates the recurring trigger, message delivery, worker logic, and monitoring model clearly. That makes retries, delayed processing, and operational visibility easier to reason about than in a one-piece script.

## Provider Services

- Cloud Scheduler
- Pub/Sub
- Cloud Functions
- Cloud Storage
- Secret Manager
- Cloud Monitoring

## Design Considerations

### Security

Protect external credentials, keep service account access narrow, and review who can modify the schedule or subscriptions.

### Reliability

Design for safe retries, visible failures, and idempotent processing so recurring jobs do not silently drift.

### Observability

Track scheduled runs, queue backlog, processing failures, and data freshness together.

### Cost

The base schedule is inexpensive, but repeated processing, storage, and telemetry can add up over time.

### Deployment

Treat the scheduler, messaging, runtime, and monitoring definitions as one deployable system.

## Common Mistakes

- Monitoring only the schedule and not the worker or the data result.
- Ignoring duplicate deliveries or retries.
- Letting service accounts or topic permissions become too broad.
- Storing external secrets unsafely.
- Treating a scheduled job like a throwaway automation instead of a production workload.

## Related Projects

- [Project 03: Scheduled API Ingestion](../projects/project-03-scheduled-api-ingestion.md)

## How This Fits Into Cloud Engineering

This pattern matters because recurring automation is part of normal cloud operations. Good cloud engineering makes these jobs secure, observable, and easy to reason about when they fail or drift.

## Official References

- [Cloud Scheduler documentation](https://cloud.google.com/scheduler/docs)
- [Pub/Sub documentation](https://cloud.google.com/pubsub/docs)
- [Cloud Functions documentation](https://cloud.google.com/functions/docs)
