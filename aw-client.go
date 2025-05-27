package aw_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ActivityWatchClientConfig struct {
	Protocol string
	Hostname string
	Port     string
}

func (awconf * ActivityWatchClientConfig) Address() string {
	return fmt.Sprintf("%s://{%s}:{%s}", awconf.Protocol, awconf.Hostname, awconf.Port)
}

type ActivityWatchClient struct {
	Testing        bool
	ClientName     string
	ClientHostname string
	ServerAddress  string
	Config         ActivityWatchClientConfig
}

func (awc *ActivityWatchClient) Init() error {
	var err error
	awc.ClientHostname, err = os.Hostname()
	awc.ServerAddress = awc.Config.Address()
	return err
}

func (awc *ActivityWatchClient) url(endpoint string) string {
  return fmt.Sprintf("%s/api/0/%s", awc.ServerAddress, endpoint)
}

func (awc *ActivityWatchClient) post(endpoint string, data any) (*http.Response,error) {
  jsonData, err := json.Marshal(data)
  if err != nil {
    return nil, err
  }
  req, err := http.NewRequest("POST", awc.url(endpoint), bytes.NewBuffer(jsonData))
  if err != nil {
    return nil, err
  }
  req.Header.Set("Content-Type", "application/json")
  return http.DefaultClient.Do(req)
}

func (awc *ActivityWatchClient) InsertEvent(bucket_id string, event Event) {
  endpoint := fmt.Sprintf("buckets/%s/events", bucket_id)
  awc.post(endpoint, event)
  
}
