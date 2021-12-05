module github.com/yedf/dtmdriver-clients

go 1.15

require (
	github.com/kevwan/gozero-dtm v0.0.0-20211201122454-5007fb2b35a1
	github.com/tal-tech/go-zero v1.2.4
	github.com/yedf/dtmcli v1.6.0
	github.com/yedf/dtmdriver v0.0.0-20211203060147-29426c663b6e
	github.com/yedf/dtmdriver-gozero v0.0.0-20211202141526-9fe6083e31d2
	github.com/yedf/dtmdriver-protocol1 v0.0.0-20211205104407-f50a46978fb3
	github.com/yedf/dtmgrpc v1.6.0
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => google.golang.org/grpc v1.42.0
