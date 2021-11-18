package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
	kfkAddr = flag.String("kfkaddr", ":9092", "kafka service address")
	name = flag.String("name", "Default Name", "instance name")
	groupId = flag.String("gid", "0", "kafka group id")
)

func main() {
	flag.Parse()
	log.Printf("SYSTEM %s Start: \nAddress: %s\nKafka Address: %s GroupId: %s\n", *name, *addr, *kfkAddr, *groupId)
	hub := newHub()
	go hub.run()
	kfkClient := InitKfaClient(hub)

	go kfkClient.Consume()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, kfkClient, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}