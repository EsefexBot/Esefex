FROM golang:1.21.5

# Install ffmpeg
RUN apt-get update && apt-get install -y ffmpeg

# Install go packages
WORKDIR /api
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build and run
RUN go build -o main .
CMD ["/api/main"]