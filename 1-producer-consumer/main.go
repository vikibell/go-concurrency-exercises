//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, dataChanel chan<- *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(dataChanel)
			return
		}

		dataChanel <- tweet
	}
}

func consumer(dataChanel <-chan *Tweet) {
	for t := range dataChanel {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	dataChannel := make(chan *Tweet)

	go producer(stream, dataChannel)
	go consumer(dataChannel)

	fmt.Printf("Process took %s\n", time.Since(start))
}
