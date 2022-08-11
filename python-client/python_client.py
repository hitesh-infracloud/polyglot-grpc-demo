import grpc
import addition_pb2
import addition_pb2_grpc

def send_addition_request(stub):
    response = stub.Add(addition_pb2.AddRequest(number=10, anotherNumber=20))
    return response


def run():
    with grpc.insecure_channel('localhost:10000') as channel:
        stub = addition_pb2_grpc.AdditionStub(channel)
        additionResponse = send_addition_request(stub)
        if not additionResponse.sum:
            print("server returned incomplete response")
            return
        print("sum is %d" % additionResponse.sum)


if __name__ == '__main__':
    run()