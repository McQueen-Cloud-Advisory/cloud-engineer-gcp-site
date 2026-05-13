# Amazon Elastic Container Registry (ECR)

## Definition

Amazon ECR is AWS's managed registry for container images and OCI artifacts.

It matters because container delivery does not start at the runtime. Teams need a controlled place to publish artifacts, manage versions, scan images, and decide which workloads are allowed to pull approved images.

## How It Is Commonly Used

Amazon ECR is commonly used with containerized workloads such as App Runner or other AWS container platforms. A normal flow is simple: a pipeline builds an image, tags it, pushes it to ECR, and the runtime pulls the selected version during deployment.

This service becomes especially important when a team wants to separate build permissions from runtime permissions and keep the release process reproducible.

## What To Pay Attention To

### Identity and Access

Limit who can push and pull images and which runtimes are allowed to use production repositories. Treat repository policies and runtime roles as part of the deployment boundary, not just as setup details.

### Networking

Review how build systems and runtimes reach the registry, especially in restricted network environments. Registry access failures often look like application deployment issues even though the problem is really in the supply chain.

### Security

Enable scanning and lifecycle controls, and review how base images are chosen and maintained. ECR can help with image hygiene, but it does not replace patching discipline or dependency review.

### Operations and Observability

Track image publishing activity, pull failures, and retention behavior so image provenance stays visible. If the team cannot tell which image version is running and how it got there, operations will degrade quickly.

### Cost

Stored images and data transfer determine cost, so lifecycle cleanup matters.

## Common Mistakes

- Letting many users or pipelines publish directly to production repositories.
- Keeping old tags and unused images indefinitely.
- Treating image scanning as the entire security story.
- Ignoring the relationship between the registry, runtime role, and deployment pipeline.

## How This Fits Into Cloud Engineering

ECR sits at the intersection of platform engineering, deployment automation, and security review. Cloud engineers need to understand that image registries are part of the production system, not just a place to park artifacts.

## Related Projects

- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)

## Related Patterns

- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

## Official References

- [Amazon ECR documentation](https://docs.aws.amazon.com/ecr/)
- [Amazon ECR user guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/what-is-ecr.html)
