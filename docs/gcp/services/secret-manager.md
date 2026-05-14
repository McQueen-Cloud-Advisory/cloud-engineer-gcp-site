# Secret Manager

## Purpose

Secret Manager stores and controls access to sensitive runtime values such as API keys, tokens, passwords, certificates, and integration credentials.

It is used when teams need a safer alternative to scattering secrets across source code, local config files, CI variables, and long-lived environment settings.

## Definition

Secret Manager is Google Cloud's managed service for storing and controlling access to secrets such as API keys, tokens, and credentials.

It matters because secret handling is not only about hiding values. Teams also need versioning, access control, rotation, and auditability so application credentials do not become scattered across code, pipelines, and runtime configuration.

Secret Manager is not a replacement for identity design, and it is not a generic configuration store for every application setting. It is specifically for sensitive values that need controlled storage and retrieval. That boundary matters because weak secret handling often begins when teams stop distinguishing between ordinary configuration and high-risk credentials.

In simple terms:

> Secret Manager is where sensitive runtime values live when you want access to be deliberate, reviewable, and easier to rotate.

## What Problem It Solves

Secret Manager solves the problem of secret sprawl. Without a central secret-management service, credentials end up embedded in code, copied into build pipelines, pasted into environment variables, or reused across too many systems.

It gives teams a managed way to store secrets, control which identities can access them, track versions, and rotate values more safely.

That does not remove engineering responsibility. Engineers still need to decide which secrets should exist at all, which workloads should read them, how rotation works, and how applications behave when secret access fails.

## How It Is Commonly Used

Secret Manager is commonly used by Cloud Run services, Cloud Functions, scheduled jobs, and other Google Cloud workloads that need sensitive configuration at runtime. A service account is granted access to a narrow set of secrets, and the workload retrieves those values only when needed.

This service is most effective when it is combined with careful IAM design and a clear decision about which secrets should exist at all.

Common use cases include:

- application credentials for third-party APIs,
- database passwords or connection secrets,
- webhook secrets and signing keys,
- internal service integration tokens,
- deployment-time or runtime values that should not live in source control.

## Foundational Concepts Connected to Secret Manager

Secret Manager connects directly to several cloud engineering foundations.

### Identity and Access

Secret access should be defined around workload identity, not around convenience. The main control question is which human or runtime identity should be able to read which secret.

### Application Configuration

Sensitive configuration needs different handling than ordinary configuration. Secret Manager helps teams separate high-risk values from the rest of the runtime settings.

### Security

Secret storage is part of the security architecture because it directly affects credential hygiene, incident response, and blast radius.

### Auditability

Sensitive systems need more than storage. They need visibility into who accessed secrets, which versions exist, and how secret changes are controlled.

### Reliability

Applications often fail indirectly when secret retrieval breaks. That makes secret availability and access behavior part of runtime reliability, not just part of administration.

## When to Use It

Use Secret Manager when an application or automation path needs sensitive values that should not live directly in code or unmanaged configuration.

Good use cases include:

- runtime credentials for applications and scheduled jobs,
- secrets shared across controlled deployment paths,
- values that need versioning and rotation,
- workloads that should read only a narrow set of secrets through a service account.

Secret Manager is a strong choice when central control, versioning, auditability, and least-privilege access matter.

## When Not to Use It

Do not use Secret Manager as a dumping ground for every application setting. Non-sensitive configuration usually belongs elsewhere.

Secret Manager is also not a substitute for least privilege or secret rotation discipline. A secret stored centrally is still risky if too many identities can read it or if it never changes.

Do not treat secret retrieval as an invisible implementation detail. The runtime still needs a failure plan for missing permissions, deleted versions, or broken rotation.

## Compare To

### Secret Manager vs. Environment Variables

Environment variables are a delivery mechanism inside a runtime. They are not a secret-management system by themselves.

Secret Manager provides central storage, versioning, access control, and auditability. Many workloads still expose a secret to code through configuration, but Secret Manager governs how that secret is stored and retrieved.

### Secret Manager vs. Cloud KMS

Cloud KMS manages cryptographic keys and encryption operations. It is about protecting keys and performing encryption-related functions.

Secret Manager manages application secrets such as tokens, passwords, and API keys. Teams often use both in the same environment, but they solve different problems.

## Tradeoffs

Secret Manager's biggest advantage is centralization. Teams get a safer place to store and manage sensitive values instead of spreading them through code and tooling.

The tradeoff is that secret handling still needs architecture. Rotation, version pinning, retrieval timing, and permission scope all need deliberate decisions.

Secret Manager also makes it easier to improve auditability, but it can create false confidence if teams treat storage as the whole problem. Secrets can still be overexposed in logs, environment dumps, or overly broad IAM policies.

Another tradeoff is runtime dependency. If secret retrieval is required for startup or request handling, the workload needs a clear plan for how retrieval failures are surfaced and handled.

## Common Mistakes

- Giving one service account access to too many secrets.
- Treating secret storage as a substitute for rotation discipline.
- Leaving unused secret versions and old integration credentials in place indefinitely.
- Forgetting that secret retrieval failures can become application outages.
- Storing ordinary configuration alongside high-risk secrets without clear boundaries.

## Cloud Engineering Considerations

### Identity and Access

Use IAM and service accounts so each workload can only read the secrets it actually needs. Secret access should be as narrow as the application's real dependency set.

### Networking and Runtime Retrieval

Review how workloads retrieve secrets and whether network restrictions or private service access affect that path. Secret retrieval is part of runtime behavior, not a separate administrative action.

### Security and Governance

Treat secret lifecycle, rotation, and access logging as part of the security design. The question is not only where a secret is stored but also how it is changed and who can observe that change.

### Observability

Track secret access failures and unusual access patterns so credential issues are visible quickly. Authentication problems often appear first as application errors rather than explicit secret-manager incidents.

### Reliability

Secret retrieval should be part of the reliability model. Engineers should know what happens when a workload starts without access, when a secret version changes, and how applications recover from misconfigured permissions.

### Cost

Stored versions and access operations add cost, so clean up unused secrets and avoid noisy access loops.

## Project and Pattern Connections

Secret Manager is most directly connected to:

- [Project 02: Serverless Contact Form](../projects/project-02-serverless-contact-form.md)
- [Project 03: Scheduled API Ingestion](../projects/project-03-scheduled-api-ingestion.md)
- [Project 05: Agentic RAG Assistant](../projects/project-05-agentic-rag-assistant.md)
- [Serverless API](../patterns/serverless-api.md)
- [Scheduled Job](../patterns/scheduled-job.md)
- [Agentic RAG Assistant](../patterns/agentic-rag-assistant.md)

It also appears in most serious Google Cloud application paths because any runtime with credentials, tokens, or integration secrets needs a safer storage model than hard-coded configuration.

## Official References

- [Secret Manager documentation](https://cloud.google.com/secret-manager/docs)
- [Secret Manager overview](https://cloud.google.com/secret-manager/docs/overview)
- [Access control with IAM](https://cloud.google.com/secret-manager/docs/access-control)
