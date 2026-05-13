# Analytics Platform on Google Cloud

## Purpose

This pattern explains how to build a first-pass analytics platform on Google Cloud using managed storage, messaging, and analytical query services.

## Pattern Summary

A small analytics platform on Google Cloud often combines Cloud Storage as a landing zone with BigQuery as the analytical destination. Pub/Sub and a serverless processing layer can be used to move data into the right shape as it enters the platform.

This pattern matters because analytics systems need more than one storage choice. They require data movement, curation, access control, and operational visibility across the full pipeline.

## When This Pattern Fits

Use this pattern when:

- data needs to land from multiple sources,
- the team wants clear separation between ingestion and analytical serving,
- BigQuery-style managed analytics is a good fit,
- and the architecture should grow by layering capabilities instead of starting overly complex.

## When Not to Use It

Do not use this pattern when a single transactional database already answers the business question or when the team has not yet clarified data ownership, sensitivity, or freshness needs.

## Common Use Cases

- Reporting pipelines
- Batch or event-driven analytics ingestion
- Portfolio analytics projects

## Reference Architecture

```text
Data source
-> Pub/Sub or batch load
-> Cloud Storage raw zone
-> BigQuery curated layer
-> Monitoring
```

## Why This Pattern Works

It works because Google Cloud gives the platform a clean split between landing storage, event-driven or batch ingestion, analytical serving, and operations visibility. That makes it easier to reason about freshness, cost, and permissions than in a loosely assembled collection of point tools.

## Provider Services

- Cloud Storage
- Pub/Sub
- Cloud Functions or Cloud Run
- BigQuery
- Cloud Monitoring

## Design Considerations

### Security

Separate access to raw and curated data where appropriate and review how service accounts move data through the pipeline.

### Reliability

Treat ingestion and transformation as production workloads with alerts, retries, and data-quality review.

### Observability

Track ingestion success, storage growth, and analytical freshness so issues surface before users notice them.

### Cost

Storage growth, repeated data movement, and query volume are the main cost areas to monitor.

### Deployment

Add layers incrementally so the data flow stays explainable and maintainable.

## Common Mistakes

- Loading data directly into the serving layer without a raw zone.
- Treating BigQuery as if schema and access decisions can wait indefinitely.
- Ignoring freshness and data-quality monitoring.
- Allowing service accounts broader access than the pipeline needs.
- Moving data too many times before validating the analytical use case.

## Related Projects

- [Project 04: Analytics Platform](../projects/project-04-analytics-platform.md)

## How This Fits Into Cloud Engineering

This pattern matters because analytics architecture is really platform design: storage, identity, orchestration, monitoring, and cost control all have to work together. That is squarely cloud-engineering work.

## Official References

- [BigQuery documentation](https://cloud.google.com/bigquery/docs)
- [Cloud Storage documentation](https://cloud.google.com/storage/docs)
- [Pub/Sub documentation](https://cloud.google.com/pubsub/docs)
