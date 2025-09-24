FROM scratch
COPY gofafo /app
EXPOSE 8080
ENTRYPOINT ["/app"]
