package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/apex/gateway"
)

// func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
// 	return &events.APIGatewayProxyResponse{
// 		StatusCode: 200,
// 		Headers: map[string]string{
// 			"Content-Type": "text",
// 		},
// 		Body: "Hello World",
// 	}, nil
// }

// func main() {
// 	// Make the handler available for Remote Procedure Call by AWS Lambda
// 	lambda.Start(handler)
// }

func handler(w http.ResponseWriter, r *http.Request) {
	// example retrieving values from the api gateway proxy request context.
	requestContext, ok := gateway.RequestContext(r.Context())
	if !ok || requestContext.Authorizer["sub"] == nil {
		fmt.Fprint(w, "Hello World from Go")
		return
	}
}

func main() {
	port := flag.Int("port", -1, "specify a port to use http rather than AWS Lambda")
	flag.Parse()
	listener := gateway.ListenAndServe
	portStr := "n/a"
	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		listener = http.ListenAndServe
	}
	http.HandleFunc("/api/helloworld", handler)

	log.Printf("Server running on port %s...", portStr)
	log.Fatal(listener(portStr, nil))
}
