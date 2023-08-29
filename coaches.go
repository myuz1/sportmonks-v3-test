package sportmonks_v3_test

import (
	"context"
	"fmt"
	"net/url"
)

// Coach provides a struct representation of a Coach resource
type Coach struct {
	ID           int     `json:"coach_id"`
	TeamID       int     `json:"team_id"`
	CountryID    int     `json:"country_id"`
	CommonName   string  `json:"common_name"`
	FullName     string  `json:"fullname"`
	FirstName    string  `json:"firstname"`
	LastName     string  `json:"lastname"`
	Nationality  string  `json:"nationality"`
	BirthDate    string  `json:"birthdate"`
	BirthCountry string  `json:"birthcountry"`
	BirthPlace   *string `json:"birthplace"`
	ImagePath    string  `json:"image_path"`
}

// CoachByID fetches a Coach resource by ID.
func (c *HTTPClient) CoachByID(ctx context.Context, id int) (*Coach, *Meta, error) {
	path := fmt.Sprintf(coachesURI+"/%d", id)

	response := struct {
		Data *Coach `json:"data"`
		Meta *Meta  `json:"meta"`
	}{}

	err := c.getResource(ctx, path, url.Values{}, &response)

	if err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, err
}
