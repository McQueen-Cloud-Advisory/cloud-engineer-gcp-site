# Pub/Sub

## Purpose

Pub/Sub is Google Cloud's managed messaging service for asynchronous event delivery between producers and consumers.

It is used when systems need to decouple producers from consumers, buffer work, or let multiple downstream services react to the same event without tight coordination.

## Definition

Pub/Sub is a managed messaging layer built around topics and subscriptions. Producers publish messages to a topic, and subscribers receive those messages asynchronously through configured subscriptions.

Its value is that systems do not need to know every consumer directly. That makes it a common building block for event-driven architectures, ingestion pipelines, and decoupled integrations.

Pub/Sub is not a database and it is not the same thing as a task queue tuned for one worker to do one piece of work. It is a managed messaging backbone for distributing events and asynchronous work. That boundary matters because many integration problems start when teams use messaging without deciding whether they need fan-out events, explicit work queues, or direct synchronous calls.

In simple terms:

> Pub/Sub lets one part of a system publish work or events without hard-coding every other part that needs to react.

## What Problem It Solves

Pub/Sub solves the problem of tightly coupled integrations. Without a messaging layer, services often depend on direct calls, shared timing assumptions, or brittle chains of integration logic.

Pub/Sub gives teams a managed way to decouple producers from consumers, absorb bursts, fan out events to multiple subscribers, and separate message production from message processing.

That does not remove engineering responsibility. Engineers still need to design message schemas, ordering assumptions, retry behavior, dead-letter handling, subscriber ownership, and backlog visibility.

## How It Is Commonly Used

Pub/Sub is commonly used for:

- event-driven architectures,
- asynchronous processing and work queues,
- ingestion pipelines and decoupled transformations,
- systems where multiple consumers need to react to the same event,
- buffering work between services with different scaling profiles.

In many Google Cloud systems, Pub/Sub sits between producers and downstream runtimes such as Cloud Functions, Cloud Run, Dataflow, or custom subscribers that process work independently.

## Foundational Concepts Connected to Pub/Sub

Pub/Sub connects directly to several cloud engineering foundations.

### Events and Messaging

Pub/Sub is one of the main event-distribution layers in Google Cloud. Topic and subscription design shape how systems communicate asynchronously.

### Decoupling and Backpressure

Messaging matters because producers and consumers rarely scale at the same rate. Pub/Sub creates a buffer and boundary that helps systems evolve more independently.

### Identity and Access

Publishing and subscribing are separate permissions. Producers, consumers, and administrators should not all receive the same trust level.

### Observability

Messaging layers need visibility into backlog, failed delivery, acknowledgment behavior, and subscriber health. Quiet message systems can fail for a long time before humans notice.

### Cost Management

Message volume, retention, replay behavior, and downstream processing all affect total cost. Messaging cost is not only the broker cost. It is the full system cost created by the events.

## When to Use It

Use Pub/Sub when the system needs asynchronous communication rather than a direct request-response dependency.

Good use cases include:

- event-driven architectures,
- ingestion pipelines,
- decoupled integrations between services,
- fan-out event delivery to multiple consumers,
- buffering work between components with different scaling profiles.

Pub/Sub is a strong choice when loose coupling, asynchronous delivery, and multi-consumer patterns are more important than synchronous immediacy.

## When Not to Use It

Do not use Pub/Sub without a plan for message acknowledgment, retries, dead-letter handling, and schema ownership. Messaging is powerful, but it shifts complexity rather than eliminating it.

Pub/Sub is also a poor fit when the workflow truly needs a synchronous answer in the request path or when the team has not defined what subscribers should do when they fall behind.

Do not treat asynchronous messaging as automatically cleaner than direct integration. It is only cleaner when event boundaries, ownership, and operational signals are defined clearly.

## Compare To

### Pub/Sub vs. Direct HTTP Service Calls

Direct service calls are useful when one service needs an immediate answer from another service.

Pub/Sub is useful when producers and consumers should be decoupled, when processing can happen asynchronously, or when multiple consumers may need the same event.

### Pub/Sub vs. Cloud Tasks

Cloud Tasks is designed for explicit task delivery to a worker or HTTP endpoint with a more queue-like model.

Pub/Sub is designed for broader event distribution and asynchronous messaging across multiple subscribers. Use Cloud Tasks when the problem is targeted task execution. Use Pub/Sub when the problem is event propagation and decoupling.

## Tradeoffs

Pub/Sub's biggest advantage is decoupling. Producers can publish events without knowing every consumer, and consumers can evolve independently.

The tradeoff is operational complexity. Message ordering, duplicate delivery, backlog growth, schema evolution, and dead-letter handling all need explicit design.

Pub/Sub also makes it easy to add more subscribers over time. That flexibility is useful, but it can also create unclear ownership if teams subscribe to critical topics without operational responsibility.

Another tradeoff is that asynchronous systems can hide failure. A producer may succeed while the downstream backlog quietly grows or consumers repeatedly fail.

## Common Mistakes

- Publishing messages without a stable schema or versioning approach.
- Ignoring duplicate delivery or retry behavior.
- Creating subscriptions without clear ownership.
- Granting broad publish or subscribe permissions across unrelated services.
- Letting backlog grow without alerts or capacity planning for consumers.
- Treating message publication as proof that the business workflow completed successfully.

## Cloud Engineering Considerations

### Identity and Access

Review which service accounts can publish and subscribe and keep topic or subscription access narrow. Messaging layers should not become shared trust zones by default.

### Networking and Message Flow

Plan how publishers and subscribers reach Pub/Sub and what cross-project or hybrid paths are involved. Message flow is still system flow, even when it is asynchronous.

### Security and Governance

Treat published data as part of the application surface and review who can read messages at each stage. Sensitive data in transit through messaging still needs strong controls.

### Observability

Track backlog, acknowledgment failures, delivery latency, replay behavior, and subscriber health so message flow problems are visible.

### Reliability

Pub/Sub reliability depends on subscriber design as much as the messaging service itself. Consumers need idempotency, retry tolerance, and a plan for how they recover when backlog grows or downstream systems fail.

### Cost

Message volume, retention, egress, and downstream processing determine cost. Noisy producers can make the entire system more expensive, even if the messaging layer itself looks affordable.

## Project and Pattern Connections

Pub/Sub is most directly connected to:

- [Project 03: Scheduled API Ingestion](../projects/project-03-scheduled-api-ingestion.md)
- [Project 04: Analytics Platform](../projects/project-04-analytics-platform.md)
- [Scheduled Job](../patterns/scheduled-job.md)
- [Analytics Platform](../patterns/analytics-platform.md)

It also appears anywhere a Google Cloud system needs asynchronous work distribution, event fan-out, or a buffer between producers and consumers.

## Official References

- [Pub/Sub documentation](https://cloud.google.com/pubsub/docs)
- [Pub/Sub overview](https://cloud.google.com/pubsub/docs/overview)
- [Subscription overview](https://cloud.google.com/pubsub/docs/subscription-overview)
