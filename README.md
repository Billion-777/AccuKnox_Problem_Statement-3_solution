# AccuKnox_Problem_Statement-3_solution




# Kubernetes Deployment of a Go Web App

## Problem Statement 3 Solution
- A Go web app that displays current date/time and hostname.
- Dockerized and pushed to Docker Hub.
- Deployed on Kubernetes with 2 replicas.
- Exposed via a LoadBalancer service.

### Steps to Run

1. **Build & Push Docker Image**
   ```bash
   docker build -t <your-dockerhub-username>/datetime-app .
   docker push <your-dockerhub-username>/datetime-app
