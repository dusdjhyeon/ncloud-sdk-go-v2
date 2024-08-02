/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type DisableFlowLogRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 네트워크인터페이스번호
NetworkInterfaceNo *string `json:"networkInterfaceNo"`
}
