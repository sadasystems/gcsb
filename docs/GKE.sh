#!/bin/bash

PROJECT_ID=gcsb-324823
SPANNER_INSTANCE=test-instance
SPANNER_DATABASE=test-database
GKE_CLUSTER_NAME=test-cluster

# Required services
declare -a GCPServices=("spanner.googleapis.com" "cloudbuild.googleapis.com" "container.googleapis.com" "artifactregistry.googleapis.com" )
 
# Enable each required service
for svc in ${GCPServices[@]}; do
  gcloud services enable $svc --project $PROJECT_ID
done