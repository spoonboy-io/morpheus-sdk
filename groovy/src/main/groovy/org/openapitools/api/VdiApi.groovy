package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.AnyOfobject200Success
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject257
import org.openapitools.model.InlineObject258
import org.openapitools.model.InlineObject259
import org.openapitools.model.InlineObject260
import org.openapitools.model.InlineObject261
import org.openapitools.model.InlineObject262
import org.openapitools.model.InlineResponse200160
import org.openapitools.model.InlineResponse200161
import org.openapitools.model.InlineResponse200162
import org.openapitools.model.InlineResponse200163
import org.openapitools.model.InlineResponse200164
import org.openapitools.model.Model200Success

class VdiApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addVDIApps ( InlineObject257 inlineObject257, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi-apps"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject257


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    AnyOfobject200Success.class )

    }

    def addVDIGateways ( InlineObject259 inlineObject259, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi-gateways"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject259


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    AnyOfobject200Success.class )

    }

    def addVDIPools ( InlineObject261 inlineObject261, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi-pools"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject261


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    AnyOfobject200Success.class )

    }

    def addVdiAllocation ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi/${id}/allocate"

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
                    "POST", "",
                    Object.class )

    }

    def getVDIAllocations ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi-allocations/${id}"

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
                    InlineResponse200163.class )

    }

    def getVDIApps ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi-apps/${id}"

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
                    InlineResponse200160.class )

    }

    def getVDIGateways ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi-gateways/${id}"

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
                    InlineResponse200161.class )

    }

    def getVDIPools ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi-pools/${id}"

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
                    InlineResponse200162.class )

    }

    def getVdi ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi/${id}"

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
                    InlineResponse200164.class )

    }

    def listVDIAllocations ( Long max, Long offset, String sort, String direction, String phrase, String name, String id, String status, Long poolId, Long userId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi-allocations"

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
        if (name != null) {
            queryParams.put("name", name)
        }
        if (id != null) {
            queryParams.put("id", id)
        }
        if (status != null) {
            queryParams.put("status", status)
        }
        if (poolId != null) {
            queryParams.put("poolId", poolId)
        }
        if (userId != null) {
            queryParams.put("userId", userId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listVDIApps ( Long max, Long offset, String sort, String direction, String phrase, String name, String description, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi-apps"

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
        if (name != null) {
            queryParams.put("name", name)
        }
        if (description != null) {
            queryParams.put("description", description)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listVDIGateways ( Long max, Long offset, String sort, String direction, String phrase, String name, String description, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi-gateways"

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
        if (name != null) {
            queryParams.put("name", name)
        }
        if (description != null) {
            queryParams.put("description", description)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listVDIPools ( Long max, Long offset, String sort, String direction, String phrase, String name, String description, Boolean enabled, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi-pools"

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
        if (name != null) {
            queryParams.put("name", name)
        }
        if (description != null) {
            queryParams.put("description", description)
        }
        if (enabled != null) {
            queryParams.put("enabled", enabled)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listVdi ( String phrase, Long max, Long offset, String sort, String direction, String name, String description, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


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
        if (name != null) {
            queryParams.put("name", name)
        }
        if (description != null) {
            queryParams.put("description", description)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def removeVDIApps ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi-apps/${id}"

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

    def removeVDIGateways ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi-gateways/${id}"

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

    def removeVDIPools ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi-pools/${id}"

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

    def updateVDIApps ( Long id, InlineObject258 inlineObject258, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi-apps/${id}"

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
        bodyParams = inlineObject258


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    AnyOfobject200Success.class )

    }

    def updateVDIGateways ( Long id, InlineObject260 inlineObject260, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi-gateways/${id}"

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
        bodyParams = inlineObject260


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    AnyOfobject200Success.class )

    }

    def updateVDIPools ( Long id, InlineObject262 inlineObject262, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/vdi-pools/${id}"

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
        bodyParams = inlineObject262


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    AnyOfobject200Success.class )

    }

}
