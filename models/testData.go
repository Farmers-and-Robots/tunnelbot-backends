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

var NewValidAddress = Address{
	City:   "Seattle",
	State:  "WA",
	Street: "5050 Shiras Ave",
	Zip:    "98101",
}

var ValidPersonAssociation = PersonAssociation{
	DisplayName: "Test Person",
	Email:       "foo@far.com",
	PersonId:    1,
	Role:        "owner",
}

var ValidFarmName = "TestFarm"
var ValidCreateDate = "2020-01-0100:00:00"

var NewValidFarmName = "TestFarm2"
var NewValidCreateDate = "2020-02-0100:00:00"
var NewValidPersonAssociation = PersonAssociation{
	DisplayName: "Test Person 2",
}

var NewValidFarm = Farm{
	Address:    &NewValidAddress,
	CreateDate: NewValidCreateDate,
	Id:         2,
	Kind:       "frm",
	Name:       NewValidFarmName,
	Owners: []PersonAssociation{
		NewValidPersonAssociation,
	},
	People: []PersonAssociation{
		NewValidPersonAssociation,
	},
}

var ValidDevice = Device{
	Category: DeviceCategory("heating"),
	Id:       1,
	Kind:     "dev",
	Name:     "Test Device",
}

var NewValidDevice = Device{
	Category: NewValidDeviceCategory,
	Id:       2,
	Kind:     "dev",
	Name:     "Test Device 2",
}

var NewValidDeviceCategory = DeviceCategory("irrigation")

var ValidEvent = Event{
	Date:        "2020-01-0100:00:00",
	Id:          1,
	Description: "Test Event",
	Kind:        "evt",
	Operator:    "Test Operator",
}

var NewValidEvent = Event{
	Date:        "2020-02-0100:00:00",
	Id:          2,
	Description: "Test Event 2",
	Kind:        "evt",
	Operator:    "Test Operator 2",
}

var ValidPerson = Person{
	FirstName:   "Test",
	LastName:    "Person",
	DisplayName: "Test Person",
	Email:       "foo@far.com",
	Id:          1,
	Kind:        "per",
	Phone:       "555-555-5555",
}

var NewValidPerson = Person{
	FirstName:   "Test",
	LastName:    "Person 2",
	DisplayName: "Test Person 2",
	Email:       "fool@far.com",
	Id:          2,
	Kind:        "per",
	Phone:       "777-777-7777",
}
