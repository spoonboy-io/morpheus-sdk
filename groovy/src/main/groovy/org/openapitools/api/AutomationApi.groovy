package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.AddExecuteSchedules200Response
import org.openapitools.model.AddExecuteSchedulesRequest
import org.openapitools.model.DefaultError
import org.openapitools.model.ExecuteExecutionRequest200Response
import org.openapitools.model.ExecuteExecutionRequestRequest
import org.openapitools.model.GetExecuteSchedules200Response
import org.openapitools.model.GetExecutionRequest200Response
import org.openapitools.model.ListExecuteSchedules200Response
import org.openapitools.model.Model200Success
import org.openapitools.model.UpdateExecuteSchedulesRequest

class AutomationApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addExecuteSchedules ( AddExecuteSchedulesRequest addExecuteSchedulesRequest, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/execute-schedules"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = addExecuteSchedulesRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    AddExecuteSchedules200Response.class )

    }

    def executeExecutionRequest ( Long instanceId, Long containerId, Long serverId, ExecuteExecutionRequestRequest executeExecutionRequestRequest, Closure onSuccess, Closure onFailure)  {
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
        bodyParams = executeExecutionRequestRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    ExecuteExecutionRequest200Response.class )

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
                    GetExecuteSchedules200Response.class )

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
                    GetExecutionRequest200Response.class )

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
                    ListExecuteSchedules200Response.class )

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

    def updateExecuteSchedules ( Long id, UpdateExecuteSchedulesRequest updateExecuteSchedulesRequest, Closure onSuccess, Closure onFailure)  {
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
        bodyParams = updateExecuteSchedulesRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    AddExecuteSchedules200Response.class )

    }

}
