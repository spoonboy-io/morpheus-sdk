package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject2
import org.openapitools.model.InlineResponse200
import org.openapitools.model.Model200Success

class ApplianceSettingsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def listApplianceSettings ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/appliance-settings"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200.class )

    }

    def reindex ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/appliance-settings/reindex"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def setApplianceSettingsMaintenanceMode ( Boolean enabled, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/appliance-settings/maintenance"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (enabled != null) {
            queryParams.put("enabled", enabled)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def updateApplianceSettings ( InlineObject2 inlineObject2, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/appliance-settings"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject2


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

}
