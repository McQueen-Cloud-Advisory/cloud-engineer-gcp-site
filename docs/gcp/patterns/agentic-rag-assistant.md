# Agentic RAG Assistant on Google Cloud

## Purpose

This pattern explains how to build an AI assistant on Google Cloud that combines model access, retrieval, safety controls, and optionally custom agent logic.

## Pattern Summary

An agentic RAG assistant on Google Cloud typically combines an application runtime such as Cloud Run, model access through Vertex AI, retrieval or agent capabilities through Agent Builder, and safety controls such as Model Armor. Supporting services for identity, secrets, and monitoring are part of the architecture, not side details.

This pattern matters because AI applications are operational systems. They need identity, observability, cost visibility, data governance, and clear failure behavior in addition to model quality.

## When This Pattern Fits

Use this pattern when:

- users need answers grounded in approved content,
- the team wants managed model access and retrieval tooling,
- the runtime should stay separate from the model and retrieval layers,
- and safety, monitoring, and cost visibility are part of the initial design.

## When Not to Use It

Do not use this pattern when plain search is enough, when the source corpus is low quality, or when the team is trying to hide weak application design behind agent terminology.

## Common Use Cases

- Internal knowledge assistants
- Retrieval-backed support tools
- Guided AI application workflows

## Reference Architecture

```text
User
-> Cloud Run
-> Vertex AI and Agent Builder
-> Retrieval and safety layers
-> Secrets and monitoring
```

## Why This Pattern Works

It works because the application boundary, model calls, retrieval, safety controls, and runtime operations are separate enough to govern and observe. That gives the team room to improve quality without losing sight of permissions, cost, and failure handling.

## Provider Services

- Cloud Run
- Vertex AI
- Vertex AI Agent Builder
- Agent Development Kit
- Model Garden
- Model Armor
- Secret Manager
- Cloud Monitoring

## Design Considerations

### Security

Review prompt injection, source data trust, runtime identity, and safety controls as part of the initial design.

### Reliability

Define how the system should behave when retrieval fails, models time out, or an agent step produces poor results.

### Observability

Track request volume, retrieval quality, latency, and user-facing outcomes across the full request path.

### Cost

Model usage, retrieval, safety checks, and runtime hosting can all compound quickly, so cost visibility is essential.

### Deployment

Start with a narrow corpus and simple workflow before expanding the surface area of the assistant.

## Common Mistakes

- Treating the assistant like a model endpoint instead of an application system.
- Granting broad service-account or tool permissions.
- Ignoring low-confidence retrieval or low-quality user outcomes.
- Expanding the workflow before monitoring and cost controls are in place.
- Assuming agent terminology automatically means the design is production-ready.

## Related Projects

- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)

## How This Fits Into Cloud Engineering

This pattern matters because AI systems still rely on the same foundations as other cloud workloads: runtime identity, secret handling, observability, and cost control. Good cloud engineering makes those concerns explicit instead of letting the AI layer hide them.

## Official References

- [Vertex AI documentation](https://cloud.google.com/vertex-ai/docs)
- [Agent Builder documentation](https://docs.cloud.google.com/agent-builder)
- [Model Armor overview](https://cloud.google.com/security/products/model-armor)
