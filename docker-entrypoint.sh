#!/bin/bash

migrate -source /migrations -database $PG_CONN up
./main