FROM golang:1.13

# Copy project files
WORKDIR /go/src/github.com/RaniSputnik/lovedist

# Build the binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o lovedist .

# /tmp not writable FROM scratch :(
FROM golang:1.10-alpine

# Copy binary
WORKDIR /root/
COPY --from=0 /go/src/github.com/RaniSputnik/lovedist/lovedist .

# Copy love files
COPY ./love ./love

# Run the API
EXPOSE 8080
CMD ["./lovedist"]
