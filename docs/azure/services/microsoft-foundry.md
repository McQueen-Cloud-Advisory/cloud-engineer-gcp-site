# Microsoft Foundry

## Purpose

Microsoft Foundry, through Azure AI Foundry, provides a platform surface for building, evaluating, and operating AI applications on Azure.

## Definition

Azure AI Foundry is Microsoft's broader AI platform layer for organizing model usage, evaluation, orchestration, safety, and application-building workflows inside Azure. It is meant to provide more structure than a single model endpoint by giving teams a place to manage the wider lifecycle of AI application work.

That makes it different from Azure OpenAI alone. Azure OpenAI is model access. Foundry is closer to the broader platform context around AI application development and operations.

In simple terms:

> Foundry is the Azure platform layer that helps AI work behave more like an engineered product and less like a pile of model experiments.

## What Problem It Solves

It gives teams a managed environment for organizing model access, safety, evaluation, and application workflows in one platform context.

## How It Is Commonly Used

Foundry is commonly used for:

- evaluating and comparing AI application behavior,
- organizing AI projects beyond a single model endpoint,
- combining models, retrieval, safety, and orchestration components,
- giving teams a more structured AI development experience inside Azure,
- aligning AI application work with Azure governance and operational patterns.

## When to Use It

- Use it when building AI workloads that need more than a single model endpoint.
- Use it when teams need a clearer platform for AI evaluation, management, and integration.
- Use it when AI work needs to fit into Azure governance and operational patterns.

## When Not to Use It

- Do not treat the platform as a substitute for defining retrieval, safety, and application boundaries.
- Do not expand tool choice faster than the team can operate and evaluate it.
- Do not assume a richer platform automatically simplifies the surrounding application architecture.

## Common Mistakes

- Introducing a broad AI platform before the problem and evaluation criteria are clear.
- Treating Foundry features as a replacement for runtime design or secure access patterns.
- Allowing too many people to create or change AI project resources without ownership discipline.
- Tracking model access while ignoring end-user quality and retrieval behavior.
- Expanding platform scope without cost or governance controls.

## Cloud Engineering Considerations

### Identity and Access

Review who can create projects, use models, and manage associated data or search resources.

### Networking

Plan how the AI platform interacts with search, storage, application runtimes, and governance controls.

### Security

Use the platform as part of a wider AI safety and governance approach rather than as a stand-alone control.

### Observability

Track usage, evaluation results, latency, and integration failures as part of the workload lifecycle.

### Cost

Model usage and connected platform services can increase cost quickly during experimentation.

## How This Fits Into Cloud Engineering

Foundry matters because modern AI systems are not just model calls. They are governed application platforms with evaluation, retrieval, safety, and runtime dependencies. Cloud engineering on Azure includes deciding when that broader platform layer helps and when it is more than the workload actually needs.

## Related Projects

- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)

## Related Patterns

- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

## Official References

- [Azure AI Foundry documentation](https://learn.microsoft.com/en-us/azure/ai-foundry/)
- [What is Azure AI Foundry](https://learn.microsoft.com/en-us/azure/ai-foundry/what-is-azure-ai-foundry)
