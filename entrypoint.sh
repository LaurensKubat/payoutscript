#!/usr/bin/env bash
redis-server &
service postgresql start
while ! pg_isready -h 0.0.0.0 -p 5432 > /dev/null 2> /dev/null; do
        sleep 1
done

# creating a postgresql user. See https://github.com/ArkEcosystem/core/blob/develop/packages/core/lib/config/testnet/plugins.js
# for the origin of the username and password.
su -c "psql -c \"CREATE USER ark WITH PASSWORD 'password' \"" postgres
su -c "psql -c \"CREATE DATABASE ark_testnet WITH OWNER ark \"" postgres

(cd core/packages/core && yarn full:testnet)