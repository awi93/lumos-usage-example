package http

import (
	"dimall.id/standard-template/bootstrap"
	"dimall.id/standard-template/http/controller"
	"dimall.id/standard-template/repo"
	"github.com/dimall-id/lumos/http"
	"log"
)

func Routes () {
	pr := repo.ProductRepo{
		DB: bootstrap.Db,
	}
	pc := controller.ProductController{
		Repo: pr,
	}
	err := http.AddRoute(http.Route{
		Name: "List of Product",
		HttpMethod: "GET",
		Url: "/products",
		Roles: []string{
			"USER",
		},
		Func: pc.Index,
	})
	if err != nil {
		log.Fatal(err)
	}
}