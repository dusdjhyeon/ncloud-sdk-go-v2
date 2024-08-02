/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type BlockStorageSnapshotInstance struct {

	// 블록스토리지스냅샷인스턴스번호
BlockStorageSnapshotInstanceNo *string `json:"blockStorageSnapshotInstanceNo,omitempty"`

	// 블록스토리지스냅샷이름
BlockStorageSnapshotName *string `json:"blockStorageSnapshotName,omitempty"`

	// 블록스토리지스냅샷볼륨사이즈
BlockStorageSnapshotVolumeSize *int64 `json:"blockStorageSnapshotVolumeSize,omitempty"`

	// 원본블록스토리지인스턴스번호
OriginalBlockStorageInstanceNo *string `json:"originalBlockStorageInstanceNo,omitempty"`

	// 블록스토리지스냅샷인스턴스상태
BlockStorageSnapshotInstanceStatus *CommonCode `json:"blockStorageSnapshotInstanceStatus,omitempty"`

	// 블록스토리지스냅샷인스턴스OP
BlockStorageSnapshotInstanceOperation *CommonCode `json:"blockStorageSnapshotInstanceOperation,omitempty"`

	// 블록스토리지스냅샷인스턴스상태이름
BlockStorageSnapshotInstanceStatusName *string `json:"blockStorageSnapshotInstanceStatusName,omitempty"`

	// 생성일시
CreateDate *string `json:"createDate,omitempty"`

	// 원본블록스토리지볼륨암호화여부
IsEncryptedOriginalBlockStorageVolume *bool `json:"isEncryptedOriginalBlockStorageVolume,omitempty"`

	// 블록스토리지스냅샷설명
BlockStorageSnapshotDescription *string `json:"blockStorageSnapshotDescription,omitempty"`

	// 스냅샷유형
SnapshotType *CommonCode `json:"snapshotType,omitempty"`

	// 베이스스냅샷인스턴스번호
BaseSnapshotInstanceNo *string `json:"baseSnapshotInstanceNo,omitempty"`

	// 스냅샷일련Depth
SnapshotChainDepth *int32 `json:"snapshotChainDepth,omitempty"`

	// 하이퍼바이저타입
HypervisorType *string `json:"hypervisorType,omitempty"`

	// 부팅가능 여부
IsBootable *bool `json:"isBootable,omitempty"`
}
