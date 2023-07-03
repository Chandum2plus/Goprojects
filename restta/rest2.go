package main

import (
	"github.com/emicklei/go-restful/v3"
	"io"
	"log"
	"net/http"
)

// This example shows how to create a (route) filter that performs basic authentication on the http request.
//
// GET http://localhost:8080/secret
// and use admin,admin for the credentials

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/secrets").Filter(basicAuthenticates).To(secrets))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func basicAuthenticates(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	//user/pwd = admin/admin
	u, p, ok := req.Request.BasicAuth()
	if !ok || u != "admin" || p != "admin" {
		resp.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		resp.WriteErrorString(401, "401: Not Authorized")
		return
	}
	chain.ProcessFilter(req, resp)
}
func secrets(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "401")
}
