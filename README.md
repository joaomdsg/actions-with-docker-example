# actions-with-docker-example
An example github actions pipeline using docker that builds, tests and pushes a go application container image to github container registry (ghcr.io)

## Go Application
The application is a very simple go rest api that provides a singe endpoint `/healthcheck`. `GET`requests to this endpoint return this JSON message: 
```json
{
    "status": "healthy"
}
```

## CI/CD pipeline
The pipeline configuration consists of a yml file in `./.github/workflows/ci.yml`. Tis file is automatically picked up by github and executes the tasks below on every new git push.
