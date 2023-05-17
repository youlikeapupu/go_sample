package controller
import (
    "net/http"
    "github.com/gin-gonic/gin"
    "learning-go/database"
    "fmt"
    "time"
    // "reflect"
    "encoding/json"
)

type Product struct {
    ID int `form:"id" gorm:"column:id"`
    ProductName string `form:"productname" gorm:"column:product_name"`
    Spec string `form:"spec" gorm:"column:spec" json:"spec"`
    Quantity  int `form:"quantity" gorm:"column:quantity"`
}

type ProductReq struct {
    ID int `form:"id" json:"id"`
    ProductName string `form:"productname" json:"productname"`
    Spec interface{} `form:"spec" json:"spec"`
    Quantity  int `form:"quantity" json:"quantity"`
}

type ProductRes struct {
    ID int `form:"id" gorm:"column:id"`
    ProductName string `form:"productname" gorm:"column:product_name"`
    Spec string`form:"spec" gorm:"column:spec" json:"spec"`
    Quantity  int `form:"quantity" gorm:"column:quantity"`
    CreateTime time.Time  `form:"createat" gorm:"column:createat"`
}

func CreateProduct(c *gin.Context) {
    db, err := database.Connect()
    if err != nil {
        fmt.Println(err)
    }

    // 接收post參數
    var product ProductReq
    c.BindJSON(&product)

    // 轉jason格式
    specJson, err := json.Marshal(product.Spec)
    if err != nil {
        fmt.Println(err)
    }

    // 沒什麼意義就想留著
    // var appendSpec []interface{}
    // for _, v := range product.Spec.([]interface{}) {
    //     appendSpec = append(appendSpec, v);
    // }

    data := Product{
        ProductName: product.ProductName,
        Spec: string(specJson),
        Quantity: product.Quantity,
    }

    result := db.Table("products").Create(&data)

    c.JSON(http.StatusOK, gin.H{
        "insert":  product,
        "result":  "ok",
        "error":  result.Error,
        "recordsCount": result.RowsAffected,
    })
}

func UpdateProduct(c *gin.Context) {
    db, err := database.Connect()
    if err != nil {
        fmt.Println(err)
    }
    
    // 接收post參數
    var product ProductReq
    c.BindJSON(&product)

    fmt.Println(product)

    // 轉jason格式
    specJson, err := json.Marshal(product.Spec)
    if err != nil {
        fmt.Println(err)
    }

    data := Product{
        ProductName: product.ProductName,
        Spec: string(specJson),
        Quantity: product.Quantity,
    }

    db.Table("products").Where("id = (?)", product.ID).Update(&data)

    c.JSON(http.StatusOK, gin.H{
        "data":  product,
        "result":  "ok",
    })
}