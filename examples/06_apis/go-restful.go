package main

import (
	"log"
	"net/http"
	"strconv"

	restful "github.com/emicklei/go-restful"
)

type Deploy struct {
	Id    int
	Image string
}

type DeployResource struct {
	deploys map[int]Deploy
}

func (dr DeployResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/deploys").
		Doc("Deploy images").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/{deploy-id}").To(dr.findDeploy).
		Doc("get a deploy").
		Writes(Deploy{}))

	ws.Route(ws.POST("").To(dr.createDeploy).
		Doc("create a deploy").
		Reads(Deploy{}))

	container.Add(ws)
}

func (dr DeployResource) findDeploy(req *restful.Request, resp *restful.Response) {
	deployID := req.PathParameter("deploy-id")
	id, err := strconv.Atoi(deployID)
	if err != nil {
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	if deploy, ok := dr.deploys[id]; ok {
		resp.WriteEntity(deploy)
		return
	}

	resp.WriteErrorString(http.StatusNotFound, "")
	return
}

func (dr DeployResource) createDeploy(req *restful.Request, resp *restful.Response) {
	deploy := new(Deploy)
	if err := req.ReadEntity(deploy); err != nil {
		resp.WriteError(http.StatusBadRequest, err)
		return
	}
	deploy.Id = len(dr.deploys) + 1
	dr.deploys[deploy.Id] = *deploy
	resp.WriteHeaderAndEntity(http.StatusCreated, deploy)
}

func main() {
	wsContainer := restful.NewContainer()
	dr := DeployResource{map[int]Deploy{}}
	dr.Register(wsContainer)

	log.Printf("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
