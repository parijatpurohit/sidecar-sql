package transport

import "github.com/golang/protobuf/ptypes/timestamp"

type ErrResponse struct {
	Message string
}
type GetUserDetailsRequest struct {
	UserUUID string `json:"userUUID"`
}

type GetUserDetailsResponse struct {
	UserDetails []UserDetails
}

type UserDetails struct {
	UserID     string
	UserIDType string
}

type GetUserVaccineDetailsRequest struct {
	UserUUID string
}

type GetUserVaccineDetailsResponse struct {
	VaccineDetails []VaccineDetails
}

type VaccineDetails struct {
	Name       string
	UserID     string
	UserIDType string
	Authority  string
	Country    string
	DOV        *timestamp.Timestamp // DateOfVaccine
}

type GetRequiredVaccineRequest struct {
	Country string
}

type GetRequiredVaccineResponse struct {
	RequiredVaccineDetails []RequiredVaccineDetails
}

type RequiredVaccineDetails struct {
	Name                   string
	Description            string                   // includes disease, and info section for that vaccine
	AuthorisedBodiesNearby []AuthorisedBodiesNearby // extension: Can include booking slot from the app itself
}

type AuthorisedBodiesNearby struct {
	Name     string
	Location string
	Call     string
}
