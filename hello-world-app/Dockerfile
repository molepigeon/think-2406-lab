## Sample Dockerfile to create the Go binary and create the vulnerable image used in the registry tutorial

FROM golang:latest as gobuild
WORKDIR /go/src/github.com/molepigeon/think-2406-lab/hello-world-app
ADD cmd cmd
RUN CGO_ENABLED=0 GOOS=linux /usr/local/go/bin/go build -a ./cmd/hello-world

FROM gliderlabs/alpine:3.6
COPY --from=gobuild /go/src/github.com/molepigeon/think-2406-lab/hello-world-app/hello-world /hello-world
CMD ["/hello-world"]
