# Model Garden

## Purpose

Model Garden is Google Cloud's catalog of models and related options available through Vertex AI.

It is used when teams need to compare, choose, and govern model options without treating model selection as an ad hoc decision made from memory or hype.

## Definition

Model Garden is the model-discovery and selection layer within Vertex AI. It helps teams browse, compare, and choose model options without building a separate procurement or integration path for each model source.

That sounds simple, but model selection is not a cosmetic decision. The model chosen affects latency, safety behavior, evaluation effort, integration style, and cost.

Model Garden is not the entire AI platform. It is the discovery and selection surface within a larger platform and workload architecture. That boundary matters because many teams treat choosing a model as if it were the whole AI design decision.

In simple terms:

> Model Garden is where Google Cloud teams decide which models are actually worth putting into an application.

## What Problem It Solves

Model Garden solves the problem of model sprawl and shallow selection. Without a structured discovery surface, teams often choose models based on reputation, isolated demos, or whichever API they touched first.

It helps teams compare and choose model options without building a completely separate discovery path for each provider or runtime.

That does not remove the need for evaluation. It makes evaluation more practical by giving teams a clear place to explore the options available through Vertex AI.

## How It Is Commonly Used

It is commonly used for:

- comparing model options for a workload,
- evaluating quality, latency, and cost tradeoffs,
- selecting models for RAG or agentic applications,
- giving teams a clearer inventory of available models and capabilities,
- reducing the friction of model experimentation inside Google Cloud.

In many teams, Model Garden becomes the shortlist tool that turns vague model exploration into a more disciplined selection process tied to workload needs.

## Foundational Concepts Connected to Model Garden

Model Garden connects directly to several cloud engineering foundations.

### Model Selection

Model choice affects far more than output quality. It also affects latency, safety posture, availability expectations, and how the runtime should be designed.

### Evaluation

Model Garden only creates value when teams use it as part of a comparison and evaluation loop rather than as a catalog to browse casually.

### Governance

Organizations need to know which models are allowed, who can use them, and how production choices are reviewed. Model access without governance turns into platform sprawl quickly.

### Observability

Model usage should be visible enough for teams to tell which models are being used, how they behave, and whether they still match the application need over time.

### Cost Management

Different models create different cost profiles. Selection work is incomplete if it ignores what production usage will cost.

## When to Use It

Use Model Garden when selecting models for a Google Cloud AI workload.

Good use cases include:

- comparing multiple model options for the same application,
- creating an informed shortlist before deeper evaluation,
- selecting models for retrieval-backed or agentic systems,
- giving teams a clearer inventory of available model capabilities inside Vertex AI.

Model Garden is most useful when the organization wants model choice to be intentional rather than accidental.

## When Not to Use It

Do not choose models only on brand or benchmark assumptions without workload-specific evaluation.

Do not expand model choice faster than the team can test and govern it. More available models can create more confusion if there is no selection discipline.

Do not treat model selection as a one-time technical choice with no operational impact. Production traffic often changes which tradeoffs matter most.

## Compare To

### Model Garden vs. Vertex AI

Vertex AI is the broader AI platform for model access, tooling, evaluation, and application integration.

Model Garden is the model-discovery and selection surface within that broader platform. It helps teams decide which models deserve deeper evaluation or use.

### Model Garden vs. Choosing Models from Benchmarks Alone

Benchmarks and reputation can be helpful inputs, but they do not tell the whole story for a specific workload.

Model Garden is most useful when it leads into workload-specific evaluation of quality, latency, safety, and cost rather than replacing that evaluation.

## Tradeoffs

Model Garden's biggest advantage is visibility. It reduces the friction of discovering and comparing model options inside Google Cloud.

The tradeoff is that more visible options can encourage more experimentation than the team can govern well. A wide catalog is useful only when the organization has a clear model-selection process.

Model Garden also helps teams move faster in early exploration. That is useful, but it can create weak decisions if teams stop at "it worked once" rather than evaluating how the model behaves under real workload conditions.

Another tradeoff is that model choice can look reversible on paper while still being costly to change in a live application because prompts, evaluation baselines, and safety settings may all need revision.

## Common Mistakes

- Choosing a model based on reputation instead of workload fit.
- Ignoring how cost changes at production volume.
- Letting too many models into the platform without governance.
- Measuring prompt quality but not safety, latency, or failure behavior.
- Treating evaluation as optional because the first demo looks good.
- Assuming the best benchmark result will be the best operational choice.

## Cloud Engineering Considerations

### Identity and Access

Review which teams or runtimes are allowed to use different models and how model access is governed.

### Networking and Runtime Integration

Plan how application runtimes reach the chosen models and how that affects latency and integration design.

### Security and Governance

Model choice influences data handling and safety controls, so governance should start before deployment.

### Observability and Evaluation

Track which models are used, how they perform, and whether outcomes meet the application need. The right model is the one that behaves acceptably in the live workload, not just in early testing.

### Reliability

Teams should understand how model availability, latency variation, and failure behavior affect the application. A model choice is part of the reliability profile, not just part of the feature set.

### Cost

Different models can have very different cost profiles, so evaluation should include spend as well as quality.

## Project and Pattern Connections

Model Garden is most directly connected to:

- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)
- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

It matters anywhere teams need model choice to be governed by workload fit instead of by whichever option happened to be tried first.

## Official References

- [Model Garden overview](https://cloud.google.com/vertex-ai/generative-ai/docs/model-garden/overview)
- [Vertex AI documentation](https://cloud.google.com/vertex-ai/docs)
