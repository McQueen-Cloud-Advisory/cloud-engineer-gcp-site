# Agent Development Kit

## Purpose

The Agent Development Kit page covers the more custom end of agent design on Google Cloud, where teams own more of the orchestration, runtime behavior, and evaluation model directly.

It is used when a team needs agent behavior that goes beyond narrow managed configuration and is willing to take on more engineering responsibility for how that agent works.

## Definition

Agent Development Kit, in the context of this site, represents the custom end of agent design on Google Cloud. Instead of relying only on managed agent configuration, it points toward building agent logic that the team owns more directly.

That increases control, but it also increases responsibility. The team now owns more of the orchestration logic, evaluation model, runtime behavior, tool usage, and integration discipline.

In simple terms:

> The Agent Development Kit path is for teams that need custom agent behavior and are willing to own more of the engineering around it.

## What Problem It Solves

It solves the problem of managed agent features not being flexible enough for the workload. Some teams need custom tool logic, more explicit orchestration, specialized evaluation flows, or deeper control over how the agent behaves across steps.

It helps teams think about the application layer of agent development instead of treating managed platform features as the entire solution.

That does not mean custom is better by default. It means custom becomes reasonable only when the use case truly needs it.

## How It Is Commonly Used

It is commonly used for:

- custom agent workflows that exceed narrow managed configuration,
- agent systems with non-trivial tool logic,
- experimental or advanced orchestration that still needs to fit into Google Cloud runtimes,
- workloads where the team wants more direct control over behavior and evaluation,
- cases where managed agent features are useful but not sufficient.

In practice, this path usually sits on top of other Google Cloud services. The runtime still needs model access, service accounts, secrets, monitoring, and a clear deployment target such as Cloud Run.

## Foundational Concepts Connected to Agent Development Kit

Agent Development Kit connects directly to several cloud engineering foundations.

### Orchestration and Tool Use

Custom agent development is mostly about orchestrating model calls, tool invocations, retrieval steps, and decision logic in a way the team can explain and evaluate.

### Runtime Ownership

Once the team owns the agent logic more directly, it also owns deployment, scaling, rollback behavior, and failure handling more directly.

### Identity and Access

Custom agent paths often interact with models, internal APIs, search indexes, data stores, and third-party systems. Those access paths should be kept narrow.

### Evaluation

Custom behavior increases the need for evaluation because the team has fewer managed guardrails. The question is not only whether the agent responds, but whether it behaves acceptably.

### Safety and Cost

Custom orchestration can create higher risk and higher spend if tool use, retry behavior, and model-call chaining are left ungoverned.

## When to Use It

Use the Agent Development Kit path when a custom agent workflow needs more control than narrow managed configuration provides.

Good use cases include:

- agents with non-trivial tool selection or sequencing,
- workflows that need custom state handling or branching behavior,
- systems where evaluation and runtime logic need to be tightly aligned,
- teams that need to integrate agent behavior deeply into an existing application architecture.

This path makes sense when the need for customization is real and defensible, not when it is only interesting.

## When Not to Use It

Do not adopt custom agent tooling before validating that the use case needs it. Extra control is expensive if it does not produce a clear product or operational gain.

Do not separate agent code from the identity, safety, and observability concerns of the surrounding system. Custom logic does not remove the need for platform discipline.

Do not add custom complexity where a simpler managed path would meet the requirement. A custom agent is not automatically a more mature architecture.

## Compare To

### Agent Development Kit vs. Vertex AI Agent Builder

Vertex AI Agent Builder is the more managed path. It reduces some of the custom integration and orchestration work for search, retrieval, and agent-like experiences.

The Agent Development Kit path is better when the team needs more direct control over behavior, tooling, evaluation, or runtime logic than the managed path exposes.

### Agent Development Kit vs. Raw Model Calls

Raw model calls can be enough for narrow AI features with limited orchestration.

Agent Development Kit becomes relevant when the workload needs multi-step logic, tool usage, retrieval coordination, or custom control flow that is more than a single prompt-response pattern.

## Tradeoffs

The biggest advantage of the Agent Development Kit path is flexibility. Teams can design the agent behavior around the actual workload instead of forcing the workload into a narrower managed shape.

The tradeoff is engineering burden. More control means more responsibility for evaluation, safety, observability, debugging, and long-term maintainability.

Custom agent logic can also create hidden complexity. A system that looks intelligent in a demo may be brittle in production if tool failures, ambiguous prompts, or retrieval gaps are not handled explicitly.

Another tradeoff is cost. More orchestration often means more model calls, more tool calls, and more runtime time per user request.

## Common Mistakes

- Building custom orchestration before the basic workflow has been proven.
- Underestimating the operational burden of owning more agent logic.
- Leaving service-account and tool permissions too broad.
- Skipping evaluation design because the system is still "experimental."
- Ignoring how custom logic affects latency, failure handling, and cost.
- Treating a successful demo as proof the agent is safe to operate.

## Cloud Engineering Considerations

### Identity and Access

Review which service accounts can reach models, tools, and data sources used by the agent workflow.

### Networking and Integration Paths

Plan how the agent runtime communicates with models, retrieval systems, and application backends.

### Security and Governance

Custom agent behavior increases the surface area for prompt injection, unsafe tool use, and data leakage, so controls need to be explicit.

### Observability and Evaluation

Track tool usage, step failures, latency, and end-user outcomes as part of the application lifecycle. A custom agent needs evidence that it behaves acceptably, not just evidence that it runs.

### Reliability

Custom agent systems need clear failure behavior. Engineers should know what happens when a tool call fails, when retrieval returns weak content, or when the model produces an unusable step.

### Cost

Custom agent flows can increase model and infrastructure usage, so monitor the whole request path rather than only the final model call.

## Project and Pattern Connections

Agent Development Kit is most directly connected to:

- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)
- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

It matters most when a team moves past simple AI features and starts owning more of the agent behavior as application code.

## Official References

- [Agent Engine overview](https://docs.cloud.google.com/vertex-ai/generative-ai/docs/agent-engine/overview)
- [Vertex AI generative AI overview](https://cloud.google.com/vertex-ai/generative-ai/docs/overview)
