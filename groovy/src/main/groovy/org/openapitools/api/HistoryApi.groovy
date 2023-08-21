package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineResponse20048
import org.openapitools.model.InlineResponse20049

class HistoryApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def getHistory ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/processes/${id}"

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
                    InlineResponse20048.class )

    }

    def getHistoryEvent ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/processes/events/${id}"

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
                    InlineResponse20049.class )

    }

    def listHistory ( Long instanceId, Long containerId, Long serverId, Long zoneId, Long appId, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/processes"

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
        if (zoneId != null) {
            queryParams.put("zoneId", zoneId)
        }
        if (appId != null) {
            queryParams.put("appId", appId)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

}
