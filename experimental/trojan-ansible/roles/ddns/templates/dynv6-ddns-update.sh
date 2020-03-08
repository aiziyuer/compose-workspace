#!/bin/bash

curl -X GET \
    http://ipv4.dynv6.com/api/update?hostname={{ SITE_DOMAIN }}&ipv4=auto&token={{ DDNS_TOKEN }}