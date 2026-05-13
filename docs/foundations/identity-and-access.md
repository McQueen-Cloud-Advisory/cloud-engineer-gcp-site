# Identity and Access

## Purpose

This page explains how identities, roles, policies, and trust relationships control access in cloud environments and why identity design is one of the most important architecture decisions.

## Core Concepts

Identity answers two different questions. Authentication asks who or what is making a request. Authorization asks what that identity is allowed to do. Cloud systems need both.

There are also different identity types.

- Human identities are for administrators, developers, operators, and analysts.
- Workload identities are for applications, services, functions, jobs, or containers.
- External identities may come from partners, customers, or federated enterprise directories.

Good cloud design keeps these separate. Humans should not share long-lived admin credentials with workloads, and workloads should not carry broad permissions just because it is convenient during early development.

## What Good Access Design Looks Like

- Grant permissions to groups or roles where possible instead of managing users one by one.
- Prefer short-lived credentials or managed identity paths over static keys.
- Scope permissions to the smallest useful boundary.
- Separate deployment permissions from runtime permissions.
- Make privileged access reviewable and auditable.

Identity design also shapes incident response. If permissions are too broad or unclear, a compromise spreads faster and investigation becomes harder.

## Common Mistakes

- Using shared accounts or embedding keys directly in code or configuration.
- Giving workloads broad administrative permissions they do not need.
- Treating identity as separate from networking, pipelines, and data access.
- Forgetting to review who can assume roles, impersonate service accounts, or manage secrets.
- Building manually around exceptions instead of designing a clean access model early.

## How This Fits Into Cloud Engineering

Identity and access management sits underneath every other cloud concern. Network boundaries, data protection, CI/CD, and application runtimes all depend on a trustworthy permission model. A cloud engineer who understands identity well can reduce blast radius, simplify audits, and make systems safer to operate.

## Official References

- [Google Cloud IAM overview](https://cloud.google.com/iam/docs/overview)
- [AWS IAM documentation](https://docs.aws.amazon.com/iam/)
- [Azure role-based access control overview](https://learn.microsoft.com/en-us/azure/role-based-access-control/overview)
