# Start from golang base image
FROM golang:1.17-alpine as builder

RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true

# Set the current working directory inside the container
WORKDIR /usr/src/app

# Copy go.mod, go.sum files and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy sources to the working directory
COPY . .

## Set the Go environment
#ENV GOOS linux
#ENV CGO_ENABLED 0
#ENV GOARCH amd64

# Build the app
ARG PROJECT
#WORKDIR $PROJECT/cmd/app
#RUN go build -a -v -o server $project
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./$PROJECT/cmd/app

# Start a new stage from scratch
FROM scratch

ARG PROJECT

WORKDIR /dist

# Copy build artifacts from the previous stage
COPY --from=builder /usr/src/app/server /server

COPY --from=builder /usr/src/app/$PROJECT/config ../config


# copy the ca-certificate.crt from the build stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Run the executable
CMD ["/server"]
