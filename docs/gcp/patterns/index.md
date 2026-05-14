# Google Cloud Patterns

<div class="section-intro">
	<p class="section-kicker">Reusable architecture</p>
	<p class="section-lede">Use the Google Cloud patterns pages to understand workload shape before you focus on one implementation.</p>
	<p class="section-summary">These pages make the design intent visible so you can compare alternatives, explain tradeoffs, and separate the pattern from the platform details.</p>
</div>

<div class="section-grid">
	<div class="section-card">
		<h2>Static Site</h2>
		<p>Understand simple public delivery with low operational overhead and clear deployment boundaries.</p>
		<p class="section-card-links"><a href="static-site/">Open Pattern</a></p>
	</div>
	<div class="section-card">
		<h2>Serverless API</h2>
		<p>See how managed runtimes, secrets, and application data fit together around an API surface.</p>
		<p class="section-card-links"><a href="serverless-api/">Open Pattern</a></p>
	</div>
	<div class="section-card">
		<h2>Scheduled Job</h2>
		<p>Frame recurring automation and asynchronous execution as a repeatable system design.</p>
		<p class="section-card-links"><a href="scheduled-job/">Open Pattern</a></p>
	</div>
	<div class="section-card">
		<h2>Analytics Platform</h2>
		<p>Connect ingestion, storage, and analytical queries into one architecture pattern.</p>
		<p class="section-card-links"><a href="analytics-platform/">Open Pattern</a></p>
	</div>
	<div class="section-card">
		<h2>Agentic RAG Assistant</h2>
		<p>Study AI-oriented system shape without losing sight of application and platform design.</p>
		<p class="section-card-links"><a href="agentic-rag-assistant/">Open Pattern</a></p>
	</div>
</div>

## Purpose

This section explains reusable Google Cloud architecture patterns and helps you connect workload shape to service choice.

## What Pattern Pages Are For

Pattern pages help you understand why a design is assembled a certain way before you focus on the details of one implementation. They are useful when you want to see the system shape first and the project steps second.

Patterns describe the design. Projects show one implementation of that design.

## Current Patterns

- [Static Site](static-site.md) shows a simple public delivery pattern with low operational overhead.
- [Serverless API](serverless-api.md) explains how managed runtimes, secrets, and application data fit together.
- [Scheduled Job](scheduled-job.md) covers recurring automation and asynchronous execution.
- [Analytics Platform](analytics-platform.md) frames ingestion, storage, and analytical query as one architecture pattern.
- [Agentic RAG Assistant](agentic-rag-assistant.md) shows how AI systems still depend on strong application and platform design.

## How To Use These Pages

Read a pattern page before building when you want the architecture intent to be clear. Return to it after implementation when you want to explain tradeoffs, compare providers, or summarize why the design works.

Pattern pages are especially useful for separating the shared architecture from the Google Cloud-specific implementation details.

## What Good Pattern Learning Looks Like

You are using these pages well if they help you answer questions like these:

- What recurring workload shape is this page describing?
- Why does Google Cloud make this pattern feel the way it does?
- Which services are required and which are optional?
- What are the main operational risks?
- What changes if the workload grows or the access model tightens?

That is the kind of thinking that makes patterns more useful than memorized product mappings.

## How This Fits Into Cloud Engineering

Cloud engineers need reusable thinking, not only service familiarity. Pattern pages build that skill by helping you recognize common system shapes and explain them clearly across provider contexts.

## Official References

- [Google Cloud documentation](https://cloud.google.com/docs)
- [Google Cloud architecture framework](https://cloud.google.com/architecture/framework)
