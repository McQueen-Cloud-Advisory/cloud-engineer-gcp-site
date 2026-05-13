# Observability

## Purpose

This page explains how logs, metrics, traces, dashboards, and alerts help engineers understand system behavior and respond effectively when that behavior changes.

## Core Concepts

Observability is the ability to infer what is happening inside a system from the signals it emits. Monitoring tools are part of that, but observability is broader. The goal is not only to know that something is broken. The goal is to understand why.

The main signal types answer different questions.

- Metrics show trends and health indicators such as latency, error rate, throughput, or resource usage.
- Logs capture event detail and are useful when you need to inspect what happened in a specific request or task.
- Traces show how a request moved across services and where time or errors accumulated.

Strong observability combines these instead of relying on one signal alone.

## What Good Observability Looks Like

- Important user journeys and background jobs are instrumented on purpose.
- Alerts are tied to symptoms that matter, not every possible event.
- Dashboards help people see service health quickly during normal operation and incidents.
- Engineers can correlate logs, metrics, and traces for the same failure path.
- Runbooks or response notes exist for the alerts the team expects to receive.

Observability should be designed with the system. If it is added only after incidents, the hardest failure paths usually remain invisible.

## Common Mistakes

- Sending huge amounts of telemetry without deciding what questions it should answer.
- Creating noisy alerts that train people to ignore them.
- Measuring infrastructure only and ignoring user-facing symptoms.
- Retaining logs or traces indefinitely without cost review.
- Assuming a dashboard is enough without testing whether it helps during failure.

## How This Fits Into Cloud Engineering

Cloud engineers rely on observability for debugging, incident response, performance tuning, deployment confidence, and cost visibility. Weak observability slows every other engineering activity because teams cannot prove what the system is doing or whether a change actually improved it.

## Official References

- [Google Cloud Monitoring overview](https://cloud.google.com/monitoring/docs/monitoring-overview)
- [Amazon CloudWatch documentation](https://docs.aws.amazon.com/cloudwatch/)
- [Azure Monitor overview](https://learn.microsoft.com/en-us/azure/azure-monitor/overview)
