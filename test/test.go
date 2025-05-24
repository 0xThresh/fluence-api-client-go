package main

import (
    "fmt"
    "os"

    fluenceapi "github.com/0xthresh/fluence-api-client-go"
)

func main() {
    apiKey := os.Getenv("FLUENCE_API_KEY")
    if apiKey == "" {
        fmt.Println("Please set the FLUENCE_API_KEY environment variable.")
        return
    }

    client, err := fluenceapi.NewClient(nil, &apiKey)
    if err != nil {
        fmt.Printf("Error creating client: %v\n", err)
        return
    }

    // List SSH keys
    sshKeys, err := client.ListSshKeys()
    if err != nil {
        fmt.Printf("Error listing SSH keys: %v\n", err)
    } else {
        fmt.Printf("SSH Keys: %+v\n", sshKeys)
    }

    // List running VMs
    vms, err := client.ListVmsV3()
    if err != nil {
        fmt.Printf("Error listing VMs: %v\n", err)
    } else {
        fmt.Printf("Running VMs: %+v\n", vms)
    }
}