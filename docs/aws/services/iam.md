# AWS Identity and Access Management (IAM)

## Purpose

IAM is the core access-control system for AWS. It determines who can sign in or assume access, what actions they can perform, and under what conditions those actions are allowed.

## Definition

AWS Identity and Access Management (IAM) is the policy and role system that controls access to AWS accounts, services, and resources. It works through a small set of building blocks: identities, policies, roles, and trust relationships.

The important point is that IAM is not just a list of permissions. It is the model that separates human access, workload access, and administrative delegation. In practice, that means engineers, deployment pipelines, Lambda functions, EC2 instances, and third-party integrations should rarely all share the same privilege path.

In simple terms:

> IAM answers two questions for every AWS action: who is acting, and are they allowed to do it?

## What Problem It Solves

Without IAM discipline, AWS environments drift toward shared credentials, broad administrator access, and unclear ownership. IAM gives teams a way to use temporary credentials, least-privilege roles, and explicit delegation instead of relying on a few oversized access paths.

## How It Is Commonly Used

IAM commonly appears in work such as:

- federated human access through an identity provider or IAM Identity Center,
- execution roles for services such as Lambda, ECS, and EC2,
- deployment roles for CI/CD pipelines,
- cross-account access for shared platform, logging, or security functions,
- resource and trust policies that allow one service to act on another service safely.

## When to Use It

- Use it any time a user, service, or automation path needs access to AWS resources.
- Use roles and short-lived credentials for workloads and deployment systems.
- Use it to separate deployment permissions from runtime permissions.
- Use it to control cross-account access explicitly instead of sharing credentials.

## When Not to Use It

- Do not rely on long-lived access keys for normal workloads when roles are available.
- Do not treat IAM as the same thing as end-user application authentication. Your application's customer identity model is often separate.
- Do not use broad administrator policies as a substitute for clear access design.

## Common Mistakes

- Giving many engineers or workloads `AdministratorAccess` because it is faster in the short term.
- Reusing one role across multiple applications with different access needs.
- Reviewing permission policies but ignoring trust policies, which decide who can assume a role.
- Leaving old access keys, unused roles, or stale federation mappings in place.
- Treating root account credentials as a normal operational access path.

## Cloud Engineering Considerations

### Identity and Access

Good IAM design starts by distinguishing human users, deployers, and runtime identities. Those paths should have different permissions, approval expectations, and monitoring.

### Networking

IAM is a control-plane service, but network design still affects access patterns. VPC endpoints, source IP conditions, and cross-account boundaries can change how safely identities reach resources.

### Security

Use MFA for human access, prefer federation and temporary credentials, and review sensitive permissions such as `iam:PassRole`, KMS administration, and organization-wide actions carefully.

### Observability

CloudTrail, Access Analyzer, and authorization failure investigation matter because access problems are often operational problems first and security problems second.

### Cost

IAM itself is not usually a direct cost driver, but weak access design often creates cost risk elsewhere by allowing oversized resources, unmanaged experiments, or accidental data exposure.

## How This Fits Into Cloud Engineering

AWS systems are only as safe and operable as their identity model. If you cannot explain who deploys a system, what identity runs it, and how privileged access is controlled, the system is not well engineered yet.

## Related Projects

- [Project 01: Static Site](../projects/project-01-static-site.md)
- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)
- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)

## Related Patterns

- [Static Site](../patterns/static-site.md)
- [Serverless API](../patterns/serverless-api.md)

## Official References

- [AWS IAM documentation](https://docs.aws.amazon.com/iam/)
- [IAM best practices](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html)
