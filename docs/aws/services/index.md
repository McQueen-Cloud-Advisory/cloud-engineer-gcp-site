# AWS Services

## Purpose

This section explains AWS services from a practical cloud engineering perspective and keeps the focus on the services that appear in the learning path, projects, and reusable patterns.

## How To Read This Section

These pages are not meant to be a complete AWS catalog. They are meant to help you answer a smaller set of questions.

- What problem does this service solve?
- When is it a good fit for an AWS design?
- What identity, networking, and observability concerns come with it?
- How does it connect to the rest of a real workload?

That is why the service pages work best when you read them in support of a project or pattern instead of trying to memorize them all in isolation.

## What This Service Set Is Designed To Teach

The AWS path emphasizes the services that help learners understand public delivery, serverless compute, object storage, application data, event-driven workflows, monitoring, and AI platform extension.

## Service Categories

### Identity and Access

- [IAM](iam.md)

### Storage and Data

- [S3](s3.md)
- [DynamoDB](dynamodb.md)

### Compute and Application Hosting

- [Lambda](lambda.md)
- [App Runner](app-runner.md)
- [Elastic Container Registry (ECR)](ecr.md)

### Integration and Scheduling

- [API Gateway](api-gateway.md)
- [EventBridge](eventbridge.md)

### Observability and Operations

- [CloudWatch](cloudwatch.md)

### AI and Agentic Platforms

- [AI and Agentic Workloads](ai-and-agents.md)
- [Amazon Bedrock](bedrock.md)
- [Agents for Amazon Bedrock](bedrock-agents.md)
- [Knowledge Bases for Amazon Bedrock](knowledge-bases-for-bedrock.md)
- [Guardrails for Amazon Bedrock](guardrails-for-bedrock.md)

### Security and Secrets

- [Secrets Manager](secrets-manager.md)

## How The Categories Connect

In practice, AWS systems are assembled across categories. IAM shapes who can deploy or run the workload. API Gateway or EventBridge introduces traffic or events. Lambda or App Runner handles compute. S3 or DynamoDB stores data. CloudWatch tells you whether the whole path is healthy. Later, Bedrock services extend that same platform path into AI workloads instead of replacing normal engineering concerns.

## How This Fits Into Cloud Engineering

Cloud engineers use service knowledge to make architecture choices, not to build flash cards. This section is designed to help you explain why a service belongs in a system, what responsibility comes with it, and how it affects security, delivery, and operations.

## Official References

- [AWS documentation](https://docs.aws.amazon.com/)
- [AWS Well-Architected Framework](https://docs.aws.amazon.com/wellarchitected/latest/framework/welcome.html)
