# Unary

Unary RPCs where the client sends a single request to the server and gets a single response back, just like a normal function call.

#### Unary RPC

First let’s look at the simplest type of RPC, where the client sends a single request and gets back a single response.

Once the client calls the method on the stub/client object, the server is notified that the RPC has been invoked with the client’s metadata for this call, the method name, and the specified deadline if applicable.
The server can then either send back its own initial metadata (which must be sent before any response) straight away, or wait for the client’s request message - which happens first is application-specific.
Once the server has the client’s request message, it does whatever work is necessary to create and populate its response. The response is then returned (if successful) to the client together with status details (status code and optional status message) and optional trailing metadata.
If the status is OK, the client then gets the response, which completes the call on the client side.

#### Getting Started

##### Define proto

Proto3 reference <https://developers.google.com/protocol-buffers/docs/proto3>

```protobuf
syntax = "proto3";
package v1;

import "google/protobuf/timestamp.proto";

// define type
message ToDo {
  int64 id = 1;
  string title = 2;
  string descripton = 3;
  google.protobuf.Timestamp reminder = 4;
}

// define request
message CreateRequest {
  string api = 1;
  ToDo todo = 2;
}

// define response
message CreateResponse {
  string api = 1;
  ToDo todo = 2;
}

// define service
service TodoService {
  rpc Create(CreateRequest) returns (CreateResponse) {}
}

```

##### Compile stub

```sh
./compile
```

##### Server

```sh
go run server.go
```

##### Client

```sh
go run client.go
```

