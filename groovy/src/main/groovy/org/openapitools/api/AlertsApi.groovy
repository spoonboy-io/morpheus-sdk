package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.AddAlerts200Response
import org.openapitools.model.AddAlertsRequest
import org.openapitools.model.DefaultError
import org.openapitools.model.GetAlerts200Response
import org.openapitools.model.ListAlerts200Response
import org.openapitools.model.Model200Success
import org.openapitools.model.UpdateAlerts200Response
import org.openapitools.model.UpdateAlertsRequest

class AlertsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addAlerts ( AddAlertsRequest addAlertsRequest, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/alerts"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = addAlertsRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    AddAlerts200Response.class )

    }

    def deleteAlerts ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/alerts/${id}"

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

    def getAlerts ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/alerts/${id}"

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
                    GetAlerts200Response.class )

    }

    def listAlerts ( Long max, Long offset, Date lastUpdated, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/alerts"

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
        if (lastUpdated != null) {
            queryParams.put("lastUpdated", lastUpdated)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    ListAlerts200Response.class )

    }

    def updateAlerts ( Long id, UpdateAlertsRequest updateAlertsRequest, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/alerts/${id}"

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
        bodyParams = updateAlertsRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    UpdateAlerts200Response.class )

    }

}
