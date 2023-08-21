package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.Dashboard
import org.openapitools.model.DefaultError

class DashboardApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def dashboard ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/dashboard"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Dashboard.class )

    }

}
