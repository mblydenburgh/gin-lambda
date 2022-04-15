package repository

type UserItem struct {
	UserId         string `dynamo:"UserId,hash"`
	ModelTypeAndId string `index:"modelTypeAndId,range"`
	UserName       string `dynamo:"User Name"`
	FirstName      string `dynamo:"First Name"`
	LastName       string `dynamo:"Last Name"`
	DateOfBirth    string `dynamo:"Date Of Birth"`
}

type CarItem struct {
	UserId         string `dynamo:"UserId, hash"`
	ModelTypeAndId string `dynamo:"ModelTypeAndId, range"`
	Manufacturer   string `dynamo:"Manufacturer"`
	Model          string `dynamo:"Model"`
	Trim           string `dynamo:"Trim"`
	Year           int32  `dynamo:"Year"`
	VehicleType    string `dynamo:"Vehicle Type"`
	Color          string `dynamo:"Color"`
	VIN            string `dynamo:"VIN"`
}
