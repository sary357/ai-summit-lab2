#!/bin/bash

# verify the input parameters
if [ $# -eq 0 ]; then
	echo "Missing necessary folder name (AWS CDK FOLDER)"
	exit 1
elif [ $# -ne 1 ]; then
	echo "Too many input parameters"
	exit 1
fi	

# generate aws cdk files
echo "1. deploying cdk app..."
cd $1
echo "go to folder: $1"
source .venv/bin/activate
cdk bootstrap
cdk deploy  --require-approval never
echo "1. deployed cdk app..."

