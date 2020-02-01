package dto 

type Payload struct {
	ExchangeName string   `json:"exchangeName"`
	Topic string   `json:"topic"`
	Data map[string]interface{} `json:"data"`
}