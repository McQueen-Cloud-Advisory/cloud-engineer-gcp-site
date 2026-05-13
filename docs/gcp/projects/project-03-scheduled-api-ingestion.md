# Project 03: Scheduled API Ingestion on Google Cloud

## Purpose

Build a scheduled ingestion workflow so you practice recurring automation, messaging, secret handling, and landing data in storage.

## Scenario

Assume you need to pull data from an external API on a schedule and land it in Google Cloud for later analysis or downstream processing. The interesting challenge is not only scheduling the job. It is handling service-account permissions, retries, storage structure, and the visibility needed to detect silent ingestion failures.

This project is useful because it introduces recurring automation as a real cloud workload rather than a one-off script.

## Architecture

```text
Cloud Scheduler
-> Pub/Sub
-> Cloud Functions
-> External API call
-> Cloud Storage
-> Cloud Monitoring
```

## What You Will Build

- A scheduled trigger path for ingestion.
- A function that pulls external data and stores results.
- A landing zone for raw data in storage.
- Monitoring and failure visibility for the recurring job.

## Why This Architecture Works

Cloud Scheduler gives you the schedule, Pub/Sub provides clean decoupling, Cloud Functions handles the logic, and Cloud Storage provides a simple landing zone for raw data. Secret Manager and Cloud Monitoring complete the operational model by handling credentials and making failures visible.

## Services Used

- Pub/Sub
- [Cloud Functions](../services/cloud-functions.md)
- [Cloud Storage](../services/cloud-storage.md)
- [Secret Manager](../services/secret-manager.md)
- [Cloud Monitoring](../services/cloud-monitoring.md)
- Cloud Scheduler

## Skills Practiced

- Scheduled automation
- Event-driven messaging
- External API integration
- Data landing zone design
- Thinking about freshness and retries

## Implementation Steps

1. Define the schedule, source API contract, and the raw storage structure you want in Cloud Storage.
2. Configure Cloud Scheduler and Pub/Sub so the recurring trigger is explicit and easy to inspect.
3. Build the Cloud Function that pulls the external data and stores it in the landing zone.
4. Store credentials in Secret Manager and keep service-account permissions narrow.
5. Add monitoring and alerts for function failures, missing runs, or unusual result patterns.
6. Document how the job is triggered, how duplicates are handled, and how an operator would detect stale data.

## Security and Operations Considerations

Store external credentials in Secret Manager, keep service-account access narrow, and review who can trigger or modify the schedule. Think about duplicate deliveries, malformed source responses, and what should happen when the upstream API is slow or unavailable.

Operationally, the project becomes stronger when you can show how the team would detect missing data rather than only failed code execution.

## Cost Considerations

The base pattern is usually low cost, but repeated runs, storage growth, telemetry volume, and downstream processing can add up over time.

## How to Extend This Project

- Add retry or dead-letter handling.
- Add data quality checks before storage.
- Add a downstream transformation step into analytics.
- Add freshness reporting or delayed-run alerts.

## Portfolio Value

This project shows recurring integration work, event-driven design, and the operational thinking needed for scheduled jobs rather than simple one-time automation.

## Official References

- [Cloud Scheduler documentation](https://cloud.google.com/scheduler/docs)
- [Pub/Sub documentation](https://cloud.google.com/pubsub/docs)
- [Cloud Functions documentation](https://cloud.google.com/functions/docs)
- [Cloud Storage documentation](https://cloud.google.com/storage/docs)
