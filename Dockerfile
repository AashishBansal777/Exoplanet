# Start with a base Go image
FROM golang:1.22

# Set the Current Working Directory inside the container
RUN mkdir /app
ADD . /app
WORKDIR /app

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]