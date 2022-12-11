package main
 
import (
  "fmt"
  "net"
  "log"
)
 
 
func main() {
 
  ip_list, err := net.LookupIP("techieindoor.com")
 
  if err == nil {
 
      fmt.Println("IPs are: ")
 
      for _, ip := range ip_list {
 
 
          fmt.Println(ip)
 
      }
 
  } else {
 
      log.Fatal("IP lookup failed. Error is: ", err)
 
  }
}
