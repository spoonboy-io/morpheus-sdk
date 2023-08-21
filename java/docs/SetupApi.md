# SetupApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**setup**](SetupApi.md#setup) | **POST** /api/setup | Setup appliance


<a name="setup"></a>
# **setup**
> Setup setup(UNKNOWN_BASE_TYPE)

Setup appliance

Initialize a freshly installed appliance to create the master tenant and System Admin user.  Authorization is not required.  This operation can only be executed successfully once. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.SetupApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");

    SetupApi apiInstance = new SetupApi(defaultClient);
    UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE = {"applianceName":"The Matrix","accountName":"Meta Cortex Corporation","firstName":"Thomas","lastName":"Anderson","email":"tanderson@mccorp.com","username":"tanderson","password":"QnW}cg}8}<~:P9YU"}; // UNKNOWN_BASE_TYPE | 
    try {
      Setup result = apiInstance.setup(UNKNOWN_BASE_TYPE);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling SetupApi#setup");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional]

### Return type

[**Setup**](Setup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Appliance setup complete. The master tenant and system admin user were created successfully. |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

