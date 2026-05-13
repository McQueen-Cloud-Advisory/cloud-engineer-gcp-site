# AWS Getting Started

## Purpose

This page shows how to begin learning AWS through a practical cloud engineering path instead of a broad, unstructured tour of the AWS catalog.

## Understand The AWS Operating Model First

- Accounts are the main environment and billing boundary.
- IAM users, roles, and policies control how people and workloads interact with AWS.
- Regions matter because service placement, latency, and cost are tied to them.
- The root account is not your normal working identity and should be treated carefully.

If these basics are unclear, the rest of AWS will feel much harder than it needs to.

## Recommended First Service Set

Start with a compact group of services that teaches the AWS delivery model well.

- IAM for access control.
- S3 for object storage and static delivery.
- Lambda for event-driven compute.
- API Gateway for HTTP entry points.
- DynamoDB for application data.
- Secrets Manager for runtime secrets.
- CloudWatch for logs, metrics, and alerts.
- EventBridge for scheduled and event-driven workflows.

These services support several of the early projects and cover a large part of the everyday AWS operating model.

## Suggested First Steps

1. Set up a sandbox account or environment and review billing visibility before building anything expensive.
2. Learn how IAM roles, policies, and least privilege shape both human and workload access.
3. Pick a region deliberately and keep your early experiments consistent there.
4. Work through the storage, compute, API, data, and monitoring pages that support the first projects.
5. Add scheduled processing, analytics, and Bedrock topics only after the core application path feels comfortable.

## What To Delay Until Later

Do not try to learn every AWS service category at the start. You can safely defer some topics until the base platform path is clear.

- Broad multi-account governance design beyond the basics.
- Advanced networking patterns.
- Large-scale analytics design.
- Bedrock and agentic workloads.

Delaying these is not avoidance. It is sequencing. You learn them better once IAM, storage, compute, observability, and cost are already familiar.

## Common Mistakes

- Using the root account for day-to-day work.
- Creating resources in several regions without realizing it.
- Granting broad IAM permissions because it is faster during setup.
- Treating the AWS catalog as a checklist instead of a workload-driven design space.

## How This Fits Into Cloud Engineering

Use this page to identify the first AWS concepts and services you need before moving into project work. The goal is to make AWS feel operable and explainable, not overwhelming.

## Official References

- [AWS getting started](https://aws.amazon.com/getting-started/)
- [AWS global infrastructure](https://aws.amazon.com/about-aws/global-infrastructure/)
- [AWS Well-Architected Framework](https://docs.aws.amazon.com/wellarchitected/latest/framework/welcome.html)
