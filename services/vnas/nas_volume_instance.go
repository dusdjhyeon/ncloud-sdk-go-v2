/*
 * vnas
 *
 * VPC NAS 관련 API<br/>https://ncloud.apigw.ntruss.com/vnas/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vnas

type NasVolumeInstance struct {

	// NAS볼륨인스턴스번호
NasVolumeInstanceNo *string `json:"nasVolumeInstanceNo,omitempty"`

	// NAS볼륨인스턴스상태
NasVolumeInstanceStatus *CommonCode `json:"nasVolumeInstanceStatus,omitempty"`

	// 서버설명
NasVolumeInstanceOperation *CommonCode `json:"nasVolumeInstanceOperation,omitempty"`

	// NAS볼륨인스턴스상태이름
NasVolumeInstanceStatusName *string `json:"nasVolumeInstanceStatusName,omitempty"`

	// 생성일시
CreateDate *string `json:"createDate,omitempty"`

	// NAS볼륨설명
NasVolumeDescription *string `json:"nasVolumeDescription,omitempty"`

	// 마운트정보
MountInformation *string `json:"mountInformation,omitempty"`

	// 볼륨할당프로토콜유형
VolumeAllotmentProtocolType *CommonCode `json:"volumeAllotmentProtocolType,omitempty"`

	// 볼륨이름
VolumeName *string `json:"volumeName,omitempty"`

	// 볼륨총사이즈
VolumeTotalSize *int64 `json:"volumeTotalSize,omitempty"`

	// 볼륨사이즈
VolumeSize *int64 `json:"volumeSize,omitempty"`

	// 스냅샷볼륨설정비율
SnapshotVolumeConfigurationRatio *float32 `json:"snapshotVolumeConfigurationRatio,omitempty"`

	// 스냅샷볼륨설정기간유형
SnapshotVolumeConfigPeriodType *CommonCode `json:"snapshotVolumeConfigPeriodType,omitempty"`

	// 스냅샷자동생성주기요일유형
SnapshotVolumeConfigDayOfWeekType *CommonCode `json:"snapshotVolumeConfigDayOfWeekType,omitempty"`

	// 스냅샷볼륨설정시간
SnapshotVolumeConfigTime *int32 `json:"snapshotVolumeConfigTime,omitempty"`

	// 스냅샷볼륨사이즈
SnapshotVolumeSize *int64 `json:"snapshotVolumeSize,omitempty"`

	// 스냅샷설정여부
IsSnapshotConfiguration *bool `json:"isSnapshotConfiguration,omitempty"`

	// 이벤트설정여부
IsEventConfiguration *bool `json:"isEventConfiguration,omitempty"`

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// ZONE코드
ZoneCode *string `json:"zoneCode,omitempty"`

	// 초기화스크립트번호
InitScriptNo *string `json:"initScriptNo,omitempty"`

	// NAS볼륨서버인스턴스번호리스트
NasVolumeServerInstanceNoList []*string `json:"nasVolumeServerInstanceNoList,omitempty"`

	// 볼륨암호화여부
IsEncryptedVolume *bool `json:"isEncryptedVolume,omitempty"`

NasVolumeInstanceCustomIpList []*NasVolumeInstanceCustomIp `json:"nasVolumeInstanceCustomIpList,omitempty"`

	// 반납보호여부
IsReturnProtection *bool `json:"isReturnProtection,omitempty"`
}
