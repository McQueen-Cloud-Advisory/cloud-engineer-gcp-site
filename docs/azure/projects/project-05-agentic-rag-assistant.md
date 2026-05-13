# Project 05: Agentic RAG Assistant on Azure

## Purpose

Build an agentic retrieval-augmented assistant so you practice AI application delivery, retrieval, safety controls, and platform operations on Azure.

## Scenario

Assume an internal team wants an assistant that can answer questions from approved content, possibly use tools, and remain safe enough for real organizational use. A model endpoint alone is not enough. The system needs retrieval, runtime identity, safety controls, secret management, telemetry, and a controlled application surface.

This project is useful because it forces you to treat AI as a governed Azure workload rather than a standalone model demo.

## Architecture

```text
User request
-> Azure Container Apps or API Management
-> Azure OpenAI and Foundry Agent Service
-> Azure AI Search
-> Azure AI Content Safety
-> Application Insights, Key Vault, and managed identities
```

## What You Will Build

- A user-facing application endpoint.
- A retrieval-backed assistant over approved content.
- Safety, identity, and telemetry controls around the AI path.
- A clear explanation of how retrieval, safety, and runtime identity fit together.

## Why This Architecture Works

Azure OpenAI provides model access, Foundry Agent Service supports orchestration, Azure AI Search gives retrieval over approved content, and Content Safety provides a managed moderation layer. Container Apps or API Management defines the application boundary, while Key Vault, managed identities, and Application Insights complete the security and observability model.

## Services Used

- [Azure Container Apps](../services/container-apps.md)
- [Azure API Management](../services/api-management.md)
- [Azure OpenAI](../services/azure-openai.md)
- [Foundry Agent Service](../services/foundry-agent-service.md)
- [Azure AI Search](../services/azure-ai-search.md)
- [Azure AI Content Safety](../services/content-safety.md)
- [Azure Key Vault](../services/key-vault.md)
- [Managed Identities](../services/managed-identities.md)
- [Application Insights](../services/application-insights.md)
- [Microsoft Foundry](../services/microsoft-foundry.md)

## Skills Practiced

- AI application integration
- Retrieval design
- Agent workflow planning
- AI safety and observability
- Explaining AI as a governed Azure workload

## Implementation Steps

1. Define the assistant's scope, user type, source content, and evaluation criteria.
2. Choose the user-facing runtime and API boundary for the application.
3. Configure model access and retrieval over the approved content corpus.
4. Add agent behavior only where it adds clear value over a simpler workflow.
5. Apply managed identities, Key Vault, Content Safety, and telemetry before calling the system ready.
6. Document how prompts, retrieval, safety controls, and operational ownership fit together.

## Security and Operations Considerations

Review prompt injection, document governance, runtime identities, secret storage, and content safety checks as part of the initial implementation. The workload becomes credible when you can explain what happens if retrieval is poor, the model fails, or safety checks block or flag a result.

## Cost Considerations

Model usage, search, safety checks, and container runtime costs can increase quickly, so keep testing scope controlled and monitor usage patterns early.

## How to Extend This Project

- Add user authentication and role-based access.
- Add evaluation and feedback capture.
- Add CI/CD and usage dashboards.
- Add more explicit tool-calling or workflow automation with narrow permissions.

## Portfolio Value

This project shows that you can frame AI work as a cloud engineering system with security, identity, observability, and deployment concerns rather than as a simple prompt demo.

## Official References

- [Azure OpenAI documentation](https://learn.microsoft.com/en-us/azure/ai-services/openai/)
- [Azure AI Search documentation](https://learn.microsoft.com/en-us/azure/search/)
- [Azure AI Content Safety documentation](https://learn.microsoft.com/en-us/azure/ai-services/content-safety/)
- [Azure Container Apps documentation](https://learn.microsoft.com/en-us/azure/container-apps/)
- [Azure AI Foundry documentation](https://learn.microsoft.com/en-us/azure/ai-foundry/)
