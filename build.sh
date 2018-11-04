#!/usr/bin/env bash

# build views into binary and then deploy
echo "===== Generating assets file ======="
go-assets-builder views -o assets.go

env GOOS=linux GOARCH=amd64 go build -tags 'bindatafs' -o homef-gin

rsync -azP public/ root@homefbase:/home/apps/homef/public/
#rsync -azP views/ root@homefbase:/home/apps/homef/views/
#rsync -azP vendor/ root@homefbase:/home/apps/homef/vendor/

ssh -l root homefbase "systemctl stop homef.service; systemctl status homef.service; rm /home/apps/homef/homef-gin"
scp homef-gin root@homefbase:/home/apps/homef/

ssh -l root homefbase "systemctl start homef.service; systemctl status homef.service;"

echo "Cleaning Up"
rm homef-gin

echo "Finshed build/deploy"
