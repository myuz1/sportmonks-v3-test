package sportmonks_v3_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var bookmakersResponse = `{
	"data": [
		{
		  "id": 1,
		  "name": "10Bet",
		  "logo": null
		},
		{
		  "id": 2,
		  "name": "bet365",
		  "logo": null
		}
	]
}
`

var bookmakerResponse = `{
	"data": {
		"id": 1,
    	"name": "10Bet",
    	"logo": null
	}
}
`

func TestBookmakers(t *testing.T) {
	url := defaultBaseURL + "/bookmakers?api_token=api-key"

	t.Run("returns bookmaker struct slice", func(t *testing.T) {
		server := mockResponseServer(t, bookmakersResponse, 200, url)

		client := newTestHTTPClient(server)

		bookmakers, _, err := client.Bookmakers(context.Background())

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertBookmaker(t, &bookmakers[0])
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		bookmakers, _, err := client.Bookmakers(context.Background())

		if bookmakers != nil {
			t.Fatalf("Test failed, expected nil, got %+v", bookmakers)
		}

		assert.Equal(
			t,
			"Request failed with message: The requested endpoint does not exists!, code: 404",
			err.Error(),
		)
	})
}

func TestBookMakerByID(t *testing.T) {
	url := defaultBaseURL + "/bookmakers/1?api_token=api-key"

	t.Run("returns a single bookmaker struct", func(t *testing.T) {
		server := mockResponseServer(t, bookmakerResponse, 200, url)

		client := newTestHTTPClient(server)

		bookmaker, _, err := client.BookmakerByID(context.Background(), 1)

		if err != nil {
			t.Fatalf("Test failed, expected nil, got %s", err.Error())
		}

		assertBookmaker(t, bookmaker)
	})

	t.Run("returns bad status code error", func(t *testing.T) {
		server := mockResponseServer(t, errorResponse, 400, url)

		client := newTestHTTPClient(server)

		bookmaker, _, err := client.BookmakerByID(context.Background(), 1)

		if bookmaker != nil {
			t.Fatalf("Test failed, expected nil, got %+v", bookmaker)
		}

		assert.Equal(
			t,
			"Request failed with message: The requested endpoint does not exists!, code: 404",
			err.Error(),
		)
	})
}

func assertBookmaker(t *testing.T, bookmaker *Bookmaker) {
	assert.Equal(t, 1, bookmaker.ID)
	assert.Equal(t, "10Bet", bookmaker.Name)
	assert.Nil(t, bookmaker.Logo)
}
