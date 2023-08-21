# BlueprintsApi

All URIs are relative to *https://CHANGEME*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addBlueprint**](BlueprintsApi.md#addBlueprint) | **POST** /api/blueprints | Create a Blueprint
[**deleteBlueprint**](BlueprintsApi.md#deleteBlueprint) | **DELETE** /api/blueprints/{id} | Delete a Blueprint
[**getBlueprint**](BlueprintsApi.md#getBlueprint) | **GET** /api/blueprints/{id} | Get a Specific Blueprint
[**listBlueprints**](BlueprintsApi.md#listBlueprints) | **GET** /api/blueprints | Get All Blueprints
[**updateBlueprint**](BlueprintsApi.md#updateBlueprint) | **PUT** /api/blueprints/{id} | Updating a Blueprint
[**updateBlueprintImage**](BlueprintsApi.md#updateBlueprintImage) | **POST** /api/blueprints/{id}/image | Update Blueprint Image
[**updateBlueprintPermissions**](BlueprintsApi.md#updateBlueprintPermissions) | **PUT** /api/blueprints/{id}/update-permissions | Update Blueprint Permissions


<a name="addBlueprint"></a>
# **addBlueprint**
> Object addBlueprint(UNKNOWN_BASE_TYPE)

Create a Blueprint

Create a Blueprint

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BlueprintsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BlueprintsApi apiInstance = new BlueprintsApi(defaultClient);
    UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE = {"image":"/assets/apps/template.png","name":"ARM: Ubuntu 16Test","type":"arm","arm":{"json":"{\n    \"$schema\": \"http://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#\",\n    \"contentVersion\": \"1.0.0.0\",\n    \"parameters\": {\n        \"virtualMachineName\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"DMUBARM<%=sequence%>\",\n            \"metadata\": {\n                \"description\": \"Name of Virtual Machine\"\n            }\n        },\n        \"networkInterfaceName\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"demo-arm-ubuntu-nic\",\n            \"metadata\": {\n                \"description\": \"Name of Network Interface\"\n            }\n        },\n        \"publicIpAddressName\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"demo-arm-ubuntu-ip\",\n            \"metadata\": {\n                \"description\": \"Name of IP Address\"\n            }\n        },\n        \"adminUsername\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"morpheus\",\n            \"metadata\": {\n                \"description\": \"OS Username\"\n            }\n        },\n        \"adminPassword\": {\n            \"type\": \"securestring\",\n            \"metadata\": {\n                \"description\": \"OS Password\"\n            }\n        },\n        \"location\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"westus\"\n        },\n        \"virtualMachineSize\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"Standard_A1\"\n        },\n        \"virtualNetworkName\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"morphdemoVnet\"\n        },    \n        \"networkSecurityGroupName\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"morph-demo-sg\"\n        },\n        \"diagnosticsStorageAccountName\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"morphdemosa\"\n        },\n        \"subnetName\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"morphDemoSubnet\"\n        },\n        \"publicIpAddressType\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"Dynamic\"\n        },\n        \"publicIpAddressSku\": {\n            \"type\": \"string\",\n            \"defaultValue\": \"Basic\"\n        }\n    },\n    \"variables\": {\n        \"vnetId\": \"[resourceId('morph-demo','Microsoft.Network/virtualNetworks', parameters('virtualNetworkName'))]\",\n        \"subnetRef\": \"[concat(variables('vnetId'), '/subnets/', parameters('subnetName'))]\"\n    },\n    \"resources\": [\n        {\n            \"name\": \"[parameters('virtualMachineName')]\",\n            \"type\": \"Microsoft.Compute/virtualMachines\",\n            \"apiVersion\": \"2016-04-30-preview\",\n            \"location\": \"[parameters('location')]\",\n            \"dependsOn\": [\n                \"[concat('Microsoft.Network/networkInterfaces/', parameters('networkInterfaceName'))]\"\n            ],\n            \"properties\": {\n                \"osProfile\": {\n                    \"computerName\": \"[parameters('virtualMachineName')]\",\n                    \"adminUsername\": \"[parameters('adminUsername')]\",\n                    \"adminPassword\": \"[parameters('adminPassword')]\"\n                },\n                \"hardwareProfile\": {\n                    \"vmSize\": \"[parameters('virtualMachineSize')]\"\n                },\n                \"storageProfile\": {\n                    \"imageReference\": {\n                        \"publisher\": \"Canonical\",\n                        \"offer\": \"UbuntuServer\",\n                        \"sku\": \"16.04-LTS\",\n                        \"version\": \"latest\"\n                    },\n                    \"osDisk\": {\n                        \"createOption\": \"fromImage\",\n                        \"managedDisk\": {\n                            \"storageAccountType\": \"Standard_LRS\"\n                        }\n                    },\n                    \"dataDisks\": []\n                },\n                \"networkProfile\": {\n                    \"networkInterfaces\": [\n                        {\n                            \"id\": \"[resourceId('Microsoft.Network/networkInterfaces', parameters('networkInterfaceName'))]\"\n                        }\n                    ]\n                },\n                \"diagnosticsProfile\": {\n                    \"bootDiagnostics\": {\n                        \"enabled\": true,\n                        \"storageUri\": \"[reference(resourceId('morph-demo', 'Microsoft.Storage/storageAccounts', parameters('diagnosticsStorageAccountName')), '2015-06-15').primaryEndpoints['blob']]\"\n                    }\n                }\n            }\n        },\n        {\n            \"name\": \"[parameters('networkInterfaceName')]\",\n            \"type\": \"Microsoft.Network/networkInterfaces\",\n            \"apiVersion\": \"2016-09-01\",\n            \"location\": \"[parameters('location')]\",\n            \"dependsOn\": [\n                \"[concat('Microsoft.Network/publicIpAddresses/', parameters('publicIpAddressName'))]\"\n            ],\n            \"properties\": {\n                \"ipConfigurations\": [\n                    {\n                        \"name\": \"ipconfig1\",\n                        \"properties\": {\n                            \"subnet\": {\n                                \"id\": \"[variables('subnetRef')]\"\n                            },\n                            \"privateIPAllocationMethod\": \"Dynamic\",\n                            \"publicIpAddress\": {\n                                \"id\": \"[resourceId('morph-demo','Microsoft.Network/publicIpAddresses', parameters('publicIpAddressName'))]\"\n                            }\n                        }\n                    }\n                ],\n                \"networkSecurityGroup\": {\n                    \"id\": \"[resourceId('morph-demo', 'Microsoft.Network/networkSecurityGroups', parameters('networkSecurityGroupName'))]\"\n                }\n            }\n        },\n        {\n            \"name\": \"[parameters('publicIpAddressName')]\",\n            \"type\": \"Microsoft.Network/publicIpAddresses\",\n            \"apiVersion\": \"2017-08-01\",\n            \"location\": \"[parameters('location')]\",\n            \"properties\": {\n                \"publicIpAllocationMethod\": \"[parameters('publicIpAddressType')]\"\n            },\n            \"sku\": {\n                \"name\": \"[parameters('publicIpAddressSku')]\"\n            }\n        }\n    ],\n    \"outputs\": {\n        \"adminUsername\": {\n            \"type\": \"string\",\n            \"value\": \"[parameters('adminUsername')]\"\n        }\n    }\n}","configType":"json","installAgent":true}}; // UNKNOWN_BASE_TYPE | 
    try {
      Object result = apiInstance.addBlueprint(UNKNOWN_BASE_TYPE);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BlueprintsApi#addBlueprint");
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

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="deleteBlueprint"></a>
# **deleteBlueprint**
> Model200Success deleteBlueprint(id)

Delete a Blueprint

Delete a Blueprint

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BlueprintsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BlueprintsApi apiInstance = new BlueprintsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      Model200Success result = apiInstance.deleteBlueprint(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BlueprintsApi#deleteBlueprint");
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
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**Model200Success**](Model200Success.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="getBlueprint"></a>
# **getBlueprint**
> InlineResponse20014 getBlueprint(id)

Get a Specific Blueprint

This endpoint retrieves a specific blueprint.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BlueprintsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BlueprintsApi apiInstance = new BlueprintsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    try {
      InlineResponse20014 result = apiInstance.getBlueprint(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BlueprintsApi#getBlueprint");
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
 **id** | **Long**| Morpheus ID of the Object being referenced |

### Return type

[**InlineResponse20014**](InlineResponse20014.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="listBlueprints"></a>
# **listBlueprints**
> Object listBlueprints(max, offset, name, phrase, labels, allLabels)

Get All Blueprints

This endpoint retrieves all blueprints.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BlueprintsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BlueprintsApi apiInstance = new BlueprintsApi(defaultClient);
    Long max = 25l; // Long | Maximum number of records to return, -1 can be used to fetch all records
    Long offset = 0l; // Long | Offset records, the number of records to skip, for paginating requests
    String name = "example-%"; // String | Filter by name, wildcard may be specified as %, eg. example-%
    String phrase = "phrase_example"; // String | Search phrase for partial matches on name or description
    String labels = "labels_example"; // String | Filter by label(s), matches records that contain any of the specified labels
    String allLabels = "allLabels_example"; // String | Filter by label(s), matches records that contain all of the specified labels
    try {
      Object result = apiInstance.listBlueprints(max, offset, name, phrase, labels, allLabels);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BlueprintsApi#listBlueprints");
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
 **max** | **Long**| Maximum number of records to return, -1 can be used to fetch all records | [optional] [default to 25l]
 **offset** | **Long**| Offset records, the number of records to skip, for paginating requests | [optional] [default to 0l]
 **name** | **String**| Filter by name, wildcard may be specified as %, eg. example-% | [optional]
 **phrase** | **String**| Search phrase for partial matches on name or description | [optional]
 **labels** | **String**| Filter by label(s), matches records that contain any of the specified labels | [optional]
 **allLabels** | **String**| Filter by label(s), matches records that contain all of the specified labels | [optional]

### Return type

**Object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Response |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateBlueprint"></a>
# **updateBlueprint**
> InlineResponse20014 updateBlueprint(id, UNKNOWN_BASE_TYPE)

Updating a Blueprint

Update a Blueprint. This overwrites the entire config, so the entire blueprint config should be passed.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BlueprintsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BlueprintsApi apiInstance = new BlueprintsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE = {"name":"sample","description":"A sample nginx blueprint","type":"morpheus","tiers":{"Web":{"linkedTiers":[],"tierIndex":1,"instances":[{"instance":{"type":"nginx"},"groups":{"My Group":{"clouds":{"My Cloud":{"instance":{"layout":{"code":"nginx-vmware-1.9-single","id":179},"name":"test-nginx-${sequence}","allowExisting":false,"createUser":"on","type":"nginx","userGroup":{"id":""}},"networkInterfaces":[{"ipMode":"","primaryInterface":true,"network":{"id":"","hasPool":false},"networkInterfaceTypeId":4,"networkInterfaceTypeIdName":"VMXNET 3"}],"volumes":[{"vId":255,"controllerMountPoint":"46:0:4:0","size":10,"maxIOPS":null,"name":"root","rootVolume":true,"storageType":1,"datastoreId":"autoCluster","maxStorage":0}],"config":{"resourcePoolId":"resgroup-123","createUser":true},"plan":{"code":"vm-1024","id":76}}}}}}]}}}; // UNKNOWN_BASE_TYPE | 
    try {
      InlineResponse20014 result = apiInstance.updateBlueprint(id, UNKNOWN_BASE_TYPE);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BlueprintsApi#updateBlueprint");
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
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional]

### Return type

[**InlineResponse20014**](InlineResponse20014.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateBlueprintImage"></a>
# **updateBlueprintImage**
> InlineResponse20014 updateBlueprintImage(id, templateImage)

Update Blueprint Image

Update Blueprint Image

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BlueprintsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BlueprintsApi apiInstance = new BlueprintsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    File templateImage = new File("/path/to/file"); // File | 
    try {
      InlineResponse20014 result = apiInstance.updateBlueprintImage(id, templateImage);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BlueprintsApi#updateBlueprintImage");
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
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **templateImage** | **File**|  | [optional]

### Return type

[**InlineResponse20014**](InlineResponse20014.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

<a name="updateBlueprintPermissions"></a>
# **updateBlueprintPermissions**
> InlineResponse20014 updateBlueprintPermissions(id, inlineObject21)

Update Blueprint Permissions

Update Blueprint Permissions

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.BlueprintsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://CHANGEME");
    
    // Configure HTTP bearer authorization: bearerAuth
    HttpBearerAuth bearerAuth = (HttpBearerAuth) defaultClient.getAuthentication("bearerAuth");
    bearerAuth.setBearerToken("BEARER TOKEN");

    BlueprintsApi apiInstance = new BlueprintsApi(defaultClient);
    Long id = 1L; // Long | Morpheus ID of the Object being referenced
    InlineObject21 inlineObject21 = new InlineObject21(); // InlineObject21 | 
    try {
      InlineResponse20014 result = apiInstance.updateBlueprintPermissions(id, inlineObject21);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling BlueprintsApi#updateBlueprintPermissions");
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
 **id** | **Long**| Morpheus ID of the Object being referenced |
 **inlineObject21** | [**InlineObject21**](InlineObject21.md)|  | [optional]

### Return type

[**InlineResponse20014**](InlineResponse20014.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful Request |  -  |
**4XX** | Error Codes |  -  |
**5XX** | Error Codes |  -  |

