# MorpheusApi.IntegrationConfigIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Name, a unique identifier for the integration | 
**type** | **String** | Integration Type Code | 
**enabled** | **Boolean** | Set &#x60;true&#x60; to enable integration | [optional] 
**refresh** | **Boolean** | Pass &#x60;false&#x60; to skip refresh.  By default, refresh is done on update, when it is supported by the integration type.  | [optional] [default to true]
**credential** | [**OneOfobjectobject**](OneOfobjectobject.md) | Map containing Credential ID or the default {\&quot;type\&quot;: \&quot;local\&quot;}  which means use the values set in the local task options username and password instead of associating a credential.  | [optional] 



## Enum: TypeEnum


* `ansible` (value: `"ansible"`)

* `ansibleTower` (value: `"ansibleTower"`)

* `bindDns` (value: `"bindDns"`)

* `chef` (value: `"chef"`)

* `cherwell` (value: `"cherwell"`)

* `cypher` (value: `"cypher"`)

* `docker.registry` (value: `"docker.registry"`)

* `git` (value: `"git"`)

* `github` (value: `"github"`)

* `microsoft.dns` (value: `"microsoft.dns"`)

* `powerDns` (value: `"powerDns"`)

* `puppet` (value: `"puppet"`)

* `remedy` (value: `"remedy"`)

* `amazonDns` (value: `"amazonDns"`)

* `saltMaster` (value: `"saltMaster"`)

* `serviceNow` (value: `"serviceNow"`)

* `vro` (value: `"vro"`)




