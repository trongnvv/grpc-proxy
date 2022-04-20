package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-proxy/proto"
	"log"
	"net/http"
	"time"
)

const (
	defaultName = "world"
)

type client struct {
	testServiceClient pb.TestServiceClient
}

var (
	addr = flag.String("addr", "localhost:8000", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)
var a client

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	//defer conn.Close()
	a.testServiceClient = pb.NewTestServiceClient(conn)
	//sendRequest := func() {
	//	r, err := c.CallTest(context.Background(), &pb.Request{Name: "world"})
	//	if err != nil {
	//		log.Fatalf("could not greet: %v", err)
	//	}
	//	log.Printf("Greeting: %s", r.GetNumber())
	//}
	//
	//go func() {
	//	sendRequest()
	//}()
	//
	//go func() {
	//	time.Sleep(11 * time.Second)
	//
	//	sendRequest()
	//}()
	// Contact the server and print out its response.

	server()
}

func server() {

	// handle `/` route to `http.DefaultServeMux`
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		// get response headers
		header := res.Header()

		// set content type header
		header.Set("Content-Type", "application/json")

		// reset date header (inline call)
		res.Header().Set("Date", "01/01/2020")

		// set status header
		res.WriteHeader(http.StatusBadRequest) // http.StatusBadRequest == 400
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		r, err := a.testServiceClient.CallTest(ctx, &pb.Request{Name: *name})

		if err != nil {
			fmt.Println("log error **********************************", err)
		} else {
			log.Println("run")
		}
		log.Printf("Greeting: %s", r.GetNumber())

		log.Println("call")
		defer cancel()
		// respond with a JSON string
		fmt.Fprint(res, `{"status":"FAILURE"}`)
	})

	// listen and serve using `http.DefaultServeMux`
	log.Fatal(http.ListenAndServe(":9000", nil))

}
