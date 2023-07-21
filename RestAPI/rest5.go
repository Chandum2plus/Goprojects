package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/chandu").To(res))
	restful.Add(ws)

	go func() {
		log.Fatal(http.ListenAndServe(":8081", nil))
	}()
	container2 := restful.NewContainer()

	ws2 := new(restful.WebService)
	ws2.Route(ws2.GET("/chandu2").To(res2))
	container2.Add(ws2)
	server := &http.Server{Addr: ":8080", Handler: container2}
	log.Fatal(server.ListenAndServe())
}
func res(req *restful.Request, resp *restful.Response) {
	resp.Write([]byte("Hello this is Chandu here !"))
}
func res2(req *restful.Request, resp *restful.Response) {
	resp.Write([]byte("I am in m2-plus"))
}
