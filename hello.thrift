struct HelloRequest {
  1: optional string message;
}

struct HelloResponse {
  1: optional string message;
}

service Hello {
  HelloResponse hello(1: HelloRequest request);
}
