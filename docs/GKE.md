# Distributed load testing with GKE

- [Distributed load testing with GKE](#distributed-load-testing-with-gke)
  - [Setup Environment](#setup-environment)
    - [Create a GCP Project](#create-a-gcp-project)
    - [Enable services](#enable-services)
    - [Create a service account](#create-a-service-account)
    - [Grant access to spanner](#grant-access-to-spanner)
    - [Create key for service account](#create-key-for-service-account)
  - [Setup Spanner Database](#setup-spanner-database)
    - [Create a Spanner Instance](#create-a-spanner-instance)
    - [Create a database](#create-a-database)
  - [Setup GKE](#setup-gke)
    - [Create GKE Cluster](#create-gke-cluster)
    - [Import Service Account Key](#import-service-account-key)
    - [Build Docker Container](#build-docker-container)
  - [Run the tool](#run-the-tool)
    - [Single instance load operation](#single-instance-load-operation)
    - [Multi instance load operation](#multi-instance-load-operation)

## Setup Environment

```sh
export PROJECT_ID=spanner-test
export SPANNER_INSTANCE=test-instance
export SPANNER_DATABASE=test-database
export GKE_CLUSTER_NAME=test-cluster
export GCP_REGION=us-west2
```

### Create a GCP Project

```sh
gcloud projects create $PROJECT_ID
```

### Enable services

```sh
gcloud services enable spanner.googleapis.com --project $PROJECT_ID
gcloud services enable cloudbuild.googleapis.com --project $PROJECT_ID
gcloud services enable container.googleapis.com --project $PROJECT_ID
gcloud services enable artifactregistry.googleapis.com --project $PROJECT_ID
```

### Create a service account

```sh
gcloud iam service-accounts create gcsb-test-sa \
    --description="GCSB Test Account" \
    --display-name="gcsb" \
    --project $PROJECT_ID
```

### Grant access to spanner

```sh
gcloud projects add-iam-policy-binding $PROJECT_ID \
    --member="serviceAccount:gcsb-test-sa@${PROJECT_ID}.iam.gserviceaccount.com" \
    --role="roles/spanner.databaseUser"
```

### Create key for service account

```sh
gcloud iam service-accounts keys create key.json --iam-account=gcsb-test-sa@${PROJECT_ID}.iam.gserviceaccount.com
```

## Setup Spanner Database

### Create a Spanner Instance

```sh
gcloud alpha spanner instances create $SPANNER_INSTANCE --config=regional-us-east1 --processing-units=1000 --project $PROJECT_ID
```

### Create a database

```sh
gcloud spanner databases create $SPANNER_DATABASE --instance=$SPANNER_INSTANCE --project $PROJECT_ID
```

## Setup GKE

### Create GKE Cluster

```sh
gcloud container clusters create $GKE_CLUSTER_NAME \
  --project $PROJECT_ID \
  --region $GCP_REGION \
  --num-nodes 3
```

### Import Service Account Key

```sh
kubectl create secret generic gcsb-sa-key --from-file=key.json=./key.json
```

### Build Docker Container

```sh
gcloud builds submit --tag gcr.io/$PROJECT_ID/gcsb .
```

## Run the tool

To run inside Kubernetes, you will need to mount the ServiceAccount key secret into the container. Additionally, you will need to set the environment variable `GOOGLE_APPLICATION_CREDENTIALS` pointing to this key file. Below, there are instructions for single container launches versus parallel launches.

### Single instance load operation

If you want to simply run a single instance of gcsb to perform a load operation, you can launch a POD via the `kubectl run` command.

```sh
kubectl run gcsb-load \
  --image=gcr.io/$PROJECT_ID/gcsb \
  --replicas=1 \
  --restart=Never \
  --overrides='{
     "apiVersion": "v1",
     "spec": {
        "containers": [{
          "name": "gcsb-load",
          "image": "gcr.io/'$PROJECT_ID'/gcsb",
          "command": [ "/gcsb" ],
          "args": [ "load", "--project='$PROJECT_ID'", "--instance='$SPANNER_INSTANCE'", "--database='$SPANNER_DATABASE'", "--table=SingleSingers", "--operations=1000000"],
          "volumeMounts": [{"mountPath": "/var/secrets/google", "name": "google-cloud-key"}],
          "env": [{ "name": "GOOGLE_APPLICATION_CREDENTIALS", "value": "/var/secrets/google/key.json" }]
        }],
        "volumes": [ { "name": "google-cloud-key", "secret": { "secretName": "gcsb-sa-key" } } ]
     }
  }'
```

### Multi instance load operation

```sh
kubectl apply -f docs/gke_load.yaml
```
