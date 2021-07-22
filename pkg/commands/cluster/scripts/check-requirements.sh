#!/bin/sh
if [ ! -f /etc/os-release ]; then
  echo ERROR cannot detect the OS release because of the absence of /etc/os-release
  exit 1;
fi

if cat /etc/*release | grep ^NAME | grep Ubuntu; then
  if cat /etc/*release | grep ^VERSION | grep -F "20."; then
    echo "OS type is checked"
  else
    echo ERROR Support only Ubuntu 20.xx or above
  fi
else
  echo ERROR - support only Ubuntu OS
  exit 1;
fi