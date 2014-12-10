package main

import (
	"flag"
	"fmt"
	"github.com/kolo/xmlrpc"
	"io"
	"log"
	"os"
	"strings"
	"time"
	"encoding/csv"
)

func main() {
	flag.Parse() // Scan the arguments list

	csvFiled, err := os.Open("domains.csv")
	defer csvFiled.Close()
	if err != nil {
		panic(err)
	}

	csvReaderd := csv.NewReader(csvFiled)
	csvReaderd.TrailingComma = true

	var domainsarr []string

	for {
		fields, err := csvReaderd.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		if !strings.HasPrefix(fields[0], "#") {

			domainsarr = append(domainsarr, fields[0])

		}

	}

	client, err := xmlrpc.NewClient("http://rpc.pingomatic.com", nil)
	//	client, err := xmlrpc.NewClient("http://blogsearch.google.fi/ping/RPC2", nil)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	
	client2, err := xmlrpc.NewClient("http://blogsearch.google.fi/ping/RPC2", nil)
	if err != nil {
		log.Fatal("arith error:", err)
	}	
	
	result := struct {
		Message string `xmlrpc:"message"`
	}{}

	for _, domain := range domainsarr {
		
		client.Call("weblogUpdates.extendedPing", []interface{}{"http://" + domain + "/rss.xml", "http://" + domain + "/rss.xml"}, &result)
		fmt.Println("client",domain, result)
		client2.Call("weblogUpdates", []interface{}{"http://" + domain + "/rss.xml", "http://" + domain + "/rss.xml"}, &result)
		fmt.Println("client2",domain, result)
		time.Sleep(60000 * 6 * time.Millisecond)


	}

}
