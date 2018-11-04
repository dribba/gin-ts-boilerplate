package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"./controllers"
	"github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"
)

var isProd = os.Getenv("GIN_MODE") == "release" || os.Getenv("PROD") == "true"

func renderHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"isProd": isProd,
	})
}

func loadTemplates() multitemplate.Render {
	templateBox, err := rice.FindBox("views")
	if err != nil {
		fmt.Println(err)
	}

	render := multitemplate.New()
	list := [...]string{"index.html"}
	for _, x := range list {
		templateString, err := templateBox.String(x)
		if err != nil {
			fmt.Println(err)
		}
		tmplMessage, err := template.New(x).Parse(templateString)
		if err != nil {
			fmt.Println(err)
		}
		render.Add(x, tmplMessage)
	}
	return render
}

func main() {
	r := gin.Default()

	assetBox := rice.MustFindBox("public")
	r.StaticFS("/assets", assetBox.HTTPBox())

	r.HTMLRender = loadTemplates()
	// router.HTMLRender = gorice.New(rice.MustFindBox("views"))
	// r.Use(static.Serve("/assets", static.LocalFile("./assets", true)))
	// r.LoadHTMLGlob("server/views/*")

	r.GET("/api/home", controllers.Home)

	r.GET("/", renderHome)
	r.NoRoute(renderHome)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	r.Run(":" + port)
}
