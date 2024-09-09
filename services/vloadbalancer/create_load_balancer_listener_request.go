/*
 * vloadbalancer
 *
 * VPC Load Balancer 관련 API<br/>https://ncloud.apigw.ntruss.com/vloadbalancer/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vloadbalancer

type CreateLoadBalancerListenerRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// SSL인증서번호
SslCertificateNo *string `json:"sslCertificateNo,omitempty"`

	// HTTP2사용여부
UseHttp2 *bool `json:"useHttp2,omitempty"`

	// 로드밸런서인스턴스번호
LoadBalancerInstanceNo *string `json:"loadBalancerInstanceNo"`

	// 포트
Port *int32 `json:"port"`

	// 프로토콜유형코드
ProtocolTypeCode *string `json:"protocolTypeCode"`

	// TLS최소지원버전유형코드
TlsMinVersionTypeCode *string `json:"tlsMinVersionTypeCode,omitempty"`

	// 타겟그룹번호
TargetGroupNo *string `json:"targetGroupNo"`

	// 암호화스위트리스트
CipherSuiteList []*string `json:"cipherSuiteList,omitempty"`
}
