#!/bin/bash

http POST http://192.168.178.101:3000/api/robots/pump/commands/run_pump --raw='{ "id": 0 }'
