#!/usr/bin/env bash

#set -o errexit
#set -o nounset
#set -o pipefail

########################### common env - start  ###############################
export set TZ="Asia/Shanghai"
export set UTC=true
########################### common env - stop   ###############################

########################### common function - start ###########################
info() { printf "【info】%b\n" "$*"; }
execute() { printf "【info】command: %s\n" "$*"; eval "$*"; }
########################### common function - stop ############################

# remove output rule
command="iptables -t nat -n -L OUTPUT --line-numbers | awk '\$2==\"REDSOCKS\" {print \$1}' | sort -n -r "
info "command: $command"
rules=`eval "$command"`
for rulenum in $rules; do
  execute "iptables -t nat -D OUTPUT $rulenum"
done

# remove input rule
command="iptables -t nat -n -L INPUT --line-numbers | awk '\$2==\"REDSOCKS\" {print \$1}' | sort -n -r "
info "command: $command"
rules=`eval "$command"`
for rulenum in $rules; do
  execute "iptables -t nat -D INPUT $rulenum"
done

# remove mangle rule
command="iptables -t nat -n -L PREROUTING --line-numbers | awk '\$2==\"REDSOCKS\" {print \$1}' | sort -n -r "
info "command: $command"
rules=`eval "$command"`
for rulenum in $rules; do
  execute "iptables -t nat -D PREROUTING $rulenum"
done

# remove REDSOCKS rule
command="iptables -t nat -n -L REDSOCKS --line-numbers | grep -v Chain| grep -v num |awk '{print \$1}' | sort -n -r "
info "command: $command"
rules=`eval "$command"`
for rulenum in $rules; do
  execute "iptables -t nat -D REDSOCKS $rulenum"
done

# remove old chain
execute "iptables -t nat -X REDSOCKS || true"

# add new chain
execute "iptables -t nat -N REDSOCKS"
execute "iptables -t nat -A REDSOCKS -p tcp -j REDIRECT --to-ports 1080"

# white list 
# eg. export WHITELIST_IP=100.95.0.0/16,100.94.0.0/16
WHITELIST_IP=${WHITELIST_IP:-'100.95.0.0/16'}
IFS=',' read -ra IPS <<< "$WHITELIST_IP"
for ip in "${IPS[@]}"; do
  execute "iptables -t nat -A OUTPUT -p tcp -d $ip -j REDSOCKS"
done

