/*
 * cdn
 *
 * <br/>https://ncloud.apigw.ntruss.com/cdn/v2
 *
 * API version: 2018-07-04T02:02:10Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cdn

type GetGlobalCdnInstanceListRequest struct {

	// CDN 인스턴스 번호
	CdnInstanceNo *string `json:"cdnInstanceNo,omitempty"`

	// 페이지 번호
	PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지 사이즈
	PageSize *int32 `json:"pageSize,omitempty"`
}
