# port-scanning
Lightweight, Fast, and improoved port scanning module written out of golang, assembly is in testing.
<br>
<br>
# installs
```
go get -d github.com/ArkAngeL43/port-scanning

```
<br>
<br>
```go
package main

import (
        "flag"
        "fmt"
        "time"

        "github.com/ArkAngeL43/port-scanning/port"
)

var (
        flagTarget = flag.String("t", "", `Target host(s). Provide a single IP: "1.2.3.4", a CIDR block ">
        flagPort   = flag.Int("sp", 1, `Target Start port | Provide a port to start from EX -> 1`)
        flagPortn  = flag.Int("ep", 8090, `Target End Port   | Provide a port to stop ascanning at ex -> >
)

func main() {
        flag.Parse()
        t := time.Now()
        port.GetOpenPorts(*flagTarget, port.PortRange{Start: *flagPort, End: *flagPortn})
        fmt.Println("\033[31m\n[*] Script ended at -> ", time.Since(t))
}

```
<br>
