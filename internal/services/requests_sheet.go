package services

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type RequestsSheetService struct {
	spreadsheetId string
	readRange string
	srv *sheets.Service
}

type RequestsResponse struct {
	Id string `json:"id" mapstructure:"id"`
}

func NewRequestsSheets(ctx context.Context, googleClient *http.Client, spreadsheetId, readRange string) (*RequestsSheetService, error) {
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(googleClient))
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return &RequestsSheetService{
		spreadsheetId,
		readRange,
		srv,
	}, nil
}

func (ESS *RequestsSheetService) CreateRequest(ctx context.Context, request []string) (RequestsResponse, error) {
	response := RequestsResponse {}

	res, err := ESS.srv.Spreadsheets.Values.Get(ESS.spreadsheetId, ESS.readRange).Do()
	if err != nil || res.HTTPStatusCode != 200 {
		return response, fmt.Errorf("%v", err)
	}

	tableLen := len(res.Values) + 1

	newReadRange := fmt.Sprintf("A%v:Z%v", tableLen, tableLen)

	values := [][]interface{}{}

	values = append(values, make([]interface{}, 0))

	for _, v := range request {
		values[0] = append(values[0], v)
	}

	row := &sheets.ValueRange{
		Values: values,
	}

	_, err = ESS.srv.Spreadsheets.Values.Update(ESS.spreadsheetId, newReadRange, row).ValueInputOption("USER_ENTERED").Context(ctx).Do()
	if err != nil {
		return response, fmt.Errorf("%v", err)
	}

	return response, nil
}