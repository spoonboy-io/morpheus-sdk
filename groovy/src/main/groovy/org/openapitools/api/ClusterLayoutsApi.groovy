package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject50
import org.openapitools.model.InlineObject51
import org.openapitools.model.InlineResponse20025
import org.openapitools.model.Model200Success
import org.openapitools.model.SuccessId

class ClusterLayoutsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addClusterLayoutClone ( Long id, String name, String description, String computeVersion, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/cluster-layouts/${id}/clone"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (name != null) {
            queryParams.put("name", name)
        }
        if (description != null) {
            queryParams.put("description", description)
        }
        if (computeVersion != null) {
            queryParams.put("computeVersion", computeVersion)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessId.class )

    }

    def addClusterLayouts ( InlineObject50 inlineObject50, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/cluster-layouts"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject50


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessId.class )

    }

    def deleteClusterLayout ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/cluster-layouts/${id}"

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

    def getClusterLayout ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/cluster-layouts/${id}"

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
                    InlineResponse20025.class )

    }

    def listClusterLayouts ( Long max, Long offset, String sort, String direction, String phrase, String provisionType, String labels, String allLabels, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/cluster-layouts"

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
        if (provisionType != null) {
            queryParams.put("provisionType", provisionType)
        }
        if (labels != null) {
            queryParams.put("labels", labels)
        }
        if (allLabels != null) {
            queryParams.put("allLabels", allLabels)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def updateClusterLayout ( Long id, InlineObject51 inlineObject51, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/cluster-layouts/${id}"

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
        bodyParams = inlineObject51


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    SuccessId.class )

    }

}
