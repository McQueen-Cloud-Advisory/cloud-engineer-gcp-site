# Cloud Engineer GCP Site

## Purpose

This repository contains the **GCP-hosted static website implementation** for **Cloud Engineer by McQueen Cloud Advisory**.

Cloud Engineer is a public documentation and learning site that teaches practical cloud engineering fundamentals through provider-neutral foundations, provider-specific service guides, hands-on projects, reusable architecture patterns, and career-focused guidance.

This repo turns the documentation site into a cloud engineering portfolio project by building the MkDocs source into static site files and deploying them through **Firebase Hosting**.

Planned production URL:

```text
https://cloudengineer.mcqueencloud.com/
```

## What This Repo Demonstrates

This repository is intended to demonstrate:

- Static site generation with MkDocs Material
- Source-controlled technical documentation
- Firebase Hosting deployment
- Custom domain readiness
- GCP project configuration
- Repeatable local build validation
- Cloud engineering documentation practices

The site content is important, but this repo is also meant to show how a technical documentation product can be built, deployed, and maintained as a cloud-hosted asset.

## Site Structure

- `docs/foundations/` contains provider-neutral cloud engineering concepts.
- `docs/gcp/` contains Google Cloud services, projects, and patterns.
- `docs/aws/` contains AWS services, projects, and patterns.
- `docs/azure/` contains Azure services, projects, and patterns.
- `docs/career/` contains portfolio, resume, and interview guidance.

## Technology Stack

- MkDocs Material
- Python virtual environment
- Firebase Hosting
- Google Cloud Platform
- Git and GitHub

## Local Development

Create a virtual environment:

```bash
python -m venv .venv
```

If `python` is not on your Windows `PATH`, use:

```bash
py -m venv .venv
```

Activate the environment on Windows:

```bash
.venv\Scripts\activate
```

Install dependencies:

```bash
pip install -r requirements.txt
```

Run the site locally:

```bash
mkdocs serve
```

The local site will usually be available at:

```text
http://127.0.0.1:8000
```

## Docker

Build the container image:

```bash
docker build -t cloud-engineer-site .
```

Run the container locally:

```bash
docker run --rm -p 8080:8080 cloud-engineer-site
```

The container builds the MkDocs site during the image build and serves the generated static files from a minimal runtime image at:

```text
http://127.0.0.1:8080
```

If you want to run the image on a platform that injects a `PORT` environment variable, the container will listen on that port automatically.

The builder stages in the Dockerfile are pinned to immutable digests for repeatable builds. The container image is also built and vulnerability-scanned automatically in the `Container Security` GitHub Actions workflow on pull requests, pushes to `main`, weekly schedule, and manual runs.

## Build Validation

Build the site strictly before publishing changes:

```bash
mkdocs build --strict
```

Current versions of `mkdocs-material` may print an upstream banner about MkDocs 2.0 unless `NO_MKDOCS_2_WARNING=true` is set. For a quiet local PowerShell build, run:

```powershell
$env:NO_MKDOCS_2_WARNING = 'true'
mkdocs build --strict
```

## Firebase Hosting Deployment

Firebase Hosting serves the generated MkDocs output from the `site/` directory.

Build the site:

```bash
mkdocs build --strict
```

Deploy to Firebase Hosting:

```bash
firebase deploy --only hosting
```

The `site/` directory is generated build output and should not be committed to Git.

## Firebase Configuration

This repo includes Firebase Hosting configuration files:

```text
firebase.json
.firebaserc
```

The Firebase public directory should be:

```text
site
```

For this MkDocs site, Firebase Hosting should not be configured as a single-page app.

## Custom Domain

The intended custom domain is:

```text
cloudengineer.mcqueencloud.com
```

The Firebase default hosting URL may remain available as a fallback, but the preferred public URL should use the McQueen Cloud Advisory subdomain once DNS and SSL are configured.

## Content Guidelines

Write in a practical, direct, educational tone.

Good pages should:

- define the concept clearly,
- explain why it matters,
- show what it looks like in practice,
- connect the topic to real cloud engineering work,
- include official documentation references.

Avoid placeholder text, vague service descriptions, hype language, and certification-cram style summaries.

## Reference Policy

Use official documentation only.

Approved reference sources include:

- Google Cloud official documentation
- AWS official documentation
- Microsoft Learn and Azure official documentation
- Official tooling documentation such as GitHub, Terraform, Docker, Python, Firebase, and MkDocs

Do not use blogs, Reddit, Stack Overflow, Medium, YouTube, or third-party tutorials as page references for now.

## Contributing or Extending the Documentation

Keep provider sections structurally consistent where possible.

When adding a service, project, or pattern page:

1. Create the Markdown file.
2. Add useful first-pass content.
3. Add official references.
4. Update the relevant index page.
5. Add the page to `mkdocs.yml`.
6. Run `mkdocs build --strict`.

Do not add files to navigation before confirming that they exist.

## Git Ignore Notes

The following should not be committed:

```text
site/
.venv/
.firebase/
firebase-debug.log
```

## Publisher

Cloud Engineer is created and maintained by **McQueen Cloud Advisory**.
