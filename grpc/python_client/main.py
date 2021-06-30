import os
import sys
import grpc
work_dir = os.path.dirname(os.path.abspath(__file__))
sys.path.append(os.path.join(work_dir, '..'))
from hello import hello_pb2
from hello import hello_pb2_grpc

def main():
    with grpc.insecure_channel('localhost:8080') as channel:
        stub = hello_pb2_grpc.GreeterStub(channel)
        resp = stub.SayHello(hello_pb2.HelloRequest(name = "zhangsan"))
        print("resp=>{}".format(resp.message))

if __name__ == '__main__':
    main()
