package client

import (
	"bytes"
	"net/http"

	"github.com/BoseCorp/pester"
)

type HttpClient struct {
	client *pester.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: pester.New(),
	}
}

//
func (c *HttpClient) send(dataType string, queue SendQueue) {
	url := ""
	data, err := queue.Marshal()
	if err != nil {
		//logp.Err("Failed to send data due to marshal error: %s", err)
		return
	}
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))

	response, err := c.client.Do(request)
	if err != nil {
		//logp.Err("Error happened when sending data: DataTypes=%s, err=%s",dataType,err)
	}
	if response.StatusCode >= 400 {
		//logp.Err("StatusCode=%d which is >400 when sending data: DataType=%s",response.StatusCode,dataType)
	}

}
