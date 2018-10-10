#!/usr/bin/env bash

rsync -avvzP onajome@homef.org:/home/onajome/homef.org/sites/default/files/ /Users/jome/projects/homef/files/

go run migration/migration.go