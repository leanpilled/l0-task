package nats

import (
	"encoding/json"
	"wb-0/entity"
	"wb-0/repository"
	"wb-0/repository/cache"

	"github.com/nats-io/stan.go"
)

type NatsSub struct {
	client stan.Conn
}

func NewNatsSub(ClusterID string, ClientID string) *NatsSub {
	sc, err := stan.Connect(ClusterID, ClientID)
	if err != nil {
		panic("error connection nats-streaming")
	}
	return &NatsSub{client: sc}
}

func (n *NatsSub) Run(repo repository.OrderRepo, c *cache.Cache) {
	n.client.Subscribe("foo", func(m *stan.Msg) {
		var order entity.Order
		json.Unmarshal(m.Data, &order)
		order, err := repo.CreateOrder(order)
		if err != nil {
			panic(err)
		}
		c.Set(order)
	}, stan.DurableName("my-durable"))
}

func (n *NatsSub) ShutDown() error {
	return n.client.Close()
}
