# Infrastructure as Code

## Purpose

This page explains infrastructure as code as the practice of declaring cloud resources in versioned definitions so environments can be reviewed, reproduced, and changed safely.

## Core Concepts

Infrastructure as code replaces manual, click-driven environment changes with files that describe the desired state of the system. Those definitions can be reviewed in pull requests, tested in pipelines, and reapplied consistently across environments.

This matters because cloud environments change constantly. Networks, identities, storage, compute, policies, and DNS are all part of the system. If those changes happen manually, drift appears quickly and the true environment becomes hard to understand.

Useful IaC practices usually include:

- Keeping infrastructure definitions in version control.
- Using reusable modules or templates for common patterns.
- Reviewing proposed changes before apply or deploy.
- Managing state carefully when the tool requires it.
- Separating environments intentionally rather than copying definitions ad hoc.

## Operational Discipline Still Matters

IaC does not remove risk by itself. Teams still need good access control, plan review, change sequencing, and rollback strategy. They also need to decide how secrets are referenced, how shared modules are versioned, and how drift will be detected when manual changes do happen.

## Common Mistakes

- Treating IaC as a one-time provisioning tool instead of the ongoing source of truth.
- Allowing manual changes in production without reconciling them back into code.
- Using highly coupled templates that are hard to review or reuse.
- Ignoring state protection, remote state access, or locking requirements.
- Granting broad apply permissions to pipelines or users without proper review.

## How This Fits Into Cloud Engineering

Cloud engineering depends on repeatability. Infrastructure as code makes environment change part of normal engineering practice instead of an undocumented operational side channel. It is one of the main ways teams reduce drift, improve auditability, and make cloud systems easier to evolve.

## Official References

- [Terraform introduction](https://developer.hashicorp.com/terraform/intro)
- [AWS CloudFormation documentation](https://docs.aws.amazon.com/cloudformation/)
- [Azure Bicep overview](https://learn.microsoft.com/en-us/azure/azure-resource-manager/bicep/overview)
