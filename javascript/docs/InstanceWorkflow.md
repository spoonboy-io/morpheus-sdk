# MorpheusApi.InstanceWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**taskSet** | [**InstanceWorkflowTaskSet**](InstanceWorkflowTaskSet.md) |  | [optional] 
**taskPhase** | **String** | Task Phase to run for Provisioning workflows. The default is &#x60;provision&#x60;. | [optional] [default to &#39;provision&#39;]



## Enum: TaskPhaseEnum


* `start` (value: `"start"`)

* `stop` (value: `"stop"`)

* `preProvision` (value: `"preProvision"`)

* `provision` (value: `"provision"`)

* `postProvision` (value: `"postProvision"`)

* `preDeploy` (value: `"preDeploy"`)

* `deploy` (value: `"deploy"`)

* `reconfigure` (value: `"reconfigure"`)

* `teardown` (value: `"teardown"`)

* `startup` (value: `"startup"`)

* `shutdown` (value: `"shutdown"`)




