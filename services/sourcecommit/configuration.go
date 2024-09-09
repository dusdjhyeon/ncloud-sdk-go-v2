/*
 * api
 *
 * <br/>https://sourcecommit.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package sourcecommit

import (
	"os"
	"strings"

	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

func NewConfiguration(region string, apiKeys ...*ncloud.APIKey) *ncloud.Configuration {
	cfg := &ncloud.Configuration{
		BasePath:      "https://sourcecommit.apigw.ntruss.com/api/v1",
		DefaultHeader: make(map[string]string),
		UserAgent:     "sourcecommit/1.0.0/go",
	}
	if len(apiKeys) > 0 {
		cfg.APIKey = apiKeys[0]
	}

	cfg.InitCredentials()
	if os.Getenv("NCLOUD_API_GW") != "" {
		cfg.BasePath = os.Getenv("NCLOUD_API_GW") + "/api/v1"
		cfg.BasePath = strings.Replace(cfg.BasePath, "fin-ncloud", "sourcecommit", 1)
		cfg.BasePath = strings.Replace(cfg.BasePath, "ncloud", "sourcecommit", 1)
	}
	switch region {
	case "SGN":
		cfg.BasePath = strings.Replace(cfg.BasePath, "/v1", "/sg-v1", 1)
	}

	return cfg
}
