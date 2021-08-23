#!/bin/bash

function cleanup {
    kill "$BASE_PID"
    kill "$HISTORY_PID"
    kill "$REVIEWS_PID"
}
trap cleanup EXIT

go build -o /tmp/srv-base ./base
go build -o /tmp/srv-history ./history
go build -o /tmp/srv-gateway ./gateway

/tmp/srv-base &
BASE_PID=$!

/tmp/srv-history &
HISTORY_PID=$!

sleep 1

/tmp/srv-gateway
