## TODO
- [x] Access and Refresh token impl
- [x] Split Auth and Payment service into 2 
	- How will common deps be shared?
- [ ] Dockerize auth and payment service
- [ ] Write tests for Auth

### Service-to-Service Communication
- [ ] Gateway Service
- [ ] Commons Package
- [ ] Understanding Service-to-Service Communication
- [ ] Leveraging gRPC for Efficient Intra-Service Communication
- [ ] Payload validation

### Service Discovery
- [ ] Service Discovery
- [ ] Creating the Registry
- [ ] Adopting Service Discovery with Consul

- [ ] Write Kubernetes for auth and payment service deployment - use EKS
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
- OAuth https://youtu.be/iHFQyd__2A0
- Architecture https://youtube.com/playlist?list=PL7yAAGMOat_Fn8sAXIk0WyBfK_sT1pohu&si=rQtGmmh-4K9mjN31