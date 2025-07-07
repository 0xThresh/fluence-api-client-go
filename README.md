# Fluence Golang API Client

A Go client package that provides connections to the [Fluence API](https://api.fluence.dev/docs) (v0.8.64).

Using this module, the Fluence provider establishes a new client and sends HTTP(s) requests to the Fluence API to perform CRUD operations. It also handles data mapping from user's inputs to `models.go`. The Fluence URL defaults to `https://api.fluence.dev/`, and is set in the `client.go` file.

## Features

This client supports all Fluence API endpoints including:

- **SSH Key Management**: List, create, and remove SSH keys
- **Virtual Machine Management**: Create, list, update, remove, and get status of VMs
- **Marketplace Operations**: Browse available configurations, countries, hardware, and offers
- **Resource Management**: Estimate deposits and get default OS images

## Usage
Using the client requires a valid Fluence API key. Keys can be created in the [API keys](https://console.fluence.network/settings/api-keys) page in the Fluence console.

Once you have an API key, set the `FLUENCE_API_KEY` environment variable:
`export FLUENCE_API_KEY=<your key>`

Finally, you can run `go run test/test.go` to validate your authentication, which should return any SSH keys you've uploaded, running VMs, available configurations, and other marketplace data. 