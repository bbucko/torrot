package tracker

import (
	"log"
	"net/http"
)

type rqEvent string

const (
	started rqEvent = "started"
	stopped  rqEvent = "stopped"
	completed rqEvent = "completed"
)


//Request represents request made to the tracker
type Request struct {

	infoHash   string
	peerId     string
	port       uint8
	uploaded   uint64
	downloaded uint64
	left       int
	compact    bool
	noPeerId   bool
	event      rqEvent

	ip         string
	numWant    int
	key        int
	trackerId  int

}

type Response struct {
	warningMessage string
	interval       int64
	minInterval    int64
	trackerId      string
	complete       int64
	incomplete     int64


}

type FailureResponse struct {
	failureReason string
}

func sendReq() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
