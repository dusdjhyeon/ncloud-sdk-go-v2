/*
 * vnas
 *
 * VPC NAS 관련 API<br/>https://ncloud.apigw.ntruss.com/vnas/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vnas

type ChangeNasVolumeSnapshotConfigurationRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// NAS볼륨인스턴스번호
NasVolumeInstanceNo *string `json:"nasVolumeInstanceNo"`

	// 스냅샷볼륨설정비율
SnapshotVolumeConfigurationRatio *int32 `json:"snapshotVolumeConfigurationRatio"`

	// 스냅샷자동생성주기요일
SnapshotVolumeConfigDayOfWeekTypeCode *string `json:"snapshotVolumeConfigDayOfWeekTypeCode,omitempty"`

	// 스냅샷자동생성주기시각
SnapshotVolumeConfigTime *int32 `json:"snapshotVolumeConfigTime,omitempty"`
}
