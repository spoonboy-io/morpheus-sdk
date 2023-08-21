package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject87
import org.openapitools.model.InlineObject88
import org.openapitools.model.InlineObject89
import org.openapitools.model.InlineObject90
import org.openapitools.model.InlineResponse20055
import org.openapitools.model.Model200Success
import org.openapitools.model.SuccessMessage

class IncidentsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addIncident ( InlineObject87 inlineObject87, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/incidents"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject87


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def deleteIncidents ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/incidents/${id}"

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

    def getIncidents ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/incidents/${id}"

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
                    InlineResponse20055.class )

    }

    def listIncidents ( Long max, Long offset, String status, String severity, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/incidents"

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
        if (status != null) {
            queryParams.put("status", status)
        }
        if (severity != null) {
            queryParams.put("severity", severity)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def updateIncidents ( Long id, InlineObject88 inlineObject88, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/incidents/${id}"

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
        bodyParams = inlineObject88


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateIncidentsReopen ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/incidents/${id}/reopen"

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
                    SuccessMessage.class )

    }

    def updateMuteAllIncidents ( InlineObject90 inlineObject90, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/incidents/mute-all"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject90


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateMuteIncidents ( Long id, InlineObject89 inlineObject89, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/incidents/${id}/mute"

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
        bodyParams = inlineObject89


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

}
