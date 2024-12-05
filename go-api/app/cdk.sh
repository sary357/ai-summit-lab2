#!/bin/bash

cd $1

echo "go to folder: $1"

cdk init app --language python

source .venv/bin/activate

python -m pip install -r requirements.txt


