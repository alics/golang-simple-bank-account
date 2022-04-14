package presentation

import (
	"context"
	"github.com/gin-gonic/gin"
	"mondu-challenge-alihamedani/application"
	"mondu-challenge-alihamedani/application/contracts"
	"net/http"
	"strconv"
)

type ResponseData struct {
	Status bool
	Result interface{}
	Error  string
}

type AccountController struct {
	AccountService application.BankAccountService
}

func (a AccountController) Create(ctx *gin.Context) {
	responseDto := ResponseData{
		Status: true,
	}

	var req contracts.CreateAccountModel
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responseDto.Status = false
		responseDto.Error = err.Error()
		ctx.JSON(http.StatusBadRequest, responseDto)
		return
	}

	account, err := a.AccountService.AddAccount(context.Background(), &req)
	if err != nil {
		responseDto.Status = false
		responseDto.Error = err.Error()
		ctx.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	responseDto.Result = account
	ctx.JSON(http.StatusOK, responseDto)
}

func (a AccountController) Deposit(ctx *gin.Context) {
	responseDto := ResponseData{
		Status: true,
	}

	var req contracts.CreateTransferModel
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responseDto.Status = false
		responseDto.Error = err.Error()
		ctx.JSON(http.StatusBadRequest, responseDto)
		return
	}

	err := a.AccountService.DepositMoneyToAccount(context.Background(), &req)
	if err != nil {
		responseDto.Status = false
		responseDto.Error = err.Error()
		ctx.JSON(http.StatusInternalServerError, responseDto)
		return
	}

	responseDto.Result = "success"
	ctx.JSON(http.StatusOK, responseDto)
}

func (a AccountController) Withdraw(ctx *gin.Context) {
	responseDto := ResponseData{
		Status: true,
	}
	var req contracts.CreateTransferModel
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responseDto.Status = false
		responseDto.Error = err.Error()
		ctx.JSON(http.StatusBadRequest, responseDto)
		return
	}

	err := a.AccountService.WithdrawMoneyFromAccount(context.Background(), &req)
	if err != nil {
		responseDto.Status = false
		responseDto.Error = err.Error()
		ctx.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	responseDto.Result = "success"
	ctx.JSON(http.StatusOK, responseDto)
}

func (a AccountController) GetBalance(ctx *gin.Context) {
	responseDto := ResponseData{
		Status: true,
	}

	id := ctx.Param("id")

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		responseDto.Status = false
		responseDto.Error = err.Error()
		ctx.JSON(http.StatusBadRequest, responseDto)
		return
	}

	req := contracts.GetAccountBalanceModel{}
	req.Id = i

	accountBalance, err := a.AccountService.GetBalance(context.Background(), &req)
	if err != nil {
		responseDto.Status = false
		responseDto.Error = err.Error()
		ctx.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	responseDto.Result = accountBalance
	ctx.JSON(http.StatusOK, responseDto)
}
