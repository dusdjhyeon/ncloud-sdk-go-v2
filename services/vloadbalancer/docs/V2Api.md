# \V2Api

All URIs are relative to *https://ncloud.apigw.ntruss.com/vloadbalancer/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLoadBalancerListenerCertificate**](V2Api.md#AddLoadBalancerListenerCertificate) | **Post** /addLoadBalancerListenerCertificate | 
[**AddTarget**](V2Api.md#AddTarget) | **Post** /addTarget | 
[**ChangeLoadBalancerInstanceConfiguration**](V2Api.md#ChangeLoadBalancerInstanceConfiguration) | **Post** /changeLoadBalancerInstanceConfiguration | 
[**ChangeLoadBalancerListenerConfiguration**](V2Api.md#ChangeLoadBalancerListenerConfiguration) | **Post** /changeLoadBalancerListenerConfiguration | 
[**ChangeTargetGroupConfiguration**](V2Api.md#ChangeTargetGroupConfiguration) | **Post** /changeTargetGroupConfiguration | 
[**ChangeTargetGroupHealthCheckConfiguration**](V2Api.md#ChangeTargetGroupHealthCheckConfiguration) | **Post** /changeTargetGroupHealthCheckConfiguration | 
[**CreateLoadBalancerInstance**](V2Api.md#CreateLoadBalancerInstance) | **Post** /createLoadBalancerInstance | 
[**CreateLoadBalancerListener**](V2Api.md#CreateLoadBalancerListener) | **Post** /createLoadBalancerListener | 
[**CreateTargetGroup**](V2Api.md#CreateTargetGroup) | **Post** /createTargetGroup | 
[**DeleteLoadBalancerInstances**](V2Api.md#DeleteLoadBalancerInstances) | **Post** /deleteLoadBalancerInstances | 
[**DeleteLoadBalancerListeners**](V2Api.md#DeleteLoadBalancerListeners) | **Post** /deleteLoadBalancerListeners | 
[**DeleteTargetGroups**](V2Api.md#DeleteTargetGroups) | **Post** /deleteTargetGroups | 
[**GetLoadBalancerInstanceDetail**](V2Api.md#GetLoadBalancerInstanceDetail) | **Post** /getLoadBalancerInstanceDetail | 
[**GetLoadBalancerInstanceList**](V2Api.md#GetLoadBalancerInstanceList) | **Post** /getLoadBalancerInstanceList | 
[**GetLoadBalancerListenerCertificateList**](V2Api.md#GetLoadBalancerListenerCertificateList) | **Post** /getLoadBalancerListenerCertificateList | 
[**GetLoadBalancerListenerList**](V2Api.md#GetLoadBalancerListenerList) | **Post** /getLoadBalancerListenerList | 
[**GetLoadBalancerRuleList**](V2Api.md#GetLoadBalancerRuleList) | **Post** /getLoadBalancerRuleList | 
[**GetTargetGroupDetail**](V2Api.md#GetTargetGroupDetail) | **Post** /getTargetGroupDetail | 
[**GetTargetGroupList**](V2Api.md#GetTargetGroupList) | **Post** /getTargetGroupList | 
[**GetTargetList**](V2Api.md#GetTargetList) | **Post** /getTargetList | 
[**RemoveLoadBalancerListenerCertificate**](V2Api.md#RemoveLoadBalancerListenerCertificate) | **Post** /removeLoadBalancerListenerCertificate | 
[**RemoveTarget**](V2Api.md#RemoveTarget) | **Post** /removeTarget | 
[**SetLoadBalancerDescription**](V2Api.md#SetLoadBalancerDescription) | **Post** /setLoadBalancerDescription | 
[**SetLoadBalancerInstanceSubnet**](V2Api.md#SetLoadBalancerInstanceSubnet) | **Post** /setLoadBalancerInstanceSubnet | 
[**SetTarget**](V2Api.md#SetTarget) | **Post** /setTarget | 
[**SetTargetGroupDescription**](V2Api.md#SetTargetGroupDescription) | **Post** /setTargetGroupDescription | 


# **AddLoadBalancerListenerCertificate**
> AddLoadBalancerListenerCertificateResponse AddLoadBalancerListenerCertificate(addLoadBalancerListenerCertificateRequest)


로드밸런서리스너인증서추가

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**addLoadBalancerListenerCertificateRequest** | **[\*AddLoadBalancerListenerCertificateRequest](AddLoadBalancerListenerCertificateRequest.md)** | addLoadBalancerListenerCertificateRequest | 

### Return type

*[**AddLoadBalancerListenerCertificateResponse**](AddLoadBalancerListenerCertificateResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTarget**
> AddTargetResponse AddTarget(addTargetRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**addTargetRequest** | **[\*AddTargetRequest](AddTargetRequest.md)** | addTargetRequest | 

### Return type

*[**AddTargetResponse**](AddTargetResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeLoadBalancerInstanceConfiguration**
> ChangeLoadBalancerInstanceConfigurationResponse ChangeLoadBalancerInstanceConfiguration(changeLoadBalancerInstanceConfigurationRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**changeLoadBalancerInstanceConfigurationRequest** | **[\*ChangeLoadBalancerInstanceConfigurationRequest](ChangeLoadBalancerInstanceConfigurationRequest.md)** | changeLoadBalancerInstanceConfigurationRequest | 

### Return type

*[**ChangeLoadBalancerInstanceConfigurationResponse**](ChangeLoadBalancerInstanceConfigurationResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeLoadBalancerListenerConfiguration**
> ChangeLoadBalancerListenerConfigurationResponse ChangeLoadBalancerListenerConfiguration(changeLoadBalancerListenerConfigurationRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**changeLoadBalancerListenerConfigurationRequest** | **[\*ChangeLoadBalancerListenerConfigurationRequest](ChangeLoadBalancerListenerConfigurationRequest.md)** | changeLoadBalancerListenerConfigurationRequest | 

### Return type

*[**ChangeLoadBalancerListenerConfigurationResponse**](ChangeLoadBalancerListenerConfigurationResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeTargetGroupConfiguration**
> ChangeTargetGroupConfigurationResponse ChangeTargetGroupConfiguration(changeTargetGroupConfigurationRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**changeTargetGroupConfigurationRequest** | **[\*ChangeTargetGroupConfigurationRequest](ChangeTargetGroupConfigurationRequest.md)** | changeTargetGroupConfigurationRequest | 

### Return type

*[**ChangeTargetGroupConfigurationResponse**](ChangeTargetGroupConfigurationResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeTargetGroupHealthCheckConfiguration**
> ChangeTargetGroupHealthCheckConfigurationResponse ChangeTargetGroupHealthCheckConfiguration(changeTargetGroupHealthCheckConfigurationRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**changeTargetGroupHealthCheckConfigurationRequest** | **[\*ChangeTargetGroupHealthCheckConfigurationRequest](ChangeTargetGroupHealthCheckConfigurationRequest.md)** | changeTargetGroupHealthCheckConfigurationRequest | 

### Return type

*[**ChangeTargetGroupHealthCheckConfigurationResponse**](ChangeTargetGroupHealthCheckConfigurationResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoadBalancerInstance**
> CreateLoadBalancerInstanceResponse CreateLoadBalancerInstance(createLoadBalancerInstanceRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createLoadBalancerInstanceRequest** | **[\*CreateLoadBalancerInstanceRequest](CreateLoadBalancerInstanceRequest.md)** | createLoadBalancerInstanceRequest | 

### Return type

*[**CreateLoadBalancerInstanceResponse**](CreateLoadBalancerInstanceResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoadBalancerListener**
> CreateLoadBalancerListenerResponse CreateLoadBalancerListener(createLoadBalancerListenerRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createLoadBalancerListenerRequest** | **[\*CreateLoadBalancerListenerRequest](CreateLoadBalancerListenerRequest.md)** | createLoadBalancerListenerRequest | 

### Return type

*[**CreateLoadBalancerListenerResponse**](CreateLoadBalancerListenerResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTargetGroup**
> CreateTargetGroupResponse CreateTargetGroup(createTargetGroupRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**createTargetGroupRequest** | **[\*CreateTargetGroupRequest](CreateTargetGroupRequest.md)** | createTargetGroupRequest | 

### Return type

*[**CreateTargetGroupResponse**](CreateTargetGroupResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadBalancerInstances**
> DeleteLoadBalancerInstancesResponse DeleteLoadBalancerInstances(deleteLoadBalancerInstancesRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteLoadBalancerInstancesRequest** | **[\*DeleteLoadBalancerInstancesRequest](DeleteLoadBalancerInstancesRequest.md)** | deleteLoadBalancerInstancesRequest | 

### Return type

*[**DeleteLoadBalancerInstancesResponse**](DeleteLoadBalancerInstancesResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadBalancerListeners**
> DeleteLoadBalancerListenersResponse DeleteLoadBalancerListeners(deleteLoadBalancerListenersRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteLoadBalancerListenersRequest** | **[\*DeleteLoadBalancerListenersRequest](DeleteLoadBalancerListenersRequest.md)** | deleteLoadBalancerListenersRequest | 

### Return type

*[**DeleteLoadBalancerListenersResponse**](DeleteLoadBalancerListenersResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTargetGroups**
> DeleteTargetGroupsResponse DeleteTargetGroups(deleteTargetGroupsRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deleteTargetGroupsRequest** | **[\*DeleteTargetGroupsRequest](DeleteTargetGroupsRequest.md)** | deleteTargetGroupsRequest | 

### Return type

*[**DeleteTargetGroupsResponse**](DeleteTargetGroupsResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadBalancerInstanceDetail**
> GetLoadBalancerInstanceDetailResponse GetLoadBalancerInstanceDetail(getLoadBalancerInstanceDetailRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getLoadBalancerInstanceDetailRequest** | **[\*GetLoadBalancerInstanceDetailRequest](GetLoadBalancerInstanceDetailRequest.md)** | getLoadBalancerInstanceDetailRequest | 

### Return type

*[**GetLoadBalancerInstanceDetailResponse**](GetLoadBalancerInstanceDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadBalancerInstanceList**
> GetLoadBalancerInstanceListResponse GetLoadBalancerInstanceList(getLoadBalancerInstanceListRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getLoadBalancerInstanceListRequest** | **[\*GetLoadBalancerInstanceListRequest](GetLoadBalancerInstanceListRequest.md)** | getLoadBalancerInstanceListRequest | 

### Return type

*[**GetLoadBalancerInstanceListResponse**](GetLoadBalancerInstanceListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadBalancerListenerCertificateList**
> GetLoadBalancerListenerCertificateListResponse GetLoadBalancerListenerCertificateList(getLoadBalancerListenerCertificateListRequest)


로드밸런서리스너인증서조회

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getLoadBalancerListenerCertificateListRequest** | **[\*GetLoadBalancerListenerCertificateListRequest](GetLoadBalancerListenerCertificateListRequest.md)** | getLoadBalancerListenerCertificateListRequest | 

### Return type

*[**GetLoadBalancerListenerCertificateListResponse**](GetLoadBalancerListenerCertificateListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadBalancerListenerList**
> GetLoadBalancerListenerListResponse GetLoadBalancerListenerList(getLoadBalancerListenerListRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getLoadBalancerListenerListRequest** | **[\*GetLoadBalancerListenerListRequest](GetLoadBalancerListenerListRequest.md)** | getLoadBalancerListenerListRequest | 

### Return type

*[**GetLoadBalancerListenerListResponse**](GetLoadBalancerListenerListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadBalancerRuleList**
> GetLoadBalancerRuleListResponse GetLoadBalancerRuleList(getLoadBalancerRuleListRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getLoadBalancerRuleListRequest** | **[\*GetLoadBalancerRuleListRequest](GetLoadBalancerRuleListRequest.md)** | getLoadBalancerRuleListRequest | 

### Return type

*[**GetLoadBalancerRuleListResponse**](GetLoadBalancerRuleListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTargetGroupDetail**
> GetTargetGroupDetailResponse GetTargetGroupDetail(getTargetGroupDetailRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getTargetGroupDetailRequest** | **[\*GetTargetGroupDetailRequest](GetTargetGroupDetailRequest.md)** | getTargetGroupDetailRequest | 

### Return type

*[**GetTargetGroupDetailResponse**](GetTargetGroupDetailResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTargetGroupList**
> GetTargetGroupListResponse GetTargetGroupList(getTargetGroupListRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getTargetGroupListRequest** | **[\*GetTargetGroupListRequest](GetTargetGroupListRequest.md)** | getTargetGroupListRequest | 

### Return type

*[**GetTargetGroupListResponse**](GetTargetGroupListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTargetList**
> GetTargetListResponse GetTargetList(getTargetListRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**getTargetListRequest** | **[\*GetTargetListRequest](GetTargetListRequest.md)** | getTargetListRequest | 

### Return type

*[**GetTargetListResponse**](GetTargetListResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveLoadBalancerListenerCertificate**
> RemoveLoadBalancerListenerCertificateResponse RemoveLoadBalancerListenerCertificate(removeLoadBalancerListenerCertificateRequest)


로드밸런서리스너인증서제거

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**removeLoadBalancerListenerCertificateRequest** | **[\*RemoveLoadBalancerListenerCertificateRequest](RemoveLoadBalancerListenerCertificateRequest.md)** | removeLoadBalancerListenerCertificateRequest | 

### Return type

*[**RemoveLoadBalancerListenerCertificateResponse**](RemoveLoadBalancerListenerCertificateResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveTarget**
> RemoveTargetResponse RemoveTarget(removeTargetRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**removeTargetRequest** | **[\*RemoveTargetRequest](RemoveTargetRequest.md)** | removeTargetRequest | 

### Return type

*[**RemoveTargetResponse**](RemoveTargetResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetLoadBalancerDescription**
> SetLoadBalancerDescriptionResponse SetLoadBalancerDescription(setLoadBalancerDescriptionRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**setLoadBalancerDescriptionRequest** | **[\*SetLoadBalancerDescriptionRequest](SetLoadBalancerDescriptionRequest.md)** | setLoadBalancerDescriptionRequest | 

### Return type

*[**SetLoadBalancerDescriptionResponse**](SetLoadBalancerDescriptionResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetLoadBalancerInstanceSubnet**
> SetLoadBalancerInstanceSubnetResponse SetLoadBalancerInstanceSubnet(setLoadBalancerInstanceSubnetRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**setLoadBalancerInstanceSubnetRequest** | **[\*SetLoadBalancerInstanceSubnetRequest](SetLoadBalancerInstanceSubnetRequest.md)** | setLoadBalancerInstanceSubnetRequest | 

### Return type

*[**SetLoadBalancerInstanceSubnetResponse**](SetLoadBalancerInstanceSubnetResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetTarget**
> SetTargetResponse SetTarget(setTargetRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**setTargetRequest** | **[\*SetTargetRequest](SetTargetRequest.md)** | setTargetRequest | 

### Return type

*[**SetTargetResponse**](SetTargetResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetTargetGroupDescription**
> SetTargetGroupDescriptionResponse SetTargetGroupDescription(setTargetGroupDescriptionRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**setTargetGroupDescriptionRequest** | **[\*SetTargetGroupDescriptionRequest](SetTargetGroupDescriptionRequest.md)** | setTargetGroupDescriptionRequest | 

### Return type

*[**SetTargetGroupDescriptionResponse**](SetTargetGroupDescriptionResponse.md)

### Authorization

[x-ncp-iam](../README.md#x-ncp-iam)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

