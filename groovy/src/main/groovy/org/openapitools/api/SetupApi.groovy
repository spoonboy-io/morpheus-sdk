package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.AnyOfobjectobject
import org.openapitools.model.DefaultError
import org.openapitools.model.Setup
import org.openapitools.model.UNKNOWN_BASE_TYPE

class SetupApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def setup ( UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/setup"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = UNKNOWN_BASE_TYPE


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Setup.class )

    }

}
