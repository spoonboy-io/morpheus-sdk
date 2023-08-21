package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject66
import org.openapitools.model.Model200Success
import org.openapitools.model.OneOfstringlong

class CypherApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addCypherKey ( String cypherPath, OneOfstringlong ttl, String value, String type, InlineObject66 inlineObject66, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/cypher/${cypherPath}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (cypherPath == null) {
            throw new RuntimeException("missing required params cypherPath")
        }

        if (ttl != null) {
            queryParams.put("ttl", ttl)
        }
        if (value != null) {
            queryParams.put("value", value)
        }
        if (type != null) {
            queryParams.put("type", type)
        }


        contentType = 'application/json';
        bodyParams = inlineObject66


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def getCypherKey ( String cypherPath, String leaseToken, String sort, String direction, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/cypher/${cypherPath}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (cypherPath == null) {
            throw new RuntimeException("missing required params cypherPath")
        }

        if (leaseToken != null) {
            queryParams.put("leaseToken", leaseToken)
        }
        if (sort != null) {
            queryParams.put("sort", sort)
        }
        if (direction != null) {
            queryParams.put("direction", direction)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Model200Success.class )

    }

    def listCypherKeys ( String leaseToken, Boolean list, String phrase, Long max, Long offset, String sort, String direction, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/cypher"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (leaseToken != null) {
            queryParams.put("leaseToken", leaseToken)
        }
        if (list != null) {
            queryParams.put("list", list)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
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




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def removeCypher ( String cypherPath, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/cypher/${cypherPath}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (cypherPath == null) {
            throw new RuntimeException("missing required params cypherPath")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

}
