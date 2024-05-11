package balance_test

import (
	"testing"

	"github.com/dossancto/challenge/internal/balance"
	"github.com/dossancto/challenge/internal/entities"
)

func TestRequestBalanceUseCase_FirstPlan_ShoudReturnOnlyOneItem(t *testing.T) {
	customer := entities.Person{
		Age:      41,
		Location: "SP",
		Name:     "Erikaya",
		CPF:      "123.456.789-10",
	}

	loans, err := balance.RequestBalanceUseCase(balance.RequestLoanInput{
		Customer: customer,
		Income:   3000,
	})

	if err != nil {
		t.Error(err)
	}

	loansCount := len(loans.Loans)

	if loansCount != 1 {
		t.Errorf("Must contain only one loans, contains %d", loansCount)
		t.FailNow()
	}

	loanType1 := loans.Loans[0].Type

	if loanType1 != "personal" {
		t.Errorf("The loan type must be 'personal', received %s", loanType1)
	}
}

func TestRequestBalanceUseCase_FirstPlan_ShoudReturnTwoLoans(t *testing.T) {
	customer := entities.Person{
		Age:      29,
		Location: "SP",
		Name:     "Erikaya",
		CPF:      "123.456.789-10",
	}

	loans, err := balance.RequestBalanceUseCase(balance.RequestLoanInput{
		Customer: customer,
		Income:   3000,
	})

	if err != nil {
		t.Error(err)
	}

	loansCount := len(loans.Loans)

	if loansCount != 2 {
		t.Errorf("Must contain only two loans, contains %d", loansCount)
		t.FailNow()
	}

	loanType1 := loans.Loans[0].Type

	if loanType1 != "personal" {
		t.Errorf("The loan type must be 'personal', received %s", loanType1)
	}

	loanType2 := loans.Loans[1].Type

	if loanType2 != "collateralized" {
		t.Errorf("The loan type must be 'collateralized', received %s", loanType2)
	}
}

func TestRequestBalanceUseCase_FirstPlan_ShoudReturnThreeLoans(t *testing.T) {
	customer := entities.Person{
		Age:      29,
		Location: "SP",
		Name:     "Erikaya",
		CPF:      "123.456.789-10",
	}

	loans, err := balance.RequestBalanceUseCase(balance.RequestLoanInput{
		Customer: customer,
		Income:   6000,
	})

	if err != nil {
		t.Error(err)
	}

	loansCount := len(loans.Loans)

	if loansCount != 3 {
		t.Errorf("Must contain only three loans, contains %d", loansCount)
		t.FailNow()
	}

	loanType1 := loans.Loans[0].Type

	if loanType1 != "personal" {
		t.Errorf("The loan type must be 'personal', received %s", loanType1)
	}

	loanType2 := loans.Loans[1].Type

	if loanType2 != "collateralized" {
		t.Errorf("The loan type must be 'collateralized', received %s", loanType2)
	}

	loanType3 := loans.Loans[2].Type

	if loanType3 != "payroll" {
		t.Errorf("The loan type must be 'payroll', received %s", loanType2)
	}
}
