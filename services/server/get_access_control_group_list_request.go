/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-09-28T05:05:08Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type GetAccessControlGroupListRequest struct {

	// 접근제어그룹설정번호리스트
AccessControlGroupConfigurationNoList []*string `json:"accessControlGroupConfigurationNoList,omitempty"`

	// 디폴트여부
IsDefault *bool `json:"isDefault,omitempty"`

	// 접근제어그룹명
AccessControlGroupName *string `json:"accessControlGroupName,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`
}
