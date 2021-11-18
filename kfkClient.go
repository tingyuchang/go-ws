package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

var (
	topic = "message_channel"
	partition = 0
)

type KFKClient struct {
	// The Kafaka connection.
	conn *kafka.Conn
	// Websocket Hub
	hub *Hub
}


func (c *KFKClient) Produce(data []byte) (err error){
	_, err = c.conn.WriteMessages(
		kafka.Message{Value: data},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
		return
	}

	return err

}

func (c *KFKClient) Consume()  {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		GroupID:   *groupId,
		Topic:     topic,
		MinBytes:  1e3, // 1KB
		MaxBytes:  10e6, // 10MB
		MaxWait: 500*time.Millisecond,
		//PartitionWatchInterval: 1000 * time.Microsecond,
		//HeartbeatInterval: 100 * time.Microsecond,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		c.hub.broadcast <- m.Value
		//fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func InitKfaClient(hub *Hub)*KFKClient {
	conn, err := kafka.DialLeader(context.Background(), "tcp", *kfkAddr, topic, partition)
	if err != nil {
		log.Fatal("Init Kafka client fail", err)
		return nil
	}
	client := &KFKClient{conn: conn, hub: hub}

	return client
}