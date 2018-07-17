/*
 * clouddb
 *
 * Cloud DB<br/>https://ncloud.apigw.ntruss.com/clouddb/v2
 *
 * API version: 2018-07-02T10:10:19Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package clouddb

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
CloudDB인스턴스생성
* @param ctx context.Context for authentication, logging, tracing, etc.
@param createCloudDBInstanceRequest createCloudDBInstanceRequest
@return CreateCloudDbInstanceResponse*/
func (a *V2ApiService) CreateCloudDBInstance(ctx context.Context, createCloudDBInstanceRequest CreateCloudDbInstanceRequest) (CreateCloudDbInstanceResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     CreateCloudDbInstanceResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/createCloudDBInstance"

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
	localVarPostBody = &createCloudDBInstanceRequest
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
CloudDB서버인스턴스삭제
* @param ctx context.Context for authentication, logging, tracing, etc.
@param deleteCloudDBServerInstanceRequest deleteCloudDBServerInstanceRequest
@return DeleteCloudDbServerInstanceResponse*/
func (a *V2ApiService) DeleteCloudDBServerInstance(ctx context.Context, deleteCloudDBServerInstanceRequest DeleteCloudDbServerInstanceRequest) (DeleteCloudDbServerInstanceResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     DeleteCloudDbServerInstanceResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/deleteCloudDBServerInstance"

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
	localVarPostBody = &deleteCloudDBServerInstanceRequest
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
CloudDB설정그룹리스트조회
* @param ctx context.Context for authentication, logging, tracing, etc.
@param getCloudDBConfigGroupListRequest getCloudDBConfigGroupListRequest
@return GetCloudDbConfigGroupListResponse*/
func (a *V2ApiService) GetCloudDBConfigGroupList(ctx context.Context, getCloudDBConfigGroupListRequest GetCloudDbConfigGroupListRequest) (GetCloudDbConfigGroupListResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     GetCloudDbConfigGroupListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/getCloudDBConfigGroupList"

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
	localVarPostBody = &getCloudDBConfigGroupListRequest
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
CloudDB이미지상품리스트
* @param ctx context.Context for authentication, logging, tracing, etc.
@param getCloudDBImageProductListRequest getCloudDBImageProductListRequest
@return GetCloudDbImageProductListResponse*/
func (a *V2ApiService) GetCloudDBImageProductList(ctx context.Context, getCloudDBImageProductListRequest GetCloudDbImageProductListRequest) (GetCloudDbImageProductListResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     GetCloudDbImageProductListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/getCloudDBImageProductList"

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
	localVarPostBody = &getCloudDBImageProductListRequest
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
CloudDB인스턴스리스트조회
* @param ctx context.Context for authentication, logging, tracing, etc.
@param getCloudDBInstanceListRequest getCloudDBInstanceListRequest
@return GetCloudDbInstanceListResponse*/
func (a *V2ApiService) GetCloudDBInstanceList(ctx context.Context, getCloudDBInstanceListRequest GetCloudDbInstanceListRequest) (GetCloudDbInstanceListResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     GetCloudDbInstanceListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/getCloudDBInstanceList"

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
	localVarPostBody = &getCloudDBInstanceListRequest
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
CloudDB상품리스트조회
* @param ctx context.Context for authentication, logging, tracing, etc.
@param getCloudDBProductListRequest getCloudDBProductListRequest
@return GetCloudDbProductListResponse*/
func (a *V2ApiService) GetCloudDBProductList(ctx context.Context, getCloudDBProductListRequest GetCloudDbProductListRequest) (GetCloudDbProductListResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     GetCloudDbProductListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/getCloudDBProductList"

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
	localVarPostBody = &getCloudDBProductListRequest
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
CloudDB서버인스턴스재부팅
* @param ctx context.Context for authentication, logging, tracing, etc.
@param rebootCloudDBServerInstanceRequest rebootCloudDBServerInstanceRequest
@return RebootCloudDbServerInstanceResponse*/
func (a *V2ApiService) RebootCloudDBServerInstance(ctx context.Context, rebootCloudDBServerInstanceRequest RebootCloudDbServerInstanceRequest) (RebootCloudDbServerInstanceResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     RebootCloudDbServerInstanceResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/rebootCloudDBServerInstance"

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
	localVarPostBody = &rebootCloudDBServerInstanceRequest
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
