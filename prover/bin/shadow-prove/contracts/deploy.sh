#!/bin/bash


echo "Starting npm install --production..."
# Install production dependencies
npm install --production

# Check if npm install was successful
if [ $? -eq 0 ]; then
    echo "npm install --production completed successfully."
else
    echo "npm install --production failed. Exiting..."
    exit 1
fi

echo "Executing deploy script..."
# Execute the deploy script
node ./scripts/deploy.js

# Check if the deploy script was successful
if [ $? -eq 0 ]; then
    echo "Deploy script executed successfully."
else
    echo "Deploy script execution failed."
fi