package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSingleCompany(t *testing.T) {
	company := getSingleCompany()
	expectedCompany := Company{Id: 1, Name: "Company One"}
	assert.Equal(t, expectedCompany, company)

	companyReq := getSingleCompanyReq()
	assert.Equal(t, expectedCompany, companyReq)
}

func TestGetCompanyList(t *testing.T) {
	companies := getCompanyList()
	expectedCompanies := []Company{
		{
			Id:   1,
			Name: "Company One",
		},
		{
			Id:   2,
			Name: "Company Two",
		},
	}
	assert.Equal(t, expectedCompanies, companies)
}
