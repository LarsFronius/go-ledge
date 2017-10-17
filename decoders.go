package ledge

import (
	"bufio"
	"bytes"
	"io"
)

var (
	rpcDecoderInstance = &rpcDecoder{}
)

type rpcDecoder struct{}

func (r *rpcDecoder) Decode(bufReader *bufio.Reader) ([]byte, error) {
	scanner := bufio.NewScanner(bufReader)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if atEOF {
			return len(data), data, nil
		}
		i := bytes.LastIndex(data, separator)
		if i >= 0 {
			//we found the separator!
			return i + 1, data[0:i], nil
		} else {
			// did not find it, need to read more
			return 0, nil, nil
		}
		return
	})
	found := scanner.Scan()
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	if found {
		return scanner.Bytes(), nil
	}
	return scanner.Bytes(), io.EOF
}
