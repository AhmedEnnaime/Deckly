#!/bin/bash
set -e

echo "Waiting for database to be ready..."
sleep 5

echo "Migrating dev db"
go run cmd/dbmigrate/main.go

echo "Migrating test db"
go run cmd/dbmigrate/main.go -dbName=decklytest

echo "Starting application"
go run cmd/api/main.go