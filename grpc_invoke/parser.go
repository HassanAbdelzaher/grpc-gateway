package grpc_invoke

import (
	"bufio"
	"io"

	"github.com/fullstorydev/grpcurl"
	"github.com/golang/protobuf/proto"
)

type protoBuffRequestParser struct {
	r            *bufio.Reader
	err          error
	requestCount int
}

// NewTextRequestParser returns a RequestParser that reads data in the protobuf
// text format from the given reader.
//
// Input data that contains more than one message should include an ASCII
// 'Record Separator' character (0x1E) between each message.
//
// Empty text is a valid text format and represents an empty message. So if the
// given reader has no data, the returned parser will yield an empty message
// for the first call to Next and then return io.EOF thereafter. This also means
// that if the input data ends with a record separator, then a final empty
// message will be parsed *after* the separator.
func NewProtobuffRequestParser(in io.Reader) grpcurl.RequestParser {
	return &protoBuffRequestParser{r: bufio.NewReader(in)}
}

func (f *protoBuffRequestParser) Next(msg proto.Message) error {
	if f.err != nil {
		return f.err
	}

	var b []byte
	_, f.err = f.r.Read(b)
	if f.err != nil && f.err != io.EOF {
		return f.err
	}

	f.requestCount++

	return proto.Unmarshal(b, msg)
}

func (f *protoBuffRequestParser) NumRequests() int {
	return f.requestCount
}
