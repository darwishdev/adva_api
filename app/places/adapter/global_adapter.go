package adapter

import (
	"encoding/json"

	rmsv1 "github.com/darwishdev/bzns_pro_api/common/pb/bznspro/v1"
)

//list

func (a *PlacesAdapter) PlacesListGrpcFromSql(resp *[][]byte) (*rmsv1.PlacesListResponse, error) {
	var records []*rmsv1.Place

	for _, v := range *resp {
		var record rmsv1.Place

		err := json.Unmarshal([]byte(v), &record)
		if err != nil {
			return nil, err
		}

		records = append(records, &record)
	}
	response := &rmsv1.PlacesListResponse{
		Places: records,
	}
	return response, nil
}
func neighbourhoodsInput(neighbourhoods []*rmsv1.NeighbourhoodNested) []*rmsv1.PlacesInput {
	var records []*rmsv1.PlacesInput
	for _, neighbourhood := range neighbourhoods {
		record := &rmsv1.PlacesInput{
			Value: neighbourhood.NeighbourhoodId,
			Label: neighbourhood.NeighbourhoodName,
		}

		records = append(records, record)
	}

	return records
}
func districtsInput(ditricts []*rmsv1.DistrictNested) []*rmsv1.PlacesInput {
	var records []*rmsv1.PlacesInput
	for _, district := range ditricts {
		neighbourhoods := neighbourhoodsInput(district.Neighbourhoods)
		record := &rmsv1.PlacesInput{
			Value:    district.DistrictId,
			Label:    district.DistrictName,
			Children: neighbourhoods,
		}

		records = append(records, record)
	}
	return records

}
func (a *PlacesAdapter) PlacesInputListGrpcFromSql(resp []*rmsv1.Place) (*rmsv1.PlacesInputListResponse, error) {
	var records []*rmsv1.PlacesInput
	for _, city := range resp {
		districts := districtsInput(city.Districts)
		record := &rmsv1.PlacesInput{
			Value:    city.CityId,
			Label:    city.CityName,
			Children: districts,
		}

		records = append(records, record)
	}
	response := &rmsv1.PlacesInputListResponse{
		Options: records,
	}
	return response, nil
}
