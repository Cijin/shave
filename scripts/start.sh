#!/bin/sh
set -e
./migration
exec ./server
