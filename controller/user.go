package controller
import (
    // "net/http"
    // "github.com/gin-gonic/gin"
    // "learning-go/exception"
    "net/http"
    "github.com/gin-gonic/gin"
    // "strconv"
    "learning-go/database"
    "fmt"
    "time"
)


type User struct {
    ID int `form:"id" gorm:"column:id"`
    UserName string `form:"username" gorm:"column:username"`
    Alias string `form:"alias" gorm:"column:alias"`
    Status  int `form:"status" gorm:"column:status"`
    CreateTime time.Time  `gorm:"column:createat"`
}

type InsertData struct {
    Username string `gorm:"default:galeone"`
    Alias string `gorm:"default:galeone"`
    Status int `gorm:"default:0"`
}

type UpdateData struct {
    Username string `gorm:"default:galeone"`
    Alias string `gorm:"default:galeone"`
    Status int `gorm:"default:0"`
}

// type Res struct {
//     ID     int
//     Name   string
//     Enable int
// }

// func GetErrList(c *gin.Context) {
// 	// e:=

//     c.JSON(http.StatusOK, gin.H{
//     	"code": exception.GetCode(200),
//     	"message": exception.GetMessage(200),
//         "exception": exception.GetErr(),
//     })
// }

// 定義user表名
func (User) TableName () string {
    return "user"
}

func GetPlayer(c *gin.Context) {
    db, err := database.Connect()
    if err != nil {
        fmt.Println(err)
    }

    // 接收 get 參數
    var user User
    c.Bind(&user)

    // 無指定key
    data := []User{}

    db.Table("user").Where(&user).Find(&data)

    // map塞key給User陣列
    datamap := map[int]User{}

    // 塞 key
    for key, value := range data {
        datamap[key] = value
    }

    c.JSON(http.StatusOK, gin.H{
        "result":  "ok",
        "data":  datamap,
    })
}

func CreatePlayer(c *gin.Context) {
    db, err := database.Connect()
    if err != nil {
        fmt.Println(err)
    }

    // 接收post參數
    var user User
    c.Bind(&user)

    a := InsertData{
        Username: user.UserName,
        Alias: user.Alias,
        Status: user.Status,
    }

    result := db.Table("user").Create(&a)

    c.JSON(http.StatusOK, gin.H{
        "insert":  a,
        "result":  "ok",
        "error":  result.Error,
        "recordsCount": result.RowsAffected,
    })
}

func UpdatePlayer(c *gin.Context) {
    db, err := database.Connect()
    if err != nil {
        fmt.Println(err)
    }

    var user User
    c.Bind(&user)

    b := UpdateData{
        Username: user.UserName,
        Alias: user.Alias,
        Status: user.Status,
    }

    fmt.Println(user.ID, b)

    db.Table("user").Where("id = (?)", user.ID).Update(&b)

    c.JSON(http.StatusOK, gin.H{
        "data":  user,
        "result":  "ok",
    })
}