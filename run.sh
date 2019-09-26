#!/bin/sh

DB_ADDRESS=127.0.0.1:3306 \
DB_USERNAME=user \
DB_PASSWORD=password \
DB_NAME=simple-blog \
SECRET=mysupersecret \
./go-simple-blog
