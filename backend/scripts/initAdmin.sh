#!/usr/local/bin/bash
read -p "username: " username
read -p "password: " password
sqlite3 ~/git/go-andrewwillette.com/backend/sqlite-database.db "INSERT INTO userCredentials(username, password) VALUES('$username', '$password');"