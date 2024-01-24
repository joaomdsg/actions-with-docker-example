# actions-with-docker-example
An example github actions pipeline using docker that builds, tests and pushes a go application container image to github container registry (ghcr.io)

## Go Application
The application is a Go REST API that offers a single endpoint, `/healthcheck`. When a `GET` request is sent to this endpoint, the API responds with a JSON message: 
```json
{ "status": "healthy" }
```

## CI/CD pipeline
The pipeline configuration is defined in `.github/workflows/ci.yml`. This file is automatically detected by GitHub and triggers the defined tasks on every commit. 

GitHub Actions consumes computing power, storage, and bandwidth resources. These are taken from the account's limited free resources that vary by account type and are renewed monthly. If exceeded, more resources must be purchased for the pipeline to keep running. 

To save resources, include "`no-ci`" in the commit message when making minor changes like updating the readme file. This prevents the pipeline from running unnecessarily.

If the pipeline succeeds, a new docker image sould be avaliable in the [github packages page](https://github.com/users/joaomdsg/packages).
