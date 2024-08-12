
# Use an official Node.js runtime as a parent image
FROM node:18-alpine

# Set working directory
WORKDIR /app

# Copy package.json and yarn.lock, then install dependencies
COPY package.json yarn.lock ./
RUN yarn install

# Copy the rest of the application code
COPY . .

# Build the frontend
RUN yarn build

# Expose the port for the frontend server
EXPOSE 4173

# Start the frontend server
CMD ["yarn", "preview", "--host", "0.0.0.0"]
