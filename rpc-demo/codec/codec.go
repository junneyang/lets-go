package codec

import "io"

const (
	MagicNumber = 0x3bef5c
	GobType     = "application/gob"
	JsonType    = "application/json"
)

type CodecType string

type Option struct {
	MagicNumber int
	CodecType   CodecType
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodecType:   GobType,
}

type Header struct {
	ServiceMethod string
	Seq           uint64
	Error         string
}

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type CodecFunc func(io.ReadWriteCloser) Codec

var CodeCMap map[CodecType]CodecFunc

func init() {
	CodeCMap = make(map[CodecType]CodecFunc)
	CodeCMap[GobType] = NewGobCodec
}
