package services

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type BuilderSheetService struct {
	spreadsheetId string
	readRange string
	srv *sheets.Service
}

func NewBuilderSheets(ctx context.Context, googleClient *http.Client, spreadsheetId, readRange string) (*BuilderSheetService, error) {
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(googleClient))
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return &BuilderSheetService{
		spreadsheetId,
		readRange,
		srv,
	}, nil
}

func (ESS *BuilderSheetService) GetTexts(ctx context.Context) ([]string, error) {
	res, err := ESS.srv.Spreadsheets.Values.Get(ESS.spreadsheetId, ESS.readRange).Do()
	if err != nil || res.HTTPStatusCode != 200 {
		return nil, fmt.Errorf("%v", err)
	}

	texts := []string{}

	for _, val := range res.Values {
		for _, v := range val {
			texts = append(texts, fmt.Sprintf("%v", v))
		}
	}

	return texts, nil
}