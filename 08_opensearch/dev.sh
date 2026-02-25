#!/bin/bash

echo "Checking Minikube status..."

# Check if minikube is already running
if minikube status | grep -q "Running"; then
    echo "✅ Minikube is already started."
else
    echo "🚀 Starting minikube..."
    minikube start --insecure-registry=host.minikube.internal:5000
fi

# Check if the DOCKER_HOST env var is already pointing to Minikube
if [[ "$DOCKER_HOST" == *"minikube"* ]]; then
    echo "✅ Docker environment is already configured for Minikube."
else
    echo "🔧 Configuring docker to use minikube..."
    eval $(minikube docker-env)
fi

echo "📦 Starting skaffold..."
skaffold dev -p dev -f skaffold.yaml --namespace=default
