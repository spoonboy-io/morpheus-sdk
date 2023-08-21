package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError

class SecurityScansApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def getSecurityScans ( Long id, Boolean results, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/security-scans/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (results != null) {
            queryParams.put("results", results)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listSecurityScans ( Long max, Long offset, String sort, String direction, String phrase, Long securityPackageId, Long serverId, Boolean results, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/security-scans"

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
        if (direction != null) {
            queryParams.put("direction", direction)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (securityPackageId != null) {
            queryParams.put("securityPackageId", securityPackageId)
        }
        if (serverId != null) {
            queryParams.put("serverId", serverId)
        }
        if (results != null) {
            queryParams.put("results", results)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

}
