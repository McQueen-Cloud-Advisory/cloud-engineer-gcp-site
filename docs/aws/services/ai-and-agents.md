# AI and Agentic Workloads on AWS

## Purpose

This page frames modern AI and agentic systems on AWS as cloud engineering workloads that need the same rigor as any other production system: clear identity, secure data access, observable runtime behavior, and explicit cost controls.

## Definition

AI workloads on AWS are not only about calling a model. They are systems that combine model access, prompts, retrieval, application logic, security controls, deployment paths, and operational visibility.

That matters because a prompt demo can look successful while the surrounding system is still weak. Teams still need to answer questions about who can access models, where source data comes from, how unsafe outputs are handled, and how model failures affect the application.

In simple terms:

> AI on AWS becomes cloud engineering the moment it has users, data, permissions, and operational consequences.

## What Problem It Solves

This page helps teams connect Bedrock, retrieval, tools, guardrails, runtimes, and surrounding AWS services into one operating model instead of treating AI as a separate experimental track.

## How It Is Commonly Used

On AWS, AI systems commonly combine:

- Amazon Bedrock for managed model access,
- Knowledge Bases for retrieval over approved content,
- Agents for Amazon Bedrock for tool use or multi-step orchestration,
- Guardrails for Amazon Bedrock for policy checks,
- application runtimes such as App Runner, Lambda, or API Gateway,
- supporting services for secrets, logging, and access control.

## When to Use It

- Use it when planning how Bedrock-based applications fit into existing cloud systems.
- Use it when a workload needs model access, retrieval, tools, or multi-step agent behavior.
- Use it when you need a clear operating model for AI security, monitoring, and cost control.

## When Not to Use It

- Do not start with agent orchestration if a simple prompt workflow is enough.
- Do not treat model access as separate from IAM, secrets, data access, and logging design.
- Do not assume a successful prototype is already production-ready.

## Common Mistakes

- Starting with agents before the retrieval, prompt, and evaluation problems are understood.
- Giving one runtime broad access to models, tools, and data sources without clear boundaries.
- Treating safety as a model setting instead of a system design concern.
- Measuring only token latency instead of the full user-facing workflow.
- Ignoring how fast experimentation can become real spend.

## Cloud Engineering Considerations

### Identity and Access

Review which roles can call models, retrieve source data, invoke tools, and operate surrounding runtimes or pipelines. Runtime identity should be narrower than developer or platform-operator access.

### Networking

Plan how AI runtimes, data sources, and user-facing APIs communicate, especially when private connectivity or VPC controls are required.

### Security

Address prompt injection, sensitive data exposure, tool abuse, unsafe output handling, and guardrail requirements before production rollout.

### Observability

Track latency, model failures, retrieval quality, tool behavior, and user-facing outcomes so the workload is measurable end to end.

### Cost

Costs can increase quickly through token usage, retrieval, orchestration, and repeated experimentation, so budgets and usage reviews matter.

## How This Fits Into Cloud Engineering

AWS AI workloads are a strong example of why cloud engineering is broader than infrastructure provisioning. The hard part is connecting model access, runtime identity, data boundaries, observability, and deployment discipline into a system that keeps working over time.

## Related Projects

- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)

## Related Patterns

- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

## Official References

- [AWS generative AI overview](https://aws.amazon.com/generative-ai/)
- [Amazon Bedrock documentation](https://docs.aws.amazon.com/bedrock/)
