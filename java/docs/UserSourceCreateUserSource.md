

# UserSourceCreateUserSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A name for the Identity Source | 
**type** | **String** | IDM type code | 
**description** | **String** | description |  [optional]
**defaultAccountRole** | [**UserSourceCreateUserSourceDefaultAccountRole**](UserSourceCreateUserSourceDefaultAccountRole.md) |  | 
**roleMappings** | [**OneOfarraymap**](OneOfarraymap.md) |  |  [optional]
**roleMappingNames** | **Map&lt;String, String&gt;** | Map of Morpheus &#39;&#x60;Role ID&#x60;&#39;:&#39;&#x60;Role Name&#x60;&#39;.  |  [optional]
**allowCustomMappings** | **Boolean** | Enable Role Mapping Permission |  [optional]
**manualRoleAssignment** | **Boolean** | Manual Role Assignment |  [optional]
**config** | [**OneOfuserSourceCreateLDAPuserSourceCreateJumpClouduserSourceCreateActiveDirectoryuserSourceCreateOktauserSourceCreateOneLoginuserSourceCreateSamluserSourceCreateCustomExternaluserSourceCreateCustomApi**](OneOfuserSourceCreateLDAPuserSourceCreateJumpClouduserSourceCreateActiveDirectoryuserSourceCreateOktauserSourceCreateOneLoginuserSourceCreateSamluserSourceCreateCustomExternaluserSourceCreateCustomApi.md) |  |  [optional]



