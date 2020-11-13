package main

import (  
    "fmt"
    "net"
)

func main() {  
    addr, err := net.LookupHost("golangbot123.com")
    if err != nil {
        if dnsErr, ok := err.(*net.DNSError); ok {
            if dnsErr.Timeout() {
                fmt.Println("operation timed out")
                return
            }
            if dnsErr.Temporary() {
                fmt.Println("temporary error")
                return
            }
            fmt.Println("Generic DNS error", err)
            return
        }
        fmt.Println("Generic error", err)
        return
    }
    fmt.Println(addr)
}