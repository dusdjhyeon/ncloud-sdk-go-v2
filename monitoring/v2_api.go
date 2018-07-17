/*
 * monitoring
 *
 * <br/>https://ncloud.apigw.ntruss.com/monitoring/v2
 *
 * API version: 2018-06-25T02:38:27Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package monitoring

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"golang.org/x/net/context"
)

// Linger please
var (
	_ context.Context
)

type V2ApiService service

/* V2ApiService
B.메트릭 리스트 조회
* @param ctx context.Context for authentication, logging, tracing, etc.
@param getListMetricsRequest getListMetricsRequest
@return GetListMetricsResponse*/
func (a *V2ApiService) GetListMetrics(ctx context.Context, getListMetricsRequest GetListMetricsRequest) (GetListMetricsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     GetListMetricsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/getListMetrics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &getListMetricsRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* V2ApiService
A.메트릭 통계 조회
* @param ctx context.Context for authentication, logging, tracing, etc.
@param getMetricStatisticsRequest getMetricStatisticsRequest
@return GetMetricStatisticsResponse*/
func (a *V2ApiService) GetMetricStatistics(ctx context.Context, getMetricStatisticsRequest GetMetricStatisticsRequest) (GetMetricStatisticsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     GetMetricStatisticsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/getMetricStatistics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &getMetricStatisticsRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}
