/*
 * clouddb
 *
 * Cloud DB<br/>https://ncloud.apigw.ntruss.com/clouddb/v2
 *
 * API version: 2018-07-02T10:10:19Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package clouddb

type CreateCloudDbInstanceRequest struct {

	// CloudDB이미지상품코드
	CloudDBImageProductCode *string `json:"cloudDBImageProductCode,omitempty"`

	// CloudDB상품코드
	CloudDBProductCode *string `json:"cloudDBProductCode,omitempty"`

	// DB유형코드
	DbKindCode *string `json:"dbKindCode"`

	// Collation
	Collation *string `json:"collation,omitempty"`

	// 데이터스토리지타입
	DataStorageTypeCode *string `json:"dataStorageTypeCode,omitempty"`

	// HA여부
	IsHa *bool `json:"isHa,omitempty"`

	// 호스트IP
	HostIp *string `json:"hostIp,omitempty"`

	// CloudDB서버이름
	CloudDBServerName *string `json:"cloudDBServerName,omitempty"`

	// CloudDB서비스이름
	CloudDBServiceName *string `json:"cloudDBServiceName"`

	// CloudDB기본이름
	CloudDBBasicName *string `json:"cloudDBBasicName,omitempty"`

	// CloudDB유저이름
	CloudDBUserName *string `json:"cloudDBUserName,omitempty"`

	// CloudDB유저패스워드
	CloudDBUserPassword *string `json:"cloudDBUserPassword,omitempty"`

	// CloudDB포트
	CloudDBPort *int32 `json:"cloudDBPort,omitempty"`

	// CloudDB설정그룹번호
	CloudDBConfigGroupNo *string `json:"cloudDBConfigGroupNo,omitempty"`

	// 백업여부
	IsBackup *bool `json:"isBackup,omitempty"`

	// 백업파일보관기간
	BackupFileRetentionPeriod *int32 `json:"backupFileRetentionPeriod,omitempty"`

	// 자동Backup여부
	IsAutomaticBackup *bool `json:"isAutomaticBackup,omitempty"`

	// 백업시간
	BackupTime *string `json:"backupTime,omitempty"`

	// 리전번호
	RegionNo *string `json:"regionNo,omitempty"`

	// zone번호
	ZoneNo *string `json:"zoneNo,omitempty"`
}
