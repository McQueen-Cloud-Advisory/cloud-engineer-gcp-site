# Compute

## Purpose

This page explains the main cloud compute models and how to choose between them based on control, scale behavior, operational ownership, and workload shape.

## Core Concepts

Compute is the part of the platform that runs your code. In cloud environments, that can mean virtual machines, containers, serverless functions, managed application runtimes, or scheduled jobs. Each model makes a different tradeoff between flexibility and operational burden.

- Virtual machines give the most control over the operating system and runtime, but require more patching, hardening, and capacity management.
- Containers package the application more consistently and can run on managed container platforms or full orchestrators.
- Serverless functions reduce host management and scale well for event-driven workloads, but impose runtime constraints and encourage stateless design.
- Managed application runtimes sit between raw infrastructure and fully serverless execution by handling more of the platform while still supporting longer-lived services.

No compute model is universally best. The goal is to match the platform to the workload rather than forcing the workload into the newest service.

## How To Choose a Compute Model

Ask a few practical questions first.

- Does the workload need full operating system control or a custom runtime?
- Is traffic steady, spiky, event-driven, or batch-oriented?
- How much startup latency can the application tolerate?
- Does the team want to manage patching, autoscaling, and host-level troubleshooting?
- Where does state live, and can the workload be treated as disposable?

A public web API may fit well on containers or a managed runtime. A scheduled data pull may fit a function platform. A legacy application with operating system dependencies may still need virtual machines. Good compute choices are based on workload reality, not preference alone.

## Common Mistakes

- Choosing compute based on hype instead of runtime and operational requirements.
- Treating stateful workloads as if they were easy to move between models without redesign.
- Ignoring cold start behavior, concurrency limits, or connection management.
- Overlooking the surrounding needs for secrets, networking, logging, and rollout strategy.
- Assuming less infrastructure visibility means less operational responsibility.

## How This Fits Into Cloud Engineering

Cloud engineers select compute not just for execution, but for operability. The compute choice influences deployment design, identity handling, autoscaling behavior, patching, troubleshooting, and cost. In real systems, the right answer is usually the model that the team can operate well over time.

## Official References

- [Google Cloud compute options](https://cloud.google.com/docs/compute)
- [AWS compute services](https://aws.amazon.com/products/compute/)
- [Azure compute services](https://learn.microsoft.com/en-us/azure/architecture/guide/technology-choices/compute-decision-tree)
