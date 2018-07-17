/*
 * loadbalancer
 *
 * <br/>https://ncloud.apigw.ntruss.com/loadbalancer/v2
 *
 * API version: 2018-06-21T02:19:18Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package loadbalancer

type SslCertificate struct {

	// 인증서명
	CertificateName *string `json:"certificateName,omitempty"`

	// 비밀키
	PrivateKey *string `json:"privateKey,omitempty"`

	// 공개키인증서
	PublicKeyCertificate *string `json:"publicKeyCertificate,omitempty"`

	// chainca
	CertificateChain *string `json:"certificateChain,omitempty"`
}
