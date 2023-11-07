FROM golang:1.21.3 as builder

LABEL maintainer="danvitrinh@outlook.com"
LABEL stage=builder


# Copy relevant folders into container
# Project files
COPY ./cmd /go/src/app/cmd
COPY ./handlers /go/src/app/handlers
COPY ./constants /go/src/app/constants
COPY ./structs /go/src/app/structs

# Go mod and sum 
COPY ./go.mod /go/src/app/go.mod
COPY ./go.sum /go/src/app/go.sum

# Set up execution environment in container's GOPATH
WORKDIR /go/src/app

# Compile binary
RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd -o executable -a -ldflags '-extldflags "-static"'

# Define exposed port
EXPOSE 8080

# Instantiate binary
ENTRYPOINT ["./cmd/executable"]


