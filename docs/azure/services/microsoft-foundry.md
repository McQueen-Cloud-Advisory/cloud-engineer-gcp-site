# Microsoft Foundry

## Purpose

Microsoft Foundry, through Azure AI Foundry, provides a platform surface for building, evaluating, and operating AI applications on Azure.

It is used when a team needs more structure than a single model endpoint and wants AI work to live inside a governed Azure platform rather than as a loose set of experiments.

## Definition

Azure AI Foundry is Microsoft's broader AI platform layer for organizing model usage, evaluation, orchestration, safety, and application-building workflows inside Azure. It is meant to provide more structure than a single model endpoint by giving teams a place to manage the wider lifecycle of AI application work.

That makes it different from Azure OpenAI alone. Azure OpenAI is model access. Foundry is closer to the broader platform context around AI application development and operations.

Foundry is not the entire application architecture. It is the AI platform layer inside a broader system that still needs runtime boundaries, identity design, retrieval quality, safety controls, and operational discipline.

In simple terms:

> Foundry is the Azure platform layer that helps AI work behave more like an engineered product and less like a pile of model experiments.

## What Problem It Solves

Foundry solves the problem of trying to turn model access into a real workload without any shared platform surface for evaluation, governance, orchestration, and surrounding AI components.

It gives teams a managed environment for organizing model access, safety, evaluation, and application workflows in one platform context.

That does not remove engineering responsibility. Engineers still need to define retrieval boundaries, runtime ownership, access controls, and what success actually looks like for the AI system.

## How It Is Commonly Used

Foundry is commonly used for:

- evaluating and comparing AI application behavior,
- organizing AI projects beyond a single model endpoint,
- combining models, retrieval, safety, and orchestration components,
- giving teams a more structured AI development experience inside Azure,
- aligning AI application work with Azure governance and operational patterns.

In many Azure AI workloads, Foundry becomes the coordinating platform around Azure OpenAI, Azure AI Search, safety controls, and the surrounding application runtime.

## Foundational Concepts Connected to Microsoft Foundry

Foundry connects directly to several cloud engineering foundations.

### AI Platform Design

Foundry is a platform surface, not just a single capability. It is relevant when the AI workload needs a governed operating model rather than only an endpoint to call.

### Evaluation

AI systems need to be evaluated for quality, latency, grounding, and failure behavior. Foundry matters when teams need evaluation to be part of the platform rather than an afterthought.

### Retrieval and Orchestration

Many AI systems need more than raw model access. They also need retrieval, tools, and surrounding workflow structure.

### Identity and Access

AI platforms expand the number of sensitive access paths in a system. Model usage, project administration, and connected data services all need deliberate boundaries.

### Observability and Cost

AI platforms need visibility into usage, latency, failures, evaluation outcomes, and spend. Otherwise teams cannot tell whether the system is useful or merely expensive.

## When to Use It

Use Foundry when building AI workloads that need more than a single model endpoint.

Good use cases include:

- AI projects that need evaluation and management structure,
- workloads that combine model access with retrieval or agent behavior,
- teams that want AI work to fit into Azure governance and operational patterns,
- moving from prototype experimentation toward a more operable AI platform model.

Foundry is strongest when the challenge is operating AI well, not only getting an answer from a model.

## When Not to Use It

Do not treat the platform as a substitute for defining retrieval, safety, and application boundaries.

Do not expand tool choice faster than the team can operate and evaluate it.

Do not assume a richer platform automatically simplifies the surrounding application architecture.

## Compare To

### Microsoft Foundry vs. Azure OpenAI

Azure OpenAI provides model access.

Foundry is the broader platform context for organizing model usage, evaluation, orchestration, and surrounding AI workflows.

### Microsoft Foundry vs. Foundry Agent Service

Foundry is the broader platform layer.

Foundry Agent Service sits higher in the stack and focuses more specifically on agent-style orchestration and tool use.

## Tradeoffs

Foundry's biggest advantage is structure. It gives teams a managed place to treat AI work as a governed application platform instead of a scattered set of model experiments.

The tradeoff is platform breadth. Broader capability can improve consistency, but it also introduces more moving parts that teams need to understand and operate.

Foundry also lowers the friction of experimentation with more platform support. That is useful, but it can encourage feature growth faster than evaluation, safety, and ownership mature.

Another tradeoff is that a strong platform surface still cannot rescue a weak application boundary or unclear product goal.

## Common Mistakes

- Introducing a broad AI platform before the problem and evaluation criteria are clear.
- Treating Foundry features as a replacement for runtime design or secure access patterns.
- Allowing too many people to create or change AI project resources without ownership discipline.
- Tracking model access while ignoring end-user quality and retrieval behavior.
- Expanding platform scope without cost or governance controls.
- Assuming that having a platform means the AI system is ready for production.

## Cloud Engineering Considerations

### Identity and Access

Review who can create projects, use models, and manage associated data or search resources.

### Networking

Plan how the AI platform interacts with search, storage, application runtimes, and governance controls.

### Security

Use the platform as part of a wider AI safety and governance approach rather than as a stand-alone control.

### Observability

Track usage, evaluation results, latency, and integration failures as part of the workload lifecycle.

### Reliability

Reliable AI platforms need fallback thinking, clear behavior for failed dependencies, and an operating model that explains what happens when model, retrieval, or safety layers degrade.

### Cost

Model usage and connected platform services can increase cost quickly during experimentation.

## Project and Pattern Connections

Microsoft Foundry is most directly connected to:

- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)
- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

It matters when the Azure learning path moves from simple model access into a broader AI platform operating model.

## Official References

- [Azure AI Foundry documentation](https://learn.microsoft.com/en-us/azure/ai-foundry/)
- [What is Azure AI Foundry](https://learn.microsoft.com/en-us/azure/ai-foundry/what-is-azure-ai-foundry)
- [Azure AI Foundry concepts](https://learn.microsoft.com/en-us/azure/ai-foundry/concepts/overview)
