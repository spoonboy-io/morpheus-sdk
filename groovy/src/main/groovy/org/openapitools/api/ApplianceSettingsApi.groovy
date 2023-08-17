package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.ListApplianceSettings200Response
import org.openapitools.model.Model200Success
import org.openapitools.model.UpdateApplianceSettingsRequest

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
                    ListApplianceSettings200Response.class )

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

    def updateApplianceSettings ( UpdateApplianceSettingsRequest updateApplianceSettingsRequest, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/appliance-settings"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = updateApplianceSettingsRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

}
