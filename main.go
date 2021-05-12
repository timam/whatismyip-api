package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)



func applicationPort() string {
	if port, ok := os.LookupEnv("PORT"); ok {
		return port
	}
	return "8080"
}



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
	portNumber :=  applicationPort()
	fmt.Println(fmt.Sprintf("Application started at port %s", portNumber))
	http.HandleFunc("/", UserInternetProtocolDetails)
	log.Fatal(http.ListenAndServe(":" +portNumber , nil))
}
