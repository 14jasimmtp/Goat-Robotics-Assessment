FROM golang:1.22-alpine3.19 AS build-stage

WORKDIR /app
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./timetracker ./cmd

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM scratch AS build-release-stage
WORKDIR /
COPY --from=build-stage /app/timetracker /timetracker
COPY --from=build-stage /app/config.env /
EXPOSE 3000

ENTRYPOINT ["/timetracker"]
