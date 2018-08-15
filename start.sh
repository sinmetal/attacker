#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o attacker.bin
gsutil mv attacker.bin $STORAGE_PATH
gcloud compute instances create "datastore-heavy" --zone "us-central1-b" --machine-type "n1-highcpu-8" --scopes "https://www.googleapis.com/auth/cloud-platform" --metadata binpath=$STORAGE_PATH,task=DATASTORE_PutHeavyEntity --metadata-from-file startup-script=ephemeral.sh
gcloud compute instances create "datastore-light" --zone "us-central1-b" --machine-type "n1-highcpu-8" --scopes "https://www.googleapis.com/auth/cloud-platform" --metadata binpath=$STORAGE_PATH,task=DATASTORE_PutLightEntity --metadata-from-file startup-script=ephemeral.sh