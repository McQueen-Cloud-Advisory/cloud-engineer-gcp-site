# Firebase Hosting Deployment with Workload Identity Federation

## Purpose

This note documents the current GitHub Actions authentication model for Firebase Hosting deployments in this repository.

It is intentionally high level and excludes secret values, provider resource strings, and other environment-specific credential details.

## Current Deployment Model

- The site is built with `mkdocs build --strict` inside GitHub Actions.
- The deploy workflow uses GitHub OIDC and Google Workload Identity Federation instead of a long-lived Firebase token or a downloaded service account key.
- The workflow requests `id-token: write` so GitHub can mint an OIDC token for the run.
- Google authentication is handled by `google-github-actions/auth`.
- The workflow hardwires the deploy service account email because that identity is not sensitive and is tied to a single repository/project deployment path.
- The workflow reads a single GitHub Actions secret for the Workload Identity Provider resource name.

## Repo-Side Configuration

The live workflow is:

- `.github/workflows/deploy-firebase.yml`

Important implementation details:

- `permissions` includes `contents: read` and `id-token: write`.
- `workload_identity_provider` is loaded from a GitHub Actions secret.
- `service_account` is hardwired in the workflow.
- `project_id` is hardwired in the workflow.

## Google Cloud Side Requirements

The Google Cloud side of this setup requires:

- a deploy service account with Firebase Hosting deployment permissions,
- a Workload Identity Pool,
- an OIDC Workload Identity Provider for GitHub Actions,
- an attribute condition restricting trust to this repository and branch,
- an IAM binding granting the GitHub federated principal `roles/iam.workloadIdentityUser` on the deploy service account.

## GitHub Side Requirements

The GitHub side of this setup requires:

- one repository secret containing the Workload Identity Provider resource name,
- the deploy workflow stored in the repository,
- no checked-in credential files.

## Credential Hygiene

Generated credential files from `google-github-actions/auth` should never be committed or shipped in container build contexts.

This repository already excludes those files through:

- `.gitignore`
- `.dockerignore`

The ignored pattern is:

```text
gha-creds-*.json
```

## Operational Notes

- After changing IAM bindings, Workload Identity Pools, or Providers, allow a short propagation window before rerunning the workflow.
- If deployment fails, the first place to inspect is the `Authenticate to Google Cloud` step in the GitHub Actions run.
- The previous long-lived token path should not be reintroduced unless there is a strong operational reason to do so.

## Current Steady State

As of the successful WIF rollout:

- Docker-based local/runtime validation is in place.
- Container scanning is in place.
- Firebase Hosting deployment uses short-lived federated credentials.
- The fallback GitHub Pages workflow still exists as a secondary publishing path.