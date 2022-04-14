package tests

import (
	"log"
	"mondu-challenge-alihamedani/application"
	"mondu-challenge-alihamedani/infrastructure/di"
	util "mondu-challenge-alihamedani/infrastructure/utils"
	"os"
	"testing"
)

var accountService application.BankAccountService

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	accountService = di.NewAccountServiceResolve(config)

	os.Exit(m.Run())
}
