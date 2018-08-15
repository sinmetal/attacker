#!/bin/bash

# Install google-fluentd
curl -sSO https://dl.google.com/cloudagents/install-logging-agent.sh
sha256sum install-logging-agent.sh
sudo bash install-logging-agent.sh
# Restart google-fluentd
service google-fluentd restart

BIN_PATH=$(curl http://metadata/computeMetadata/v1/instance/attributes/binpath -H "Metadata-Flavor: Google")
gsutil cp $BIN_PATH attacker.bin
sudo chmod +x attacker.bin

# Your Task
echo "Start Task!"
./attacker.bin
echo "DONE Task!"

# Delete Me
INSTANCE_NAME=$(curl http://metadata/computeMetadata/v1/instance/name -H "Metadata-Flavor: Google")
INSTANCE_ZONE=$(curl http://metadata/computeMetadata/v1/instance/zone -H "Metadata-Flavor: Google")

IFS='/'
set -- $INSTANCE_ZONE
INSTANCE_ZONE=$4
echo $INSTANCE_ZONE
yes | gcloud compute instances delete $INSTANCE_NAME --zone $INSTANCE_ZONE