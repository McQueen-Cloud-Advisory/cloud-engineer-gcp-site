# Vertex AI

## Purpose

Vertex AI is Google Cloud's managed AI platform for model access, development, and deployment workflows.

It is used when a team wants AI and machine learning capabilities to live inside a managed Google Cloud platform instead of being bolted onto the side of the architecture.

## Definition

Vertex AI is the main Google Cloud platform for AI and machine learning workloads. In the context of this site, it matters most as the managed platform layer for model access, evaluation, related AI tooling, and integration into wider Google Cloud systems.

Vertex AI is broader than a single model endpoint. It is closer to a platform surface for AI application development and operation.

That boundary matters because many teams talk about "using AI" when they are really adopting a platform that influences runtime design, evaluation, access control, observability, and cost.

In simple terms:

> Vertex AI is the Google Cloud platform layer that helps AI workloads live inside a real cloud architecture instead of beside it.

## What Problem It Solves

Vertex AI solves the problem of assembling AI capabilities into a platform manually. Teams need model access, evaluation, integration points, and operational controls, but they often do not want to build every part of that platform from scratch.

It gives teams a managed platform for building AI applications without assembling every model and tooling component themselves.

That does not make AI architecture automatic. Engineers still need to decide how models are chosen, what data they can access, how outputs are evaluated, and where the application runtime boundary lives.

## How It Is Commonly Used

Vertex AI is commonly used for:

- managed model access,
- generative AI application features,
- evaluation and comparison workflows,
- retrieval-backed assistants and agent-like systems when combined with surrounding tools,
- AI workloads that need to fit into Google Cloud governance and operations,
- broader machine learning platform workflows where managed AI infrastructure is useful.

In many Google Cloud architectures, Vertex AI is not the entire application. It is the AI platform surface behind a runtime such as Cloud Run and around supporting services such as Model Garden, Agent Builder, Model Armor, Secret Manager, and Cloud Monitoring.

## Foundational Concepts Connected to Vertex AI

Vertex AI connects directly to several cloud engineering foundations.

### Models and Endpoints

Model access is the most visible part of Vertex AI, but it is only useful when teams understand which models to use, how they are invoked, and what kinds of latency and failure behavior they create.

### Evaluation

AI quality cannot be judged only by whether a request returns an answer. Teams need evaluation workflows that compare models, prompts, grounding quality, and failure cases.

### Runtime Integration

Most AI features still live inside a broader application. Vertex AI needs a surrounding runtime, security model, and operational design.

### Identity and Access

Models, datasets, supporting AI resources, and connected services all need controlled access. AI platforms expand the number of sensitive access paths in a system.

### Observability and Cost

AI platforms need visibility into usage, latency, failure rate, and spend. Otherwise teams cannot tell whether the system is useful or merely expensive.

## When to Use It

Use Vertex AI when the workload needs managed AI platform capabilities inside Google Cloud.

Good use cases include:

- applications that need model access under Google Cloud governance,
- generative AI features with managed platform support,
- evaluation and model comparison work,
- agentic or retrieval-backed systems that need a larger AI platform around them,
- teams that want AI workloads to fit normal cloud operations rather than remain isolated experiments.

Vertex AI is strongest when the challenge is operating AI well, not just calling a model once.

## When Not to Use It

Do not treat model access as the whole architecture when retrieval, safety, or runtime concerns are still undefined.

Do not expand platform usage faster than the team can evaluate and operate it safely. Platform richness can create the illusion of progress while increasing operational risk.

Do not confuse platform breadth with production readiness. A managed AI platform still needs workload-specific evaluation, security, and ownership.

## Compare To

### Vertex AI vs. Model Garden

Model Garden is the model-discovery and selection surface inside Vertex AI.

Vertex AI is broader. It is the overall managed platform for model access, evaluation, and surrounding AI workflows.

### Vertex AI vs. Vertex AI Agent Builder

Vertex AI gives the broader AI platform surface.

Vertex AI Agent Builder sits higher in the stack and focuses more specifically on managed retrieval, search, and agent-like experiences built on top of that platform.

## Tradeoffs

Vertex AI's biggest advantage is platform integration. Teams can work with managed AI capabilities inside a cloud environment that already has identity, monitoring, and related services.

The tradeoff is platform complexity. Once teams move beyond a simple demo, they need to reason about model choice, evaluation, access control, latency, and cost as one operating system rather than as separate concerns.

Vertex AI also lowers the friction of experimentation. That is valuable, but it can encourage growth in model usage and AI features faster than the team can govern or evaluate properly.

Another tradeoff is that the platform makes AI easier to adopt, not automatically easier to operate well.

## Common Mistakes

- Choosing platform features before defining the workload and evaluation criteria.
- Leaving service-account and data-access boundaries vague.
- Focusing on model quality while ignoring application latency and runtime behavior.
- Treating platform telemetry as a substitute for end-to-end observability.
- Letting experimentation scale without cost guardrails.
- Assuming model availability is the same thing as application readiness.

## Cloud Engineering Considerations

### Identity and Access

Review which service accounts can access models, data, and surrounding AI platform resources. AI systems often create more privileged access paths than teams expect initially.

### Networking and Data Movement

Plan how AI runtimes, retrieval layers, and application services connect to Vertex AI endpoints. Model access is only part of the network path; retrieval systems and application services add their own data-movement concerns.

### Security and Governance

Address data classification, prompt safety, and output handling as part of the design, not after deployment.

### Observability and Evaluation

Track model usage, failures, latency, and outcome quality as part of the application lifecycle. The core operational question is whether the system produces acceptable results, not just whether the platform responded.

### Reliability

Reliable AI platforms need fallback thinking, clear handling for failed model calls, and defined behavior when related systems such as retrieval or safety controls are unavailable. AI workloads should fail in explainable ways.

### Cost

Model usage and supporting AI platform features can increase cost quickly during experimentation and production. Cost controls should be designed before high-volume adoption, not after surprise bills.

## Project and Pattern Connections

Vertex AI is most directly connected to:

- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)
- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

It is the main AI platform layer behind the other AI-focused pages in this section.

## Official References

- [Vertex AI documentation](https://cloud.google.com/vertex-ai/docs)
- [Vertex AI overview](https://cloud.google.com/vertex-ai/docs/start/introduction-unified-platform)
- [Generative AI on Vertex AI overview](https://cloud.google.com/vertex-ai/generative-ai/docs/overview)
