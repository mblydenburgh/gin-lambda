package repository

type UserItem struct {
	UserId         string `dynamo:"UserId,hash"`
	ModelTypeAndId string `dynamo:",range"`
	FirstName      string `dynamo:"First Name"`
	LastName       string `dynamo:"Last Name"`
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
