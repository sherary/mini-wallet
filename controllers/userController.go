package controllers

import (
	helpers "mini-wallet/helpers"
	models "mini-wallet/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Signup(ctx *gin.Context) {
	customerXID := ctx.PostForm("customer_xid")

	newUser := &models.User{
		Customer_XID: []byte(customerXID),
	}
	err := db.Create(newUser).Error

	if err != nil {
		errMessage := map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err,
			},
		}

		ctx.JSON(http.StatusInternalServerError, errMessage)
		return
	}

	token, err := helpers.GenerateToken(customerXID, false)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	data := map[string]interface{}{
		"status": "success",
		"data": map[string]interface{}{
			"token": token,
		},
	}

	ctx.JSON(200, data)
}
