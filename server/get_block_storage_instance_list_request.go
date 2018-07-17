/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-07-02T09:06:17Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type GetBlockStorageInstanceListRequest struct {

	// 서버인스턴스번호
	ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`

	// 블록스토리지인스턴스번호리스트
	BlockStorageInstanceNoList *[]string `json:"blockStorageInstanceNoList,omitempty"`

	// 검색할필터명
	SearchFilterName *string `json:"searchFilterName,omitempty"`

	// 검색할필터값
	SearchFilterValue *string `json:"searchFilterValue,omitempty"`

	// 블록스토리지구분코드리스트
	BlockStorageTypeCodeList *[]string `json:"blockStorageTypeCodeList,omitempty"`

	// 페이지번호
	PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
	PageSize *int32 `json:"pageSize,omitempty"`

	// 블록스토리지인스턴스상태코드
	BlockStorageInstanceStatusCode *string `json:"blockStorageInstanceStatusCode,omitempty"`

	// 디스크유형코드
	DiskTypeCode *string `json:"diskTypeCode,omitempty"`

	// 디스크유형상세코드
	DiskDetailTypeCode *string `json:"diskDetailTypeCode,omitempty"`

	// 리전번호
	RegionNo *string `json:"regionNo,omitempty"`

	// ZONE번호
	ZoneNo *string `json:"zoneNo,omitempty"`

	// 소팅대상
	SortedBy *string `json:"sortedBy,omitempty"`

	// 소팅순서
	SortingOrder *string `json:"sortingOrder,omitempty"`
}
