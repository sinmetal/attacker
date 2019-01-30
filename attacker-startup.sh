#!/bin/bash

# Install google-fluentd
curl -sSO https://dl.google.com/cloudagents/install-logging-agent.sh
sha256sum install-logging-agent.sh
sudo bash install-logging-agent.sh
# Restart google-fluentd
service google-fluentd restart

echo "Start attacker!"
gsutil cp gs://bin-kouzoh-p-sinmetal/attacker.bin .
sudo chmod +x attacker.bin
./attacker.bin
echo "DONE attacker!"