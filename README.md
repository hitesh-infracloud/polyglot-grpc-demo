# Polyglot gRPC demo

Trying out gRPC with NodeJs gRPC client and Golang gRPC server

Goal : To create a simple gRPC server and client, where the client sends a single message while invoking the remote method, and server sends a single response. This is one of the types communications between a gRPC client and server - Unary RPC

Example taken to achieve the goal is : gRPC client sends two numbers (in a single message) to gRPC server and server sends the sum of the two numbers as response.

Below are the steps (at a high level) that I followed to code the demo

## Steps to install protoc and generate 
1. `go install github.com/golang/protobuf/protoc-gen-go@latest` -> installs protoc globally. The binary resides either at $HOME/go or /usr/local/go
2. `go install github.com/golang/protobuf/protoc-gen-go-grpc@latest` -> same locations as above
3. Either follow #4 or #5
4. To generate grpc and go code from proto files (separately):
   1. protoc --proto_path=<directory to look for proto file> --go_out=<output dir> <filepath> -> this will generate go code
   2. for this case (first cd into go-grpc-server): `protoc --proto_path=addition --go_out=addition addition/addition.go`
   3. protoc --go-grpc_out=<output dir> <filepath>
   4. for this case (first cd into go-grpc-server): `protoc --go-grpc_out==addition addition/addition.go`
5. I personally prefer using buf instead of generating manually:
   1. Still have todo #1 and #2
   2. Install buf: `go install github.com/bufbuild/buf/cmd/buf@v1.7.0`
   3. Install task: `sudo snap install task --classic`
   4. Install gogofaster plugin: `go install github.com/gogo/protobuf/protoc-gen-gogofaster@latest`
   5. Create buf.gen.yaml, buf.yaml and Taskfile.yaml with given contents
   6. Generate: `task generate`

## Steps to install yarn:
1. sudo apt remove cmdtest (optional, do if it causes trouble)
2. sudo apt remove yarn
3. curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
4. echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list
5. sudo apt-get update
6. sudo apt-get install yarn

## Running the server and client

Note : You need golang, node and yarn installed for this. And the server runs at port `10000`. It's a hard coded port number in the code

To run the server, use this command

`$ go run grpc_server.go`

To run the client, first get the dependencies using this command

`$ yarn install`

Then to run the client, use this command (you can provide any two integer numbers as arguments)

`$ node grpc_client.js 10 20`

and you will get the output in the client side as

`The sum of 10 and 20 is 30`

## Steps to create python client
1. Install python3
2. Install pip
3. Install grpcio-tools: `pip install grpcio-tools`
4. Generate python stub: `python3 -m grpc_tools.protoc -I addition --python_out=addition --grpc_python_out=addition addition/addition.proto`
5. Code reference: python-client/python_client.py
6. Run python client: `cd python_client && python3 python_client.py`

## Certificate generation for secured server
1. The certificate generation has 3 parts:
   1. First, generate CA's private key and its self-signed certificate
   2. Second, create server's private key and CSR
   3. Third, use CA's private key to sign the server's CSR to get server certificate
2. Usage of certificate on server side:
   1. Load server certificate and server private key
   2. Create transport credentials out of it and secure the server
3. Usage of certificate on client side
   1. Load CA's certificate.
   2. Create transport credentials out of it and send requests to server
4. Since server's certificate was signed using CA's private key, server will be able to recognize the certificate sent by client via request.