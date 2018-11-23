# Simple go micro-service application

## To parse the data.csv file, run make command, an output is 5 (limit) users from the database
```shell
$ make start
```


## What is not done
* creates a new record in a database, or updates the existing one(was problem with new mongodb driver)
* Support for a TERM or KILL signal
* code quality checkers
* tests

## Unused Dockerfile and go mod files
Spent a lot of time on gRPC tutorials and configure dockerfile for micro-service architecture, which does not used
