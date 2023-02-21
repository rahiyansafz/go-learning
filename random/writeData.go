package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Company struct {
	CompanyName   string `json:"companyName"`
	CompanyId     string `json:"companyId"`
	CompanyType   string `json:"companyType"`
	IsMenu        bool   `json:"isMenu"`
	CompanyRating int    `json:"companyRating"`
}

// Define an array of companies
var companies = []Company{
	{CompanyName: "Apple", CompanyId: "1", CompanyType: "Technology", IsMenu: true, CompanyRating: 5},
	{CompanyName: "Microsoft", CompanyId: "2", CompanyType: "Technology", IsMenu: false, CompanyRating: 4},
	{CompanyName: "Tesla", CompanyId: "3", CompanyType: "Automotive", IsMenu: true, CompanyRating: 5},
}

func CreateCompany() {
	// Create a new JSON file for the companies data
	file, err := os.Create("companies.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Loop over the companies and write the data to the JSON file
	for _, company := range companies {
		b, err := json.Marshal(company)
		if err != nil {
			fmt.Println("Error marshalling company:", err)
			return
		}

		_, err = file.Write(b)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("Companies data written to file successfully.")
}
