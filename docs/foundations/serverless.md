# Serverless

## Purpose

This page explains serverless as an operating model in which the platform takes on more infrastructure management so teams can focus more on code, events, and application behavior.

## Core Concepts

Serverless usually means you do not manage servers directly, even though servers still exist underneath. The provider handles more of the provisioning, scaling, and runtime management. In exchange, your application has to fit the platform's constraints.

Serverless is often associated with event-driven functions, but the broader idea is managed execution with less host ownership. It tends to work best when workloads are stateless, bursty, short-lived, or triggered by HTTP requests, schedules, or platform events.

The benefits are real.

- Less operating system and host management.
- Fast scaling for uneven workloads.
- Easier alignment between usage and cost for some traffic patterns.

But serverless is not magic. Teams still need to think about concurrency, retries, timeouts, cold starts, permissions, logging, and where application state actually lives.

## When It Fits Well

- Event-driven processing.
- Lightweight APIs or backends.
- Scheduled automation.
- Systems where paying for idle capacity would be wasteful.

It fits poorly when the workload needs long-lived processes, deep operating system control, unusual runtime behavior, or consistent low-latency startup with no tolerance for platform constraints.

## Common Mistakes

- Treating functions as if they were miniature servers with hidden state.
- Ignoring idempotency and retry behavior in event-driven paths.
- Storing secrets or connection logic in ways that do not survive scaling well.
- Assuming low traffic always means low cost if surrounding services are still expensive.
- Choosing serverless before confirming the workload actually matches the model.

## How This Fits Into Cloud Engineering

Cloud engineers use serverless to reduce undifferentiated infrastructure work, but the real job is still system design. They need to shape permissions, failure handling, observability, deployments, and cost around the runtime model. Good serverless systems feel simple to operate because the hard edges were designed on purpose.

## Official References

- [Google Cloud Functions overview](https://cloud.google.com/functions/docs/concepts/overview)
- [AWS Lambda documentation](https://docs.aws.amazon.com/lambda/)
- [Azure Functions overview](https://learn.microsoft.com/en-us/azure/azure-functions/functions-overview)
