# AWS Lambda

## Purpose

AWS Lambda runs event-driven code without requiring you to manage servers directly.

It is used when a workload fits a small execution unit that reacts to an event, schedule, queue message, or simple API request rather than a long-lived application service.

## Definition

AWS Lambda is a managed function runtime. You provide code, a runtime configuration, and an event source, and AWS handles the underlying server lifecycle.

Lambda is often described as "serverless," but that should not be confused with "effortless." The service removes host management, yet the application still needs good identity design, deployment discipline, failure handling, and observability.

Lambda is not the right answer for every backend. It is designed for function-shaped execution units, not for hiding a monolithic service inside a collection of handlers.

In simple terms:

> Lambda is a way to run focused units of application logic when something happens, without owning the servers underneath.

## What Problem It Solves

Lambda gives teams a fast way to deploy APIs, event processors, background tasks, and automation steps without taking on operating system maintenance or capacity management for each small workload component.

It solves the problem of wanting code to run when something happens without also wanting to operate a full application host for every small piece of logic.

That does not remove engineering responsibility. Engineers still need to design idempotency, retries, permissions, packaging, timeout behavior, and observability.

## How It Is Commonly Used

Common Lambda use cases include:

- handlers behind API Gateway,
- background processing triggered by queues, storage events, or database changes,
- scheduled jobs,
- lightweight integrations between AWS services,
- automation tasks that do not justify a full application host.

In many AWS systems, Lambda is the narrow execution layer behind API Gateway or EventBridge while DynamoDB, S3, Secrets Manager, and CloudWatch support data, configuration, and operations.

## Foundational Concepts Connected to Lambda

Lambda connects directly to several cloud engineering foundations.

### Serverless Compute

Lambda is a serverless runtime. That reduces host management, but it still requires clear boundaries, ownership, and deployment discipline.

### Events and Messaging

Many Lambda workloads are event-driven. That means retries, duplicate delivery, and downstream dependency behavior all matter.

### Identity and Access

Each function should have an execution role that matches only what that function actually needs to do.

### Observability

Function behavior needs to be visible through logs, metrics, traces where appropriate, and alarms tied to real operational thresholds.

### Cost Management

Invocation volume, duration, memory allocation, and noisy triggers all affect cost. Function design is also cost design.

## When to Use It

- Use it for event-driven APIs, background processing, and scheduled workloads.
- Use it when the execution model matches the problem well.
- Use it when you want small, deployable units of logic that connect naturally to AWS events.
- Use it when operational simplicity matters more than full runtime control.

Lambda is strongest when the execution boundary is narrow and the event model is clear.

## When Not to Use It

- Do not use it for workloads that need long-running sessions, stable local state, or specialized host control.
- Do not assume a function runtime is the best fit for every API or background worker.
- Do not assume serverless removes the need for deployment, observability, retry design, or failure ownership.

## Compare To

### Lambda vs. App Runner

Lambda is designed for function-shaped execution.

App Runner is designed for containerized web services. Lambda is better when the workload is event-driven and narrow. App Runner is better when the workload is really an application service with a sustained HTTP shape.

### Lambda vs. EventBridge

EventBridge routes or schedules events.

Lambda runs the code that handles those events. Many systems use both together: EventBridge to trigger, Lambda to execute.

## Tradeoffs

Lambda's biggest advantage is operational simplicity for small execution units. Teams can attach code to events quickly without running servers.

The tradeoff is reduced runtime flexibility. Lambda works best when the workload fits its execution model and becomes awkward when teams try to turn it into a general application host.

Lambda also makes it easy to create many small components. That can improve modularity, but it can also create sprawl if naming, ownership, and observability are weak.

Another tradeoff is that event-driven systems can look simple in code while still being operationally complex once retries, concurrency, and downstream dependencies are involved.

## Common Mistakes

- Building oversized functions that mix unrelated responsibilities.
- Giving many functions the same broad execution role.
- Ignoring idempotency and retry behavior for event-driven processing.
- Connecting functions to VPC resources without understanding the networking and cold-start impact.
- Treating logs as the only observability layer instead of defining metrics, alarms, and operational thresholds.
- Treating deployment success as proof that runtime behavior is correct.

## Cloud Engineering Considerations

### Identity and Access

Each function should usually have its own execution role. Runtime permissions should be narrower than deployment permissions, and downstream access should be explicit.

### Networking

If a function needs private resources, understand how VPC attachment affects connectivity, startup behavior, and access to other AWS services.

### Security

Protect environment configuration, validate incoming events, keep secrets out of code packages, and review how untrusted input reaches downstream services.

### Observability

Use structured logs, metrics, tracing where appropriate, and alarms tied to failures, latency, throttling, and backlog indicators.

### Reliability

Reliable Lambda design depends on idempotency, clear retry expectations, and understanding how the function behaves when downstream services fail or events are delivered more than once.

### Cost

Invocation count, duration, memory choice, and noisy triggers all influence cost. A function that runs often and inefficiently stops being cheap very quickly.

## Project and Pattern Connections

AWS Lambda is most directly connected to:

- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)
- [Project 03: Scheduled API Ingestion](../projects/project-03-scheduled-api-ingestion.md)
- [Project 04: Analytics Platform](../projects/project-04-analytics-platform.md)
- [Serverless API](../patterns/serverless-api.md)
- [Scheduled Job](../patterns/scheduled-job.md)

It matters when a workload needs focused event-driven compute rather than a long-lived application service.

## Official References

- [AWS Lambda documentation](https://docs.aws.amazon.com/lambda/)
- [Lambda developer guide](https://docs.aws.amazon.com/lambda/latest/dg/welcome.html)
- [Best practices for working with AWS Lambda functions](https://docs.aws.amazon.com/lambda/latest/dg/best-practices.html)
