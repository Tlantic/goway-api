package main

import (
	"github.com/gin-gonic/gin"
	"github.com/andrepinto/goway-api/api"

	"github.com/andrepinto/goway-api-elasticsearch-store"
	"flag"
	"github.com/itsjamie/gin-cors"
	"time"
)


func main(){

	flag.Parse()

	var(
		elasticUrl 	= flag.String("elasticUrl", "", "")
		elasticIndex 	= flag.String("elasticIndex", "gateway", "gateway")
		elasticType 	= flag.String("elasticType", "http-logger", "http-logger")
	)

	elasticConn := goway_api_elasticsearch_store.NewElasticConn(*elasticUrl, *elasticIndex, *elasticType)
	elasticConn.Start()
	repo := goway_api_elasticsearch_store.NewElasticAnalyticRepository(&goway_api_elasticsearch_store.ElasticAnalyticRepositoryOptions{
		Elastic: elasticConn,
	})

	apiResource := api.NewApiResource(&api.ApiOptions{
		repo,
	})
	RunApi(apiResource)
}

func RunApi(api *api.ApiResource){
	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders: "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: true,
		ValidateHeaders: false,
	}))

	router.GET("v1/", api.Version)

	router.GET("v1/requests/last-hour", api.LastHour)



	router.Run(":9999")
}