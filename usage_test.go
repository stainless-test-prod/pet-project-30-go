// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package petproject30_test

import (
	"context"
	"os"
	"testing"

	"github.com/miriambudayr/pet-project-30-go"
	"github.com/miriambudayr/pet-project-30-go/internal/testutil"
	"github.com/miriambudayr/pet-project-30-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := petproject30.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	t.Skip("Prism tests are disabled")
	pets, err := client.Pets.List(context.TODO(), petproject30.PetListParams{})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", pets)
}
