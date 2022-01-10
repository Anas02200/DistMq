package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"DistQueue.com/m/client"
)

const maxN = 10000000
const maxBufferSize = 1024 * 1024

func main() {

	log.Printf("Starting Client")
	client := client.NewClient([]string{"localhost"})
	want, err := sendData(client)
	if err != nil {
		log.Fatalf("sending error : %v", err)
	}
	got, err := recieveData(client)
	if err != nil {
		log.Fatalf("recieving error : %v", err)
	}
	if want != got {
		log.Fatalf("expected sum error , want : %d , got : , %d", want, got)
	}
	log.Printf("Test passed")
}

func sendData(client *client.SimpleClient) (sum int64, err error) {

	var b bytes.Buffer
	for i := 0; i < maxN; i++ {

		sum += int64(i)
		fmt.Fprintf(&b, "%d\n", i)
		//fmt.Fprint(&b,i)
		if b.Len() >= maxBufferSize {
			if err := client.Send(b.Bytes()); err != nil {
				log.Fatalf("sending error : %v", err)
				return 0, err
			}
			b.Reset()
		}
	}
	if b.Len() != 0 {
		if err := client.Send(b.Bytes()); err != nil {
			log.Fatalf("sending error : %v", err)
			return 0, err
		}
		b.Reset()

	}
	return sum, nil

}

func recieveData(client *client.SimpleClient) (sum int64, err error) {
	scratch := make([]byte, maxBufferSize)
	for {
		res, err := client.Recieve(scratch)
		if err == io.EOF {
			return sum, nil
		} else if err != nil {
			return 0, err
		}
		ints := strings.Split(string(res), "\n")
		for _, str := range ints {
			if str == "" {
				continue
			}
			i, err := strconv.Atoi(str)
			if err != nil {
				return 0, err
			}
			sum += int64(i)
		}

	}

}
