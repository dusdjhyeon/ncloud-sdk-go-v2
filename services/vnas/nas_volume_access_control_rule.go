/*
 * vnas
 *
 * VPC NAS 관련 API<br/>https://ncloud.apigw.ntruss.com/vnas/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vnas

type NasVolumeAccessControlRule struct {

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`

	// ReadAccess여부
ReadAccess *bool `json:"readAccess,omitempty"`

	// WriteAccess여부
WriteAccess *bool `json:"writeAccess,omitempty"`
}
