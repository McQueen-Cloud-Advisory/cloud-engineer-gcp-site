# Project 05: Agentic RAG Assistant on Google Cloud

## Purpose

Build an agentic retrieval-augmented assistant so you practice AI runtime integration, retrieval, safety controls, and platform operations on Google Cloud.

## Scenario

Assume a team wants an assistant that can answer from approved content, possibly take actions, and remain safe enough for real users. A raw model call is not enough. The application needs retrieval, service-account boundaries, safety controls, secrets, monitoring, and a user-facing runtime that is easy to operate.

This project is useful because it forces you to treat AI as a Google Cloud application platform problem, not just a model selection problem.

## Architecture

```text
User request
-> Cloud Run
-> Vertex AI and Vertex AI Agent Builder
-> Model Garden and Model Armor
-> Secret Manager and Cloud Monitoring
```

## What You Will Build

- A user-facing assistant endpoint.
- A retrieval-backed AI workflow over approved content.
- Safety, secret, and monitoring controls around the AI path.
- A documented explanation of how model choice, retrieval, and safety fit together.

## Why This Architecture Works

Cloud Run gives the assistant a clean runtime boundary. Vertex AI provides model access, Agent Builder supports retrieval or higher-level orchestration, Model Garden supports informed model choice, and Model Armor adds safety controls. Secret Manager and Cloud Monitoring make the runtime more secure and operable.

## Services Used

- [Cloud Run](../services/cloud-run.md)
- [Vertex AI](../services/vertex-ai.md)
- [Vertex AI Agent Builder](../services/vertex-ai-agent-builder.md)
- [Agent Development Kit](../services/agent-development-kit.md)
- [Model Garden](../services/model-garden.md)
- [Model Armor](../services/model-armor.md)
- [Secret Manager](../services/secret-manager.md)
- [Cloud Monitoring](../services/cloud-monitoring.md)
- [IAM and Service Accounts](../services/iam-and-service-accounts.md)

## Skills Practiced

- AI application integration
- Retrieval design
- Agent workflow planning
- AI safety and observability
- Explaining AI as a Google Cloud system

## Implementation Steps

1. Define the assistant's scope, users, content boundaries, and evaluation criteria.
2. Choose the Cloud Run application boundary and the service-account model for the runtime.
3. Configure model access and retrieval over approved content using Vertex AI tooling.
4. Add agent behavior or custom orchestration only where it improves the workload clearly.
5. Apply safety controls, secrets handling, and monitoring before treating the system as operationally ready.
6. Document how the assistant handles retrieval, model failure, unsafe content, and end-user feedback.

## Security and Operations Considerations

Review prompt injection, document governance, safety controls, and service-account scope from the beginning of the build. The key operational question is whether the team can detect poor retrieval, unsafe outputs, high latency, or unexpected cost growth before the assistant becomes unreliable.

## Cost Considerations

Model usage, retrieval, safety controls, and runtime hosting can compound quickly, so limit scope and monitor usage patterns early.

## How to Extend This Project

- Add user authentication and role-based access.
- Add evaluation workflows and feedback capture.
- Add deployment automation and usage dashboards.
- Add more custom tool use with explicit, narrow permissions.

## Portfolio Value

This project shows that you can frame AI work as a cloud engineering system with deployment, identity, safety, retrieval, and observability requirements rather than as a simple prompt demo.

## Official References

- [Vertex AI documentation](https://cloud.google.com/vertex-ai/docs)
- [Agent Builder documentation](https://docs.cloud.google.com/agent-builder)
- [Model Armor overview](https://cloud.google.com/security/products/model-armor)
- [Cloud Run documentation](https://cloud.google.com/run/docs)
