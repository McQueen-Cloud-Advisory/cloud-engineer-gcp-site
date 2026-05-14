# Amazon Bedrock

## Purpose

Amazon Bedrock provides managed access to foundation models and generative AI building blocks on AWS.

It is used when a team wants AI capabilities to live inside a managed AWS platform instead of being bolted onto the side of the architecture.

## Definition

Amazon Bedrock is AWS's managed platform for using foundation models without operating your own model-serving infrastructure. In the context of this site, it matters most as the managed platform layer for model access, related AI tooling, and integration into wider AWS systems.

Bedrock is broader than a single model endpoint. It is closer to a platform surface for AI application development and operation.

That boundary matters because many teams talk about "using AI" when they are really adopting a platform that influences runtime design, evaluation, access control, observability, and cost.

In simple terms:

> Bedrock is the AWS platform layer that helps AI workloads live inside a real cloud architecture instead of beside it.

## What Problem It Solves

Bedrock solves the problem of assembling AI capabilities into a platform manually. Teams need model access, integration points, and operational controls, but they often do not want to build every part of that platform from scratch.

It gives teams a managed platform for building AI applications without assembling every model and tooling component themselves.

That does not make AI architecture automatic. Engineers still need to decide how models are chosen, what data they can access, how outputs are evaluated, and where the application runtime boundary lives.

## How It Is Commonly Used

Bedrock is commonly used for:

- prompt-driven application features,
- summarization, generation, extraction, and classification workflows,
- RAG systems with retrieval layered around model calls,
- agentic systems that combine models with tools and policy controls,
- enterprise AI workloads that need AWS-native identity and operations integration.

In many AWS architectures, Bedrock is not the entire application. It is the AI platform surface behind a runtime such as App Runner or Lambda and around supporting services such as Knowledge Bases, Agents, Guardrails, Secrets Manager, and CloudWatch.

## Foundational Concepts Connected to Bedrock

Bedrock connects directly to several cloud engineering foundations.

### Models and Endpoints

Model access is the most visible part of Bedrock, but it is only useful when teams understand which models to use, how they are invoked, and what kinds of latency and failure behavior they create.

### Evaluation

AI quality cannot be judged only by whether a request returns an answer. Teams need evaluation workflows that compare models, prompts, grounding quality, and failure cases.

### Runtime Integration

Most AI features still live inside a broader application. Bedrock needs a surrounding runtime, security model, and operational design.

### Identity and Access

Models, supporting AI resources, and connected services all need controlled access. AI platforms expand the number of sensitive access paths in a system.

### Observability and Cost

AI platforms need visibility into usage, latency, failure rate, and spend. Otherwise teams cannot tell whether the system is useful or merely expensive.

## When to Use It

- Use it when an application needs managed foundation model access on AWS.
- Use it when you want to integrate prompts, model calls, and AI application features into existing AWS systems.
- Use it as the base model layer for retrieval, guardrails, or agent workflows.

Bedrock is strongest when the challenge is operating AI well, not just calling a model once.

## When Not to Use It

- Do not treat model access as the whole architecture when the workload also needs retrieval, tools, and governance.
- Do not move sensitive prompts or data into production without access review and output validation.
- Do not assume the easiest model integration path is automatically the best production design.

Do not expand platform usage faster than the team can evaluate and operate it safely. Platform richness can create the illusion of progress while increasing operational risk.

## Compare To

### Bedrock vs. Bedrock Knowledge Bases

Knowledge Bases is the managed retrieval layer inside the Bedrock ecosystem.

Bedrock is broader. It is the overall managed platform for model access and surrounding AI workflows.

### Bedrock vs. Agents for Amazon Bedrock

Bedrock gives the broader AI platform surface.

Agents for Amazon Bedrock sits higher in the stack and focuses more specifically on managed orchestration, tool use, and multi-step AI experiences built on top of that platform.

## Tradeoffs

Bedrock's biggest advantage is platform integration. Teams can work with managed AI capabilities inside a cloud environment that already has identity, logging, and related services.

The tradeoff is platform complexity. Once teams move beyond a simple demo, they need to reason about model choice, evaluation, access control, latency, and cost as one operating system rather than as separate concerns.

Bedrock also lowers the friction of experimentation. That is valuable, but it can encourage growth in model usage and AI features faster than the team can govern or evaluate properly.

Another tradeoff is that the platform makes AI easier to adopt, not automatically easier to operate well.

## Common Mistakes

- Selecting models without evaluating quality, latency, and cost against a real workload.
- Granting broad model access to many runtimes without clear ownership.
- Treating prompt design as the only safety control.
- Ignoring how model behavior affects the rest of the application.
- Measuring success by first response quality instead of long-term operational fit.
- Assuming model availability is the same thing as application readiness.

## Cloud Engineering Considerations

### Identity and Access

Control which roles can call models and separate application runtime permissions from operator access.

### Networking

Review how Bedrock is reached from application runtimes and how those runtimes reach protected data or downstream tools.

### Security

Plan for data classification, prompt safety, output review, and guardrails around model behavior.

### Observability

Track model usage, latency, failures, and downstream application quality together.

### Reliability

Reliable AI platforms need fallback thinking, clear handling for failed model calls, and defined behavior when related systems such as retrieval or safety controls are unavailable. AI workloads should fail in explainable ways.

### Cost

Model invocation patterns, tokens, and repeated experimentation can quickly affect cost.

## Project and Pattern Connections

Amazon Bedrock is most directly connected to:

- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)
- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

It is the main AI platform layer behind the other AI-focused pages in this section.

## Official References

- [Amazon Bedrock documentation](https://docs.aws.amazon.com/bedrock/)
- [Amazon Bedrock user guide](https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-bedrock.html)
- [Amazon Bedrock foundation models overview](https://docs.aws.amazon.com/bedrock/latest/userguide/foundation-models-reference.html)
