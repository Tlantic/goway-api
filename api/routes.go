package api

import (
	"github.com/gin-gonic/gin"
	"github.com/andrepinto/goway-api/domain"
	"fmt"
	"time"
)

func(api *ApiResource) Version(c *gin.Context){
	c.JSON(200, api.HttpResponse("ok", "API_VERSION", "v1"))
}


func(api *ApiResource) LastHour(c *gin.Context){
	t0 := time.Now()

	fmt.Println(c.Query("client"))
	fmt.Println(c.Query("product"))
	_, data := api.AnalyticRespository.GetLastHourRequest((*domain.QueryOptions)(nil))

	t1 := time.Now()


	// Get duration.
	d := t1.Sub(t0)
	fmt.Println("Duration", d.Seconds())
	c.JSON(200, api.HttpResponse("ok", "LAST_HOUR_REQUEST", data))
}
