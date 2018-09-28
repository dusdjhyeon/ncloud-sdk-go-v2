/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-09-28T05:05:08Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type NasVolumeInstanceRating struct {

	// 측정시간
RatingTime *string `json:"ratingTime,omitempty"`

	// 볼륨사이즈
VolumeSize *int64 `json:"volumeSize,omitempty"`

	// 볼륨사용사이즈
VolumeUseSize *int64 `json:"volumeUseSize,omitempty"`

	// 볼륨사용비율
VolumeUseRatio *float32 `json:"volumeUseRatio,omitempty"`

	// 평균볼륨사이즈
AverageVolumeSize *int64 `json:"averageVolumeSize,omitempty"`

	// 평균볼륨사용사이즈
AverageVolumeUseSize *int64 `json:"averageVolumeUseSize,omitempty"`

	// 평균볼륨사용비율
AverageVolumeUseRatio *float32 `json:"averageVolumeUseRatio,omitempty"`

	// 최대볼륨사용사이즈
MaxVolumeUseSize *int64 `json:"maxVolumeUseSize,omitempty"`

	// 최대볼륨사용비율
MaxVolumeUseRatio *float32 `json:"maxVolumeUseRatio,omitempty"`

	// 최소볼륨사용사이즈
MinVolumeUseSize *int64 `json:"minVolumeUseSize,omitempty"`

	// 최소볼륨사용비율
MinVolumeUseRatio *float32 `json:"minVolumeUseRatio,omitempty"`

	// 스냅샷볼륨사이즈
SnapshotVolumeSize *int64 `json:"snapshotVolumeSize,omitempty"`

	// 스냅샷볼륨사용사이즈
SnapshotVolumeUseSize *int64 `json:"snapshotVolumeUseSize,omitempty"`

	// 스냅샷볼륨사용비율
SnapshotVolumeUseRatio *float32 `json:"snapshotVolumeUseRatio,omitempty"`

	// 평균스냅샷볼륨사이즈
SnapshotAverageVolumeSize *int64 `json:"snapshotAverageVolumeSize,omitempty"`

	// 평균스냅샷볼륨사용사이즈
SnapshotAverageVolumeUseSize *int64 `json:"snapshotAverageVolumeUseSize,omitempty"`

	// 평균스냅샷볼륨사용비율
SnapshotAverageVolumeUseRatio *float32 `json:"snapshotAverageVolumeUseRatio,omitempty"`

	// 최대스냅샷볼륨사용사이즈
SnapshotMaxVolumeUseSize *int64 `json:"snapshotMaxVolumeUseSize,omitempty"`

	// 최대스냅샷볼륨사용비율
SnapshotMaxVolumeUseRatio *float32 `json:"snapshotMaxVolumeUseRatio,omitempty"`

	// 최소스냅샷볼륨사용사이즈
SnapshotMinVolumeUseSize *int64 `json:"snapshotMinVolumeUseSize,omitempty"`

	// 최소스냅샷볼륨사용비율
SnapshotMinVolumeUseRatio *float32 `json:"snapshotMinVolumeUseRatio,omitempty"`
}
