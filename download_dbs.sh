#!/bin/sh

# https://db-ip.com/db/download/ip-to-city-lite
curl -# -fL https://download.db-ip.com/free/dbip-city-lite-2026-05.mmdb.gz | gzip -dc > ./dbs/city.mmdb

# https://db-ip.com/db/download/ip-to-asn-lite
curl -# -fL https://download.db-ip.com/free/dbip-asn-lite-2026-05.mmdb.gz | gzip -dc > ./dbs/asn.mmdb
