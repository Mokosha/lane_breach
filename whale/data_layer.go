package main

//
// Code for accessing data backend
//
type DataLayer interface {
	// Take arguments and return data from backend
	GetData(args *GetDataArgs) (DataResponse, error)
}

//
// GetDataArgs holds arguments for fetching data
//
type GetDataArgs struct {
	Table string
	Key   string
}

//
// DataResponse holds data fetched from backend
//
type DataResponse struct {
	Data interface{}
}

//
// dataFetcher _implements_ DataLayer
//
type dataFetcher struct {
	// Probably has attached dynamo client, configuration, etc
}

func NewDataLayer() (DataLayer, error) {
	return &dataFetcher{}, nil
}

//
// GetData -- TODO actually implement this for dynamo
//
func (d *dataFetcher) GetData(args *GetDataArgs) (DataResponse, error) {
	return DataResponse{}, nil
}
