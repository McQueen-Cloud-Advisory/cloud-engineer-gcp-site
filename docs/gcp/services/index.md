# Google Cloud Services

## Purpose

This section explains Google Cloud services from a practical cloud engineering perspective and focuses on the services that power the current learning path.

It is meant to help you understand how services fit into real systems, not just memorize short definitions in isolation.

## Definition

This section is a curated guide to the Google Cloud services used most directly across the site's projects, patterns, and provider-specific learning path.

It is not a full Google Cloud catalog. It is a working set of services chosen because they help explain the main architectural decisions a cloud engineer has to make: runtime, storage, messaging, identity, observability, analytics, secrets, and AI platform design.

In simple terms:

> This section is the Google Cloud service map for the systems this site teaches you to build.

## What Problem This Section Solves

Cloud service catalogs are broad, and that breadth can make early learning messy. People often read isolated service descriptions without understanding how those services work together in a real workload.

This section narrows the scope. It helps you answer practical questions such as:

- which service solves this problem,
- where that service belongs in an architecture,
- what tradeoffs come with choosing it,
- and what identity, networking, observability, and cost implications follow from that choice.

## How To Read This Section

These pages are intended to make the current provider path easier to understand by answering a focused set of questions.

- What problem does this service solve in Google Cloud?
- How does it fit into a project or pattern?
- What identity, networking, and observability assumptions come with it?
- How does it compare to the surrounding runtime and data choices?

This section is most useful when you read it alongside a project or pattern rather than as a disconnected glossary.

A practical reading order is:

1. Start with a project or pattern.
2. Use the relevant service pages to understand the main building blocks.
3. Compare nearby services when a design choice is not obvious.
4. Return to the foundations pages when a service keeps depending on the same core concept.

## What This Service Set Is Designed To Teach

The Google Cloud path emphasizes the services most useful for learning:

- workload identity and access control,
- application hosting and serverless runtime choice,
- storage and analytics boundaries,
- event-driven integration,
- monitoring and operational visibility,
- secret handling and safer runtime configuration,
- and how AI features fit into normal cloud engineering rather than sitting outside it.

## Service Categories

### Identity and Access

Google Cloud systems start with trust boundaries. These pages explain how human access and workload identity should be separated.

- [IAM and Service Accounts](iam-and-service-accounts.md)

### Application Hosting and Delivery

These services define how code or web assets are delivered. The main question here is whether the workload is a static site, a function-sized execution unit, or a containerized application service.

- [Firebase Hosting](firebase-hosting.md)
- [Cloud Run](cloud-run.md)
- [Cloud Functions](cloud-functions.md)

### Storage and Data

These services hold files, raw data, and analytical data. The key boundary is whether the system needs durable objects, analytical querying, or both.

- [Cloud Storage](cloud-storage.md)
- [BigQuery](bigquery.md)

### Messaging and Integration

These services help systems communicate asynchronously instead of through tightly coupled direct calls.

- [Pub/Sub](pubsub.md)

### Observability and Operations

These services help teams detect failures, watch health, and operate workloads like systems instead of demos.

- [Cloud Monitoring](cloud-monitoring.md)

### AI and Agentic Platforms

These pages explain how AI features become part of a real application architecture, including model access, retrieval, safety controls, and custom agent behavior.

- [AI and Agentic Workloads](ai-and-agents.md)
- [Vertex AI](vertex-ai.md)
- [Vertex AI Agent Builder](vertex-ai-agent-builder.md)
- [Agent Development Kit](agent-development-kit.md)
- [Model Garden](model-garden.md)
- [Model Armor](model-armor.md)

### Security and Secrets

These services focus on sensitive runtime configuration and the operational discipline around secret access.

- [Secret Manager](secret-manager.md)

## How The Categories Connect

Typical Google Cloud workloads in this site follow a pattern like this:

1. IAM and service accounts define who or what can act.
2. Firebase Hosting, Cloud Run, or Cloud Functions provide the delivery or runtime boundary.
3. Cloud Storage, Pub/Sub, and Secret Manager support data movement, configuration, and integration.
4. Cloud Monitoring gives the workload an operational view.
5. BigQuery extends the architecture into analytics when historical data needs to be queried.
6. Vertex AI and the related AI pages build on top of that same application foundation rather than replacing it.

The important lesson is that Google Cloud service choice is not a list-building exercise. It is a system-design exercise.

## What This Section Does Not Try To Do

This section does not try to summarize every Google Cloud product or give certification-style one-liners for the whole platform.

It is intentionally narrower than that. The goal is to build judgment about which services belong in a practical architecture and why.

## Project and Pattern Connections

Use this section together with the Google Cloud project and pattern pages:

- [Projects Overview](../projects/index.md)
- [Patterns Overview](../patterns/index.md)

## How This Fits Into Cloud Engineering

Cloud engineers need service knowledge that supports design, deployment, and operations. This section helps you explain how Google Cloud services work together in a system and what tradeoffs they introduce.

## Official References

- [Google Cloud documentation](https://cloud.google.com/docs)
- [Google Cloud architecture framework](https://cloud.google.com/architecture/framework)
