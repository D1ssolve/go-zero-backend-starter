FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/app ./cmd/api

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /
COPY --from=build /out/app /app
COPY etc/starter-api.yaml /etc/starter-api.yaml
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app", "-f", "/etc/starter-api.yaml"]
