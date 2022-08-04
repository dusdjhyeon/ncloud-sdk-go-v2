/*
 * sourcepipeline
 *
 * <br/>https://sourcepipeline.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package sourcepipeline

type ChangeProjectTasks struct {
	Name *string `json:"name"`

	Type_ *string `json:"type"`

	Config *ChangeProjectConfig `json:"config"`

	LinkedTasks []*string `json:"linkedTasks"`
}
