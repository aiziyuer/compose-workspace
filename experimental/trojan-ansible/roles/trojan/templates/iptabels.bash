#!/bin/bash

# Remove old rules
iptables -t nat -D OUTPUT -p tcp -j TROJAN
iptables -t nat -F TROJAN
iptables -t nat -X TROJAN

iptables -t nat -N TROJAN
# Allow connection to reserved networks
iptables -t nat -A TROJAN -d 80.251.216.179/32 -j RETURN
iptables -t nat -A TROJAN -d 0.0.0.0/8 -j RETURN
iptables -t nat -A TROJAN -d 10.0.0.0/8 -j RETURN
iptables -t nat -A TROJAN -d 127.0.0.0/8 -j RETURN
iptables -t nat -A TROJAN -d 169.254.0.0/16 -j RETURN
iptables -t nat -A TROJAN -d 172.16.0.0/12 -j RETURN
iptables -t nat -A TROJAN -d 192.168.0.0/16 -j RETURN
iptables -t nat -A TROJAN -d 224.0.0.0/4 -j RETURN
iptables -t nat -A TROJAN -d 240.0.0.0/4 -j RETURN

iptables -t nat -A TROJAN -p tcp -j REDIRECT --to-ports 1081

# Redirect tcp to PROXY 
iptables -t nat -A OUTPUT -p tcp -j TROJAN