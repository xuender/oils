package nets

import (
	"google.golang.org/protobuf/proto"
)

type Handler interface {
	GET() []byte
	POST(data []byte, isJSON bool) proto.Message
	Recover(any) proto.Message
}
