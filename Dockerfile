# ---- Build Stage ----
# Use an official Go runtime as a builder
FROM golang:1.19 as builder

# Set the working directory
WORKDIR /workspace

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Build the application
RUN go build -o main .

# ---- JavaScript Compile Stage ----
# Use Node.js for ESBuild
FROM node:18 as esbuilder

# Set the working directory
WORKDIR /workspace

# Copy package files and install dependencies
COPY package*.json ./
RUN npm install

# Copy all files
COPY . .

# Use ESBuild to compile JavaScript
RUN npx esbuild app.js --bundle --minify --sourcemap --target=es2015 --outfile=bundle.js

# ---- Final Stage ----
FROM golang:1.19

WORKDIR /workspace

# Copy the binary file from builder stage
COPY --from=builder /workspace/main .

# Copy the JavaScript bundle from esbuilder stage
COPY --from=esbuilder /app/bundle.js ./public/dist/app.js

# Expose port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
