FROM jokil/esefexapi:base

# Install go packages
WORKDIR /api
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build and run
RUN go build -o main .
CMD ["/api/main"]
