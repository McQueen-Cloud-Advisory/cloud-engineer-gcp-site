# AWS Identity and Access Management (IAM)

## Purpose

IAM is the core access-control system for AWS. It determines who can sign in or assume access, what actions they can perform, and under what conditions those actions are allowed.

## Definition

AWS Identity and Access Management (IAM) is the policy and role system that controls access to AWS accounts, services, and resources. It works through a small set of building blocks: identities, policies, roles, and trust relationships.

The important point is that IAM is not just a list of permissions. It is the model that separates human access, workload access, and administrative delegation. In practice, that means engineers, deployment pipelines, Lambda functions, EC2 instances, and third-party integrations should rarely all share the same privilege path.

In simple terms:

> IAM answers two questions for every AWS action: who is acting, and are they allowed to do it?

That sounds administrative, but it is actually architectural. IAM shapes how deployments happen, how applications reach other services, and how teams control blast radius when something goes wrong.

## What Problem It Solves

Without IAM discipline, AWS environments drift toward shared credentials, broad administrator access, and unclear ownership. IAM gives teams a way to use temporary credentials, least-privilege roles, and explicit delegation instead of relying on a few oversized access paths.

It solves the problem of needing a scalable trust model for humans, workloads, and automation.

That does not make access design easy. Engineers still need to reason about trust boundaries, role assumption paths, sensitive permissions, and how real workflows map to least privilege.

## How It Is Commonly Used

IAM commonly appears in work such as:

- federated human access through an identity provider or IAM Identity Center,
- execution roles for services such as Lambda, ECS, and EC2,
- deployment roles for CI/CD pipelines,
- cross-account access for shared platform, logging, or security functions,
- resource and trust policies that allow one service to act on another service safely.

IAM appears in nearly every AWS design because every runtime, deployment system, and operator needs an identity path.

## Foundational Concepts Connected to IAM

IAM connects directly to several cloud engineering foundations.

### Identity and Access

IAM is the main access-control foundation in AWS. It separates human access, workload identities, and delegated administrative paths.

### Security Architecture

Least privilege, temporary credentials, and role trust are security architecture decisions, not only console settings.

### Platform Engineering

Deployment systems, shared services, and cross-account architectures depend on IAM being understandable and consistent.

### Observability and Audit

Access changes and authorization failures need to be visible because identity problems often surface first as operational issues.

### Governance

IAM sits close to account structure, separation of duties, and how teams prove that access is intentional rather than accidental.

## When to Use It

- Use it any time a user, service, or automation path needs access to AWS resources.
- Use roles and short-lived credentials for workloads and deployment systems.
- Use it to separate deployment permissions from runtime permissions.
- Use it to control cross-account access explicitly instead of sharing credentials.

IAM is strongest when it is treated as a design system for trust, not as a growing pile of one-off permission documents.

## When Not to Use It

- Do not rely on long-lived access keys for normal workloads when roles are available.
- Do not treat IAM as the same thing as end-user application authentication. Your application's customer identity model is often separate.
- Do not use broad administrator policies as a substitute for clear access design.

## Compare To

### IAM vs. Application Authentication

IAM controls access to AWS resources and AWS control-plane or runtime actions.

Application authentication controls how your users sign in to your own application. These are related security concerns, but they are not the same system.

### IAM Roles vs. Long-Lived Access Keys

IAM roles provide temporary credentials and clearer trust boundaries.

Long-lived access keys are harder to manage safely and should be minimized where roles are possible.

## Tradeoffs

IAM's biggest advantage is precision. It gives AWS teams the tools to create explicit, least-privilege access paths for humans, services, and automation.

The tradeoff is complexity. Fine-grained access control can become hard to reason about when policies, trust relationships, and condition logic accumulate without structure.

IAM also makes delegation safer when designed well. That is valuable, but weak role reuse and oversized policies quickly undermine the benefits.

Another tradeoff is that access models age. IAM designs that start clean often drift unless they are reviewed as systems evolve.

## Common Mistakes

- Giving many engineers or workloads `AdministratorAccess` because it is faster in the short term.
- Reusing one role across multiple applications with different access needs.
- Reviewing permission policies but ignoring trust policies, which decide who can assume a role.
- Leaving old access keys, unused roles, or stale federation mappings in place.
- Treating root account credentials as a normal operational access path.
- Using one powerful shared role because it seems easier than modeling separate service responsibilities.

## Cloud Engineering Considerations

### Identity and Access

Good IAM design starts by distinguishing human users, deployers, and runtime identities. Those paths should have different permissions, approval expectations, and monitoring.

### Networking

IAM is a control-plane service, but network design still affects access patterns. VPC endpoints, source IP conditions, and cross-account boundaries can change how safely identities reach resources.

### Security

Use MFA for human access, prefer federation and temporary credentials, and review sensitive permissions such as `iam:PassRole`, KMS administration, and organization-wide actions carefully.

### Observability

CloudTrail, Access Analyzer, and authorization failure investigation matter because access problems are often operational problems first and security problems second.

### Reliability

Reliable AWS operations depend on identities working predictably. Deployments fail, services break, and incident response slows down when the trust model is unclear or fragile.

### Cost

IAM itself is not usually a direct cost driver, but weak access design often creates cost risk elsewhere by allowing oversized resources, unmanaged experiments, or accidental data exposure.

## Project and Pattern Connections

IAM is foundational across the AWS learning path and is most directly connected to:

- [Project 01: Static Site](../projects/project-01-static-site.md)
- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)
- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)
- [Static Site](../patterns/static-site.md)
- [Serverless API](../patterns/serverless-api.md)

It matters in every AWS project because every deployment path and every runtime permission set depends on it.

## Official References

- [AWS IAM documentation](https://docs.aws.amazon.com/iam/)
- [IAM best practices](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html)
- [IAM roles guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html)
