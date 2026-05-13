# Amazon EventBridge

## Purpose

Amazon EventBridge routes events between services and can also trigger workloads on a schedule.

## Definition

Amazon EventBridge is a managed event-routing service. It allows systems to publish events, filter them with rules, and send them to downstream targets such as Lambda functions, queues, workflows, or other services.

Its value is not only technical convenience. EventBridge helps systems become less tightly coupled because producers do not need to know every consumer in advance.

In simple terms:

> EventBridge is a managed event bus that lets AWS systems react to changes and schedules without hard-wiring every integration together.

## What Problem It Solves

Without event routing, systems often depend on direct service-to-service calls, custom schedulers, or brittle integration code. EventBridge gives teams a managed way to build event-driven and scheduled workflows with clearer separation between producers and consumers.

## How It Is Commonly Used

EventBridge is commonly used for:

- scheduled tasks and recurring jobs,
- routing application or platform events to multiple consumers,
- triggering automation and maintenance workflows,
- decoupling ingestion pipelines,
- connecting services that should react to the same state change independently.

## When to Use It

- Use it for scheduled tasks, event routing, and service integration flows.
- Use it when multiple systems should react to the same event without direct coupling.
- Use it to trigger ingestion or maintenance workloads at controlled intervals.
- Use it when the architecture benefits from a shared event bus rather than point-to-point integration.

## When Not to Use It

- Do not use it as a substitute for durable workflow design when the process requires stronger orchestration guarantees.
- Do not publish high-volume noisy events without a clear consumer and retention strategy.
- Do not treat "event-driven" as automatically simpler than direct integration.

## Common Mistakes

- Publishing events without a stable schema or ownership model.
- Creating broad rules that match more traffic than intended.
- Forgetting that downstream targets still need failure handling and retry review.
- Using events to hide unclear service boundaries instead of clarifying them.
- Ignoring cost and operational noise from unnecessary fan-out.

## Cloud Engineering Considerations

### Identity and Access

Review which services can publish events, which rules can invoke targets, and what permissions those targets need. The event bus should not become an uncontrolled privilege path.

### Networking

Understand how EventBridge triggers downstream services and whether those targets also depend on private connectivity or cross-account integration.

### Security

Validate event sources, review event payload sensitivity, and avoid rules that allow overly broad or unintended triggers.

### Observability

Monitor failed invocations, rule matches, and downstream workload behavior together. Event systems are harder to reason about when the trace of cause and effect is not visible.

### Cost

Event volume and downstream target usage determine cost, so filter aggressively and avoid unnecessary fan-out.

## How This Fits Into Cloud Engineering

Cloud engineering often involves deciding whether systems should call each other directly, exchange files, or communicate through events. EventBridge matters because it turns that architectural decision into a managed, repeatable integration layer.

## Related Projects

- [Project 03: Scheduled API Ingestion](../projects/project-03-scheduled-api-ingestion.md)
- [Project 04: Analytics Platform](../projects/project-04-analytics-platform.md)

## Related Patterns

- [Scheduled Job](../patterns/scheduled-job.md)
- [Analytics Platform](../patterns/analytics-platform.md)

## Official References

- [Amazon EventBridge documentation](https://docs.aws.amazon.com/eventbridge/)
- [EventBridge user guide](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-what-is.html)
