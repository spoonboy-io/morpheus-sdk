

# ApiTaskSetsTaskSetTasks

List of task objects in order
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**taskId** | **Long** | Task ID |  [optional]
**taskPhase** | [**TaskPhaseEnum**](#TaskPhaseEnum) | Task Phase. Pass operation for &#x60;operational&#x60; workflows. |  [optional]



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
OPERATION | &quot;operation&quot;



