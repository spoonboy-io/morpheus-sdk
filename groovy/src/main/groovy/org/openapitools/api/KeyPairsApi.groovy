package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject105
import org.openapitools.model.InlineObject106
import org.openapitools.model.InlineResponse20066
import org.openapitools.model.InlineResponse20067
import org.openapitools.model.Model200Success

class KeyPairsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addKeyPairs ( InlineObject105 inlineObject105, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/key-pairs"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject105


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def generateKeyPairs ( InlineObject106 inlineObject106, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/key-pairs/generate"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject106


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse20066.class )

    }

    def getKeyPairs ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/key-pairs/${id}"

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
                    InlineResponse20067.class )

    }

    def removeKeyPairs ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/key-pairs/${id}"

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
                    "DELETE", "",
                    Model200Success.class )

    }

}
