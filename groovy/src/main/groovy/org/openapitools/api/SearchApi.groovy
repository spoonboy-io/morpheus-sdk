package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.Search

class SearchApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def search ( Long max, Long offset, String phrase, String query, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/search"

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
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (query != null) {
            queryParams.put("query", query)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Search.class )

    }

}
