# Model Armor

## Purpose

Model Armor helps apply safety and content controls around AI interactions on Google Cloud.

## Definition

Model Armor is Google's managed safety-control layer for AI interactions. It is designed to help teams apply policy and risk controls around prompts and outputs before those interactions move further through the application.

Like other safety services, it should be treated as one layer in a broader system. It does not remove the need for prompt discipline, secure retrieval, runtime authorization, or human review where required.

In simple terms:

> Model Armor is the managed safety checkpoint in a Google Cloud AI request path.

## What Problem It Solves

It provides a managed safety layer for prompts and outputs so AI applications can enforce policy before results reach users or downstream systems.

## How It Is Commonly Used

It is commonly used for:

- prompt and response safety checks,
- production-facing AI experiences with policy requirements,
- workloads that need explicit moderation or control points,
- AI applications where risky inputs and outputs should be logged and reviewed,
- broader risk-mitigation designs around Vertex AI-based systems.

## When to Use It

- Use it when AI applications need prompt or response safety controls.
- Use it when policy enforcement should be explicit in the AI request path.
- Use it as part of a broader AI risk mitigation approach on Google Cloud.

## When Not to Use It

- Do not assume safety controls replace application-specific validation or human review.
- Do not wait until late testing to define acceptable model behavior.
- Do not treat policy thresholds as fixed forever once users are live.

## Common Mistakes

- Turning on safety controls without defining the actual policy objective.
- Ignoring false positives and false negatives.
- Forgetting that retrieval quality and tool behavior change safety outcomes.
- Failing to monitor how often safety rules trigger.
- Using one safety control where the system needs layered defenses.

## Cloud Engineering Considerations

### Identity and Access

Control who can configure safety behavior and which workloads can use it.

### Networking

Plan where safety checks sit in the request path and how application runtimes reach them.

### Security

Use it as one control in a broader design that also includes retrieval governance, prompt design, and access control.

### Observability

Track safety events so teams can review patterns, false positives, and risky inputs.

### Cost

Safety processing adds cost to the AI path, so monitor where and how often it is used.

## How This Fits Into Cloud Engineering

Model Armor matters because safe AI behavior must be operationalized. Teams need to know what is being filtered, who set the policy, how often it triggers, and what the application does when a request is blocked or flagged.

## Related Projects

- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)

## Related Patterns

- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

## Official References

- [Model Armor overview](https://cloud.google.com/security/products/model-armor)
- [Vertex AI generative AI overview](https://cloud.google.com/vertex-ai/generative-ai/docs/overview)
