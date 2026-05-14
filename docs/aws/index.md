# AWS

<div class="section-intro">
	<p class="section-kicker">Provider path</p>
	<p class="section-lede">Learn AWS through its account model, IAM boundaries, regional design choices, and broad service tradeoffs.</p>
	<p class="section-summary">This section helps you make sense of AWS by connecting service choices back to workload fit, operating boundaries, and architecture judgment.</p>
</div>

<div class="section-grid">
	<div class="section-card">
		<h2>Getting Started</h2>
		<p>Get clear on accounts, IAM, regions, and the first services worth learning before widening scope.</p>
		<p class="section-card-links"><a href="getting-started/">Open Getting Started</a></p>
	</div>
	<div class="section-card">
		<h2>Roadmap</h2>
		<p>Use the roadmap to keep AWS breadth from turning into shallow, unstructured learning.</p>
		<p class="section-card-links"><a href="roadmap/">View Roadmap</a></p>
	</div>
	<div class="section-card">
		<h2>Services</h2>
		<p>Read the service guides when you need clear fit, tradeoffs, and comparisons inside real systems.</p>
		<p class="section-card-links"><a href="services/">Browse Services</a></p>
	</div>
	<div class="section-card">
		<h2>Projects And Patterns</h2>
		<p>Turn the service model into deployable applications and reusable design explanations.</p>
		<p class="section-card-links"><a href="projects/">Projects</a><a href="patterns/">Patterns</a></p>
	</div>
</div>

## Purpose

This section applies cloud engineering fundamentals using AWS and explains how common cloud patterns show up in AWS's account-centric operating model.

## How AWS Feels In Practice

AWS gives you a very broad service catalog and a lot of configuration depth. That is powerful, but it also means you need a clear mental model early.

- Accounts are major governance and billing boundaries.
- IAM roles and policies shape how humans, workloads, and automation gain access.
- Regions matter because many services are region-scoped and architecture decisions often depend on where a workload runs.
- AWS frequently offers several ways to solve the same problem, so clarity about the workload matters more than memorizing options.

Learning AWS well means learning how to work with those tradeoffs, not just learning product names.

## What This Section Focuses On

The first pass through AWS focuses on the services and patterns most useful for learning practical delivery work.

- Identity and access with IAM.
- Object storage with S3.
- APIs and serverless compute with API Gateway and Lambda.
- Application data with DynamoDB.
- Secrets, monitoring, and scheduled workloads.
- Later expansion into analytics and Bedrock-based AI workloads.

This gives you a path from simple public delivery to recurring automation, data work, and AI-oriented systems.

## Recommended Path Through This Section

1. Start with [Getting Started](getting-started.md) to understand accounts, IAM, regions, and the first service set.
2. Use [Roadmap](roadmap.md) to sequence the work instead of treating every service page as equally urgent.
3. Study [Services](services/index.md) as you build the [Projects](projects/index.md).
4. Return to [Patterns](patterns/index.md) when you want to explain why the architecture works.

## Section Links

- [Getting Started](getting-started.md)
- [Roadmap](roadmap.md)
- [Services](services/index.md)
- [Projects](projects/index.md)
- [Patterns](patterns/index.md)

## How This Fits Into Cloud Engineering

AWS is useful for learning cloud engineering because it forces you to think clearly about accounts, identity, regions, service boundaries, and deployment tradeoffs. The goal is not to know every AWS service. The goal is to understand how to build and operate complete systems inside the AWS model.

## Official References

- [AWS documentation](https://docs.aws.amazon.com/)
- [AWS Well-Architected Framework](https://docs.aws.amazon.com/wellarchitected/latest/framework/welcome.html)
