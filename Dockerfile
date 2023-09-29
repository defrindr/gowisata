# Use Ubuntu as base image
FROM ubuntu:latest

# Set the working directory inside the container
WORKDIR /app

# Install necessary packages using apk (Alpine package manager)
RUN apt-get update \
    && apt-get install -y golang \
    && apt-get install -y golang \
    && apt-get install -y ca-certificates
# Copy the entire project into the container
COPY . .

RUN export GOINSECURE="proxy.golang.org/*,github.com,github.com/*"

# Install Fiber package
# RUN GOPROXY=https://goproxy.cn go get
RUN GOPROXY=https://goproxy.cn go get

# Build the Go application
RUN go build -o main .

VOLUME /app/

# Expose the port used by the application
EXPOSE 3000

# Define the default command to run the application
CMD ["./main"]
