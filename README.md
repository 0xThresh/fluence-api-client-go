# Fluence Golang API Client

A Go client package that provides connections to the [Fluence API](https://api.fluence.dev/docs).

Using this module, the Fluence provider establishes a new client and sends HTTP(s) requests to the Fluence API to perform CRUD operations. It also handles data mapping from user's inputs to `models.go`. The Fluence URL defaults to `https://api.fluence.dev/`, and is set in the `client.go` file.

## Usage
Using the client requires a valid Fluence API key. Keys can be created in the [API keys](https://console.fluence.network/settings/api-keys) page in the Fluence console.

Once you have an API key, set the `FLUENCE_API_KEY` environment variable:
`export FLUENCE_API_KEY=<your key>`

Finally, you can run `go run test/test.go` to validate your authentication, which should return any SSH keys you've uploaded, and any running VMs in your account. 