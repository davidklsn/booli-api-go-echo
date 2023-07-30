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
COPY . ./

# Build the application
RUN go build -buildvcs=false -o main . 

# ---- JavaScript Compile Stage ----
# Use Node.js for ESBuild
FROM node:18 as frontendBuilder

# Set the working directory
WORKDIR /app

# Copy package files and install dependencies
COPY package*.json ./
RUN yarn install

# Copy all files
COPY . .

# Use ESBuild to compile JavaScript
RUN npx esbuild assets/javascript/app.js --bundle --minify --sourcemap --target=es2015 --outfile=bundle.js
RUN npx tailwindcss -i ./assets/css/style.css -o main.css

# ---- Final Stage ----
FROM golang:1.19

WORKDIR /app

# Copy the binary file from builder stage
COPY --from=builder /workspace/main .
COPY --from=builder /workspace/views/*.html ./views/

# Copy the JavaScript bundle from frontendBuilder stage
COPY --from=frontendBuilder /app/bundle.js ./public/dist/app.js
COPY --from=frontendBuilder /app/main.css ./public/dist/main.css

# Expose port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
