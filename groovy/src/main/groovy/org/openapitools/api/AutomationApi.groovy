package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject12
import org.openapitools.model.InlineObject13
import org.openapitools.model.InlineObject14
import org.openapitools.model.InlineObject190
import org.openapitools.model.InlineObject191
import org.openapitools.model.InlineObject192
import org.openapitools.model.InlineObject193
import org.openapitools.model.InlineObject194
import org.openapitools.model.InlineObject195
import org.openapitools.model.InlineObject211
import org.openapitools.model.InlineObject212
import org.openapitools.model.InlineObject246
import org.openapitools.model.InlineObject247
import org.openapitools.model.InlineObject248
import org.openapitools.model.InlineObject249
import org.openapitools.model.InlineObject250
import org.openapitools.model.InlineObject251
import org.openapitools.model.InlineResponse200155
import org.openapitools.model.InlineResponse200156
import org.openapitools.model.InlineResponse2008
import org.openapitools.model.Model200Success
import org.openapitools.model.Model200SuccessExpanded

class AutomationApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addExecuteSchedules ( InlineObject12 inlineObject12, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/execute-schedules"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject12


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addPowerScheduleInstances ( Long id, InlineObject192 inlineObject192, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/power-schedules/${id}/add-instances"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject192


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200SuccessExpanded.class )

    }

    def addPowerScheduleServers ( Long id, InlineObject193 inlineObject193, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/power-schedules/${id}/add-servers"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject193


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200SuccessExpanded.class )

    }

    def addPowerSchedules ( InlineObject190 inlineObject190, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/power-schedules"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject190


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addScaleThresholds ( InlineObject211 inlineObject211, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/scale-thresholds"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject211


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addTasks ( InlineObject246 inlineObject246, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/tasks"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject246


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addWorkflows ( InlineObject249 inlineObject249, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/task-sets"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject249


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def deletePowerScheduleInstances ( Long id, InlineObject194 inlineObject194, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/power-schedules/${id}/remove-instances"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject194


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200SuccessExpanded.class )

    }

    def deletePowerScheduleServers ( Long id, InlineObject195 inlineObject195, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/power-schedules/${id}/remove-servers"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject195


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200SuccessExpanded.class )

    }

    def executeExecutionRequest ( Long instanceId, Long containerId, Long serverId, InlineObject14 inlineObject14, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/execution-request/execute"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (instanceId != null) {
            queryParams.put("instanceId", instanceId)
        }
        if (containerId != null) {
            queryParams.put("containerId", containerId)
        }
        if (serverId != null) {
            queryParams.put("serverId", serverId)
        }


        contentType = 'application/json';
        bodyParams = inlineObject14


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse2008.class )

    }

    def executeTasks ( Long id, InlineObject248 inlineObject248, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/tasks/${id}/execute"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject248


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def executeWorkflows ( Long id, InlineObject251 inlineObject251, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/task-sets/${id}/execute"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject251


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def getExecuteSchedules ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/execute-schedules/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getExecutionRequest ( String uniqueId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/execution-request/${uniqueId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (uniqueId == null) {
            throw new RuntimeException("missing required params uniqueId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getPowerSchedules ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/power-schedules/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getScaleThresholds ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/scale-thresholds/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getTaskTypes ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/task-types/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200156.class )

    }

    def getTasks ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/tasks/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getWorkflows ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/task-sets/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listExecuteSchedules ( Long max, Long offset, String sort, String direction, String phrase, String name, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/execute-schedules"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (max != null) {
            queryParams.put("max", max)
        }
        if (offset != null) {
            queryParams.put("offset", offset)
        }
        if (sort != null) {
            queryParams.put("sort", sort)
        }
        if (direction != null) {
            queryParams.put("direction", direction)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (name != null) {
            queryParams.put("name", name)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listPowerSchedules ( Long max, Long offset, String sort, String direction, String phrase, String name, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/power-schedules"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (max != null) {
            queryParams.put("max", max)
        }
        if (offset != null) {
            queryParams.put("offset", offset)
        }
        if (sort != null) {
            queryParams.put("sort", sort)
        }
        if (direction != null) {
            queryParams.put("direction", direction)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (name != null) {
            queryParams.put("name", name)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listScaleThresholds ( Long max, Long offset, String sort, String direction, String phrase, String name, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/scale-thresholds"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (max != null) {
            queryParams.put("max", max)
        }
        if (offset != null) {
            queryParams.put("offset", offset)
        }
        if (sort != null) {
            queryParams.put("sort", sort)
        }
        if (direction != null) {
            queryParams.put("direction", direction)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (name != null) {
            queryParams.put("name", name)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listTaskTypes ( String name, String code, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/task-types"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (name != null) {
            queryParams.put("name", name)
        }
        if (code != null) {
            queryParams.put("code", code)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200155.class )

    }

    def listTasks ( String type, Long max, Long offset, String sort, String direction, String phrase, String name, String labels, String allLabels, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/tasks"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (type != null) {
            queryParams.put("type", type)
        }
        if (max != null) {
            queryParams.put("max", max)
        }
        if (offset != null) {
            queryParams.put("offset", offset)
        }
        if (sort != null) {
            queryParams.put("sort", sort)
        }
        if (direction != null) {
            queryParams.put("direction", direction)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (name != null) {
            queryParams.put("name", name)
        }
        if (labels != null) {
            queryParams.put("labels", labels)
        }
        if (allLabels != null) {
            queryParams.put("allLabels", allLabels)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listWorkflows ( Long max, Long offset, String sort, String direction, String phrase, String name, String labels, String allLabels, String type, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/task-sets"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (max != null) {
            queryParams.put("max", max)
        }
        if (offset != null) {
            queryParams.put("offset", offset)
        }
        if (sort != null) {
            queryParams.put("sort", sort)
        }
        if (direction != null) {
            queryParams.put("direction", direction)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (name != null) {
            queryParams.put("name", name)
        }
        if (labels != null) {
            queryParams.put("labels", labels)
        }
        if (allLabels != null) {
            queryParams.put("allLabels", allLabels)
        }
        if (type != null) {
            queryParams.put("type", type)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def removeExecuteSchedules ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/execute-schedules/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def removePowerSchedules ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/power-schedules/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def removeScaleThresholds ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/scale-thresholds/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def removeTasks ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/tasks/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def removeWorkflows ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/task-sets/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def updateExecuteSchedules ( Long id, InlineObject13 inlineObject13, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/execute-schedules/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject13


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updatePowerSchedules ( Long id, InlineObject191 inlineObject191, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/power-schedules/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject191


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateScaleThresholds ( Long id, InlineObject212 inlineObject212, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/scale-thresholds/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject212


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateTasks ( Long id, InlineObject247 inlineObject247, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/tasks/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject247


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateWorkflows ( Long id, InlineObject250 inlineObject250, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/task-sets/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject250


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

}
