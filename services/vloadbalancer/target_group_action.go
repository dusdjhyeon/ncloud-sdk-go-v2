/*
 * vloadbalancer
 *
 * VPC Load Balancer 관련 API<br/>https://ncloud.apigw.ntruss.com/vloadbalancer/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vloadbalancer

type TargetGroupAction struct {

	// 타겟그룹가중치리스트
TargetGroupWeightList []*TargetGroupWeight `json:"targetGroupWeightList,omitempty"`

	// 세션별접근사용여부
UseStickySession *bool `json:"useStickySession,omitempty"`
}
