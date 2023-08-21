package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject206
import org.openapitools.model.InlineObject207
import org.openapitools.model.InlineResponse200131
import org.openapitools.model.InlineResponse200132
import org.openapitools.model.Model200Success

class ResourcePoolsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def createResourcePoolGroup ( InlineObject206 inlineObject206, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/resource-pools/groups"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject206


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse200132.class )

    }

    def deleteResourcePoolGroup ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/resource-pools/groups/${id}"

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

    def getResourcePoolGroups ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/resource-pools/groups"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200131.class )

    }

    def getresourcePoolGroup ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/resource-pools/groups/${id}"

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
                    InlineResponse200132.class )

    }

    def updateResourcePoolGroup ( Long id, InlineObject207 inlineObject207, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/resource-pools/groups/${id}"

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
        bodyParams = inlineObject207


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    InlineResponse200132.class )

    }

}
