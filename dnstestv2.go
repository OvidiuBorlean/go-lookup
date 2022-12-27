package main

import (
        "net"
        "fmt"
        "os"
        "time"
        "strings"
        "strconv"
        "context"
)

var host string = ""

func lookUp(host string) bool {
   start := time.Now()
   ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Millisecond)
   //ips, err := net.LookupIP(host)
   defer cancel() // important to avoid a resource leak
   var r net.Resolver
   ips, err := r.LookupIP(ctx, "ip4",host)
   if err != nil {
      fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
      //os.Exit(1)
      f, err := os.OpenFile("dnsreq.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
      if err != nil {
         fmt.Println(err)
      }
      defer f.Close()
      timeString := start.Format("2006-01-02 15:04:05")
      str := []string{timeString, " ", host, "\n"}
      if _, err := f.WriteString(strings.Join(str,"")); err != nil {
          fmt.Println(err)
      }
    }

   for _, ip := range ips {
       fmt.Println(host, " - ",ip.String())
   }
   elapsed := time.Since(start)
   fmt.Println("Hostname resolved in ", elapsed)
return true
}

func main() {
  arg_len:= len(os.Args[1:])
  if arg_len == 0 {
     fmt.Println("Golang DNS Tester v1.0")
     fmt.Println("Usage: dnstest <delay> <host1> <host2> <host3> ...\n")
  } else {
        timeDelay, _ := strconv.Atoi(os.Args[1])
        fmt.Println("Testing for ", arg_len - 1, "hostnames\n")
        for {
        globalStart := time.Now()
        for _, r := range os.Args[2:] {
          lookUp(r)
        }
        globalElapsed := time.Since(globalStart)
        fmt.Println("Global Elapsed Time", globalElapsed, "\n")
        time.Sleep(time.Duration(timeDelay) * time.Second)
        }
   }
}
