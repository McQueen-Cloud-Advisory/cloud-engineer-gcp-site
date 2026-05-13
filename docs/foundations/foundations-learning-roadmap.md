# Foundations Learning Roadmap

## Purpose

This page gives a practical order for learning cloud engineering so you build usable mental models first and service-specific detail second.

## Stage 1: Learn the Shared Primitives

Start with the concepts that appear in every cloud environment.

- Networking: how systems connect, expose services, and isolate traffic.
- Identity and access: who can do what, from where, and with which credentials.
- Compute: how code runs and what operational responsibility comes with each model.
- Storage and databases: where data lives and how it is protected and queried.
- Observability and cost: how systems are monitored and what they cost to operate.

At this stage, the goal is not certification trivia. The goal is to be able to explain what a system needs before discussing which service implements it.

## Stage 2: Build Small Systems

Once the primitives make sense, build projects that force them to work together.

- A static site teaches storage, delivery, permissions, and deployment.
- A serverless form or API teaches runtime choice, secrets, and monitoring.
- A scheduled ingestion workflow teaches automation, failure handling, and data movement.

Hands-on work matters because cloud engineering is about operating systems, not only naming services. If you cannot describe how a simple workload is exposed, secured, deployed, and observed, add more projects before chasing advanced topics.

## Stage 3: Learn Operations, Not Just Deployment

Many learners stop after they can provision resources. Real cloud engineering starts when the system has to change, fail, scale, or be audited.

- Practice reading logs and debugging failures.
- Review permissions and narrow them intentionally.
- Add CI/CD so the system can be changed safely.
- Track cost and understand the major drivers.
- Revisit designs after you understand how they behave in operation.

This stage is where architecture starts to feel real.

## Stage 4: Specialize Deliberately

After the basics are solid, go deeper into an area such as data platforms, platform engineering, security, AI systems, or multi-cloud delivery. Specialization works best when it extends strong foundations rather than replacing them.

## How To Study Effectively

- Alternate between concept study and hands-on implementation.
- Prefer a few complete projects over dozens of disconnected demos.
- Write down why you chose a service, not just which service you used.
- Rebuild the same pattern on more than one provider to separate concepts from vendor naming.

## How This Fits Into Cloud Engineering

This roadmap exists to prevent shallow learning. Cloud engineers need durable mental models that survive service renames and platform differences. Building that base first makes later provider-specific work faster and much more meaningful.

## Official References

- [Google Cloud architecture framework](https://cloud.google.com/architecture/framework)
- [AWS Well-Architected Framework](https://docs.aws.amazon.com/wellarchitected/latest/framework/welcome.html)
- [Azure Well-Architected Framework](https://learn.microsoft.com/en-us/azure/well-architected/)
