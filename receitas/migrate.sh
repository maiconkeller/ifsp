#!/bin/bash

docker run -v /home/keller/src/receitas/migrations/:/migrations --network host migrate/migrate \
  -path=/migrations/ -database postgres://postgres:postgres@localhost:7557/receitas_db?sslmode=disable $1 $2