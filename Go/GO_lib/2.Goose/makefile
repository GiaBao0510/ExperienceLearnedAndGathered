# Chỉ định shell là cmd.exe cho Windows
SHELL := cmd.exe
.SHELLFLAGS := /C

GOOSE_DRIVER ?= postgres
GOOSE_DBSTRING = "host=localhost port=5432 user=admin password=admin123 dbname=test sslmode=disable"
GOOSE_MIGRATION_DIR ?= sql

up:
    goose -driver=$(GOOSE_DRIVER) -dbstring="$(GOOSE_DBSTRING)" -dir=$(GOOSE_MIGRATION_DIR) up

down:
    goose -driver=$(GOOSE_DRIVER) -dbstring="$(GOOSE_DBSTRING)" -dir=$(GOOSE_MIGRATION_DIR) down

reset:
    goose -driver=$(GOOSE_DRIVER) -dbstring="$(GOOSE_DBSTRING)" -dir=$(GOOSE_MIGRATION_DIR) reset

.PHONY: up down reset