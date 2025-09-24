# gofafo

Tiny Go app + GoReleaser pipeline that creates GitHub Releases and publishes a Docker image to GHCR.

## One-time setup

1. In your repo, go to **Settings → Actions → General → Workflow permissions** and enable **Read and write permissions**.
2. No extra secrets required—`GITHUB_TOKEN` is used for Releases and GHCR.

## Release flow

Tag and push:

```bash
git tag v0.1.0
git push origin v0.1.0
