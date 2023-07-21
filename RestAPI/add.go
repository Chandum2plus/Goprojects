package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"strings"
)

func Operation(ws *restful.WebService, httpMethodName, apiPath, doc string, handler func(req *restful.Request, res *restful.Response)) {
	log.Println("Registering the API", apiPath, "with the httpMethod", httpMethodName)
	routeBuilder := ws.Doc(doc).Method(strings.ToUpper(httpMethodName)).Path(string(apiPath)).To(handler)
	ws.Route(routeBuilder)
}
func setRootpath(ws *restful.WebService, path string) {
	log.Println("setting the root path", path, "for webservice", ws.Documentation())
	ws.Path(path)
}
func Webservice() *restful.WebService {
	ws := new(restful.WebService)
	ws.Doc("operator")
	setRootpath(ws, "/op")
	ws.Consumes(restful.MIME_JSON, restful.MIME_XML).Produces(restful.MIME_JSON, restful.MIME_JSON)
	Operation(ws, "POST", "/add", "Adding the number", Add)
	return ws
}
func Add(req *restful.Request, res *restful.Response) {
	Data := make(map[string]int)
	err := req.ReadEntity(&Data)
	fmt.Println(Data)
	sum := 0
	for _, er := range Data {
		if Data != nil {
			sum = sum + er
		} else {
			break
		}
		if err == nil {
			res.WriteEntity(sum)
		}
	}
}

func main() {
	restful.DefaultContainer(Webservice())
	log.Fatal()

}
