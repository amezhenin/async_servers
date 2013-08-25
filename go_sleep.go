
package main
import (
  "net/http"
//  "io"
  "runtime"
  "time"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
  //io.WriteString(w, "hello, world!\n")
  //ok := make(chan bool)
  //go func() {
    //fmt.Printf("Writer2: %v\n", w)
  time.Sleep(1 * time.Second)
  w.Write([]byte("Hello async"))
    //io.WriteString(w, "hello, world!\n")
    //time.Sleep(100 * time.Second)
   // ok <- true
  //} ()
  //<- ok
}

func main() {
  runtime.GOMAXPROCS(4)
  http.HandleFunc("/", HelloServer)
  http.ListenAndServe(":8000", nil)
} 
