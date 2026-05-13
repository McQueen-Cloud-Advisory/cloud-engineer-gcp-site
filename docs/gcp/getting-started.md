# Google Cloud Getting Started

## Purpose

This page shows how to begin learning Google Cloud through a practical cloud engineering path centered on projects, service accounts, managed services, and clear delivery patterns.

## Understand The Google Cloud Operating Model First

- Projects are the core unit for administration, billing, and many resource boundaries.
- IAM and service accounts define how people and workloads gain access.
- Regions and service availability still matter, even when the platform feels highly managed.
- Some services require APIs to be enabled deliberately as part of setup and automation.

If these basics are clear, the rest of the platform becomes much easier to reason about.

## Recommended First Service Set

Start with a compact group of services that teaches Google Cloud delivery patterns well.

- IAM and service accounts.
- Cloud Storage.
- Cloud Run.
- Cloud Functions.
- Pub/Sub.
- Secret Manager.
- Cloud Monitoring.
- BigQuery.

This set covers public delivery, event-driven work, secret handling, monitoring, and data platform basics.

## Suggested First Steps

1. Create a sandbox project and review billing, IAM, logging, and enabled APIs.
2. Learn how service accounts differ from human access and why that matters for deployments and runtimes.
3. Work through the storage, compute, eventing, secret, and monitoring services used in the early projects.
4. Add analytics and Vertex AI topics after the base application path feels comfortable.

## What To Delay Until Later

You do not need to learn the full Google Cloud portfolio on the first pass. It is reasonable to defer some areas until the fundamentals are stable.

- Broader multi-project governance design beyond the basics.
- Deeper analytics optimization work.
- Vertex AI and agentic workflow topics.
- More advanced network integration patterns.

The idea is to keep your first pass focused enough that you actually finish projects and retain the operating model.

## Common Mistakes

- Mixing human and workload access without clear service-account boundaries.
- Creating resources in several projects or regions without a reason.
- Ignoring which APIs are enabled and how that affects automation.
- Jumping to AI topics before the storage, compute, and observability path is clear.

## How This Fits Into Cloud Engineering

Use this page to identify the first Google Cloud concepts and services you need before moving into project work. The goal is to make Google Cloud feel coherent and buildable, not abstract.

## Official References

- [Google Cloud get started documentation](https://cloud.google.com/docs/get-started)
- [Google Cloud architecture framework](https://cloud.google.com/architecture/framework)
- [Google Cloud global infrastructure](https://cloud.google.com/about/locations)
