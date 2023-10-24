# Set the base Image
FROM golang:alpine3.17 

# Set the working directory
WORKDIR /app

# Copy files and folders from source to folder in the Image 
COPY . .

# Build the go app
RUN go build -o ascii-art-web

#  Exposed port that the containar will listen to
EXPOSE 8080

# Run the app when the contanir start
CMD ["./ascii-art-web"]
