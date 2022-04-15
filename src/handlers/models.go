package handlers

import "time"

type UserResponse struct {
	UserId      string         `json:"userId"`
	FirstName   string         `json:"firstName"`
	LastName    string         `json:"lastName"`
	UserName    string         `json:"userName"`
	DateOfBirth time.Time      `json:"dateOfBirth"`
	Cars        *[]CarResponse `json:"cars"`
}

type AddUserPayload struct {
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	UserName    string    `json:"userName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

type CarResponse struct {
	VIN          string `json:"vin"`
	UserId       string `json:"userId"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Trim         string `json:"trim"`
	Year         int32  `json:"year"`
	VehicleTyle  string `json:"string"`
	Color        string `json:"color"`
}

type AddCarPayload struct {
	UserId       string `json:"userId"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Trim         string `json:"trim"`
	Year         int32  `json:"year"`
	VehicleTyle  string `json:"vehicleType"`
	Color        string `json:"color"`
	VIN          string `json:"vin"`
}
