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
	UserId         string
	ModelTypeAndId string
	Manufacturer   string
	Model          string
	Trim           string
	Year           int32
	VehicleType    string
	Color          string
	VIN            string
}
