package balance

func RequestBalanceUseCase(input RequestLoanInput) (RequestLoanOutput, error) {
	customer := input.Customer
	plans := UserPlans{}
	var loans []ProvidedLoan

	plans.hasPersonalPlan(input)
	plans.hasCollateralizedPlan(input)
	plans.hasPayrollPlan(input)

	if plans.Personal {
		loans = append(loans, ProvidedLoan{
			Type:  "personal",
			Taxes: 1,
		})
	}

	if plans.Collateralized {
		loans = append(loans, ProvidedLoan{
			Type:  "collateralized",
			Taxes: 2,
		})
	}

	if plans.Payroll {
		loans = append(loans, ProvidedLoan{
			Type:  "payroll",
			Taxes: 3,
		})
	}

	return RequestLoanOutput{
		Customer: customer,
		Loans:    loans,
    MatchedLoans: plans,
	}, nil
}

func (plans *UserPlans) hasPersonalPlan(RequestLoanInput) {
	plans.Personal = true
}

func (plans *UserPlans) hasCollateralizedPlan(input RequestLoanInput) {
	customer := input.Customer
	income := input.Income

	if income >= 5000 && customer.Age < 30 {
		plans.Collateralized = true
	} else if income > 3000 && income < 5000 && customer.Location == "SP" {
		plans.Collateralized = true
	} else if income <= 3000 && customer.Location == "SP" && customer.Age < 30 {
		plans.Collateralized = true
	}
}

func (plans *UserPlans) hasPayrollPlan(input RequestLoanInput) {
	customer := input.Customer
	income := input.Income

	if income >= 5000 && customer.Age < 30 {
		plans.Payroll = true
	}
}
