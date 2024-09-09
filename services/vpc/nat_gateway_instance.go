/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.apigw.ntruss.com/vpc/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type NatGatewayInstance struct {

	// VPC번호
VpcNo *string `json:"vpcNo,omitempty"`

	// VPC이름
VpcName *string `json:"vpcName,omitempty"`

	// NATGateway인스턴스번호
NatGatewayInstanceNo *string `json:"natGatewayInstanceNo,omitempty"`

	// NATGateway이름
NatGatewayName *string `json:"natGatewayName,omitempty"`

	// 공인IP주소
PublicIp *string `json:"publicIp,omitempty"`

	// NATGateway인스턴스상태
NatGatewayInstanceStatus *CommonCode `json:"natGatewayInstanceStatus,omitempty"`

	// NATGateway인스턴스상태이름
NatGatewayInstanceStatusName *string `json:"natGatewayInstanceStatusName,omitempty"`

	// NATGateway인스턴스OP
NatGatewayInstanceOperation *CommonCode `json:"natGatewayInstanceOperation,omitempty"`

	// 생성일시
CreateDate *string `json:"createDate,omitempty"`

	// NATGateway설명
NatGatewayDescription *string `json:"natGatewayDescription,omitempty"`

	// ZONE코드
ZoneCode *string `json:"zoneCode,omitempty"`

	// NATGateay 유형
NatGatewayType *CommonCode `json:"natGatewayType,omitempty"`

	// Subnet번호
SubnetNo *string `json:"subnetNo,omitempty"`

	// Subnet이름
SubnetName *string `json:"subnetName,omitempty"`

	// 공인아이피인스턴스번호
PublicIpInstanceNo *string `json:"publicIpInstanceNo,omitempty"`

	// 사설IP주소
PrivateIp *string `json:"privateIp,omitempty"`
}
