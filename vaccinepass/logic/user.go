package logic

import (
	"context"

	"github.com/parijatpurohit/sidecar-sql/vaccinepass/server/transport"
	"github.com/parijatpurohit/sidecar-sql/zz_generated/go/protogen"
)

func (l *logicImpl) GetUserDetails(req *transport.GetUserDetailsRequest) (*transport.GetUserDetailsResponse, error) {
	data, err := l.client.UserDocs_FindByUserID(context.Background(), &protogen.UserDocs_FindByUserIDRequest{
		Query: &protogen.UserDocs_FindByUserID_Query{UserID: req.UserUUID},
	})
	if err != nil {
		return nil, err
	}
	var ud []transport.UserDetails
	for _, row := range data.UserDocss {
		ud = append(ud, transport.UserDetails{
			UserID:     row.OfficialID,
			UserIDType: row.OfficialIDType,
		})
	}
	res := &transport.GetUserDetailsResponse{
		UserDetails: ud,
	}
	return res, nil
}

func (l *logicImpl) GetUserVaccineDetails(req *transport.GetUserVaccineDetailsRequest) (*transport.GetUserVaccineDetailsResponse, error) {
	userDetails, err := l.GetUserDetails(&transport.GetUserDetailsRequest{
		UserUUID: req.UserUUID,
	})

	if err != nil {
		return nil, err
	}

	var userDocIDs []string
	userDocDetails := map[string]string{}
	for _, userDetail := range userDetails.UserDetails {
		userDocIDs = append(userDocIDs, userDetail.UserID)
		userDocDetails[userDetail.UserID] = userDetail.UserIDType
	}

	data, err := l.client.UserVaccines_FindByUserDocID(context.Background(), &protogen.UserVaccines_FindByUserDocIDRequest{
		Query: &protogen.UserVaccines_FindByUserDocID_Query{
			UserDocID: userDocIDs,
		},
	})
	if err != nil {
		return nil, err
	}

	var uv []transport.VaccineDetails
	for _, row := range data.UserVacciness {
		vaccineName, err := l.getVaccineName(row.VaccineID)
		if err != nil {
			return nil, err
		}
		uv = append(uv, transport.VaccineDetails{
			Name:       vaccineName,
			UserID:     row.UserDocID,
			UserIDType: userDocDetails[row.UserDocID],
			Authority:  row.VaccineAuthorityID,
			// DOV:        row.VaccinationDate,
		})
	}
	return nil, nil
}

func (l *logicImpl) getVaccineName(vaccineID string) (string, error) {
	vaccine, err := l.client.Vaccines_FindByVaccineID(context.Background(), &protogen.Vaccines_FindByVaccineIDRequest{
		Query: &protogen.Vaccines_FindByVaccineID_Query{
			UUID: vaccineID,
		},
	})
	if err != nil {
		return "", err
	}
	vaccineName := vaccine.Vacciness[0].Name
	return vaccineName, nil
}

func (l *logicImpl) GetRequiredVaccines(req *transport.GetRequiredVaccineRequest) (*transport.GetRequiredVaccineResponse, error) {
	data, err := l.client.Countries_FindByCountryID(context.Background(), &protogen.Countries_FindByCountryIDRequest{
		Query: &protogen.Countries_FindByCountryID_Query{
			UUID: req.Country,
		},
	})
	if err != nil {
		return nil, err
	}
	var vd []transport.RequiredVaccineDetails
	for _, row := range data.Countriess {
		vd = append(vd, transport.RequiredVaccineDetails{
			Name: row.Name,
		})
	}
	res := &transport.GetRequiredVaccineResponse{
		RequiredVaccineDetails: vd,
	}
	return res, nil
}
