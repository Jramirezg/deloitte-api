# Deloitte Challenge API

API for Deloitte challenge

## Overview

API written in Golang that receives a string in text/plain and prints back a modified string

## Assumptions

* Reading the document with the assignment it was unclear if the result of Google Cloud should be Google© Cloud or Google Cloud© (as in the example). I noticed that in more cases it was mentioned Google© as a valid output and decided for it.
* A simplistic approach was followed. It is an academic exercise, additional error handling, testing and validations should be done in production.
* It was assumed that multiple occurrences of the different strings might happen in the same input

## Out of  scope

* SSL termination to be handled at a load balancer or Kubernetes ingress controller
* Authentication of the API via keys not done as it was not requested.
* No rate limiting done, since it it as academic exercise.
* Input validation

## Architecture

The code is comprised of:

* main.go containing the entrypoint of the application and the creation of the router. The router calls different functions depending of the API endpoint hit.
* API package containing the functions to be called when the API endpoints are hit. One endpoint for homepage and another for the /string URI accepting POST method. When /string is hit, the function un-marshalls the body of the request into a string and passes that string to the replace function in the replace package.
* Replace package with one function to perform the string replacement based on a string given as parameter.
* Simple testing functions provided at main level and package level for the replace function only.
* Dockerfile with a multi-stage build to produce a Docker image of the code.

## Build instructions

Execute the following commands to build the go binary

```bash
CGO_ENABLED=0 GOOS=linux go build -v -o /dapi 
```

Execute the following commands to build the docker image

```bash
docker build -t <REPO_IMAGENAME_AND_TAGS> . 
```

## Testing

```bash
curl --location --request POST '<URL>/string' \
--header 'Content-Type: text/plain' \
--data-raw '"We really like the new security features of Google Cloud and Amazon "'
```
