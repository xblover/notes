package client

type LogClient struct{}

func NewLogClient() *LogClient {
	return &LogClient{}
}

func (c *LogClient) send(dataType string, queue SendQueue) {
	data, err := queue.Marshal()
	if err != nil {
		//logp.Err("Failed to send data due to marshal error :%s",err)
		return
	}
	//logp.Info("DataType=%s, Content=%s",dataType,string(data))
}
