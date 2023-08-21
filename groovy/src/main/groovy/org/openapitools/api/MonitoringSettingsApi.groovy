package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject141
import org.openapitools.model.InlineResponse20085
import org.openapitools.model.Model200Success

class MonitoringSettingsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def getMonitoringSettings ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring-settings"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20085.class )

    }

    def updateMonitoringSettings ( InlineObject141 inlineObject141, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring-settings"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject141


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

}
