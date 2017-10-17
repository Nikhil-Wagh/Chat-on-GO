# Chat-on-a-GO
A demonstration of TCP server in GO.
### Usage
1. Open 3 terminals one server and two clients
2. In first terminal (server) type the following command
```go
go run Server.go
```
3. In second and third terminals (clients) type the following command
```go
telnet localhost 8080
```
4. Type two different names in each of the terminal and send data only after registering both the clients.
+ send message in the format `<From> <To> <Message>`
