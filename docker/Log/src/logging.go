package main

import (
	"fmt"
	"../bin/mux"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(4)
       os.MkdirAll("test", os.ModePerm);
	f, err := os.OpenFile("test/testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")
	go thread1()
	go thread2()
	go thread3()
	go thread4()

	muxRouter:=mux.NewRouter()
	subMuxRouter:= muxRouter.PathPrefix("/api/").Subrouter()
	subMuxRouter.HandleFunc("/exit",ExitFunc)
	print("Server is starting")
	server :=&http.Server{
		Handler: subMuxRouter,
		Addr: ":9090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,

	}
	print("Server is started")
	var err_server = server.ListenAndServe()
	if err != nil {
		log.Panicln("Server failed starting. Error: %s", err_server)
	}



}


func ExitFunc(response http.ResponseWriter, request *http.Request) {
	log.Print("Inside ExitFunc")
	log.Println("Just to make Server Running")

}

 func thread1( )  {
 	var threadName string="Thread1"
 	for {

		log.Println("Log From Thread "+threadName)
		log.Println("Random Int Number : "+string(rand.Int()) )
	}

}
func thread2()  {
	var threadName string="thread2"
	for {

		log.Println("Log From Thread "+threadName)
		output_float := fmt.Sprintf("%f", rand.Float32())
		log.Println("Random Float32 Number : "+ output_float )
	}


}
func thread3()  {
	var threadName string="thread3"
	for {

		log.Println("Log From Thread "+threadName)
		output_float := fmt.Sprintf("%f", rand.Float64())
		log.Println("Random Float64 Number : "+ output_float )
	}


}
func thread4()  {
	var threadName string="thread4"
	for {

		log.Println("Log From Thread "+threadName)
		output_float := fmt.Sprintf("%f", rand.ExpFloat64())
		log.Println("Random ExpFloat64  Number : "+ output_float )
	}


}
