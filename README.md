# steamcmd-scrapper
Go based steamcmd scrapper for watching for build updates for games

# Container Usage
The Docker container built from the provided Dockerfile and entrypoint.go script performs the following steps:

- Reads the contents of the file applications.yaml.
- Unmarshals the file contents into a map.
- Iterates over the map, making an HTTP GET request to https://api.steamcmd.net/v1/info/ with the appId from the map as a parameter.
- Makes an HTTP POST request to the URL specified by the WEBHOOK_URL environment variable with the response from the GET request as the body, setting the Authorization header to Bearer and the value of the WEBHOOK_TOKEN environment variable.
- Reads the response from the POST request and logs both the POST request and the response.

# Prerequisites
- Docker installed
- Access to the Docker CLI
- A WEBHOOK_URL and WEBHOOK_TOKEN environment variable set to consume the json payload. [Argo-Events](https://github.com/argoproj/argo-events)

# Steps to run the container
Run the Docker container using the command `docker run -e WEBHOOK_URL=<webhook-url> -e WEBHOOK_TOKEN=<webhook-token> -v $(pwd)/applications.yaml:/app/applications.yaml stephenshelton/steamcmd-scrapper:0.0.1`
