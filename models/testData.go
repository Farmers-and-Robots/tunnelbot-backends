package models

// ValidFarm is a valid instance of models.Farm.
var ValidFarm = Farm{
	Address:    &ValidAddress,
	CreateDate: ValidCreateDate,
	Id:         1,
	Kind:       "frm",
	Name:       ValidFarmName,
	Owners: []PersonAssociation{
		ValidPersonAssociation,
	},
	People: []PersonAssociation{
		ValidPersonAssociation,
	},
}

var ValidAddress = Address{
	City:   "Dubuque",
	State:  "IA",
	Street: "549 Almond",
	Zip:    "52001",
}

var ValidPersonAssociation = PersonAssociation{
	DisplayName: "Test Person",
	Email:       "foo@far.com",
	PersonId:    1,
	Role:        "owner",
}

var ValidFarmName = "TestFarm"
var ValidCreateDate = "2021-01-01T00:00:00Z"
