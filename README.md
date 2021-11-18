# go-ws
this repo is Go Websocket app + Kafka client

![Imgur](https://imgur.com/xETCAa2.jpg)

in above system desgin diagram, 

this repo respones to create websocket service and connect to message queue using pub/sub model,

because we distribute conenctions for multiple websocket service instances,

each instance could boardcast msg to all clients they were connecting, 

but other clients connected to other instances could't do that.

so we need a fast, low latency way to share messages,

in this demo, message queue responses to finish this job.

For personal reason, we select Kafka, but actually we think redis is a better choice in this case.


## In Kafka

we know kafka could work like both queue and pub/sub model,

so we distinguish each instance for separated groups

when a message comes in, instance would push to fixed topic,

thus, other group subscribled same topic will receive message immediately

thst's what we want to do.


## TO-DO

each instance all subscribe same topic, so even the msg was sent by this instance

we still received it, we should skip this duplicate msg, it's not hard, just leave it for future.
 

## Reference 
most websocket related code is from gorilla/websocket example for chat room, appreciate!
- https://github.com/gorilla/websocket
- https://github.com/segmentio/kafka-go
- https://github.com/Eficode/wait-for
