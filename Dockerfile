FROM golang:1.26 AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/kargo-demo ./cmd/kargo-demo

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/kargo-demo /kargo-demo

ENV GIN_MODE=release
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/kargo-demo"]
