
package main
import (
  "net/http"
//  "io"
  "runtime"
  "time"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
  time.Sleep(100 * time.Second)
  w.Write([]byte("Hello async"))
}

func main() {
  runtime.GOMAXPROCS(4)
  http.HandleFunc("/", HelloServer)
  http.ListenAndServe(":8000", nil)
} 
