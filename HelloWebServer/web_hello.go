package main
/* 
	All GO programs start running from a function 'main' in package
*/
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		name :=req.URL.Query().Get("name")
		rw.Write([]byte(fmt.Sprintf("Hey, %s", name)))
	})
	http.ListenAndServe(":8080", nil)
}
