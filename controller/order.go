package controller
import (
    "net/http"
    "github.com/gin-gonic/gin"
    "learning-go/database"
    "fmt"
    // "time"
    // "reflect"
    "encoding/json"
)

type Order struct {
    ID int `form:"id" gorm:"column:id"`
    Type  int `form:"type" gorm:"column:type"`
    UserName string `form:"username" gorm:"column:username"`
    UserID int `form:"userid" gorm:"column:userid"`
    ProductID  int `form:"productid" gorm:"column:productid"`
    Spec  string `form:"spec" gorm:"column:spec"`
    Phone  string `form:"phone" gorm:"column:phone"`
    Mark  string `form:"mark" gorm:"column:mark"`
    Address  string `form:"address" gorm:"column:address"`
    Shop  string `form:"shop" gorm:"column:shop"`
    Quantity  int `form:"quantity" gorm:"column:quantity"`
    Amount  float64 `form:"amount" gorm:"column:amount"`
    Status  int `form:"status" gorm:"column:status"`
}

type OrderReq struct {
    ID int `form:"id" json:"id"`
    Type  int `form:"type" json:"type"`
    UserName string `form:"username" json:"username"`
    UserID int `form:"userid" json:"userid"`
    ProductID  int `form:"productid" json:"productid"`
    Spec  []Spec `form:"spec" json:"spec"`
    Phone  string `form:"phone" json:"phone"`
    Mark  string `form:"mark" json:"mark"`
    Address  string `form:"address" json:"address"`
    Shop  string `form:"shop" json:"shop"`
    Quantity  int `form:"quantity" json:"quantity"`
    Amount  float64 `form:"amount" json:"amount"`
    Status  int `form:"status" json:"status"`
}

type Spec struct {
    Color string `json:"color"`
    Size string `json:"size"`
    Quan int `json:"quan"`
}

func CreateOrder(c *gin.Context) {
	db, err := database.Connect()
    if err != nil {
        fmt.Println(err)
    }
	// 接收post參數
    var order OrderReq
    c.BindJSON(&order)

    // 計算產品總數量
    amount := 0
    for _, v := range order.Spec {
    	amount += v.Quan
    }

    specJson, err := json.Marshal(order.Spec)
    if err != nil {
        fmt.Println(err)
        return
    }

    data := Order{
        Type: order.Type,
        UserName: order.UserName,
        UserID: order.UserID,
        ProductID: order.ProductID,
        Spec: string(specJson),
        Phone: order.Phone,
        Mark: order.Mark,
        Address: order.Address,
        Shop: order.Shop,
        Quantity: amount,
        Amount: order.Amount,
        Status: order.Status,
    }

	result := db.Table("orders").Create(&data)

	c.JSON(http.StatusOK, gin.H{
        "result":  "ok",
        "error":  result.Error,
        "recordsCount": result.RowsAffected,
    })
}