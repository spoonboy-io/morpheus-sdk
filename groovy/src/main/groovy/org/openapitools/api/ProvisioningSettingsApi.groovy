package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject204
import org.openapitools.model.InlineResponse200128
import org.openapitools.model.Model200Success

class ProvisioningSettingsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def listProvisioningSettings ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/provisioning-settings"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200128.class )

    }

    def updateProvisioningSettings ( InlineObject204 inlineObject204, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/provisioning-settings"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject204


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

}
