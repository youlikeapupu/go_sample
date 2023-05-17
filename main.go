package main

import (
	"learning-go/controller"
 //    "github.com/gin-gonic/contrib/static"
    "net/http"
    "github.com/gin-gonic/gin"
    "strconv"
    // "learning-go/database"
    // "fmt"
    // "time"
)

var balance = 1000

func main() {
    router := gin.Default()

    // router.Use(static.Serve("/", static.LocalFile("./fronted/build", true)))

    // router.RedirectFixedPath = true

    router.Group("/controller")

	router.GET("/balance/", getBalance)

	router.GET("/count", simpleCount)

	router.GET("/user/player", controller.GetPlayer)

	router.POST("/user/createplayer", controller.CreatePlayer)

	router.POST("/user/updateplayer", controller.UpdatePlayer)

	router.POST("/product/createproduct", controller.CreateProduct)

	router.POST("/product/updateproduct", controller.UpdateProduct)

	router.POST("/product/delproduct", controller.DelProduct)

    router.Run("localhost:5001")
}

func getBalance(c *gin.Context) {
	var msg = "您的帳戶內有:" + strconv.Itoa(balance) + "元"
	c.JSON(http.StatusOK, gin.H{
		"amount":  balance,
		"status":  "ok",
		"message": msg,
	})
}

func simpleCount(c *gin.Context) {
	var amount float64
	nubmer1, _ := strconv.ParseFloat(c.Query("nubmer1"), 64)
	nubmer2, _ := strconv.ParseFloat(c.Query("nubmer2"), 64)
	status, _ := strconv.Atoi(c.Query("status"))

	switch {
		case status == 1:
			amount = nubmer1+nubmer2
		case status == 2:
			amount = nubmer1-nubmer2
		case status == 3:
			amount = nubmer1*nubmer2
		case status == 4:
			amount = nubmer1/nubmer2
	}

	// msg := []byte("您輸入的文字為: \n" + input)                     // 純文字(text/plain)中的換行是\n，網頁格式(html)中的換行才是<br />
	// c.Data(http.StatusOK, "text/plain; charset=utf-8;", msg) // 如果沒有指定文字編碼、拿掉`charset=utf-8;`的話，中文會變亂碼。

	c.JSON(http.StatusOK, gin.H{
		"amount":  amount,
		"result":  "ok",
	})

}