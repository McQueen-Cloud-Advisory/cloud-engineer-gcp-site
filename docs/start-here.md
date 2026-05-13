# Start Here

## Purpose

This page explains how to use the site as a practical learning path so you build cloud engineering judgment instead of collecting disconnected service names.

## What This Site Is Designed To Teach

- Foundations come first because provider services make more sense when you already understand networking, identity, compute, storage, observability, and cost.
- Provider sections matter because cloud engineering is always implemented through a real platform with real defaults, limits, and operating habits.
- Projects matter because architecture decisions only become clear when you have to deploy, secure, monitor, and explain a working system.
- Patterns matter because they teach why a design is shaped a certain way, not just which services appear in it.

## Recommended First Pass

1. Read the [Foundations overview](foundations/index.md) and the core pages on networking, identity, compute, storage, observability, and cost.
2. Choose one provider path first: [Google Cloud](gcp/index.md), [AWS](aws/index.md), or [Azure](azure/index.md).
3. Read that provider's getting started page and roadmap before trying to study every service.
4. Build the project sequence in order and use service pages when a project actually needs them.
5. Use the pattern pages to explain the architecture after you can build it.
6. Use the [Career section](career/index.md) to turn the work into portfolio, resume, and interview material.

## Choose One Provider First

You do not need to learn every cloud at once. Pick one provider and go deep enough to understand its basic operating model, identity approach, deployment path, and observability surface. That depth transfers much better than shallow exposure to three platforms at the same time.

Choose a provider based on one or more of these factors.

- It matches the platform used at your current or target job.
- It supports the kind of workloads you want to build first.
- It gives you the clearest path to completing projects and explaining them well.

Once you can build and explain a few complete systems on one provider, moving to another provider becomes much easier.

## How To Use Projects And Patterns

Projects are where service knowledge becomes engineering skill. As you work through a project, focus on more than getting the deployment to succeed.

- Explain how traffic reaches the system.
- Describe which identities can deploy it and which identities run it.
- Note where secrets and data live.
- Add logs, metrics, or alerts so you can tell whether it is healthy.
- Write down the main cost drivers and tradeoffs.

Then use the pattern pages to step back and explain why the architecture works. That is what turns a small build into a reusable mental model.

## Common Mistakes

- Reading dozens of service pages before building anything.
- Switching providers too often and never learning one operating model well.
- Treating deployment as success without checking observability, permissions, and cost.
- Memorizing product names instead of explaining system behavior.

## How This Fits Into Cloud Engineering

A cloud engineer does not only provision resources. They design systems that can be deployed, secured, monitored, changed, and explained over time. This site is organized to reinforce that full operating model instead of reducing cloud learning to product recall.

## Official References

- [Google Cloud get started documentation](https://cloud.google.com/docs/get-started)
- [AWS getting started](https://aws.amazon.com/getting-started/)
- [Azure learning and training](https://learn.microsoft.com/en-us/training/azure/)
