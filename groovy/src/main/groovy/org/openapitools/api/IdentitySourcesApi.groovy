package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject82
import org.openapitools.model.InlineResponse20051
import org.openapitools.model.Model200Success
import org.openapitools.model.UserSourceCreate

class IdentitySourcesApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addIdentitySources ( Long accountId, UserSourceCreate userSourceCreate, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/user-sources"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (accountId != null) {
            queryParams.put("accountId", accountId)
        }


        contentType = 'application/json';
        bodyParams = userSourceCreate


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def getIdentitySources ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/user-sources/${id}"

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
                    InlineResponse20051.class )

    }

    def listIdentitySources ( String type, Long max, Long offset, String sort, String direction, String phrase, String name, Long accountId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/user-sources"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (type != null) {
            queryParams.put("type", type)
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
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (name != null) {
            queryParams.put("name", name)
        }
        if (accountId != null) {
            queryParams.put("accountId", accountId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def removeIdentitySources ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/user-sources/${id}"

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

    def updateIdentitySourceSubdomains ( Long id, InlineObject82 inlineObject82, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/user-sources/${id}/subdomain"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject82


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateIdentitySources ( Long id, UserSourceCreate userSourceCreate, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/user-sources/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = userSourceCreate


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

}
