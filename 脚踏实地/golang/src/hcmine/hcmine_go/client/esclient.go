package client

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"hcmine/hcmine_go/config"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esapi"
)

type EsClient struct {
	client *elasticsearch.Client
}

func NewEsClient(config config.EsClientConfig) (*EsClient, error) {
	cfg := elasticsearch.Config{
		Addresses: config.Server,
		Username:  config.Username,
		Password:  config.Password,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
			},
		},
	}

	esClient, err := elasticsearch.NewClient(cfg)
	return &EsClient{client: esClient}, err
}

func (c *EsClient) send(dataType string, queue SendQueue) {
	go func() {
		for _, item := range queue.queue {
			data, err := json.Marshal(item)
			if err != nil {
				continue
			}
			req := esapi.IndexRequest{
				Index: strings.ToLower(dataType),
				Body:  strings.NewReader(string(data)),
			}
			res, err := req.Do(context.Background(), c.client)
			if err != nil {
				//logp.Warn("Error getting response when indexing data to elasticsearch: %s", err)
				continue
			}
			if res.IsError() {
				//logp.Warn("Error indexing data to elasticsearch: %s", res.Status())
			}
			res.Body.Close()
		}
	}()
}
