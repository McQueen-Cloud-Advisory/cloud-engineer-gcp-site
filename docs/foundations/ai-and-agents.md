# AI and Agents

## Purpose

This page explains the building blocks of modern AI applications and why agentic systems should be treated as cloud systems with security, reliability, cost, and operational requirements.

## Core Concepts

AI applications are usually assembled from several separate parts. The model generates output, but the surrounding system decides which prompts are sent, what context is included, which data sources are trusted, what tools can be called, and how results are returned to users.

- A foundation model is the reasoning or generation engine, but it is only one part of the application.
- Prompting defines the task and behavior expected from the model.
- Retrieval adds external context so the model can answer with current or domain-specific information.
- Guardrails and safety controls reduce the chance of harmful, off-policy, or low-trust output.
- Evaluation measures whether the system is actually useful, not just whether the model can respond.

An agentic system goes one step further. Instead of returning a single answer, it may choose tools, call APIs, look up documents, or execute multiple steps before responding. That increases capability, but it also increases the number of failure paths. A broken search index, a weak permission model, or an unsafe tool integration can be just as serious as a poor model response.

## What Makes Production AI Difficult

The hard part of AI in cloud engineering is usually not getting a model to answer once. The hard part is operating the full system safely and predictably.

- Identity still matters because runtimes, users, and tools all need controlled access.
- Networking still matters because application services, vector stores, retrieval layers, and model endpoints must reach each other reliably.
- Observability still matters because latency, failed tool calls, unsafe prompts, and poor retrieval quality have to be visible.
- Cost still matters because model invocations, embeddings, retrieval pipelines, and attached runtimes can compound quickly.

Teams also need to distinguish between similar ideas. Retrieval-augmented generation is not the same as fine-tuning. A chatbot is not automatically an agent. A proof of concept is not the same as a production service with access control, auditability, and incident response.

## Common Mistakes

- Treating the model as the entire architecture instead of designing the surrounding application path.
- Giving the model or agent overly broad tool permissions.
- Skipping evaluation and relying on a few manual examples as proof of quality.
- Ignoring prompt injection, unsafe retrieval sources, or data leakage risks.
- Expanding scope before the team can explain cost, failure handling, and acceptable behavior.

## How This Fits Into Cloud Engineering

Cloud engineers make AI systems deployable and governable. That means selecting runtimes, controlling identities, protecting secrets, structuring retrieval paths, monitoring behavior, and defining how the system should fail when an AI component is slow, unsafe, or wrong. The value is not only model access. The value is turning AI behavior into an operable service.

## Official References

- [Google Cloud generative AI documentation](https://cloud.google.com/vertex-ai/generative-ai/docs/overview)
- [Amazon Bedrock documentation](https://docs.aws.amazon.com/bedrock/)
- [Azure AI Foundry documentation](https://learn.microsoft.com/en-us/azure/ai-foundry/what-is-azure-ai-foundry)
