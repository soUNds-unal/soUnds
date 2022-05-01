# RESOURCE
# https://www.digitalocean.com/community/tutorials/how-to-build-a-node-js-application-with-docker

# Alpine image is derived from the Alpine Linux project, and will help us keep our image size down
FROM node:10-alpine

# Creating these directories will ensure that they have the permissions we want, which will be important when we create local node modules in the container with npm install
RUN mkdir -p /sounds_inte_ms/node_modules && chown -R node:node /sounds_inte_ms

# Set the working directory of the application to /home/node/app
WORKDIR /sounds_inte_ms/src

# Copy the package.json and package-lock.json
COPY package*.json ./

# To ensure that all of the application files are owned by the non-root node user, including the contents of the node_modules directory, switch the user to node before running npm install:
USER node

# After copying the project dependencies and switching our user, we can run npm install
RUN npm install

# Copy your application code with the appropriate permissions to the application directory on the container
# This will ensure that the application files are owned by the non-root node user
COPY --chown=node:node . .

# Finally, expose port 2002 on the container and start the application
EXPOSE 2002	
CMD [ "node", "src/app.js" ]
