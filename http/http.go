package http

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const baseUrl = "https://api.getchange.io/api/v1/"

type client struct {
	h *http.Client
	k string
}

func (c *client) Run(method string, endpoint string, payload *bytes.Buffer, o interface{}) error {
	url := baseUrl + endpoint
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	req.Header.Add("authorization", fmt.Sprintf("Basic %s", c.k))
	req.Header.Set("content-type", "application/json")
	res, err := c.h.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&o); err != nil && err != io.EOF {
		return err
	}
	return nil
}

func getAuth() (string, error) {
	p := os.Getenv("CHA_PUBLIC_KEY")
	s := os.Getenv("CHA_SECRET_KEY")
	if s == "" || p == "" {
		return "", errors.New("CHA_PUBLIC_KEY or CHA_SECRET_KEY env variables not set")
	}
	k := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", p, s)))
	return k, nil
}

func New() *client {
	k, _ := getAuth()
	return &client{h: &http.Client{}, k: k}
}
