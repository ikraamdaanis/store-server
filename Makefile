#!make
include .env.local

dev:
	go run cmd/api/main.go