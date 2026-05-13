# Amazon Bedrock

## Purpose

Amazon Bedrock provides managed access to foundation models and generative AI building blocks on AWS.

## Definition

Amazon Bedrock is AWS's managed platform for using foundation models without operating your own model-serving infrastructure. It provides model access along with related AI platform capabilities that can be combined with retrieval, agents, and safety controls.

The important point is that Bedrock is not only a model endpoint. It is the model-access layer inside a wider AWS application architecture.

In simple terms:

> Bedrock lets AWS applications use managed foundation models as part of a cloud system instead of as a separate experimental stack.

## What Problem It Solves

Bedrock gives teams a managed way to integrate models without standing up their own inference infrastructure, model lifecycle platform, or provider-specific integrations from scratch.

## How It Is Commonly Used

Bedrock is commonly used for:

- prompt-driven application features,
- summarization, generation, extraction, and classification workflows,
- RAG systems with retrieval layered around model calls,
- agentic systems that combine models with tools and policy controls,
- enterprise AI workloads that need AWS-native identity and operations integration.

## When to Use It

- Use it when an application needs managed foundation model access on AWS.
- Use it when you want to integrate prompts, model calls, and AI application features into existing AWS systems.
- Use it as the base model layer for retrieval, guardrails, or agent workflows.

## When Not to Use It

- Do not treat model access as the whole architecture when the workload also needs retrieval, tools, and governance.
- Do not move sensitive prompts or data into production without access review and output validation.
- Do not assume the easiest model integration path is automatically the best production design.

## Common Mistakes

- Selecting models without evaluating quality, latency, and cost against a real workload.
- Granting broad model access to many runtimes without clear ownership.
- Treating prompt design as the only safety control.
- Ignoring how model behavior affects the rest of the application.
- Measuring success by first response quality instead of long-term operational fit.

## Cloud Engineering Considerations

### Identity and Access

Control which roles can call models and separate application runtime permissions from operator access.

### Networking

Review how Bedrock is reached from application runtimes and how those runtimes reach protected data or downstream tools.

### Security

Plan for data classification, prompt safety, output review, and guardrails around model behavior.

### Observability

Track model usage, latency, failures, and downstream application quality together.

### Cost

Model invocation patterns, tokens, and repeated experimentation can quickly affect cost.

## How This Fits Into Cloud Engineering

Bedrock matters because it turns model access into a normal part of AWS platform design. Once AI features touch real users or data, the engineering questions become the same ones that govern every other production workload: identity, networking, failure handling, monitoring, and cost.

## Related Projects

- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)

## Related Patterns

- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

## Official References

- [Amazon Bedrock documentation](https://docs.aws.amazon.com/bedrock/)
- [Amazon Bedrock user guide](https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-bedrock.html)
