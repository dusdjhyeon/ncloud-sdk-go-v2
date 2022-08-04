/*
 * sourcebuild
 *
 * <br/>https://sourcebuild.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package sourcebuild

type GetProjectDetailResponseArtifact struct {
	Use *bool `json:"use,omitempty"`

	Path []*string `json:"path,omitempty"`

	Storage *GetProjectDetailResponseArtifactStorage `json:"storage,omitempty"`

	Backup *bool `json:"backup,omitempty"`
}
