# go-grpc-example

[블로그](https://devjin-blog.com/)에 Golang으로 gRPC server 구축하기 시리즈를 작성했고, 이 repository에는 그 시리즈에 나오는 예제들을 모아놨습니다. 

## Contents

1. [Golang gRPC server 구축하기 (1) - gRPC란 무엇인가?](https://devjin-blog.com/golang-grpc-server-1/)
2. [Golang gRPC server 구축하기 (2)- gRPC server 파헤쳐보기](https://devjin-blog.com/golang-grpc-server-2/)
3. [Golang gRPC server 구축하기 (3) - RESTful하게 gRPC server와 통신하기](https://devjin-blog.com/golang-grpc-server-3/)
4. [Golang gRPC server 구축하기 (4) - gRPC middleware란? ](https://devjin-blog.com/golang-grpc-server-4/)


## Examples
- /simple - 간단한 gRPC server 예제
- /simple-user - protobuf로 User 서비스를 정의하고 이를 기반으로 구축한 gRPC server 예제
- /simple-client-server - protobuf로 Post 서비스를 새로 정의하고, User gRPC server와 통신하는 예제
- /simple-grpc-gateway - gRPC server와 RESTful하게 HTTP JSON으로 통신하는 예제
- /simple-middleware - middleware를 구현하는 간단한 예제
- /grpc-middlewares - go-grpc-middleware 패키지에서 지원하는 middleware들을 반영해본 예제