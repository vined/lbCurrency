#!/usr/bin/env bash

echo Building backend
go build -o bin/lbCurrency

echo Building frontend
webpack

echo App is running on http://localhost:8001
./bin/lbCurrency

