# Onesignal-srv

Onesignal-srv is a go microservice used to send & track push notifications using Onesignal API. Built for Golang micro.mu

## Getting started

1. Install Consul

	Consul is the default registry/discovery for go-micro apps. It's however pluggable.
	[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

2. Run Consul
	```
	$ consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul
	```
	
3. Download and start the service
	```shell
	go get github.com/suicidegang/onesignal-srv
	onesignal-srv --api_key="xxxx-xxx-xxx" --app_id="xxx_xxx"
	```

## The API
Onesignal server implements the following RPC Methods

Onesignal
- SendAll
- Send

### Onesignal.Send
```shell
$ micro query sg.micro.srv.mail Mailing.SendTemplate '{
	"message": "Hola mundo!",
	"variables": {},
	"ids": ["4cd0e86b-081f-4ea6-a0fa-68026f9612e8"]
}'
{
	"messageID": "bdff86f7-75d4-46b7-88f3-93b53a5b8448"
}
```
