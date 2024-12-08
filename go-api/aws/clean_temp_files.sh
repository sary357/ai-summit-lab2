#!/bin/bash

########
# this script is used to clean all generated folders which prefix is app-
# besides, it also clean all folders which prefix is app-
#########

app_folder_list=`ls |grep "app-"`

for i in $app_folder_list
do
	echo "clean the folder: $i"
	cd $i
	source .venv/bin/activate
	cdk destroy -f
	deactivate
	cd ..
	rm -rf $i
done

app_tmp_folder_list=`ls /tmp/ |grep app`
for i in $app_tmp_folder_list
do
	echo "clean the folder: /tmp/$i"
	rm -rf /tmp/$i

done
