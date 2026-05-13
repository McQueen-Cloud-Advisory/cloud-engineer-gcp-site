# Project 03: Scheduled API Ingestion on AWS

## Purpose

Build a scheduled ingestion workflow so you practice event scheduling, external API handling, storage design, and operational visibility.

## Scenario

Assume you need to pull data from an external API every hour or every day and land it in AWS for later reporting or downstream processing. The workload sounds simple, but recurring integrations create real operational problems: secret handling, duplicate runs, malformed data, partial failures, and silent staleness.

This project is useful because it introduces the discipline required for recurring jobs, not just one-time scripts.

## Architecture

```text
Amazon EventBridge schedule
-> AWS Lambda
-> External API call
-> Amazon S3 or Amazon DynamoDB
-> Amazon CloudWatch
```

## What You Will Build

- A scheduled ingestion trigger.
- A function that pulls external data and stores results.
- A landing zone or simple store for the retrieved data.
- Monitoring and failure visibility for the ingestion path.

## Why This Architecture Works

EventBridge gives you a managed schedule. Lambda keeps the ingestion logic lightweight and easy to deploy. S3 or DynamoDB gives you a managed place to land the data, depending on whether the result is file-oriented or application-record-oriented. Secrets Manager and CloudWatch complete the security and operations model.

## Services Used

- [Amazon EventBridge](../services/eventbridge.md)
- [AWS Lambda](../services/lambda.md)
- [Amazon S3](../services/s3.md)
- [AWS Secrets Manager](../services/secrets-manager.md)
- [Amazon CloudWatch](../services/cloudwatch.md)

## Skills Practiced

- Scheduled automation
- External API integration
- Data landing zone design
- Operational alerting
- Thinking about idempotency and freshness

## Implementation Steps

1. Define the ingestion schedule, the source API contract, and the storage shape for the retrieved data.
2. Create the EventBridge rule and the Lambda function that executes the ingestion logic.
3. Store any external API credentials in Secrets Manager and keep the function role narrow.
4. Write the data into S3 or another target with a structure that supports later analysis or recovery.
5. Add logging, metrics, and alarms for failed runs, empty results, or unusual response behavior.
6. Document how the schedule works, how duplicates are handled, and how operators would detect stale or bad data.

## Security and Operations Considerations

Store API credentials in Secrets Manager, restrict function permissions, and validate how external API responses are handled before writing data downstream. Think carefully about idempotency, retry behavior, and what should happen if the upstream API is unavailable or returns malformed data.

Operationally, recurring jobs are only as good as their visibility. If a job quietly stops collecting good data, the workload is already failing.

## Cost Considerations

The core schedule is inexpensive, but external API usage, storage growth, and log retention can become meaningful over time.

## How to Extend This Project

- Add retry and dead-letter handling.
- Add data quality checks before storing records.
- Add a downstream transformation or summary report.
- Add freshness dashboards or alerts for missing data.

## Portfolio Value

This project shows event-driven automation, integration with external systems, and the operational thinking needed for recurring jobs rather than one-off scripts.

## Official References

- [Amazon EventBridge documentation](https://docs.aws.amazon.com/eventbridge/)
- [AWS Lambda documentation](https://docs.aws.amazon.com/lambda/)
- [Amazon S3 documentation](https://docs.aws.amazon.com/s3/)
- [AWS Secrets Manager documentation](https://docs.aws.amazon.com/secretsmanager/)
