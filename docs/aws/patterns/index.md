# AWS Patterns

## Purpose

This section explains reusable AWS architecture patterns and focuses on why a design is shaped the way it is, not only which services appear in it.

## What Pattern Pages Are For

Pattern pages sit between concepts and implementation. They help you recognize a repeatable workload shape, understand the tradeoffs behind it, and connect those tradeoffs to AWS services.

Patterns describe the design. Projects show one implementation of that design.

## Current Patterns

- [Static Site](static-site.md) shows a low-operations public delivery pattern built around storage and simple hosting.
- [Serverless API](serverless-api.md) focuses on request handling, runtime identity, and managed application delivery.
- [Scheduled Job](scheduled-job.md) explains recurring automation and event-driven execution.
- [Analytics Platform](analytics-platform.md) frames ingestion, storage, and analytics as one system.
- [Agentic RAG Assistant](agentic-rag-assistant.md) shows how AI workloads still depend on identity, safety, observability, and application runtime design.

## How To Use These Pages

Read a pattern page before you build when you want the architecture intent to be clear. Read it again after you build when you want to explain why the design works, what tradeoffs it makes, and where the AWS services fit into the broader operating model.

Pattern pages are also useful for comparing the same architecture across providers. That helps you separate the pattern itself from the AWS-specific implementation.

## What Good Pattern Learning Looks Like

You are using these pages well if they help you answer questions like these:

- What problem shape does this pattern solve?
- Why does AWS make this pattern feel the way it does?
- Which service choices are essential and which are optional?
- What are the main security, reliability, and cost risks?
- When should a team choose a different pattern instead?

Pattern thinking is useful because it keeps you from learning cloud only as a list of products.

## How This Fits Into Cloud Engineering

Cloud engineers often need to explain not just what they built, but why the architecture is shaped a certain way. Pattern thinking helps you recognize reusable design ideas and communicate them more clearly.

## Official References

- [AWS documentation](https://docs.aws.amazon.com/)
- [AWS Well-Architected Framework](https://docs.aws.amazon.com/wellarchitected/latest/framework/welcome.html)
