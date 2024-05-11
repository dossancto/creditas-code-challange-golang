package balance

import (
	"github.com/dossancto/challenge/internal/entities"
)

type RequestLoanInput struct {
	Customer entities.Person
	Income   float64
}

type RequestLoanOutput struct {
	Customer     entities.Person
	Loans        []ProvidedLoan
	MatchedLoans UserPlans
}

type ProvidedLoan struct {
	Type  string
	Taxes int
}

type UserPlans struct {
	Personal       bool
	Collateralized bool
	Payroll        bool
}
