package httpserverdemo

import (
	"fmt"
	"log"
	"net/http"

	mlog "github.com/mike504110403/goutils/log"
)

//const sslCmd string = "openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt"

type Refer struct {
	handler http.Handler
	refer   string
}

func Init() {
	// Start a http server
	// service := &http.Server{
	// 	Addr: "
	// Start a http server
	// service := &http.Server{
	// 	Addr: "0.0.0.0:80",
	// }
	// Register a handler
	//http.HandleFunc("/", helloHandler)
	// refer := &Refer{
	// 	handler: http.HandlerFunc(myHandler),
	// 	refer:   "http://localhost",
	// }
	//http.HandleFunc("/hello", helloHandler)
	// Start the server
	//service.ListenAndServe()
	//http.ListenAndServe(":8080", refer)
	srv := &http.Server{
		Addr: ":8088",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mlog.Info(fmt.Sprintf("got connection: %s", r.Proto))
			w.Write([]byte("Http2~~\n"))
		}),
	}
	log.Fatal(srv.ListenAndServeTLS("/Users/linminze/server.crt", "/Users/linminze/server.key"))
}

// 這段程式碼是用來檢查 Referer 是否符合指定的網址
func (r *Refer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Referer() == r.refer {
		r.handler.ServeHTTP(w, req)
	} else {
		w.WriteHeader(403)
	}
}

func myHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("this is my handler"))
}

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//mlog.Info("Hello World!")
	fmt.Fprintf(w, "Hello World!\n")
}

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, world!\n"))
}
