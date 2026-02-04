# Use the official Bun image
FROM oven/bun:latest AS build

WORKDIR /app

# Copy package files
COPY package.json bun.lock ./

# Install dependencies
RUN bun install --frozen-lockfile

# Copy the rest of the code
COPY . .


EXPOSE 3000

# Change this line in your Dockerfile
CMD ["bun", "run", "dev", "--host", "0.0.0.0"]