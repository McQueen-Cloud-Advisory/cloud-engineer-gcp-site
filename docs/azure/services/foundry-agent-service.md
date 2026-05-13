# Foundry Agent Service

## Purpose

Foundry Agent Service provides a managed layer for building and operating agent-style AI workflows on Azure.

## Definition

Foundry Agent Service is Azure's managed orchestration layer for AI workflows that need tools, retrieval, and multi-step behavior. It sits above a simple model-call pattern and helps teams structure agent-style application behavior inside the Azure AI platform.

That makes it different from Azure OpenAI by itself. Azure OpenAI gives model access. Foundry Agent Service adds managed orchestration around that access.

In simple terms:

> Foundry Agent Service is the Azure-managed layer for AI workflows that need to plan, retrieve, call tools, and complete multi-step tasks.

## What Problem It Solves

It reduces custom orchestration work when an Azure AI application needs tool use, retrieval, or multi-step behavior.

## How It Is Commonly Used

It is commonly used for:

- internal assistants that need tool invocation,
- multi-step enterprise workflows,
- retrieval-backed experiences with structured orchestration,
- AI applications that need a managed agent runtime path,
- teams that want more than a raw model endpoint but less than a fully custom agent framework.

## When to Use It

- Use it when an AI workflow needs more than a single model prompt and response.
- Use it when Azure AI applications need a managed agent runtime path.
- Use it when tool invocation and multi-step reasoning must be part of the design.

## When Not to Use It

- Do not add agent orchestration before confirming the use case actually needs it.
- Do not rely on the agent layer alone for security, access control, or evaluation.
- Do not assume multi-step behavior is better if the task can be solved more reliably with a simpler flow.

## Common Mistakes

- Giving an agent broad access to tools or data with weak boundaries.
- Skipping evaluation of tool use and step quality.
- Ignoring latency and failure behavior across the whole orchestration path.
- Treating managed agents as a replacement for application design.
- Expanding orchestration complexity without cost controls.

## Cloud Engineering Considerations

### Identity and Access

Review which identities can configure the agent and which tools or data sources the agent can reach.

### Networking

Plan how the agent runtime reaches models, search services, and application backends.

### Security

Treat tools, retrieved content, and prompts as part of the AI attack surface.

### Observability

Track tool calls, step failures, and user-facing outcomes, not just prompt latency.

### Cost

Agent workflows can drive additional model and tool usage, so monitor full request cost rather than only model calls.

## How This Fits Into Cloud Engineering

Foundry Agent Service matters because orchestration is where AI systems become application systems. Once an assistant can take actions, retrieve content, or call backends, the workload needs strong identity, safety, telemetry, and failure-handling design.

## Related Projects

- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)

## Related Patterns

- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

## Official References

- [Azure AI Foundry documentation](https://learn.microsoft.com/en-us/azure/ai-foundry/)
- [Azure AI Foundry overview](https://learn.microsoft.com/en-us/azure/ai-foundry/what-is-azure-ai-foundry)
