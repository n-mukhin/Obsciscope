# Obsciscope

## Overview
Obsciscope is a complete monitoring and logging solution built with Prometheus, Grafana, ELK Stack (Elasticsearch, Logstash, Kibana), and a Go-based service. It provides real-time metrics, data visualization, and log aggregation, allowing for efficient incident detection and resolution.

## Stack
- **Go Service**: A simple HTTP server that exposes Prometheus metrics and generates logs.
- **Prometheus**: Collects metrics from the Go service.
- **Grafana**: Visualizes Prometheus metrics and Elasticsearch logs with dynamic dashboards.
- **Elasticsearch**: Stores aggregated logs for indexing and searching.
- **Logstash**: Ingests and parses logs from the Go service before sending them to Elasticsearch.
- **Kibana**: Searches and visualizes logs stored in Elasticsearch.

## Usage Instructions
1. Clone this repository and navigate to the **Obsciscope** directory.
2. Ensure Docker and Docker Compose are installed on your system.
3. Run `docker-compose up --build` to start all services.
4. Access the Go service at [http://localhost:8080](http://localhost:8080).
5. View metrics in Prometheus at [http://localhost:9090](http://localhost:9090).
6. Open Grafana at [http://localhost:3000](http://localhost:3000), login with `admin/admin`, and explore metrics and logs dashboards.
7. Explore logs in Kibana at [http://localhost:5601](http://localhost:5601).

This setup provides a foundational observability platform that can be adapted and extended for larger, more complex systems.
