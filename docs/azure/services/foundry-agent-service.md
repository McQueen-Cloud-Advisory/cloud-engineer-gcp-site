# Foundry Agent Service

## Purpose

Foundry Agent Service provides a managed layer for building and operating agent-style AI workflows on Azure.

It is used when a team wants managed orchestration for tool use, retrieval, and multi-step behavior without implementing every agent workflow component from scratch.

## Definition

Foundry Agent Service is Azure's managed orchestration layer for AI workflows that need tools, retrieval, and multi-step behavior. It sits above a simple model-call pattern and helps teams structure agent-style application behavior inside the Azure AI platform.

That makes it different from Azure OpenAI by itself. Azure OpenAI gives model access. Foundry Agent Service adds managed orchestration around that access.

Foundry Agent Service is not a replacement for application architecture. It helps with orchestration and tool use, but teams still need to design runtime boundaries, permissions, evaluation, and failure behavior.

In simple terms:

> Foundry Agent Service is the Azure-managed layer for AI workflows that need to plan, retrieve, call tools, and complete multi-step tasks.

## What Problem It Solves

Foundry Agent Service solves the problem of having to build too much orchestration yourself when an Azure AI application needs tool use, retrieval, or multi-step behavior.

It reduces custom orchestration work for AI workloads that need more than a single prompt-response exchange.

That does not remove engineering responsibility. Engineers still need to decide what the agent can access, how tool safety works, how quality is measured, and what the system should do when orchestration fails.

## How It Is Commonly Used

It is commonly used for:

- internal assistants that need tool invocation,
- multi-step enterprise workflows,
- retrieval-backed experiences with structured orchestration,
- AI applications that need a managed agent runtime path,
- teams that want more than a raw model endpoint but less than a fully custom agent framework.

In many Azure AI designs, Foundry Agent Service provides the orchestration layer while Azure OpenAI provides model access and Azure AI Search supports retrieval.

## Foundational Concepts Connected to Foundry Agent Service

Foundry Agent Service connects directly to several cloud engineering foundations.

### Orchestration and Tool Use

Agent workflows depend on more than model output. They also depend on safe tool invocation, stable downstream systems, and clear boundaries around what actions are allowed.

### Runtime Integration

Even managed agent features need a surrounding application runtime. User requests still need a service boundary, authentication model, and supporting operations path.

### Identity and Access

The team needs to decide which tools and data sources the agent can call and which runtimes are allowed to invoke the agent behavior.

### Evaluation

Availability is not enough. Tool quality, orchestration latency, correctness, and user success rates all need to be measured.

### Cost Management

Tool use, retrieval, and connected model usage all contribute to total AI cost.

## When to Use It

Use it when an AI workflow needs more than a single model prompt and response.

Good use cases include:

- tool-calling assistants,
- retrieval-backed multi-step workflows,
- AI applications that need managed agent orchestration,
- teams that want more structure than a custom ad hoc workflow.

Foundry Agent Service is strongest when the system needs more than a model call but not a completely custom orchestration platform.

## When Not to Use It

Do not add agent orchestration before confirming the use case actually needs it.

Do not rely on the agent layer alone for security, access control, or evaluation.

Do not assume multi-step behavior is better if the task can be solved more reliably with a simpler flow.

## Compare To

### Foundry Agent Service vs. Azure OpenAI

Azure OpenAI provides the model access.

Foundry Agent Service adds managed orchestration, tool use, and multi-step behavior around that model access.

### Foundry Agent Service vs. Custom Agent Orchestration

Foundry Agent Service is the more managed path.

Custom orchestration is better when the team needs deeper control over workflow logic, tool routing, or runtime behavior than the managed path provides.

## Tradeoffs

Foundry Agent Service's biggest advantage is speed to a more capable AI experience. Teams can adopt managed orchestration and tool use without building every component by hand.

The tradeoff is reduced control. Managed building blocks are useful, but they do not replace the need to reason about tool quality, answer quality, and runtime behavior.

Foundry Agent Service can also create false confidence because it makes it easier to stand up a working demo. A working demo is not the same as a trustworthy production system.

Another tradeoff is that orchestration quality is still a product and engineering problem. Managed tooling helps, but it cannot make weak tools or poor workflow design correct.

## Common Mistakes

- Giving an agent broad access to tools or data with weak boundaries.
- Skipping evaluation of tool use and step quality.
- Ignoring latency and failure behavior across the whole orchestration path.
- Treating managed agents as a replacement for application design.
- Expanding orchestration complexity without cost controls.
- Assuming that multi-step behavior is automatically more useful than a narrow, predictable workflow.

## Cloud Engineering Considerations

### Identity and Access

Review which identities can configure the agent and which tools or data sources the agent can reach.

### Networking

Plan how the agent runtime reaches models, search services, and application backends.

### Security

Treat tools, retrieved content, and prompts as part of the AI attack surface.

### Observability

Track tool calls, step failures, and user-facing outcomes, not just prompt latency.

### Reliability

Reliable agent-backed systems need clear behavior for tool failures, stale retrieved context, and downstream outages. The application should know what to do when orchestration is weak instead of pretending every answer is equally trustworthy.

### Cost

Agent workflows can drive additional model and tool usage, so monitor full request cost rather than only model calls.

## Project and Pattern Connections

Foundry Agent Service is most directly connected to:

- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)
- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

It matters when the Azure learning path needs orchestration and tool use to be part of the product rather than an afterthought.

## Official References

- [Azure AI Foundry documentation](https://learn.microsoft.com/en-us/azure/ai-foundry/)
- [Azure AI Foundry overview](https://learn.microsoft.com/en-us/azure/ai-foundry/what-is-azure-ai-foundry)
- [Azure AI Foundry agent service](https://learn.microsoft.com/en-us/azure/ai-foundry/agents/overview)
