# Google Cloud Patterns

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
