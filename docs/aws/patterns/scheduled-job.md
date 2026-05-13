# Scheduled Job on AWS

## Purpose

This pattern explains how to run recurring workloads on AWS without maintaining a dedicated scheduler host.

## Pattern Summary

A scheduled job pattern on AWS commonly uses Amazon EventBridge to trigger a Lambda function or another managed target on a schedule. The workload usually pulls data, performs routine maintenance, or generates periodic outputs.

This pattern matters because recurring workloads often look simple until teams need retries, secrets, observability, and idempotency.

## When This Pattern Fits

Use this pattern when:

- work needs to happen on a schedule,
- a small runtime can execute the task,
- the job should be easy to redeploy and observe,
- and there is value in keeping the scheduling layer managed.

## When Not to Use It

Do not use this pattern when the workflow needs heavy orchestration, long-running stateful execution, or more complex control flow than a simple scheduled trigger can support cleanly.

## Common Use Cases

- API ingestion
- Routine maintenance tasks
- Scheduled report generation

## Reference Architecture

```text
Schedule
-> Amazon EventBridge
-> AWS Lambda
-> Storage or downstream API
-> Logging and monitoring
```

## Why This Pattern Works

It works because the schedule, runtime logic, secrets, and telemetry can all be managed separately but deployed together. That keeps recurring automation simple enough to operate without pretending it has no failure modes.

## Provider Services

- Amazon EventBridge
- AWS Lambda
- Amazon S3
- AWS Secrets Manager
- Amazon CloudWatch

## Design Considerations

### Security

Protect any external credentials, scope the runtime role carefully, and review downstream write permissions.

### Reliability

Design the job so it can be retried safely and so partial failures are visible.

### Observability

Track scheduled invocations, failures, and downstream processing results rather than only trigger success.

### Cost

The scheduling layer is inexpensive, but repeated processing, storage, and logging can grow steadily over time.

### Deployment

Ship schedule definitions, runtime code, and monitoring together so the job is easy to reproduce.

## Common Mistakes

- Treating a scheduled job like a disposable script.
- Ignoring idempotency or duplicate execution.
- Storing external credentials unsafely.
- Monitoring only the trigger and not the data or side effects.
- Letting recurring jobs run for months without freshness checks.

## Related Projects

- [Project 03: Scheduled API Ingestion](../projects/project-03-scheduled-api-ingestion.md)

## How This Fits Into Cloud Engineering

This pattern matters because recurring automation is common and deceptively easy to underestimate. Cloud engineering here means making the job repeatable, observable, secure, and safe to retry over time.

## Official References

- [Amazon EventBridge documentation](https://docs.aws.amazon.com/eventbridge/)
- [AWS Lambda documentation](https://docs.aws.amazon.com/lambda/)
