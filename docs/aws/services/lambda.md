# AWS Lambda

## Purpose

AWS Lambda runs event-driven code without requiring you to manage servers directly.

## Definition

AWS Lambda is a managed function runtime. You provide code, a runtime configuration, and an event source, and AWS handles the underlying server lifecycle.

Lambda is often described as "serverless," but that should not be confused with "effortless." The service removes host management, yet the application still needs good identity design, deployment discipline, failure handling, and observability.

In simple terms:

> Lambda is a way to run focused units of application logic when something happens, without owning the servers underneath.

## What Problem It Solves

Lambda gives teams a fast way to deploy APIs, event processors, background tasks, and automation steps without taking on operating system maintenance or capacity management for each small workload component.

## How It Is Commonly Used

Common Lambda use cases include:

- handlers behind API Gateway,
- background processing triggered by queues, storage events, or database changes,
- scheduled jobs,
- lightweight integrations between AWS services,
- automation tasks that do not justify a full application host.

## When to Use It

- Use it for event-driven APIs, background processing, and scheduled workloads.
- Use it when the execution model matches the problem well.
- Use it when you want small, deployable units of logic that connect naturally to AWS events.
- Use it when operational simplicity matters more than full runtime control.

## When Not to Use It

- Do not use it for workloads that need long-running sessions, stable local state, or specialized host control.
- Do not assume a function runtime is the best fit for every API or background worker.
- Do not assume serverless removes the need for deployment, observability, retry design, or failure ownership.

## Common Mistakes

- Building oversized functions that mix unrelated responsibilities.
- Giving many functions the same broad execution role.
- Ignoring idempotency and retry behavior for event-driven processing.
- Connecting functions to VPC resources without understanding the networking and cold-start impact.
- Treating logs as the only observability layer instead of defining metrics, alarms, and operational thresholds.

## Cloud Engineering Considerations

### Identity and Access

Each function should usually have its own execution role. Runtime permissions should be narrower than deployment permissions, and downstream access should be explicit.

### Networking

If a function needs private resources, understand how VPC attachment affects connectivity, startup behavior, and access to other AWS services.

### Security

Protect environment configuration, validate incoming events, keep secrets out of code packages, and review how untrusted input reaches downstream services.

### Observability

Use structured logs, metrics, tracing where appropriate, and alarms tied to failures, latency, throttling, and backlog indicators.

### Cost

Invocation count, duration, memory choice, and noisy triggers all influence cost. A function that runs often and inefficiently stops being cheap very quickly.

## How This Fits Into Cloud Engineering

Lambda is often the glue that makes AWS architectures feel fast to build, but it still needs the same engineering discipline as any other runtime. Good cloud engineering with Lambda means clear boundaries, explicit permissions, predictable deployment, and visible failure handling.

## Related Projects

- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)
- [Project 03: Scheduled API Ingestion](../projects/project-03-scheduled-api-ingestion.md)

## Related Patterns

- [Serverless API](../patterns/serverless-api.md)
- [Scheduled Job](../patterns/scheduled-job.md)

## Official References

- [AWS Lambda documentation](https://docs.aws.amazon.com/lambda/)
- [Lambda developer guide](https://docs.aws.amazon.com/lambda/latest/dg/welcome.html)
