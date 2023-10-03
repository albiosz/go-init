# What is the project about?

The goal of the repository is to provide a starting point for Go projects. It contains the initial setup for:
- a development environment in a dev container
- running the app with hot reload
- debugging the app in VS Code
- running tests

# Before starting your own project

## Name your project
- `.devcontainer/devcontainer.json` - replace `"name": "app"` with the name of your app
- `docker-compose.dev.yml` - replace `name: app` with the name of your app
- `go.mod` - change module name

## .env
Create an `.env` file in the root folder following the `.env.example` convention.

## (optional) SSH keys 
If you want to interact with other services (e.g. Github) via SSH, add your SSH private key to the .ssh folder in the project root and your SSH config.

# Features

## Running the app
Run your app with hot reload:
```
air
```

## Debugging in VS Code
Three debugging modes are included:

1. Attach to dev container - attaches to the dlv process running in the dev container when the app is started with air.
2. Launch package and debug - launches the application in debug mode
3. Attach to local process - attach to the process running on the host machine. This will not work in the dev container.

## Testing
Run all tests in terminal:
```
make test
```

Tests can be run separately in VS Code.


# Description of the repository
File names with brief descriptions:
- `.devcontainer/devcontainer.json` - start Dev Container setup with all the necessary tools to start coding in Go
- `.devcontainer/Dockerfile` - Dockerfile for the Dev Container that installs all the necessary tools to work only from the Dev Container
- `.ssh` - in that folder private SSH keys should be placed eg. to interact with Github repositories
- `docker-compose.dev.yml` - initial configuration of the service that are necessary to enter the Dev Container + database service
- `Makefile` - initially includes only the command to start docker-compose.dev.yml
- `.air.toml` - setup file for the air package that is used to hot reload the application