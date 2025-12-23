# Argus

A Real-Time Economic Intelligence & Price Trend Platform

📌 Overview

Economy Platform is a cloud-native, event-driven data system that tracks real-world economic signals across categories like technology, culture, clothing, and food.
It ingests price and availability data from multiple sources, processes them as real-time events, stores long-term history in a data lake, and exposes insights such as price spikes, shortages, and market trends.

Example insights:

📈 Sudden increase in RAM prices

📉 Seasonal drop in clothing prices

🍎 Food inflation trends over time

💻 Tech hardware supply shocks

The platform is designed as a distributed backend system, not a UI-heavy app, focusing on scalability, correctness, and data integrity.

🎯 Goals

Treat economic changes as events, not database updates

Provide real-time signals and historical intelligence

Demonstrate modern cloud & system engineering practices

Be buildable by a single backend engineer, yet production-grade in design

🧠 Core Principles

Event-Driven Architecture – Kafka as the backbone

Clean Architecture – Business logic independent of infrastructure

Microservices – Each service owns one responsibility

Data Lake First – History is a first-class citizen

Typed APIs – gRPC for internal communication

🏗️ High-Level Architecture
[ Data Sources ]
      ↓
[ Scraper Services ]
      ↓
[ Kafka Topics ]
      ↓
[ Normalization / Enrichment ]
      ↓
[ Stream Processors ]
      ↓
[ Data Lake (S3/MinIO) ]
      ↓
[ Analytics & APIs ]
      ↓
[ Dashboards / Alerts ]

🔁 Data Flow (Step-by-Step)
1️⃣ Data Ingestion

Scraper services periodically fetch prices and metadata from public sources

Each observation is emitted as an immutable event

Kafka Topic

raw.market.observed

2️⃣ Normalization

Raw events are cleaned, validated, and standardized

Currency normalization, unit normalization, sanity checks

Kafka Topics

raw.market.observed → market.normalized

3️⃣ Stream Processing (Real-Time Signals)

Detects meaningful changes:

Price spikes

Sudden drops

Supply shortages

Emits higher-level events

Kafka Topics

market.normalized → market.price.changed
market.normalized → market.anomaly.detected

4️⃣ Data Lake Storage

All events are stored immutably for long-term analysis.

Storage Layout

/raw/
  /tech/
  /food/
/processed/
  normalized_prices.parquet
/aggregated/
  daily_trends.parquet
  category_index.parquet


Formats:

Parquet

Partitioned by date, category, product

5️⃣ Batch Analytics

Computes:

Inflation trends

Category indices

Volatility scores

Runs hourly/daily jobs

6️⃣ API & Consumption

gRPC APIs expose:

Price history

Trend analytics

Real-time signals

UI / dashboards are optional consumers

🧩 Tech Stack
Backend & Systems

Go – Core language

Kafka – Event streaming

gRPC + Protobuf – Service communication

Clean Architecture – Domain-driven design

Data & Storage

S3 / MinIO – Data lake

Parquet – Analytical storage

PostgreSQL (optional) – Control plane metadata

Infrastructure

Docker

Docker Compose

Terraform (optional cloud deploy)

Observability

Prometheus – Metrics

Structured logging (zap/slog)

📂 Repository Structure (Clean Architecture)
economy-platform/
├── api/                        # Protobuf contracts & generated code
│   ├── proto/
│   └── gen/
├── internal/                   # Shared infrastructure
│   ├── config/
│   ├── logger/
│   ├── kafka/
│   └── platform/
├── services/                   # Microservices
│   ├── scraper/
│   │   ├── cmd/
│   │   ├── internal/
│   │   │   ├── domain/
│   │   │   ├── usecase/
│   │   │   ├── repository/
│   │   │   └── delivery/
│   │   └── Dockerfile
│   └── normalization/
├── deployments/
│   ├── docker-compose.yml
│   └── terraform/
├── scripts/
├── Makefile
└── go.work

🧼 Clean Architecture Boundaries

Domain – Entities, value objects, interfaces

Usecase – Business rules (pure Go)

Repository – Data lake, DB, Kafka implementations

Delivery – gRPC handlers, Kafka consumers

cmd/ – Dependency injection & wiring

➡️ Business logic never depends on Kafka, S3, or gRPC.

🚀 How to Start (Development)
# Start infrastructure
docker-compose up -d kafka minio

# Generate protobufs
make proto

# Run a service
cd services/scraper
go run cmd/main.go

🧪 Failure Handling

Kafka enables replayability

Consumers are idempotent

Data lake acts as source of truth

Services can crash and recover safely

📈 Future Extensions

ML-based inflation prediction

Regional price comparison

Public economic dashboards

Open data APIs

🧑‍💻 Why This Project Matters

This project demonstrates:

Real-world relevance

System-level thinking

Cloud-native architecture

Strong backend fundamentals

It is not a CRUD app — it is a data platform.

📄 License

MIT (or your choice)

If you want, next I can:

Review this README like an interviewer

Cut it down to a 1-minute interview explanation

Help you write commit milestones

Define MVP vs stretch goals

Just say the word.