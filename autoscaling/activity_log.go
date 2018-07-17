/*
 * autoscaling
 *
 * <br/>https://ncloud.apigw.ntruss.com/autoscaling/v2
 *
 * API version: 2018-06-21T02:22:22Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package autoscaling

type ActivityLog struct {

	// 액티비티번호
	ActivityNo *string `json:"activityNo,omitempty"`

	// 오토스케일링그룹명
	AutoScalingGroupName *string `json:"autoScalingGroupName,omitempty"`

	// 상태
	Status **CommonCode `json:"status,omitempty"`

	// 상태메세지
	StatusMessage *string `json:"statusMessage,omitempty"`

	// 액션원인
	ActionCause *string `json:"actionCause,omitempty"`

	// 설명
	Description *string `json:"description,omitempty"`

	// 상세설명
	Details *string `json:"details,omitempty"`

	// 시작일시
	StartTime *string `json:"startTime,omitempty"`

	// 종료일시
	EndTime *string `json:"endTime,omitempty"`
}
