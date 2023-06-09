package controllers

import (
	"avitoTest/internal/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (co *Controller) balanceIncrease(c *gin.Context) {
	var input entities.User
	if err := c.BindJSON(&input); err != nil {
		NewResponseMessage(c, http.StatusBadRequest, err.Error())
	}
	err := co.uc.BalanceIncrease(input)
	if err != nil {
		NewResponseMessage(c, http.StatusInternalServerError, err.Error())
	}
	c.Status(http.StatusOK)
}

func (co *Controller) transferOfFunds(c *gin.Context) {
	var input entities.Transfer
	if err := c.BindJSON(&input); err != nil {
		NewResponseMessage(c, http.StatusBadRequest, err.Error())
	}
	if err := co.uc.TransferOfFunds(input); err != nil {
		NewResponseMessage(c, http.StatusInternalServerError, err.Error())
	}
	c.Status(http.StatusOK)
}

func (co *Controller) getBalance(c *gin.Context) {
	var input entities.User
	input.ID, _ = strconv.Atoi(c.Param("id"))
	balance, err := co.uc.GetBalance(input)
	if err != nil {
		NewResponseMessage(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, Response{fmt.Sprint(balance)})
}
