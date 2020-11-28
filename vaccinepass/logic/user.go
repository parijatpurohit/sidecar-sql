package logic

import (
	"github.com/parijatpurohit/sidecar-sql/vaccinepass/server/transport"
)

func GetUserDetails(req *transport.GetUserDetailsRequest) (*transport.GetUserDetailsResponse, error) {
	/*	return &transport.GetUserDetailsResponse{
		UserDetails: []transport.UserDetails{
			{
				UserID:     "123",
				UserIDType: "1234",
			},
			{
				UserID:     "123",
				UserIDType: "1234",
			},
		},
	}, nil*/

	return nil, nil
}

func GetUserVaccineDetails(req *transport.GetUserVaccineDetailsRequest) (*transport.GetRequiredVaccineResponse, error) {
	return nil, nil
}

func GetRequiredVaccines(req *transport.GetRequiredVaccineRequest) (*transport.GetRequiredVaccineResponse, error) {
	return nil, nil
}
