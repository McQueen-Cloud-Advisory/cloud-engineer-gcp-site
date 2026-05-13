# Azure Container Registry

## Definition

Azure Container Registry is Azure's managed registry for storing, organizing, scanning, and distributing container images and related OCI artifacts.

It matters because container platforms are only as reliable as the image supply chain behind them. A registry is where teams decide what gets published, what should be retained, who can deploy it, and how runtimes obtain approved versions.

## How It Is Commonly Used

Azure Container Registry is commonly used as the artifact source for Container Apps and other Azure container runtimes. In a typical delivery flow, a CI pipeline builds an image, tags it with a version or commit identifier, pushes it to a registry repository, and a runtime pulls that image during deployment.

This service is also useful when a team wants stronger control over repository organization, lifecycle cleanup, and the boundary between build permissions and runtime pull permissions.

## What To Pay Attention To

### Identity and Access

Control who can push images, who can pull them, and which runtimes are allowed to use production repositories. Prefer a model where build systems publish images and workloads only pull the specific repositories they need.

### Networking

Review how build agents and runtimes reach the registry, especially in private network designs. Registry access problems often surface as failed deployments even when the application itself is healthy.

### Security

Treat image provenance, base image selection, and vulnerability scanning as part of the deployment design. A registry stores artifacts, but it does not automatically make those artifacts safe.

### Operations and Observability

Track publishing activity, pull failures, and image retention so the container supply chain stays visible. Teams often focus heavily on runtime logs and forget that deployment failures frequently begin at the image layer.

### Cost

Storage, geo-replication, and retained image volume drive cost. Lifecycle policies matter because stale tags and unreferenced images accumulate quietly.

## Common Mistakes

- Treating the registry like passive storage instead of part of the release process.
- Allowing many users or pipelines to push directly to production repositories.
- Keeping old images indefinitely without retention rules.
- Assuming image scanning removes the need to review base images and dependency choices.

## How This Fits Into Cloud Engineering

Container registries sit between software delivery and platform operations. Cloud engineers need to understand that boundary because deployment reliability, image provenance, runtime permissions, and network design all meet here.

## Related Projects

- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)

## Related Patterns

- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

## Official References

- [Azure Container Registry documentation](https://learn.microsoft.com/en-us/azure/container-registry/)
- [Azure Container Registry overview](https://learn.microsoft.com/en-us/azure/container-registry/container-registry-intro)
