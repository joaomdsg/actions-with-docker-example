# actions-with-docker-example
An example github actions pipeline using docker that builds, tests and pushes a go application container image to github container registry (ghcr.io)

## Go Application
The application is a Go REST API that offers a single endpoint, `/healthcheck`. When a `GET` request is sent to this endpoint, the API responds with a JSON message: 
```json
{ "status": "healthy" }
```

## CI/CD pipeline
The pipeline configuration is defined in a YAML file located at `.github/workflows/ci.yml`. This file is automatically detected by GitHub and triggers the defined tasks every time a new Git push is made.

If the pipeline succeeds, a new docker image sould be avaliable in the [github packages page](https://github.com/users/joaomdsg/packages).