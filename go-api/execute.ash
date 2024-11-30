#!/bin/ash
# only for alpine image
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
VAULT_FILE=/app/.vaultenv
if [ -f "$VAULT_FILE" ]
then
    list=`cat $VAULT_FILE`
    for i in $list
    do
        export $i
    done
    app
else
    echo "File $VAULT_FILE does not exist!!"
    echo "I will exit now."
    exit 1
fi
