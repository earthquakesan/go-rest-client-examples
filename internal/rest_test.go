package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSingleCompany(t *testing.T) {
	company := GetSingleCompany()
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

func TestPutSingleCompany(t *testing.T) {
	company := Company{
		Id:   1,
		Name: "Company One",
	}
	statusCode := putSingleCompany(company)
	assert.Equal(t, 204, statusCode)
}

func TestPostCompanyList(t *testing.T) {
	companies := []Company{
		{
			Id:   1,
			Name: "Company One",
		},
		{
			Id:   2,
			Name: "Company Two",
		},
	}
	statusCode := postCompanyList(companies)
	assert.Equal(t, 204, statusCode)
}
