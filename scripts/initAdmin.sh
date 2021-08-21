#!/usr/bin/env bash
read -p "username: " username
read -p "password: " password
sqlite3 ./../sqlite-database.db "INSERT INTO userCredentials(username, password) VALUES('$username', '$password');"