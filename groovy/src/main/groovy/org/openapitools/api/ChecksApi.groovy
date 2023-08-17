package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.AddCheckApps200Response
import org.openapitools.model.AddCheckAppsRequest
import org.openapitools.model.DefaultError
import org.openapitools.model.ListCheckApps200Response

class ChecksApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addCheckApps ( AddCheckAppsRequest addCheckAppsRequest, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/apps"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = addCheckAppsRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    AddCheckApps200Response.class )

    }

    def listCheckApps ( Long max, Long offset, String sort, String name, String phrase, String status, Date lastUpdated, Boolean deleted, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/monitoring/apps"

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
        if (name != null) {
            queryParams.put("name", name)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (status != null) {
            queryParams.put("status", status)
        }
        if (lastUpdated != null) {
            queryParams.put("lastUpdated", lastUpdated)
        }
        if (deleted != null) {
            queryParams.put("deleted", deleted)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    ListCheckApps200Response.class )

    }

}
