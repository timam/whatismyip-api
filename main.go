package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type IPAddr struct {
	ClientIPAddr string `json:"client_ip_addr"`
}

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
	jsonIP, _ := json.Marshal(IPAddr{
		ClientIPAddr: ipAddr,
	})

	log.Printf("Request from :%s\n", ipAddr)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonIP)
}

func main()  {
	portNumber :=  applicationPort()
	fmt.Println(fmt.Sprintf("Application started at port %s", portNumber))
	http.HandleFunc("/", UserInternetProtocolDetails)
	log.Fatal(http.ListenAndServe(":" +portNumber , nil))
}
