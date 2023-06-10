package main

import (
	"flag"
)

func main() {

	// client := client.New("http://localhost:3000")

	// price, err := client.FetcherPrice(context.Background(), "ET")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", price)

	listenAddr := flag.String("listenAddr", ":3000", "listen address the service is running")
	flag.Parse()

	svc := NewLogginService(NewMetricService(&priceFetcher{}))

	server := NewJsonAPiServer(*listenAddr, svc)

	server.Run()

}
