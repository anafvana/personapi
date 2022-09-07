#!/bin/bash
apt install psql

user="user123"

psql -h localhost -d postgres -U "$user" -W -f ./person.sql