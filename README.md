## TODO
- [x] Access and Refresh token impl
- [x] Split Auth and Payment service into 2 
	- How will common deps be shared?
- [ ] Dockerize auth and payment service
- [ ] Write tests for Auth
- [ ] Implement Google Auth
      
### Service-to-Service Communication
- [x] Gateway Service
- [x] Commons Package
- [x] Understanding Service-to-Service Communication
- [x] Leveraging gRPC for Efficient Intra-Service Communication
- [ ] Payload validation

### Service Discovery
- [ ] Service Discovery
- [ ] Creating the Registry

- [ ] Write Kubernetes for auth and payment service deployment
- [ ] Write Github action to automatically deploy
- [ ] Configure L7 Load Balancer to auth and payment service
      
- [ ] Implement Split service
- [ ] Write Tests
- [ ] Dockerize Split service
- [ ] Kubernetes for Split service
- [ ] CI for Split service
- [ ] Test and Deploy Both services
- [ ] Performance testing 
- [ ] Load testing
- [ ] Write UI 
- [ ] Integrate UI and services
- [ ] Deploy UI using vercel
- [ ] Implement Chat service between users for transaction

### Async Communication with Message Brokers
- [ ] Asynchronous Communication with Message Brokers
- [ ] Connecting to AMQP & Creating the Exchange

### The Payment Service
- [ ] Payment Notifications

### Reliability
- [ ] Implementing Retries & Dead Letter Queues

### Observability
- [ ] Intro to Distributed Tracing with OpenTelemetry
- [ ] Setup and Send Telemetry data from the gateway
- [ ] Telemetry Middleware
- [ ] Sending Telemetry though RabbitMQ
- [ ] Structured Logging

## Resources
- Architecture https://youtu.be/KdnxzgSNLTU
- [OAuth - GOTH](https://youtu.be/iHFQyd__2A0)/https://github.com/motiv-labs/janus
- Architecture [Building Microservices in Golang/Go](https://youtube.com/playlist?list=PL7yAAGMOat_Fn8sAXIk0WyBfK_sT1pohu&si=rQtGmmh-4K9mjN31)
- [The Anatomy of an API Gateway in Golang](https://hackernoon.com/the-anatomy-of-an-api-gateway-in-golang)
- [Build a gRPC API using Go and gRPC-gateway](https://www.koyeb.com/tutorials/build-a-grpc-api-using-go-and-grpc-gateway)

## UI
https://www.figma.com/design/cNV9lPWDFrHplMdr6WXBP9/Splitgo?node-id=0-1&t=7AnvRNrIEIlJks8n-1
