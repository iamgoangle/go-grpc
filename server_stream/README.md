# Server streaming RPC

A server-streaming RPC is similar to our simple example, except the server sends back a stream of responses after getting the client’s request message. After sending back all its responses, the server’s status details (status code and optional status message) and optional trailing metadata are sent back to complete on the server side. The client completes once it has all the server’s responses.
