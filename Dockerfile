# Image built from the precompiled binary that GoReleaser provides.
# No source code (no go.mod) is needed in the context.

FROM scratch
ARG BINARY
COPY ${BINARY} /app
EXPOSE 8080
ENTRYPOINT ["/app"]
