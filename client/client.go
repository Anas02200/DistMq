package client

import "bytes"

const defaultScratchSize int64 = 1024 * 1024

type SimpleClient struct {
	addr []string
	buf  bytes.Buffer
}

//Creates new client for the server
func NewClient(addr []string) *SimpleClient {
	return &SimpleClient{
		addr: addr,
	}
}

//Sends messages to the server
func (c *SimpleClient) Send(msg []byte) error {
	_, err := c.buf.Write(msg)
	return err
}

//Recieve will wait for new messages or return an error if smth goes wrong
//A scratch buffer is used to read the data
func (c *SimpleClient) Recieve(scratch []byte) ([]byte, error) {
	if scratch == nil {
		scratch = make([]byte, defaultScratchSize)
	}
	n, err := c.buf.Read(scratch)
	if err != nil {
		return nil, err
	}
	return scratch[0:n], nil
}

// here when we read : if the size of the read buffer is smaller than the data written in the original buffer
// data will be lost when recieved
// example : with buffer size 10 : "100\n102\n103\n" we can only read "100\n102\10"
// solution => a buffer to store the unread part
