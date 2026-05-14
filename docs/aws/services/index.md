# AWS Services

## Purpose

This section explains AWS services from a practical cloud engineering perspective and focuses on the services that power the current learning path.

It is meant to help you understand how services fit into real systems, not just memorize short definitions in isolation.

## Definition

This section is a curated guide to the AWS services used most directly across the site's projects, patterns, and provider-specific learning path.

It is not a full AWS catalog. It is a working set of services chosen because they help explain the main architectural decisions a cloud engineer has to make: runtime, storage, event routing, identity, observability, secrets handling, container delivery, and AI platform design.

In simple terms:

> This section is the AWS service map for the systems this site teaches you to build.

## What Problem This Section Solves

Cloud service catalogs are broad, and that breadth can make early learning messy. People often read isolated service descriptions without understanding how those services work together in a real workload.

This section narrows the scope. It helps you answer practical questions such as:

- which service solves this problem,
- where that service belongs in an architecture,
- what tradeoffs come with choosing it,
- and what identity, networking, observability, and cost implications follow from that choice.

## How To Read This Section

These pages are intended to make the current provider path easier to understand by answering a focused set of questions.

- What problem does this service solve in AWS?
- How does it fit into a project or pattern?
- What identity, networking, and observability assumptions come with it?
- How does it compare to the surrounding runtime and data choices?

A practical reading order is:

1. Start with a project or pattern.
2. Use the relevant service pages to understand the main building blocks.
3. Compare nearby services when a design choice is not obvious.
4. Return to the foundations pages when a service keeps depending on the same core concept.

## What This Service Set Is Designed To Teach

The AWS path emphasizes the services most useful for learning:

- workload identity and access control,
- application hosting and runtime choice,
- object storage and application data boundaries,
- event-driven integration and API front doors,
- monitoring and operational visibility,
- secret handling and safer runtime configuration,
- and how AI features fit into normal cloud engineering rather than sitting outside it.

## Service Categories

### Identity and Access

AWS systems start with trust boundaries. These pages explain how human access, deployment access, and workload identity should be separated.

- [IAM](iam.md)

### Application Hosting and Delivery

These services define how code, containers, or HTTP-facing application logic are delivered. The main question here is whether the workload is a function, a managed app service, or a container delivery path.

- [Lambda](lambda.md)
- [App Runner](app-runner.md)
- [Elastic Container Registry (ECR)](ecr.md)

### Storage and Data

These services hold files and operational application data. The key boundary is whether the workload needs durable objects, a low-latency request-time database, or both.

- [S3](s3.md)
- [DynamoDB](dynamodb.md)

### Integration and Scheduling

These services help systems receive HTTP requests, publish events, or trigger work on a schedule instead of relying only on tightly coupled direct calls.

- [API Gateway](api-gateway.md)
- [EventBridge](eventbridge.md)

### Observability and Operations

These services help teams detect failures, watch health, and operate workloads like systems instead of demos.

- [CloudWatch](cloudwatch.md)

### AI and Agentic Platforms

These pages explain how AI features become part of a real AWS application architecture, including model access, retrieval, safety controls, and managed agent behavior.

- [AI and Agentic Workloads](ai-and-agents.md)
- [Amazon Bedrock](bedrock.md)
- [Agents for Amazon Bedrock](bedrock-agents.md)
- [Knowledge Bases for Amazon Bedrock](knowledge-bases-for-bedrock.md)
- [Guardrails for Amazon Bedrock](guardrails-for-bedrock.md)

### Security and Secrets

These services focus on sensitive runtime configuration and the operational discipline around secret access.

- [Secrets Manager](secrets-manager.md)

## How The Categories Connect

Typical AWS workloads in this site follow a pattern like this:

1. IAM defines who or what can act.
2. API Gateway or EventBridge introduces traffic, integration, or scheduled triggers.
3. Lambda or App Runner provides the runtime boundary, while ECR supports container delivery when needed.
4. S3, DynamoDB, and Secrets Manager support data, configuration, and application state.
5. CloudWatch gives the workload an operational view.
6. Bedrock and the related AI pages build on top of that same application foundation rather than replacing it.

The important lesson is that AWS service choice is not a list-building exercise. It is a system-design exercise.

## What This Section Does Not Try To Do

This section does not try to summarize every AWS product or give certification-style one-liners for the whole platform.

It is intentionally narrower than that. The goal is to build judgment about which services belong in a practical architecture and why.

## Project and Pattern Connections

Use this section together with the AWS project and pattern pages:

- [Projects Overview](../projects/index.md)
- [Patterns Overview](../patterns/index.md)

## How This Fits Into Cloud Engineering

Cloud engineers use service knowledge to make architecture choices, not to build flash cards. This section helps you explain how AWS services work together in a system and what tradeoffs they introduce.

## Official References

- [AWS documentation](https://docs.aws.amazon.com/)
- [AWS Well-Architected Framework](https://docs.aws.amazon.com/wellarchitected/latest/framework/welcome.html)
