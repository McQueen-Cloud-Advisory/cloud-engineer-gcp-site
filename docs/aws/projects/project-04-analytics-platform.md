# Project 04: Analytics Platform on AWS

## Purpose

Build a small analytics platform so you practice batch ingestion, layered data storage, query access, and operational oversight.

## Scenario

Assume a team needs to collect data from one or more sources, store it durably, transform it into a cleaner analytical shape, and query it for reporting or inspection. The platform does not need to be huge to be valuable. It needs to show clear movement from raw data to curated data and make that flow easy to explain.

This project is useful because it moves you from application-centric cloud work into data-platform thinking.

## Architecture

```text
Scheduled or event-driven ingestion
-> Amazon S3 landing zone
-> Transformation and cataloging
-> Query layer and reporting
-> Amazon CloudWatch
```

## What You Will Build

- A raw and curated data layout in object storage.
- A simple transformation or cataloging flow.
- A queryable analytics layer for reporting or inspection.
- Monitoring or visibility around ingestion, freshness, and query activity.

## Why This Architecture Works

S3 gives the platform a durable raw and curated storage boundary. EventBridge and Lambda can orchestrate ingestion or lightweight transformation. Glue and Athena provide a practical managed path for cataloging and querying without building a full warehouse platform from scratch. CloudWatch keeps the data flow observable.

## Services Used

- [Amazon S3](../services/s3.md)
- [Amazon EventBridge](../services/eventbridge.md)
- [AWS Lambda](../services/lambda.md)
- [Amazon CloudWatch](../services/cloudwatch.md)
- AWS Glue
- Amazon Athena

## Skills Practiced

- Analytics data flow design
- Data lake organization
- Scheduled transformation
- Monitoring data platform operations
- Explaining raw versus curated storage clearly

## Implementation Steps

1. Choose a small dataset and define the raw, curated, and queryable outcomes you want.
2. Create the raw landing zone in S3 and decide how files, partitions, or prefixes should be organized.
3. Build the ingestion and transformation steps using lightweight automation first.
4. Add cataloging and a query layer so the curated data can be inspected with SQL.
5. Add monitoring for pipeline failures, delayed data, and unexpected storage or query behavior.
6. Document how data moves through the platform and where governance or access boundaries matter.

## Security and Operations Considerations

Review who can access raw versus curated data, how credentials are handled, and whether any sensitive data needs masking or partitioned access. Analytics platforms often fail operationally through silent staleness or confusing ownership rather than through obvious runtime crashes.

## Cost Considerations

Storage growth, repeated queries, and transformation frequency can all increase cost. Keep the dataset and query scope intentionally small at first, and explain where the main spend risks would grow.

## How to Extend This Project

- Add partitioning and lifecycle rules.
- Add a dashboarding layer.
- Add data quality and freshness checks.
- Add separate raw, trusted, and curated zones.

## Portfolio Value

This project demonstrates that you can think beyond application hosting and explain the data handling, transformation, storage, and operational concerns of an analytics workload on AWS.

## Official References

- [Amazon S3 documentation](https://docs.aws.amazon.com/s3/)
- [AWS Glue documentation](https://docs.aws.amazon.com/glue/)
- [Amazon Athena documentation](https://docs.aws.amazon.com/athena/)
- [Amazon CloudWatch documentation](https://docs.aws.amazon.com/cloudwatch/)
