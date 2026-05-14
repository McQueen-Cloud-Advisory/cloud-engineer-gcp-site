# Azure Patterns

<div class="section-intro">
	<p class="section-kicker">Reusable architecture</p>
	<p class="section-lede">Use the Azure patterns pages to understand recurring workload shapes before you focus on implementation details.</p>
	<p class="section-summary">These pages connect design intent to Azure services so the underlying architecture stays visible when you compare options and tradeoffs.</p>
</div>

<div class="section-grid">
	<div class="section-card">
		<h2>Static Site</h2>
		<p>Understand simple public content delivery with low operational overhead.</p>
		<p class="section-card-links"><a href="static-site/">Open Pattern</a></p>
	</div>
	<div class="section-card">
		<h2>Serverless API</h2>
		<p>Connect request handling, identity, secrets, and managed runtime choices.</p>
		<p class="section-card-links"><a href="serverless-api/">Open Pattern</a></p>
	</div>
	<div class="section-card">
		<h2>Scheduled Job</h2>
		<p>See how recurring automation and integration workflows form one repeatable pattern.</p>
		<p class="section-card-links"><a href="scheduled-job/">Open Pattern</a></p>
	</div>
	<div class="section-card">
		<h2>Analytics Platform</h2>
		<p>Frame data movement, storage, and analytics as one operating system shape.</p>
		<p class="section-card-links"><a href="analytics-platform/">Open Pattern</a></p>
	</div>
	<div class="section-card">
		<h2>Agentic RAG Assistant</h2>
		<p>Connect AI services, retrieval, safety, and runtime operations in one architecture pattern.</p>
		<p class="section-card-links"><a href="agentic-rag-assistant/">Open Pattern</a></p>
	</div>
</div>

## Purpose

This section explains reusable Azure architecture patterns and focuses on why a design takes a certain shape inside Azure.

## What Pattern Pages Are For

Pattern pages help you recognize a repeatable workload, understand the tradeoffs behind it, and map that design to Azure services. They sit between theory and implementation.

Patterns describe the design. Projects show one implementation of that design.

## Current Patterns

- [Static Site](static-site.md) explains how to deliver simple public content with low operational overhead.
- [Serverless API](serverless-api.md) focuses on request handling, identity, secrets, and managed runtime choices.
- [Scheduled Job](scheduled-job.md) shows how Azure services can support recurring automation and integration workflows.
- [Analytics Platform](analytics-platform.md) frames data movement, storage, and analytics as one operating system.
- [Agentic RAG Assistant](agentic-rag-assistant.md) connects AI services, retrieval, safety, and runtime operations into one pattern.

## How To Use These Pages

Read a pattern page before building when you want the architecture intent to be clear. Read it again after implementation when you want to explain why the design works, which tradeoffs it makes, and how Azure services support it.

Pattern pages are also useful when comparing Azure to AWS or Google Cloud because they make the shared architecture visible underneath the provider-specific terminology.

## What Good Pattern Learning Looks Like

You are using these pages well if they help you answer questions like these:

- What recurring workload shape does this pattern represent?
- Why does Azure make this pattern feel the way it does?
- Which services are core and which are supporting?
- What are the main security, reliability, and cost risks?
- When should the team choose a different pattern or a simpler design?

That is what turns provider familiarity into architecture judgment.

## How This Fits Into Cloud Engineering

Cloud engineers need reusable ways of thinking about systems, not only lists of services. Pattern pages build that skill by helping you connect architecture shape to implementation and communication.

## Official References

- [Microsoft Learn for Azure](https://learn.microsoft.com/en-us/azure/)
- [Azure Well-Architected Framework](https://learn.microsoft.com/en-us/azure/well-architected/)
