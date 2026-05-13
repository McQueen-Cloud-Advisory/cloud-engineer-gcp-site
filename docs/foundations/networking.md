# Networking

## Purpose

This page explains the networking concepts that control how traffic enters, leaves, and moves inside cloud systems.

## Core Concepts

Cloud networking decides who can reach a system, which services can talk to each other, and how traffic is routed along the way. Even when most of the infrastructure is managed, the system still depends on network design choices.

Common building blocks include virtual networks, subnets, routing tables, firewalls or security groups, load balancers, DNS, NAT, and private connectivity options. These are not separate topics in practice. They work together to shape exposure, isolation, latency, and troubleshooting complexity.

It helps to think in three traffic directions.

- Ingress is traffic entering the system from users or other systems.
- Egress is traffic leaving the system to the internet, another region, or another platform.
- East-west traffic is service-to-service communication inside the broader environment.

Each one affects security, reliability, and cost.

## How To Think About Cloud Networking

Before choosing individual services, answer a few design questions.

- Which parts of the system must be publicly reachable?
- Which components should stay private?
- How will workloads resolve each other through DNS?
- How will outbound traffic be controlled and audited?
- What happens if the system needs to connect to on-premises networks or another cloud?

Good networking design is usually boring in the best sense. Paths are deliberate, access rules are narrow, and engineers can explain how a request reaches the right backend.

## Common Mistakes

- Exposing services publicly by default because private connectivity was not planned.
- Overcomplicating subnetting and segmentation without a clear operational reason.
- Forgetting that egress traffic can be both a security and cost issue.
- Treating DNS as an afterthought instead of part of the application path.
- Debugging application failures without checking whether the network path is even valid.

## How This Fits Into Cloud Engineering

Cloud engineers use networking knowledge to expose services safely, isolate sensitive resources, troubleshoot production issues, and understand how systems behave across environments. Many incidents that appear to be application problems are really networking assumptions surfacing under load or change.

## Official References

- [Google Cloud networking overview](https://cloud.google.com/docs/networking/network-overview)
- [AWS networking and content delivery](https://aws.amazon.com/products/networking/)
- [Azure networking documentation](https://learn.microsoft.com/en-us/azure/networking/)
