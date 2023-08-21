package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineResponse20044
import org.openapitools.model.InlineResponse20045
import org.openapitools.model.InlineResponse20046
import org.openapitools.model.Model200Success

class GuidanceApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def executeGuidances ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/guidance/${id}/execute"

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
                    "PUT", "",
                    Model200Success.class )

    }

    def getGuidanceStats ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/guidance/stats"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20045.class )

    }

    def getGuidanceTypes ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/guidance/types"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20046.class )

    }

    def getGuidances ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/guidance/${id}"

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
                    InlineResponse20044.class )

    }

    def ignoreGuidances ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/guidance/${id}/ignore"

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
                    "PUT", "",
                    Model200Success.class )

    }

    def listGuidances ( Long max, Long offset, String sort, String direction, String phrase, String severity, String type, String state, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/guidance"

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
        if (severity != null) {
            queryParams.put("severity", severity)
        }
        if (type != null) {
            queryParams.put("type", type)
        }
        if (state != null) {
            queryParams.put("state", state)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

}
