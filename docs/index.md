# Cloud Engineer

## Purpose

Cloud Engineer is a practical documentation site for learners and working engineers who want to understand how cloud systems are designed, delivered, secured, operated, and improved over time.

This site is created by McQueen Cloud Advisory and is structured as a learning path, not just a documentation shelf. It is designed to help you connect cloud concepts to deployable systems, provider-specific implementation, and career-ready communication.

In simple terms:

> This site is meant to help you move from recognizing cloud services to building and explaining real cloud systems.

## What This Site Covers

Cloud engineering is broader than provisioning infrastructure. It includes architecture, identity, networking, storage, compute, observability, automation, deployment, reliability, and cost control. A useful site about cloud engineering therefore needs to teach more than product names.

This site focuses on five connected areas:

- [Foundations](foundations/index.md): provider-neutral concepts such as networking, identity, compute, storage, observability, infrastructure as code, and cost management.
- [Google Cloud](gcp/index.md), [AWS](aws/index.md), and [Azure](azure/index.md): provider-specific paths that show how the shared concepts are implemented on real platforms.
- Projects: hands-on builds that turn abstract service knowledge into deployable systems.
- Patterns: reusable architecture explanations that show why a design takes a certain shape.
- [Career](career/index.md): guidance for turning technical work into portfolio artifacts, resume bullets, certification context, and interview-ready explanations.

The site is intended to support learning that is practical, cumulative, and explainable.

## What This Site Is Not

This site is not meant to be a complete catalog of every cloud service.

It is also not designed as a certification cram sheet, a list of isolated tutorials, or a collection of screenshots with little explanation. Those resources can be useful in the right context, but they do not usually build strong engineering judgment on their own.

The goal here is different. The goal is to help you understand questions like these:

- Why should one service be chosen over another?
- How should identities, secrets, and permissions be structured?
- How does traffic or data move through the system?
- How is the system deployed safely?
- How are failures observed and investigated?
- What cost and operational tradeoffs come with the design?

That is the difference between learning services and learning cloud engineering.

## How the Site Is Organized

The site is deliberately arranged so each section builds on the previous one.

### Foundations First

The Foundations section explains the core concepts that transfer across providers. This includes topics such as networking, identity and access, compute models, storage, databases, observability, CI/CD, infrastructure as code, and AI workloads from a cloud engineering perspective.

If you are early in your learning path, this is the best place to start. These pages make the provider sections easier to understand because they explain the problems before the services.

### Provider Paths Next

The AWS, Azure, and Google Cloud sections translate the foundational concepts into specific provider operating models. Each provider path includes:

- overview pages that explain how to approach the platform,
- getting started guides and roadmaps,
- service pages that explain the practical role of key services,
- projects that build real systems,
- and pattern pages that describe reusable architecture shapes.

The point is not to study three clouds at once. The point is to learn one provider deeply enough that the underlying engineering patterns become portable.

### Projects and Patterns Together

Projects show implementation. Patterns show design intent.

Projects help you build something concrete. Pattern pages help you step back and explain why the architecture works, what tradeoffs it makes, and how the same idea can appear across providers.

### Career Material Last, But Not Too Late

The Career section exists because technical skill only becomes visible when you can explain it clearly. Once you build projects, you should also be able to turn them into documentation, resume material, and interview stories.

## How To Use This Site

There is no single perfect sequence for every reader, but a practical path usually looks like this:

1. Start with the [Foundations](foundations/index.md) section and focus first on networking, identity, compute, storage, observability, and cost.
2. Choose one provider path: [Google Cloud](gcp/index.md), [AWS](aws/index.md), or [Azure](azure/index.md).
3. Use the provider roadmap and getting started pages to understand that platform's operating model.
4. Read service pages as they appear in projects instead of trying to study the entire service catalog up front.
5. Use the pattern pages to explain why the architecture is shaped the way it is.
6. Use the [Career](career/index.md) section to turn project work into portfolio, resume, certification, and interview material.

This sequence is designed to reinforce practical learning. You should be building, documenting, and explaining as you go, not only reading.

## Who This Site Is For

This site is especially useful for people who fall into one or more of these groups:

- learners who want a practical path into cloud engineering,
- software or data professionals moving closer to platform and operations work,
- engineers who know one provider and want a stronger cross-provider conceptual base,
- job seekers building projects, portfolio material, and interview stories,
- practitioners who want clearer explanations of service choices and operational tradeoffs.

You do not need to be an expert to use the site well. But the site assumes you want more than shallow exposure.

## What Good Progress Looks Like

Good progress is not measured only by how many services you can name. It is measured by whether you can do work like this:

- explain how a system is structured,
- justify why certain services were chosen,
- describe how access is controlled,
- explain how deployments happen,
- identify how failures are observed,
- discuss the main cost and reliability risks,
- and turn the whole design into a credible project explanation.

If the site helps you do those things more clearly, it is doing its job.

## How This Fits Into Cloud Engineering

Cloud engineering is the work of turning cloud platforms into usable, secure, observable, repeatable systems. This site is organized to support that exact goal. It starts with the transferable concepts, moves into provider-specific implementation, reinforces the material through projects and patterns, and finishes by helping you communicate the work well.

The long-term aim is not only to help you learn cloud services. It is to help you develop cloud engineering judgment.

## Official References

- [Google Cloud documentation](https://cloud.google.com/docs)
- [AWS documentation](https://docs.aws.amazon.com/)
- [Microsoft Learn for Azure](https://learn.microsoft.com/en-us/azure/)
