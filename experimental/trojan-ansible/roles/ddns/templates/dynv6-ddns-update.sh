#!/bin/bash

curl -X GET \
    http://ipv4.dynv6.com/api/update?hostname={{ ssl_domain_name }}&ipv4=auto&token={{ ddns_token }}