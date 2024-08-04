# syntax=docker/dockerfile:1

FROM golang:1.22.5

# Set working directory
WORKDIR /usr/local/go/src/api-rs

# Download Go modules
RUN go mod download

# Copy the source code
COPY . ./

# Set environment variables
ENV TZ="Asia/Jakarta"

# Install Air for live reloading
RUN go install github.com/air-verse/air@latest && go mod download

# Expose port
EXPOSE 8081

# Command to run Air
CMD ["air", "-c", ".air.toml"]
