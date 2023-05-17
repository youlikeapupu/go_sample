package controller
import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "learning-go/exception"
)

// type Res struct {
//     ID     int
//     Name   string
//     Enable int
// }

func Message(s string) string {
	var msg string = "success"
    return msg
}

func GetErrList(c *gin.Context) {
	// e:=

    c.JSON(http.StatusOK, gin.H{
    	"code": exception.GetCode(200),
    	"message": exception.GetMessage(200),
        "exception": exception.GetErr(),
    })

    // c.JSON(http.StatusOK, Res{
    // 	Code: 0,
    // 	Message: "suss",
    // 	Data: e
    // })
}

func GetSquare(c *gin.Context) {
    fmt.Printf("12331")
    var a int
    a = 6
    setDouble(&a)

    c.JSON(http.StatusOK, gin.H{
        "code": exception.GetCode(200),
        "message": exception.GetMessage(200),
        "data": a,
    })
}

func setDouble(a *int) {
    *a = *a * *a
}
