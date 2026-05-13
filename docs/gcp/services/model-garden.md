# Model Garden

## Purpose

Model Garden is Google Cloud's catalog of models and related options available through Vertex AI.

## Definition

Model Garden is the model-discovery and selection layer within Vertex AI. It helps teams browse, compare, and choose model options without building a separate procurement or integration path for each model source.

That sounds simple, but model selection is not a cosmetic decision. The model chosen affects latency, safety behavior, evaluation effort, integration style, and cost.

In simple terms:

> Model Garden is where Google Cloud teams decide which models are actually worth putting into an application.

## What Problem It Solves

It helps teams compare and choose model options without building a completely separate discovery path for each provider or runtime.

## How It Is Commonly Used

It is commonly used for:

- comparing model options for a workload,
- evaluating quality, latency, and cost tradeoffs,
- selecting models for RAG or agentic applications,
- giving teams a clearer inventory of available models and capabilities,
- reducing the friction of model experimentation inside Google Cloud.

## When to Use It

- Use it when selecting models for a Google Cloud AI workload.
- Use it when evaluation needs to compare multiple model options for the same application.
- Use it when the team needs a clearer inventory of available models and capabilities.

## When Not to Use It

- Do not choose models only on brand or benchmark assumptions without workload-specific evaluation.
- Do not expand model choice faster than the team can test and govern it.
- Do not treat model selection as a one-time technical choice with no operational impact.

## Common Mistakes

- Choosing a model based on reputation instead of workload fit.
- Ignoring how cost changes at production volume.
- Letting too many models into the platform without governance.
- Measuring prompt quality but not safety, latency, or failure behavior.
- Treating evaluation as optional because the first demo looks good.

## Cloud Engineering Considerations

### Identity and Access

Review which teams or runtimes are allowed to use different models and how model access is governed.

### Networking

Plan how application runtimes reach the chosen models and how that affects latency and integration design.

### Security

Model choice influences data handling and safety controls, so governance should start before deployment.

### Observability

Track which models are used, how they perform, and whether outcomes meet the application need.

### Cost

Different models can have very different cost profiles, so evaluation should include spend as well as quality.

## How This Fits Into Cloud Engineering

Model Garden matters because choosing a model is also choosing an operating profile. Cloud engineering in AI includes deciding which models the platform should allow, how they are evaluated, and how their cost and behavior are monitored over time.

## Related Projects

- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)

## Related Patterns

- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

## Official References

- [Model Garden overview](https://cloud.google.com/vertex-ai/generative-ai/docs/model-garden/overview)
- [Vertex AI documentation](https://cloud.google.com/vertex-ai/docs)
