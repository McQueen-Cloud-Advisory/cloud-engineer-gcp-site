# AWS Roadmap

## Purpose

This roadmap outlines a practical sequence for building AWS cloud engineering skills without getting lost in the size of the AWS service catalog.

## Stage 1: Learn Accounts, Identity, And Regions

- Read the AWS overview and getting started page first.
- Understand accounts, IAM roles, policies, billing visibility, and regional thinking.
- Make sure you can explain the difference between human access and workload access.

## Stage 2: Build The Core Application Platform Path

Study the services that support a small but realistic application path.

- IAM
- S3
- Lambda
- API Gateway
- DynamoDB
- Secrets Manager
- CloudWatch

At this stage, the goal is to understand how AWS handles public entry points, serverless compute, runtime permissions, application data, and visibility.

## Stage 3: Add Eventing And Scheduled Work

Once the core application path is comfortable, add EventBridge and related scheduled patterns.

- Build recurring automation.
- Learn how events trigger downstream processing.
- Pay attention to retry behavior, idempotency, and monitoring.

This stage is where AWS starts to feel like a broader platform rather than a collection of point services.

## Stage 4: Add Data And Analytics

Move into analytics once you can already explain the basic AWS application path.

- Learn how object storage participates in ingestion and analytical workflows.
- Connect project work to data movement and reporting concerns.
- Watch cost and data-transfer assumptions more closely.

## Stage 5: Add AI And Agentic Workloads

Use Bedrock-oriented pages after the core path is stable.

- Focus on retrieval, safety, identity, and monitoring around AI services.
- Treat AI as a cloud system, not as a separate topic outside operations.
- Reuse the same engineering habits you used for the earlier projects.

## What Success Looks Like

By the end of this roadmap, you should be able to do more than name AWS services. You should be able to explain how an AWS system is structured, how permissions are granted, how requests are processed, how failures are observed, and why the architecture uses the services it does.

## How This Fits Into Cloud Engineering

A roadmap keeps AWS learning connected to implementation. It helps you decide what to study next and makes it easier to explain your progress in project, portfolio, and interview settings.

## Official References

- [AWS Architecture Center](https://aws.amazon.com/architecture/)
- [AWS Well-Architected Framework](https://docs.aws.amazon.com/wellarchitected/latest/framework/welcome.html)
- [Amazon Bedrock documentation](https://docs.aws.amazon.com/bedrock/)
