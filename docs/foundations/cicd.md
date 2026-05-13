# CI/CD

## Purpose

This page explains continuous integration and continuous delivery as the repeatable automation path for validating changes and moving them into real environments safely.

## Core Concepts

Continuous integration is the discipline of frequently combining changes into a shared codebase and validating them quickly. Continuous delivery is the discipline of making those validated changes deployable in a controlled, repeatable way. Together, they reduce the risk that deployments become rare, manual, and unpredictable.

A useful pipeline does more than run unit tests. It often includes linting, type checks, security checks, infrastructure validation, artifact creation, and deployment steps. The exact stages vary by team, but the core idea stays the same: each change should move through the same path with clear evidence of what passed, what failed, and what was released.

## What Good Pipelines Do

- Build the same artifact once and promote it through environments when possible.
- Keep secrets out of the repo and inject them through controlled platform features.
- Make deployment steps observable so engineers can trace who changed what and when.
- Support rollback, redeployment, or forward-fix workflows without improvisation.
- Apply the same discipline to infrastructure changes as to application changes.

Cloud engineering work often involves multiple change types at once: application code, configuration, infrastructure definitions, policies, and secrets references. A good CI/CD design handles those pieces coherently instead of treating deployment as a final manual step.

## Release Strategy Matters

CI/CD is also about how change reaches production. Teams may choose direct deployment, staged promotion, blue-green releases, canary rollouts, or feature flags depending on the risk profile of the system. The right strategy depends on blast radius, rollback speed, user impact, and how well the team can observe the change after release.

## Common Mistakes

- Giving pipelines broad admin access when narrower deployment permissions would work.
- Rebuilding artifacts separately for each environment instead of promoting the same output.
- Skipping infrastructure validation and discovering deployment issues only in production.
- Storing secrets in workflow files, repo variables, or application config without proper controls.
- Treating the pipeline as a black box with poor logs, weak failure messages, or no rollback plan.

## How This Fits Into Cloud Engineering

Cloud engineers use CI/CD to make cloud change routine instead of risky. A reliable pipeline reduces manual drift, supports auditability, and gives teams confidence that infrastructure and application updates can be repeated safely. In practice, CI/CD is one of the main ways cloud systems become operable at scale.

## Official References

- [GitHub Actions documentation](https://docs.github.com/en/actions)
- [Google Cloud CI/CD overview](https://cloud.google.com/architecture/devops/devops-tech-continuous-integration-delivery-deployment)
- [Azure DevOps documentation](https://learn.microsoft.com/en-us/azure/devops/)
