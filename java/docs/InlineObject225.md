

# InlineObject225

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**taskSet** | [**ApiServersIdWorkflowTaskSet**](ApiServersIdWorkflowTaskSet.md) |  |  [optional]
**taskPhase** | [**TaskPhaseEnum**](#TaskPhaseEnum) | Task Phase to run for Provisioning workflows. The default is &#x60;provision&#x60;. |  [optional]



## Enum: TaskPhaseEnum

Name | Value
---- | -----
START | &quot;start&quot;
STOP | &quot;stop&quot;
PREPROVISION | &quot;preProvision&quot;
PROVISION | &quot;provision&quot;
POSTPROVISION | &quot;postProvision&quot;
PREDEPLOY | &quot;preDeploy&quot;
DEPLOY | &quot;deploy&quot;
RECONFIGURE | &quot;reconfigure&quot;
TEARDOWN | &quot;teardown&quot;
STARTUP | &quot;startup&quot;
SHUTDOWN | &quot;shutdown&quot;



