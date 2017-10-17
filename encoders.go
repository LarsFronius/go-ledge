package ledge

import "io"

var (
	rpcEncoderInstance = &rpcEncoder{}
	separator          = []byte{0x1e}
)

type rpcEncoder struct{}

func (r *rpcEncoder) Encode(writer io.Writer, data []byte) (int, error) {
	return writer.Write(append(data, separator...))
}
