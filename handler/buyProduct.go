package handler

import (
	"fmt"
	"log"
	"net/http"
	"ngc11/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (r *Repo) BuyProduct(c echo.Context) error {
	//get the user data
	username := c.Get("username")
	var user model.User
	r.DB.Where("username = ?", username).First(&user)

	var getP model.Product

	err := c.Bind(&getP)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	if getP.ProductID <= 0 || getP.Stock <= 0 {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "error or missing parameter")
	}

	// find id
	var p model.Product
	res := r.DB.First(&p, getP.ProductID)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, "there is no such product")
		}
		log.Println(res.Error)
		c.JSON(http.StatusInternalServerError, "Internal server Error")
	}

	//start transaction
	var t model.Transaction
	r.DB.Transaction(func(tx *gorm.DB) error {

		if getP.Stock >= p.Stock {
			log.Println("stock is not enough")
			return c.JSON(http.StatusInternalServerError, "stock is not enough")
		}
		p.Stock = p.Stock - getP.Stock
		res = r.DB.Save(&p)
		if res.Error != nil {
			log.Println(res.Error)
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}

		if user.DepositAmount <= p.Price*float64(getP.Stock) {
			log.Println("unsufficient amount of money")
			return c.JSON(http.StatusBadRequest, "unsufficient amount of money")
		}

		// subtract from user money
		user.DepositAmount -= p.Price * float64(getP.Stock)
		if res.Error != nil {
			log.Println(res.Error)
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		// update the user data
		res = r.DB.Save(&user)

		t.ProductID = p.ProductID
		t.UserID = user.UserID
		t.Quantity = getP.Stock
		t.TotalAmount = p.Price * float64(getP.Stock)
		res = r.DB.Create(&t)
		if res.Error != nil {
			log.Println(res.Error)
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}

		return nil
		// if find, add the product
	})

	// if transaction is rolled back
	if t.ProductID == 0 {
		return err
	}
	fmt.Println("here")
	return c.JSON(http.StatusCreated, t)
}
