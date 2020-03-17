FROM golang:on

# Copy over to container workspace
ADD . /go/src/github.com/ericz99/go-server-api

# Build src
RUN go install github.com/ericz99/go-server-api

# Run at entrypoint container
ENTRYPOINT /go/bin/output

# Expose port
EXPOSE 8080