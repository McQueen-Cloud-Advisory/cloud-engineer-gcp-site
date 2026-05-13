# Project 05: Agentic RAG Assistant on AWS

## Purpose

Build an agentic retrieval-augmented assistant so you practice AI runtime integration, retrieval, guardrails, and operational governance in one workload.

## Scenario

Assume a team wants an internal assistant that can answer questions from approved documents, possibly use tools, and behave safely enough for real users. A simple prompt demo is not enough. The system needs model access, retrieval over governed content, runtime identity, safety controls, secrets, and monitoring.

This project is useful because it forces you to treat AI as a real application system instead of as a standalone model experiment.

## Architecture

```text
User request
-> AWS App Runner or Amazon API Gateway
-> Amazon Bedrock or Agents for Amazon Bedrock
-> Knowledge Bases for Amazon Bedrock
-> Guardrails for Amazon Bedrock
-> Amazon CloudWatch and AWS Secrets Manager
```

## What You Will Build

- A user-facing application endpoint.
- A grounded AI workflow with retrieval over approved content.
- Guardrails, secret handling, and operational visibility for the AI path.
- A documented explanation of model access, retrieval, and safety boundaries.

## Why This Architecture Works

Bedrock provides model access, Knowledge Bases adds managed retrieval, and Bedrock Agents can handle multi-step behavior when needed. Guardrails gives the system a managed safety checkpoint, while App Runner or API Gateway provides a controlled entry path. Secrets Manager and CloudWatch complete the security and observability story.

## Services Used

- [AWS App Runner](../services/app-runner.md)
- [Amazon API Gateway](../services/api-gateway.md)
- [Amazon Bedrock](../services/bedrock.md)
- [Agents for Amazon Bedrock](../services/bedrock-agents.md)
- [Knowledge Bases for Amazon Bedrock](../services/knowledge-bases-for-bedrock.md)
- [Guardrails for Amazon Bedrock](../services/guardrails-for-bedrock.md)
- [AWS Secrets Manager](../services/secrets-manager.md)
- [Amazon CloudWatch](../services/cloudwatch.md)

## Skills Practiced

- AI application integration
- Managed retrieval design
- Agent workflow planning
- AI safety and observability
- Explaining AI workloads as cloud systems

## Implementation Steps

1. Define the assistant's scope, users, approved knowledge corpus, and success criteria.
2. Choose the application entry point and runtime model for the user-facing experience.
3. Configure Bedrock model access and connect retrieval over the approved content set.
4. Add agent behavior only if the use case truly needs tools or multi-step execution.
5. Apply guardrails, secrets handling, logging, and monitoring before calling the system production-ready.
6. Document how prompts, retrieval, tool use, safety controls, and failures are handled end to end.

## Security and Operations Considerations

Review who can access the application, what content is indexed, how secrets are stored, and how the system should respond to unsafe or adversarial input. The important operational question is not only whether the assistant answers correctly. It is whether the team can tell when retrieval quality, latency, or safety behavior has degraded.

## Cost Considerations

Model usage, retrieval, guardrails, and supporting runtime resources can add up quickly. Limit the scope of testing, define evaluation boundaries, and monitor usage patterns closely.

## How to Extend This Project

- Add user authentication and tenant separation.
- Add feedback capture and evaluation workflows.
- Add deployment automation and model usage dashboards.
- Add more explicit tool-calling workflows with narrow permissions.

## Portfolio Value

This project shows that you can treat AI workloads as engineering systems with deployment, security, observability, and governance requirements rather than as a simple prompt demo.

## Official References

- [AWS App Runner documentation](https://docs.aws.amazon.com/apprunner/)
- [Amazon Bedrock documentation](https://docs.aws.amazon.com/bedrock/)
- [Agents for Amazon Bedrock documentation](https://docs.aws.amazon.com/bedrock/latest/userguide/agents.html)
- [Knowledge Bases for Amazon Bedrock documentation](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base.html)
- [Guardrails for Amazon Bedrock documentation](https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails.html)
