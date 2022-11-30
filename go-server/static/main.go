package main

import(
	"fmt"
	"net/http"
	"log"
)
func formHandler(w http.ResponseWriter,r *http.Request){
	if err :=r.ParseForm(); err != nil{
		fmt.Printf(w ,"ParseForm() err: %v",err)
		return
	}
	fmt.Printf(w,"POST request successful")
	name :=r.FormValue("name")
	address := r.FormValue("address")
	fmt.Printf(w,"Name=%s\n", name)
	fmt.Printf(w,"Address=%s\n", address)
}

func helloHandler(w http.ResponseWriter,r *http.Request){ 
if r.URL.Path !="/Hello"{
	http.Error(w,"404 not found",http.StatusNotFound)
	return
}
if r.Method !="GET"{
	http.Error(w, "method is not supported",http.StatusNotFound)
	return
}
fmt.Printf(w,"Hello!")
}
func main()  {
	fileServer:=http.FileServer(http.Dir("/static"))
	http.Handle("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err :=http.ListenAndServe(":8080",nil);err !=nil(
		log.Fatal(err)
	)
}