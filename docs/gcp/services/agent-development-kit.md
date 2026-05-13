# Agent Development Kit

## Purpose

The Agent Development Kit page covers Google Cloud-adjacent tooling for building custom agent behavior that still needs a clear runtime, identity, and evaluation model.

## Definition

Agent Development Kit represents the more custom end of agent design on Google Cloud. Instead of relying only on managed agent configuration, it points toward building agent logic that the team owns more directly.

That increases control, but it also increases responsibility. The team now owns more of the orchestration logic, evaluation model, runtime behavior, and integration discipline.

In simple terms:

> The Agent Development Kit path is for teams that need custom agent behavior and are willing to own more of the engineering around it.

## What Problem It Solves

It helps teams think about the application layer of agent development instead of treating managed platform features as the entire solution.

## How It Is Commonly Used

It is commonly used for:

- custom agent workflows that exceed narrow managed configuration,
- agent systems with non-trivial tool logic,
- experimental or advanced orchestration that still needs to fit into Google Cloud runtimes,
- workloads where the team wants more direct control over behavior and evaluation,
- cases where managed agent features are useful but not sufficient.

## When to Use It

- Use it when a custom agent workflow needs more control than a narrow managed configuration provides.
- Use it when the team needs to reason about tools, prompts, evaluation, and runtime behavior together.
- Use it when agent logic must fit into a larger Google Cloud deployment and operations model.

## When Not to Use It

- Do not adopt custom agent tooling before validating that the use case needs it.
- Do not separate agent code from the identity, safety, and observability concerns of the surrounding system.
- Do not add custom complexity where a simpler managed path would meet the requirement.

## Common Mistakes

- Building custom orchestration before the basic workflow has been proven.
- Underestimating the operational burden of owning more agent logic.
- Leaving service-account and tool permissions too broad.
- Skipping evaluation design because the system is still "experimental."
- Ignoring how custom logic affects latency, failure handling, and cost.

## Cloud Engineering Considerations

### Identity and Access

Review which service accounts can reach models, tools, and data sources used by the agent workflow.

### Networking

Plan how the agent runtime communicates with models, retrieval systems, and application backends.

### Security

Custom agent behavior increases the surface area for prompt injection, unsafe tool use, and data leakage, so controls need to be explicit.

### Observability

Track tool usage, step failures, latency, and end-user outcomes as part of the application lifecycle.

### Cost

Custom agent flows can increase model and infrastructure usage, so monitor the whole request path.

## How This Fits Into Cloud Engineering

The Agent Development Kit path matters because it exposes the real engineering cost of agent customization. More control can be useful, but only if the team is ready to own runtime design, evaluation, and operational discipline end to end.

## Related Projects

- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)

## Related Patterns

- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

## Official References

- [Agent Engine overview](https://docs.cloud.google.com/vertex-ai/generative-ai/docs/agent-engine/overview)
- [Vertex AI generative AI overview](https://cloud.google.com/vertex-ai/generative-ai/docs/overview)
