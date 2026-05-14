# AI and Agentic Workloads on AWS

## Purpose

This page frames modern AI and agentic systems on AWS as cloud engineering workloads that need the same rigor as any other production system: clear identity, secure data access, observable runtime behavior, and explicit cost controls.

It is used when a team needs to decide how model access, retrieval, orchestration, runtimes, and safety controls fit together as one system rather than as isolated AI features.

## Definition

AI workloads on AWS are systems that combine model access, retrieval, orchestration, application runtimes, identity, safety controls, and operational telemetry. They should be treated as production systems, not as isolated model experiments.

This is not a single AWS service. It is an architectural category that spans services such as Amazon Bedrock, Knowledge Bases for Bedrock, Agents for Amazon Bedrock, Guardrails for Amazon Bedrock, Lambda, App Runner, Secrets Manager, and CloudWatch.

That distinction matters because the engineering risk often sits outside the model itself. Data access, retrieval quality, tool invocation, safety failures, runtime identity, and cost behavior usually decide whether the workload is truly production ready.

In simple terms:

> AI on AWS becomes cloud engineering when the model is only one part of a larger system that has to be secure, observable, and operable.

## What Problem It Solves

It helps teams turn model access into a real workload. Raw model calls are not enough when a system needs retrieval, tool use, runtime integration, safety controls, identity boundaries, and cost visibility.

It gives teams a framework for understanding how AI features fit inside cloud architecture instead of beside it.

That does not remove experimentation. It means experimentation should happen inside clear data, runtime, and evaluation boundaries.

## How It Is Commonly Used

On AWS, AI systems commonly combine:

- Amazon Bedrock for managed model access,
- Knowledge Bases for retrieval over approved content,
- Agents for Amazon Bedrock for tool use or multi-step orchestration,
- Guardrails for Amazon Bedrock for policy checks,
- application runtimes such as App Runner, Lambda, or API Gateway,
- supporting services for secrets, logging, and access control.

## Foundational Concepts Connected to AI and Agentic Workloads

AI and agentic workloads connect directly to several cloud engineering foundations.

### Model Access and Evaluation

Model access is only useful when teams can evaluate quality, latency, failure behavior, and fit for the real workload. Evaluation is part of engineering, not just experimentation.

### Retrieval and Data Boundaries

Many AI systems are only as good as the data they can ground on. Content ownership, indexing, freshness, and access control matter as much as prompt quality.

### Runtime and Application Architecture

An AI workload still needs a runtime boundary. Requests enter through an application service, then move through model, retrieval, and safety layers that need to be designed like any other distributed system.

### Identity and Access

AI runtimes, retrieval systems, document stores, and tool integrations should not all share the same broad permissions. Identity design is a core part of safe AI architecture.

### Safety and Governance

Prompt injection, unsafe outputs, sensitive source data, and tool misuse all need explicit controls. Safety is not an optional layer added after launch.

### Observability and Cost

AI systems need visibility into latency, model failure, retrieval quality, user outcomes, and spend. Otherwise teams learn about defects and cost problems too late.

## When to Use It

- Use it when planning how Bedrock-based applications fit into existing cloud systems.
- Use it when a workload needs model access, retrieval, tools, or multi-step agent behavior.
- Use it when you need a clear operating model for AI security, monitoring, and cost control.

This page is most useful when the challenge is architectural, not just model selection.

## When Not to Use It

- Do not start with agent orchestration if a simple prompt workflow is enough.
- Do not treat model access as separate from IAM, secrets, data access, and logging design.
- Do not assume a successful prototype is already production-ready.

Do not mistake platform breadth for architectural clarity. Using many AI services without a clear application boundary is still poor system design.

## Compare To

### AI Feature Demos vs. Production AI Systems

A demo can succeed with a prompt, a model call, and happy-path data.

A production AI system needs runtime boundaries, access control, evaluation, safety controls, incident visibility, and cost awareness. The difference is not cosmetic. It is architectural.

### Managed AI Building Blocks vs. Custom Orchestration

Managed building blocks such as Bedrock, Knowledge Bases, and Agents reduce infrastructure and integration work.

Custom orchestration gives more control, but it also shifts more responsibility to the team for evaluation, runtime behavior, security, and long-term operations.

## Tradeoffs

The biggest advantage of AWS AI building blocks is that teams can move faster on model access, retrieval, and supporting controls without assembling every piece from scratch.

The tradeoff is that AI systems are still complex. Managed services do not remove the need for strong application boundaries, data governance, and operational discipline.

AI workloads also make it easy to hide system complexity behind an apparently simple chat interface. That creates risk when teams fail to model what the system can access, invoke, or expose.

Another tradeoff is cost variability. AI systems often look affordable in early testing and then become much more expensive as usage, context size, retrieval volume, and safety layers grow.

## Common Mistakes

- Starting with agents before the retrieval, prompt, and evaluation problems are understood.
- Giving one runtime broad access to models, tools, and data sources without clear boundaries.
- Treating safety as a model setting instead of a system design concern.
- Measuring only token latency instead of the full user-facing workflow.
- Ignoring how fast experimentation can become real spend.
- Confusing model capability with application readiness.

## Cloud Engineering Considerations

### Identity and Access

Review which roles can call models, retrieve source data, invoke tools, and operate surrounding runtimes or pipelines. Runtime identity should be narrower than developer or platform-operator access.

### Networking

Plan how AI runtimes, data sources, and user-facing APIs communicate, especially when private connectivity or VPC controls are required.

### Security

Address prompt injection, sensitive data exposure, tool abuse, unsafe output handling, and guardrail requirements before production rollout.

### Observability

Track latency, model failures, retrieval quality, tool behavior, and user-facing outcomes so the workload is measurable end to end.

### Reliability

Reliable AI systems need fallback behavior, degraded-mode thinking, and clear expectations for what happens when retrieval, model calls, or safety checks fail. The workload should fail predictably rather than opaquely.

### Cost

Costs can increase quickly through token usage, retrieval, orchestration, and repeated experimentation, so budgets and usage reviews matter.

## Project and Pattern Connections

AI and agentic workloads on AWS are most directly connected to:

- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)
- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

The surrounding service pages in this section explain the main building blocks used to make that kind of workload real rather than purely conceptual.

## Official References

- [AWS generative AI overview](https://aws.amazon.com/generative-ai/)
- [Amazon Bedrock documentation](https://docs.aws.amazon.com/bedrock/)
- [Amazon Bedrock user guide](https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-bedrock.html)
