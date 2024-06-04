package handler

import (
	"log"
	"net/http"
	"ngc11/model"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func (r *Repo) Register(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	//validate user
	if user.Username == "" || user.Password == "" || user.DepositAmount <= 0 {
		return c.JSON(http.StatusBadRequest, "error or missing parameter")
	}

	// check existence of username
	var isExist model.User
	res := r.DB.Where("username = ?", user.Username).First(&isExist)
	if res.Error == nil {
		return c.JSON(http.StatusBadRequest, "user already exist")
	}

	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		log.Println("Error:", res.Error)
		return c.JSON(http.StatusInternalServerError, "Status internal error")
	}

	//hash pwd
	hashedpwd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedpwd)

	res = r.DB.Create(&user)
	if res.Error != nil {
		log.Println("Error:", res.Error)
		return c.JSON(http.StatusInternalServerError, "Status internal error")
	}

	return c.JSON(http.StatusCreated, user)
}
