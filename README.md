# Simple go micro-service application

## To parse the data.csv file, run make command
```shell
$ make start
```
## an output is 5 (limit) users from the database


## What is not done
* updates the existing user(was problem with new mongodb driver)
* support for a TERM or KILL signal
* code quality checkers
* tests

## Unused Dockerfile and go mod files
Spent a lot of time on gRPC tutorials and configure dockerfile for micro-service architecture, which does not used
