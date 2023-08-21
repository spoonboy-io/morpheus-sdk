

# IntegrationConfigIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Name, a unique identifier for the integration | 
**type** | [**TypeEnum**](#TypeEnum) | Integration Type Code | 
**enabled** | **Boolean** | Set &#x60;true&#x60; to enable integration |  [optional]
**refresh** | **Boolean** | Pass &#x60;false&#x60; to skip refresh.  By default, refresh is done on update, when it is supported by the integration type.  |  [optional]
**credential** | [**OneOfobjectobject**](OneOfobjectobject.md) | Map containing Credential ID or the default {\&quot;type\&quot;: \&quot;local\&quot;}  which means use the values set in the local task options username and password instead of associating a credential.  |  [optional]



## Enum: TypeEnum

Name | Value
---- | -----
ANSIBLE | &quot;ansible&quot;
ANSIBLETOWER | &quot;ansibleTower&quot;
BINDDNS | &quot;bindDns&quot;
CHEF | &quot;chef&quot;
CHERWELL | &quot;cherwell&quot;
CYPHER | &quot;cypher&quot;
DOCKER_REGISTRY | &quot;docker.registry&quot;
GIT | &quot;git&quot;
GITHUB | &quot;github&quot;
MICROSOFT_DNS | &quot;microsoft.dns&quot;
POWERDNS | &quot;powerDns&quot;
PUPPET | &quot;puppet&quot;
REMEDY | &quot;remedy&quot;
AMAZONDNS | &quot;amazonDns&quot;
SALTMASTER | &quot;saltMaster&quot;
SERVICENOW | &quot;serviceNow&quot;
VRO | &quot;vro&quot;



