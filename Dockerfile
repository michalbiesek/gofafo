# Build from the precompiled binary that GoReleaser puts in the context.
# With dockers_v2 the build context contains: linux/amd64/gofafo, linux/arm64/gofafo, etc.
FROM scratch
ARG TARGETPLATFORM
COPY ${TARGETPLATFORM}/gofafo /app
EXPOSE 8080
ENTRYPOINT ["/app"]
