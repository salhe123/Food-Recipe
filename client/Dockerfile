# Stage 1: Build the application
FROM node:18.16.1 AS build

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json (or yarn.lock)
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application
COPY . .

# Build the application
RUN npm run build

# Stage 2: Serve the application
FROM node:18

# Set working directory
WORKDIR /app

# Copy built assets from the build stage
COPY --from=build /app/.nuxt /app/.nuxt
COPY --from=build /app/static /app/static
COPY --from=build /app/package*.json /app/
COPY --from=build /app/nuxt.config.js /app/

# Install production dependencies
RUN npm install --only=production

# Expose the port the app runs on
EXPOSE 3000

# Start the application
CMD ["npm", "run", "start"]
