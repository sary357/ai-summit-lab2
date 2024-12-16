#!/bin/bash

# verify the input parameters
if [ $# -eq 0 ]; then
	echo "Missing necessary folder name"
	exit 1
elif [ $# -ne 1 ]; then
	echo "Too many input parameters"
	exit 1
else
	echo "start to generate AWS CDK folder: $1"
fi	

TMP_FOLDER=/tmp
REPLACE_TARGET="SampleLambdaStack"

# generate python stack file
current_dir=$(dirname "$0")
template_py=$current_dir/../templates/sample_lambda_stack.py
tmp_python_stack_file=`basename $1`
log_file=$TMP_FOLDER/$tmp_python_stack_file.log
echo "1. generating python stack file..."  >> $log_file
generated_python_stack_file_prefix="${tmp_python_stack_file//-/_}"
cp $(dirname "$0")/../templates/sample_lambda_stack.py $TMP_FOLDER/$generated_python_stack_file_prefix\_stack.py >> $log_file 2>&1
replacement_string="${tmp_python_stack_file//-/}"Stack
sed  -i "s/$REPLACE_TARGET/$replacement_string/g" $TMP_FOLDER/$generated_python_stack_file_prefix\_stack.py   >> $log_file 2>&1
sed  -i "s/class\ app/class\ App/g" $TMP_FOLDER/$generated_python_stack_file_prefix\_stack.py                 >> $log_file 2>&1
echo "1. generating python stack file: Done..."                                                               >> $log_file 

# generate aws cdk files
echo "2. generating aws cdk files..."   >> $log_file 
mkdir -p $1
cd $1
echo "go to folder: $1"                 >> $log_file 
cdk init app --language python          >> $log_file 2>&1
source .venv/bin/activate
python -m pip install -r requirements.txt  >> $log_file 2>&1
echo "2. generating aws cdk files: Done..."  >> $log_file 2>&1

# copy python stack file to aws cdk folder and replace the corresponding string ("SampleLambdaStack")
cp $TMP_FOLDER/$generated_python_stack_file_prefix\_stack.py ./$generated_python_stack_file_prefix/  >> $log_file 2>&1
mv $log_file  ./init_cdk_env.log
