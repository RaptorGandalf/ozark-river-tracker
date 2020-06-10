#!/bin/bash
read varname
migrate -source /migrations -database $PG_CONN up
./main