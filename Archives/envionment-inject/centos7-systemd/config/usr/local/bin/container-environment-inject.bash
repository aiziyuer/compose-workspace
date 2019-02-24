#!/usr/bin/env bash


/usr/bin/tr '\000' '\n' < /proc/1/environ >/etc/environment