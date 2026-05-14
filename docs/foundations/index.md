# Foundations

<div class="section-intro">
	<p class="section-kicker">Provider-neutral core</p>
	<p class="section-lede">Learn the recurring concerns underneath every cloud platform before you map them to product names.</p>
	<p class="section-summary">Use this section to separate durable concepts from vendor naming, defaults, and marketing layers so later service choices are easier to justify.</p>
</div>

<div class="section-grid">
	<div class="section-card">
		<h2>Orientation</h2>
		<p>Start with the pages that define the field, the learning path, and the main provider differences.</p>
		<p class="section-card-links"><a href="what-is-cloud-engineering/">What Is Cloud Engineering</a><a href="foundations-learning-roadmap/">Learning Roadmap</a><a href="cloud-provider-differences-overview/">Provider Differences</a></p>
	</div>
	<div class="section-card">
		<h2>Platform Basics</h2>
		<p>Understand networking, identity, compute, storage, and database choices before comparing managed services.</p>
		<p class="section-card-links"><a href="networking/">Networking</a><a href="identity-and-access/">Identity and Access</a><a href="compute/">Compute</a><a href="storage/">Storage</a><a href="databases/">Databases</a></p>
	</div>
	<div class="section-card">
		<h2>Delivery Systems</h2>
		<p>Study the building blocks that shape how workloads are packaged, deployed, and changed over time.</p>
		<p class="section-card-links"><a href="serverless/">Serverless</a><a href="containers/">Containers</a><a href="infrastructure-as-code/">Infrastructure as Code</a><a href="cicd/">CI/CD</a></p>
	</div>
	<div class="section-card">
		<h2>Operations And Scale</h2>
		<p>Connect observability, cost, and AI-oriented workloads to the rest of the operating model.</p>
		<p class="section-card-links"><a href="observability/">Observability</a><a href="cost-management/">Cost Management</a><a href="ai-and-agents/">AI and Agents</a></p>
	</div>
</div>

## Purpose

This section explains the concepts that sit underneath every cloud platform. The goal is to understand how cloud systems are designed and operated before focusing on provider-specific service names.

## How To Use This Section

Read the Foundations pages as a connected set, not as isolated definitions. Networking affects identity. Identity affects CI/CD. Compute choices affect observability and cost. Storage and database choices affect architecture for years. The more clearly you can see those relationships, the easier it becomes to design systems that work in production.

If you are early in your learning path, start with the general pages and then move into the more technical building blocks. If you already know one provider well, use this section to separate durable concepts from vendor naming and defaults.

## Recommended Reading Order

- [What Is Cloud Engineering](what-is-cloud-engineering.md)
- [Learning Roadmap](foundations-learning-roadmap.md)
- [Cloud Provider Differences](cloud-provider-differences-overview.md)
- [Networking](networking.md)
- [Identity and Access](identity-and-access.md)
- [Compute](compute.md)
- [Storage](storage.md)
- [Databases](databases.md)
- [Serverless](serverless.md)
- [Containers](containers.md)
- [Infrastructure as Code](infrastructure-as-code.md)
- [CI/CD](cicd.md)
- [Observability](observability.md)
- [Cost Management](cost-management.md)
- [AI and Agents](ai-and-agents.md)

## What Strong Foundations Look Like

Strong cloud foundations mean you can explain more than which service to click. You should be able to describe:

- how traffic reaches the system,
- who or what is allowed to use it,
- where code runs,
- where data lives,
- how changes are deployed,
- how failures are detected,
- and what cost or risk grows as the system scales.

That level of understanding transfers across providers and makes later project work much more useful.

## How This Fits Into Cloud Engineering

Cloud engineers use these ideas to choose services, set guardrails, debug failures, and explain tradeoffs to both technical and non-technical stakeholders. Provider pages make far more sense once the underlying concepts are clear.

## Official References

- [Google Cloud documentation](https://cloud.google.com/docs)
- [AWS documentation](https://docs.aws.amazon.com/)
- [Microsoft Learn for Azure](https://learn.microsoft.com/en-us/azure/)
