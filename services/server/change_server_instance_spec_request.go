/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-09-28T05:05:08Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type ChangeServerInstanceSpecRequest struct {

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo"`

	// 서버상품코드
ServerProductCode *string `json:"serverProductCode"`
}
