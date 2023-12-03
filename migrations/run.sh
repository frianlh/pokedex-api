#!/bin/sh

if [ -z $POSTGRES_DB_URL ];then
  echo POSTGRES_DB_URL is not set
  exit
fi

echo Starting migrations
echo $POSTGRES_DB_URL

/migrate -path /migrations -database $POSTGRES_DB_URL up