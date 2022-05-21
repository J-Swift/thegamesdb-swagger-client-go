package gamesdb_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	gamesdb "github.com/__GIT_USERNAME__/__GIT_REPONAME__"
)

type apiCallResult struct {
	endpoint string
	resp     *http.Response
	err      error
}

func TestApiClient(t *testing.T) {
	apikey := os.Getenv("GAMESDB_APIKEY")
	if len(apikey) == 0 {
		t.Log("No api key for TheGamesDB. Please set GAMESDB_APIKEY environment variable")
		t.FailNow()
	}

	config := gamesdb.NewConfiguration()
	client := gamesdb.NewAPIClient(config)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	apicalls := []apiCallResult{}

	// GamesAPI

	t.Log("Checking GamesAPI...")

	_, resp, err := client.GamesApi.GamesByGameID(ctx).Apikey(apikey).Id("1").Execute()
	apicalls = append(apicalls, apiCallResult{endpoint: "GamesByGameID", resp: resp, err: err})

	_, resp, err = client.GamesApi.GamesByGameName(ctx).Apikey(apikey).Name("halo").Execute()
	apicalls = append(apicalls, apiCallResult{endpoint: "GamesByGameName", resp: resp, err: err})

	_, resp, err = client.GamesApi.GamesByGameNameV1(ctx).Apikey(apikey).Name("halo").Execute()
	apicalls = append(apicalls, apiCallResult{endpoint: "GamesByGameName_v1", resp: resp, err: err})

	_, resp, err = client.GamesApi.GamesByPlatformID(ctx).Apikey(apikey).Id("1").Execute()
	apicalls = append(apicalls, apiCallResult{endpoint: "GamesByPlatformID", resp: resp, err: err})

	_, resp, err = client.GamesApi.GamesImages(ctx).GamesId("1").Apikey(apikey).Execute()
	apicalls = append(apicalls, apiCallResult{endpoint: "GamesImages", resp: resp, err: err})

	_, resp, err = client.GamesApi.GamesUpdates(ctx).Apikey(apikey).LastEditId(0).Execute()
	apicalls = append(apicalls, apiCallResult{endpoint: "GamesUpdates", resp: resp, err: err})

	// DevelopersApi

	t.Log("Checking DevelopersApi...")

	_, resp, err = client.DevelopersApi.Developers(ctx).Apikey(apikey).Execute()
	apicalls = append(apicalls, apiCallResult{endpoint: "Developers", resp: resp, err: err})

	// GenresApi

	t.Log("Checking GenresApi...")

	_, resp, err = client.GenresApi.Genres(ctx).Apikey(apikey).Execute()
	apicalls = append(apicalls, apiCallResult{endpoint: "Genres", resp: resp, err: err})

	// PlatformsApi

	t.Log("Checking PlatformsApi...")

	_, resp, err = client.PlatformsApi.Platforms(ctx).Apikey(apikey).Execute()
	apicalls = append(apicalls, apiCallResult{endpoint: "Platforms", resp: resp, err: err})

	_, resp, err = client.PlatformsApi.PlatformsByPlatformID(ctx).Apikey(apikey).Id(1).Execute()
	apicalls = append(apicalls, apiCallResult{endpoint: "PlatformsByPlatformID", resp: resp, err: err})

	_, resp, err = client.PlatformsApi.PlatformsByPlatformName(ctx).Apikey(apikey).Name("3do").Execute()
	apicalls = append(apicalls, apiCallResult{endpoint: "PlatformsByPlatformName", resp: resp, err: err})

	_, resp, err = client.PlatformsApi.PlatformsImages(ctx).Apikey(apikey).PlatformsId("1").Execute()
	apicalls = append(apicalls, apiCallResult{endpoint: "PlatformsImages", resp: resp, err: err})

	// PublishersApi

	t.Log("Checking PublishersApi...")

	_, resp, err = client.PublishersApi.Publishers(ctx).Apikey(apikey).Execute()
	apicalls = append(apicalls, apiCallResult{endpoint: "Publishers", resp: resp, err: err})

	for _, apiCall := range apicalls {
		if apiCall.err != nil || apiCall.resp.StatusCode != 200 {
			t.Logf("%s call failed with status code [%d]", apiCall.endpoint, apiCall.resp.StatusCode)
			if apiCall.err != nil {
				t.Logf("    %s", err)
			}
			t.Fail()
		}
	}
}
