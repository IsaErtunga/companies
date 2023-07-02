FROM golang:1.20 AS build-stage

WORKDIR /app

# download go modules
COPY go.mod go.sum ./
RUN go mod download

# copy source code
COPY *.go ./
COPY cmd/ ./cmd
COPY internal/ ./internal

# compile
RUN CGO_ENABLED=0 GOOS=linux go build -o /companies ./cmd

# run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /companies /companies

EXPOSE 9090

USER nonroot:nonroot

ENTRYPOINT ["/companies"]