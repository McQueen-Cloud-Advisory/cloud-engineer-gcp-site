# Cloud Engineer

## Purpose

Cloud Engineer is a public documentation and learning site by McQueen Cloud Advisory. It teaches practical cloud engineering fundamentals through provider-neutral foundations, provider-specific service guides, hands-on projects, reusable architecture patterns, and career-focused guidance.

The site is built with MkDocs Material and is organized to help learners understand how real cloud systems are designed, deployed, secured, monitored, automated, and improved over time.

## Site Structure

- `docs/foundations/` contains provider-neutral cloud engineering concepts.
- `docs/gcp/` contains Google Cloud services, projects, and patterns.
- `docs/aws/` contains AWS services, projects, and patterns.
- `docs/azure/` contains Azure services, projects, and patterns.
- `docs/career/` contains portfolio, resume, and interview guidance.

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

## Build Validation

Build the site strictly before publishing changes:

```bash
mkdocs build --strict
```

Current versions of `mkdocs-material` print an upstream banner about MkDocs 2.0 unless `NO_MKDOCS_2_WARNING=true` is set. The GitHub Actions workflow in this repository sets that environment variable automatically. For a quiet local PowerShell build, you can run:

```powershell
$env:NO_MKDOCS_2_WARNING = 'true'
mkdocs build --strict
```

## Content Guidelines

Write in a practical, direct, educational tone. Prefer explanations that connect service choices to security, operations, deployment, and cost tradeoffs. Avoid placeholder text and avoid turning the site into a certification glossary.

## Reference Policy

Use official documentation only. Service, project, and pattern pages should reference Google Cloud, AWS, Microsoft, or approved official tooling documentation only.

## Contributing or Extending the Documentation

Keep provider sections structurally consistent where possible. When you add a service, project, or pattern page, update the relevant index page and add the page to `mkdocs.yml` only after confirming the file exists.

## Publisher

Cloud Engineer is created and maintained by McQueen Cloud Advisory.
