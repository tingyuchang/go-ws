package main

import (
	"context"
	"fmt"
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
}


func (c *KFKClient) Produce(data []byte) (err error){
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	err = c.conn.SetWriteDeadline(time.Now().Add(10*time.Second))
	if err != nil {
		log.Fatal("set dead line fail", err)
		return
	}
	_, err = c.conn.WriteMessages(
		kafka.Message{Value: data},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
		return
	}

	//if err := c.conn.Close(); err != nil {
	//	log.Fatal("failed to close writer:", err)
	//	return
	//}
	return err

}

func (c *KFKClient) Consume()(err error)  {
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	err = c.conn.SetReadDeadline(time.Now().Add(10*time.Second))
	if err != nil {
		log.Fatal("set dead line fail", err)
		return
	}
	batch := c.conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	if err = batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
		return err
	}

	if err = c.conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
		return err
	}
	return err
}

func InitKfaClient()*KFKClient {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("Init Kafaka client fail", err)
		return nil
	}
	client := &KFKClient{conn: conn}

	return client
}