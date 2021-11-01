# Distributed load testing with GKE

- [Distributed load testing with GKE](#distributed-load-testing-with-gke)
  - [Setup Environment](#setup-environment)
    - [Create a GCP Project](#create-a-gcp-project)
  - [Setup Spanner Database](#setup-spanner-database)
    - [Enable Spanner API](#enable-spanner-api)
    - [Create a Spanner Instance](#create-a-spanner-instance)
    - [Create a database](#create-a-database)
  - [Setup GKE](#setup-gke)
    - [Create GKE Cluster](#create-gke-cluster)
  - [Build Docker Container](#build-docker-container)

## Setup Environment

```sh
export PROJECT_ID=spanner-test
export SPANNER_INSTANCE=test-instance
export SPANNER_DATABASE=test-database
export GKE_CLUSTER_NAME=test-cluster
```

### Create a GCP Project

```sh
gcloud projects create $PROJECT_ID
```

## Setup Spanner Database

### Enable Spanner API

```sh
gcloud services enable spanner.googleapis.com --project $PROJECT_ID
```

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
gcloud container clusters create $GKE_CLUSTER_NAME --project $PROJECT_ID
```

## Build Docker Container

```sh
docker build -f Dockerfile -t gcr.io/$PROJECT_ID/gcsb:latest .
```