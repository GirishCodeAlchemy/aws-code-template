# Use the official Node.js image
FROM node:20

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY package*.json ./
COPY . .

# Install dependencies
RUN npm install


# Expose the Lambda port (if needed)
# EXPOSE 3000

# Set the CMD to run your Lambda function
CMD [ "npm", "start" ]
