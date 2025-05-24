package fluenceapi

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

// ListVmsV3 fetches all running VM instances (v3).
func (c *Client) ListVmsV3() ([]RunningInstanceV3, error) {
    url := fmt.Sprintf("%s/vms/v3", c.HostURL)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    body, err := c.doRequest(req)
    if err != nil {
        return nil, err
    }

    var vms []RunningInstanceV3
    if err := json.Unmarshal(body, &vms); err != nil {
        return nil, err
    }
    return vms, nil
}

// CreateVmV3 creates new VM(s) with the given configuration and constraints.
func (c *Client) CreateVmV3(reqBody CreateVmV3) ([]CreatedVm, error) {
    url := fmt.Sprintf("%s/vms/v3", c.HostURL)
    payload, err := json.Marshal(reqBody)
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

    var created []CreatedVm
    if err := json.Unmarshal(body, &created); err != nil {
        return nil, err
    }
    return created, nil
}

// RemoveVms removes VMs by their IDs.
func (c *Client) RemoveVms(vmIds []string) (*VmsRemoved, error) {
    url := fmt.Sprintf("%s/vms/v3", c.HostURL)
    payload, err := json.Marshal(RemoveVms{VmIds: vmIds})
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(payload))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", "application/json")

    body, err := c.doRequest(req)
    if err != nil {
        return nil, err
    }

    var removed VmsRemoved
    if err := json.Unmarshal(body, &removed); err != nil {
        return nil, err
    }
    return &removed, nil
}

// UpdateVms updates one or more VMs (patch).
func (c *Client) UpdateVms(updates []UpdateVm) error {
    url := fmt.Sprintf("%s/vms/v3", c.HostURL)
    payload, err := json.Marshal(UpdateVms{Updates: updates})
    if err != nil {
        return err
    }

    req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(payload))
    if err != nil {
        return err
    }
	req.Header.Set("Content-Type", "application/json")

    _, err = c.doRequest(req)
    return err
}

// EstimateDeposit estimates the deposit required for a VM configuration.
func (c *Client) EstimateDeposit(reqBody EstimateDepositRequestV3) (*EstimatedDepositV3DTO, error) {
    url := fmt.Sprintf("%s/vms/v3/estimate_deposit", c.HostURL)
    payload, err := json.Marshal(reqBody)
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

    var estimate EstimatedDepositV3DTO
    if err := json.Unmarshal(body, &estimate); err != nil {
        return nil, err
    }
    return &estimate, nil
}