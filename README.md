# actions-with-docker-example
An example github actions pipeline using docker that builds, tests and pushes a go application container image to github container registry (ghcr.io)

## Go Application
The application is a Go REST API that offers a single endpoint, `/healthcheck`. When a `GET` request is sent to this endpoint, the API responds with a JSON message: 
```json
{
    "status": "healthy"
}
```

## CI/CD pipeline
The pipeline configuration consists of a yml file in `.github/workflows/ci.yml`. Tis file is automatically picked up by github and executes the tasks below on every new git push.

To push a new image to the GitHub registry, the pipeline must first authenticate itself by searching for a secret named `GHCR_TOKEN`. This secret should contain a personal access token for GitHub with write permission to packages. If the token does not exist, it must be generated and manually added as a secret in the repository settings.

If the pipeline succeeds, a new docker image sould be avaliable in [github packages](https://github.com/users/joaomdsg/packages).