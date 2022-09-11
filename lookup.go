package main

import (
	"net"
	"fmt"
	"os"
	"time"
)


var host string = ""

func lookUp(host string) bool {
   start := time.Now()
   ips, err := net.LookupIP(host)
   if err != nil {
      fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
      os.Exit(1)
    }

   for _, ip := range ips {
       fmt.Println(host, " ",ip.String())
   }
   elapsed := time.Since(start)
   fmt.Println(elapsed)
return true
}

func main() {
   globalStart := time.Now()
   hostList := []string {
 	"www.google.ro",
	"www.apple.com",
	"www.microsoft.com",
	"mcr.microsoft.com",
	"ubuntu.com",
	"tigera.io",
         "udemy.com",
   }
   for _, r := range hostList {
      lookUp(r)
    }
   globalElapsed := time.Since(globalStart)
   fmt.Println("Global Elapsed Time", globalElapsed)

}
