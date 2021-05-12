package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func UserInternetProtocolDetails(w http.ResponseWriter, r *http.Request)  {
	ipAddr := ReadUserIP(r)
	log.Printf("Request from :%s\n", ipAddr)
	_, _ = fmt.Fprintf(w, ipAddr)
}

func main()  {
	fmt.Println(fmt.Sprintf("Application started at port %s", portNumber))

	http.HandleFunc("/", UserInternetProtocolDetails)
	_ = http.ListenAndServe(portNumber, nil)
}
