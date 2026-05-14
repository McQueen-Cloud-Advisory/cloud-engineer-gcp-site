# Cloud Engineer

<div class="home-intro">
	<p class="home-kicker">Published by McQueen Cloud Advisory</p>
	<p class="home-positioning">Cloud engineering, explained through systems - not service catalogs.</p>
	<p class="home-summary">A practical documentation site focused on how cloud systems are designed, deployed, secured, operated, automated, and improved over time.</p>
	<p class="home-actions">
		<a class="md-button md-button--primary" href="start-here/">Start Here</a>
		<a class="md-button" href="roadmap/">View Roadmap</a>
	</p>
</div>

<div class="home-grid">
	<div class="home-card">
		<h2>Start Here</h2>
		<p>Understand the structure of the site and the fastest path into the material.</p>
		<p class="home-card-links"><a href="start-here/">Open Section</a></p>
	</div>
	<div class="home-card">
		<h2>Roadmap</h2>
		<p>Use a deliberate learning path instead of trying to absorb every cloud topic at once.</p>
		<p class="home-card-links"><a href="roadmap/">View Roadmap</a></p>
	</div>
	<div class="home-card">
		<h2>Foundations</h2>
		<p>Learn the provider-neutral concepts behind identity, networking, compute, storage, and operations.</p>
		<p class="home-card-links"><a href="foundations/">Explore Foundations</a></p>
	</div>
	<div class="home-card">
		<h2>Google Cloud</h2>
		<p>Follow the Google Cloud path through services, projects, patterns, and practical implementation notes.</p>
		<p class="home-card-links"><a href="gcp/">Open Google Cloud</a></p>
	</div>
	<div class="home-card">
		<h2>AWS</h2>
		<p>Study AWS through the same system-design lens: services in context, not memorized in isolation.</p>
		<p class="home-card-links"><a href="aws/">Open AWS</a></p>
	</div>
	<div class="home-card">
		<h2>Azure</h2>
		<p>Work through Azure with the same focus on identity, runtime boundaries, operations, and tradeoffs.</p>
		<p class="home-card-links"><a href="azure/">Open Azure</a></p>
	</div>
	<div class="home-card">
		<h2>Projects and Patterns</h2>
		<p>Move from service definitions into deployable systems and reusable architecture shapes.</p>
		<p class="home-card-links"><a href="gcp/projects/">Projects</a><a href="gcp/patterns/">Patterns</a></p>
	</div>
	<div class="home-card">
		<h2>Career</h2>
		<p>Turn technical work into portfolio descriptions, resume bullets, interview stories, and clearer explanations.</p>
		<p class="home-card-links"><a href="career/">Open Career</a></p>
	</div>
</div>

Cloud Engineer is a practical learning and documentation site for people who want to understand how cloud systems are designed, deployed, secured, observed, automated, and improved over time.

The site is built around a simple idea: cloud engineering is not the same thing as memorizing cloud services. A useful cloud engineer needs to understand the concepts behind the services, the tradeoffs between implementation options, and the operational responsibilities that appear after a system is deployed.

In simple terms:

> Cloud Engineer helps learners move from recognizing cloud services to building, documenting, and explaining real cloud systems.

## What You Will Find Here

This site focuses on cloud engineering through a set of connected surfaces.

- [Foundations](foundations/index.md) explains provider-neutral concepts such as identity, networking, compute, storage, databases, serverless, containers, CI/CD, observability, cost management, and AI or agentic workloads.
- [Google Cloud](gcp/index.md), [AWS](aws/index.md), and [Azure](azure/index.md) translate those concepts into provider-specific services, projects, patterns, roadmaps, and getting-started paths.
- Service pages focus on practical fit: what problem a service solves, when to use it, when not to, what it should be compared against, and what operating responsibilities come with it.
- Project pages turn concepts into deployable, explainable systems that are meant to be portfolio-ready.
- Pattern pages explain reusable architecture shapes so the site teaches design judgment rather than isolated steps.
- [Career](career/index.md) helps turn technical work into portfolio descriptions, resume bullets, interview stories, and clearer communication.

## What This Site Is Not

This site is not a complete catalog of every cloud service. The major providers already maintain official documentation for that purpose.

It is not a certification cram sheet, and it is not a collection of disconnected tutorials. The goal here is to build judgment: knowing why one service, architecture, or deployment path makes sense over another.

## How To Use This Site

A practical path through the site looks like this:

1. Start with [Foundations](foundations/index.md) to understand the concepts that transfer across providers.
2. Choose one provider path: [Google Cloud](gcp/index.md), [AWS](aws/index.md), or [Azure](azure/index.md).
3. Use roadmap and getting-started pages to understand that platform's operating model.
4. Read service pages when they appear in projects rather than trying to memorize an entire service catalog.
5. Use pattern pages to understand why an architecture is shaped the way it is.
6. Build from the project pages and document the tradeoffs, access model, deployment path, and operational decisions.
7. Use the [Career](career/index.md) section to turn the work into portfolio and interview material.

The best way to use this site is to build while reading. Cloud engineering knowledge becomes more durable when it is connected to a deployed system, a troubleshooting decision, or an architecture explanation.

## Who This Site Is For

This site is especially useful for:

- learners building a practical path into cloud engineering,
- data and analytics professionals moving toward cloud platforms,
- software professionals who want stronger infrastructure and operations context,
- engineers who know one provider and want a clearer cross-provider foundation,
- job seekers building portfolio projects and interview stories,
- practitioners who want clearer explanations of cloud service choices and tradeoffs.

You do not need to be an expert to start. But the site assumes you want more than shallow exposure. The goal is to understand enough to build, operate, troubleshoot, and explain.

## What Good Progress Looks Like

Good progress is not measured by how many cloud services you can name. It is measured by whether you can reason clearly about systems.

You are making progress when you can:

- explain what a system does,
- describe how traffic or data moves through it,
- identify which identities can deploy it and which identities can run it,
- justify why specific services were chosen,
- compare those services to reasonable alternatives,
- explain how failures would be detected,
- describe the main cost drivers,
- document the deployment path,
- and turn the work into a credible project explanation.

That is the difference between learning cloud products and developing cloud engineering judgment.

## Current Direction

The site is documentation-first and implementation-aware. It began around a real deployed MkDocs documentation site and continues to expand across Google Cloud, AWS, and Azure without pretending the providers are identical.

The goal is consistency in depth, practical design judgment, and portfolio-ready explanations rather than broad but shallow catalog coverage.

## Reference Policy

This site uses official documentation as the primary reference base. Service, project, and pattern pages should rely on vendor documentation and official tooling documentation rather than third-party tutorials.

That choice is deliberate. Official documentation reduces the risk of copying outdated implementation patterns and helps keep explanations grounded in supported platform behavior.

## Official References

- [Google Cloud documentation](https://cloud.google.com/docs)
- [AWS documentation](https://docs.aws.amazon.com/)
- [Microsoft Learn for Azure](https://learn.microsoft.com/en-us/azure/)
- [Firebase Hosting documentation](https://firebase.google.com/docs/hosting)
- [MkDocs documentation](https://www.mkdocs.org/)
