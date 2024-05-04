# base image
FROM golang:1.22.2-alpine3.19

#WORKDIR /app
WORKDIR /app

# Copy all the files from the current directory to the app directory
COPY . .

#install dependencies
RUN go mod tidy
RUN make all

# Expose port 8090 to the outside world
EXPOSE 8090

# Command to run the executable
CMD ["go", "run", "server.go"]