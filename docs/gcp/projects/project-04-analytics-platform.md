# Project 04: Analytics Platform on Google Cloud

## Purpose

Build a small analytics platform so you practice ingestion, object storage organization, analytical querying, and monitoring on Google Cloud.

## Scenario

Assume a team needs to collect data, store it durably, transform it into a more useful analytical form, and expose it through queries or reports. The goal is not to build a giant data platform. It is to show that you understand the difference between raw data landing, curated analytical data, and operational ownership.

This project is valuable because it introduces data-platform thinking in a Google Cloud-native shape.

## Architecture

```text
Sources or scheduled ingestion
-> Pub/Sub or batch load
-> Cloud Storage raw zone
-> BigQuery curated analytics layer
-> Cloud Monitoring
```

## What You Will Build

- A raw and curated data flow.
- A queryable analytics layer in BigQuery.
- Monitoring for ingestion and analytical workload health.
- A clear explanation of how data moves from raw storage to usable analytics.

## Why This Architecture Works

Cloud Storage provides a durable raw landing zone, Pub/Sub or Cloud Functions can help with ingestion flow, and BigQuery gives the platform a strong managed analytical layer for curated data. Cloud Monitoring rounds out the design with visibility into ingestion and query behavior.

## Services Used

- [Pub/Sub](../services/pubsub.md)
- [Cloud Storage](../services/cloud-storage.md)
- [BigQuery](../services/bigquery.md)
- [Cloud Monitoring](../services/cloud-monitoring.md)
- [Cloud Functions](../services/cloud-functions.md)

## Skills Practiced

- Analytics data flow design
- Object storage organization
- Data loading and query patterns
- Operational monitoring
- Explaining raw versus curated analytics boundaries

## Implementation Steps

1. Choose a small dataset and define the raw, curated, and reporting outcomes you want.
2. Design the raw landing structure in Cloud Storage so data arrival and later processing are easy to explain.
3. Build the ingestion and transformation path using lightweight eventing or scheduled automation.
4. Load curated data into BigQuery in a way that supports clear query behavior.
5. Add monitoring for ingestion failures, missing data, and unusual query or pipeline behavior.
6. Document how data moves through the system and where identity and data-governance boundaries matter.

## Security and Operations Considerations

Review who can access raw versus curated data, how service accounts move data, and how sensitive fields are handled. Analytics work often looks successful until a team realizes the wrong people can access the raw layer or that freshness issues were never being measured.

## Cost Considerations

Storage, query scans, repeated ingestion, and poor partitioning or query habits can increase cost if the dataset or schedule grows without controls.

## How to Extend This Project

- Add data quality checks.
- Add partitioning and lifecycle rules.
- Add a dashboard or scheduled analytical report.
- Add separate trusted and curated dataset layers.

## Portfolio Value

This project shows that you can move beyond application hosting and explain an end-to-end analytics workload with storage, transformation, query, and monitoring concerns.

## Official References

- [BigQuery documentation](https://cloud.google.com/bigquery/docs)
- [Cloud Storage documentation](https://cloud.google.com/storage/docs)
- [Pub/Sub documentation](https://cloud.google.com/pubsub/docs)
- [Cloud Monitoring documentation](https://cloud.google.com/monitoring/docs)
