/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type LoginKey struct {

	// 핑거프린트
Fingerprint *string `json:"fingerprint,omitempty"`

	// 키명
KeyName *string `json:"keyName,omitempty"`

	// 생성일자
CreateDate *string `json:"createDate,omitempty"`
}
