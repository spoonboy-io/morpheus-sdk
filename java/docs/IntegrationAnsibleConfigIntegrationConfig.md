

# IntegrationAnsibleConfigIntegrationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**defaultBranch** | [**DefaultBranchEnum**](#DefaultBranchEnum) | default branch |  [optional]
**ansiblePlaybooks** | **String** | Playbooks path |  [optional]
**ansibleRoles** | **String** | Roles path |  [optional]
**ansibleGroupVars** | **String** | Group variables path |  [optional]
**ansibleHostVars** | **String** | Host variables path |  [optional]
**ansibleGalaxyEnabled** | **Boolean** | Use Ansible Galaxy |  [optional]
**ansibleVerbose** | **Boolean** | Use verbose logging |  [optional]
**ansibleCommandBus** | **Boolean** | Use Morpheus Agent Command Bus |  [optional]
**cacheEnabled** | **Boolean** | Enable Git repository caching |  [optional]



## Enum: DefaultBranchEnum

Name | Value
---- | -----
MAIN | &quot;main&quot;
MASTER | &quot;master&quot;



