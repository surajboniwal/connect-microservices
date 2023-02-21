#!/bin/bash

make up
sleep 10
cd authentication && make migrate-up && cd ..