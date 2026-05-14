# Azure Container Registry

## Purpose

Azure Container Registry is Azure's managed registry for storing, organizing, scanning, and distributing container images and related OCI artifacts.

It is used when Azure workloads rely on container delivery and the team needs a controlled artifact store rather than ad hoc image publishing.

## Definition

Azure Container Registry is Azure's managed registry for storing, organizing, scanning, and distributing container images and related OCI artifacts.

It matters because container platforms are only as reliable as the image supply chain behind them. A registry is where teams decide what gets published, what should be retained, who can deploy it, and how runtimes obtain approved versions.

Container Registry is not a runtime. It stores and distributes the container artifact, but it does not execute the application. That boundary matters because teams often focus on the hosting platform while under-designing the artifact path that feeds it.

In simple terms:

> Container Registry is the managed storage and distribution layer for container images in Azure.

## What Problem It Solves

Container Registry solves the problem of needing a dependable place to store and publish container artifacts as part of Azure delivery workflows.

It gives teams a controlled registry for CI/CD pipelines, image retention, and runtime pulls so deployments can be more reproducible and auditable.

That does not remove engineering responsibility. Engineers still need tagging discipline, image review, access scoping, and lifecycle cleanup.

## How It Is Commonly Used

Azure Container Registry is commonly used as the artifact source for Container Apps and other Azure container runtimes. In a typical delivery flow, a CI pipeline builds an image, tags it with a version or commit identifier, pushes it to a registry repository, and a runtime pulls that image during deployment.

This service is also useful when a team wants stronger control over repository organization, lifecycle cleanup, and the boundary between build permissions and runtime pull permissions.

## Foundational Concepts Connected to Container Registry

Container Registry connects directly to several cloud engineering foundations.

### Containers and Artifact Management

Containerized systems depend on a reliable artifact path. The registry is part of the delivery system, not just passive storage.

### Security and Supply Chain

Image provenance, base-image hygiene, and controlled publishing all influence the risk profile before the application even starts.

### Identity and Access

The identity that publishes images should be distinct from the identity that pulls and runs them. Registry access is part of deployment security.

### Operations

When releases fail, teams need to know which image was published, which tag was deployed, and whether the runtime could retrieve the intended artifact.

### Cost Management

Retained images, geo-replication, and uncontrolled repositories create waste unless lifecycle policies are deliberate.

## When to Use It

Use Container Registry when your Azure architecture includes containerized workloads and you want a managed registry integrated with Azure identities and runtimes.

Good use cases include:

- storing images for Container Apps,
- supporting CI/CD container delivery,
- separating build permissions from runtime pull permissions,
- scanning and retaining approved artifacts.

Container Registry is strongest when the artifact path is treated as a real part of the production system.

## When Not to Use It

Do not treat Container Registry as the whole container strategy. It solves artifact storage and distribution, not runtime hosting or application operations.

Do not choose a container workflow just because a registry makes image storage convenient.

Do not assume managed image storage replaces patching, dependency review, or release controls.

## Compare To

### Container Registry vs. Container Apps

Container Registry stores the image.

Container Apps runs the application built from that image. Many Azure systems use both together.

### Container Registry vs. Blob Storage

Blob Storage can hold generic files and artifacts.

Container Registry is the specialized artifact store for container images and OCI-compatible delivery workflows.

## Tradeoffs

Container Registry's biggest advantage is integration with Azure container delivery workflows.

The tradeoff is that a registry only helps when the surrounding image workflow is disciplined. Bad tagging, weak release control, or poor image hygiene still produce unreliable deployments.

Container Registry also makes it easy to retain many image versions. That is useful for traceability, but without lifecycle cleanup it becomes clutter and cost.

Another tradeoff is that image scanning and registry controls support supply-chain health, but they do not prove the application is secure or well designed.

## Common Mistakes

- Treating the registry like passive storage instead of part of the release process.
- Allowing many users or pipelines to push directly to production repositories.
- Keeping old images indefinitely without retention rules.
- Assuming image scanning removes the need to review base images and dependency choices.
- Losing track of which image digest actually backs a production deployment.

## Cloud Engineering Considerations

### Identity and Access

Control who can push images, who can pull them, and which runtimes are allowed to use production repositories. Prefer a model where build systems publish images and workloads only pull the specific repositories they need.

### Networking

Review how build agents and runtimes reach the registry, especially in private network designs. Registry access problems often surface as failed deployments even when the application itself is healthy.

### Security

Treat image provenance, base image selection, and vulnerability scanning as part of the deployment design. A registry stores artifacts, but it does not automatically make those artifacts safe.

### Operations and Observability

Track publishing activity, pull failures, and image retention so the container supply chain stays visible. Teams often focus heavily on runtime logs and forget that deployment failures frequently begin at the image layer.

### Reliability

Reliable container delivery depends on consistent image promotion, digest-level clarity, and confidence that runtimes can retrieve the intended artifact when deployments happen.

### Cost

Storage, geo-replication, and retained image volume drive cost. Lifecycle policies matter because stale tags and unreferenced images accumulate quietly.

## Project and Pattern Connections

Azure Container Registry is most directly connected to:

- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)
- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

It matters when Azure projects move from code-only deployment into container-based delivery with a real artifact-management boundary.

## Official References

- [Azure Container Registry documentation](https://learn.microsoft.com/en-us/azure/container-registry/)
- [Azure Container Registry overview](https://learn.microsoft.com/en-us/azure/container-registry/container-registry-intro)
- [Container Registry authentication overview](https://learn.microsoft.com/en-us/azure/container-registry/container-registry-authentication)
