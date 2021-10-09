# Svr-Cli

This is a test project to mimic some Hashicorp products (nomad, waypoint, etc) which have the server and cli in a single executable

## Build app

```bash
go build .
```

## Run the Server

```bash
srv-cli server
```

this will run the server on port :4000, ready for requests

## Make requests

With making requests, there are two options

- use the cli
- use the browser.

### Requests with CLI

In a separate terminal run the following. Remember the server should be running.

```bash
srv-cli up --{stop | s} 20
srv-cli down --{down | d} 20
```
