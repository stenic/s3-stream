FROM golang:1.17 as builder

WORKDIR /workspace
COPY go.* .
RUN go mod download
COPY main.go main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /3s-stream main.go


# Use distroless as minimal base image to package the project
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot

WORKDIR /
COPY --from=builder /3s-stream .
ENTRYPOINT ["/3s-stream"]

