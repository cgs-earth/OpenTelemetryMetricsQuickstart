#!/bin/sh
# Copyright 2025 Lincoln Institute of Land Policy
# SPDX-License-Identifier: Apache-2.0


set -e

cd "$(dirname "$0")"
docker compose up -d 

open http://localhost:3000/drilldown

go run .
