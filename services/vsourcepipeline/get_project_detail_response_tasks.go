/*
 * vsourcepipeline
 *
 * <br/>https://vpcsourcepipeline.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vsourcepipeline

type GetProjectDetailResponseTasks struct {
	Id *int32 `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Type_ *string `json:"type,omitempty"`

	Config *GetProjectDetailResponseConfig `json:"config,omitempty"`

	LinkedTasks []*string `json:"linkedTasks,omitempty"`
}
