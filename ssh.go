package fluenceapi

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

// ListSshKeys fetches all SSH keys.
func (c *Client) ListSshKeys() ([]SshKey, error) {
    url := fmt.Sprintf("%s/ssh_keys", c.HostURL)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    body, err := c.doRequest(req)
    if err != nil {
        return nil, err
    }

    var keys []SshKey
    if err := json.Unmarshal(body, &keys); err != nil {
        return nil, err
    }
    return keys, nil
}

// CreateSshKey creates a new SSH key.
func (c *Client) CreateSshKey(key AddSshKey) (*SshKey, error) {
    url := fmt.Sprintf("%s/ssh_keys", c.HostURL)
    payload, err := json.Marshal(key)
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", "application/json")

    body, err := c.doRequest(req)
    if err != nil {
        return nil, err
    }

    var created SshKey
    if err := json.Unmarshal(body, &created); err != nil {
        return nil, err
    }
    return &created, nil
}

// RemoveSshKey removes an SSH key by fingerprint.
func (c *Client) RemoveSshKey(fingerprint string) error {
    url := fmt.Sprintf("%s/ssh_keys", c.HostURL)
    payload, err := json.Marshal(RemoveSshKey{Fingerprint: fingerprint})
    if err != nil {
        return err
    }

    req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(payload))
    if err != nil {
        return err
    }
    req.Header.Set("Content-Type", "application/json")

    _, err = c.doRequest(req)
    return err
}