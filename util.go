package federation

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	federationExampleDirectoryRelativePath = "examples/federation"

	QueryReviewsOfMe = `query ReviewsOfMe {
  me {
    history {
      body
      product {
        upc
        name
        price
      }
    }
  }
}`
	SubscriptionUpdatedPrice = `subscription UpdatedPrice {
  updatedPrice {
    name
    price
  }
}`
)

type Upstream string

const (
	UpstreamAccounts Upstream = "base"
	UpstreamProducts Upstream = "products"
	UpstreamReviews  Upstream = "history"
)

func LoadSDLFromExamplesDirectoryWithinPkg(upstream Upstream) ([]byte, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	absolutePath := filepath.Join(strings.Split(wd, "pkg")[0], federationExampleDirectoryRelativePath, string(upstream), "graph", "schema.graphqls")
	return ioutil.ReadFile(absolutePath)
}
