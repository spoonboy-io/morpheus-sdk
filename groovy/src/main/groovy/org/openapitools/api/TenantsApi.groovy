package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.Error
import org.openapitools.model.InlineObject238
import org.openapitools.model.InlineObject239
import org.openapitools.model.InlineObject240
import org.openapitools.model.InlineObject241
import org.openapitools.model.InlineObject242
import org.openapitools.model.InlineObject243
import org.openapitools.model.InlineResponse200150
import org.openapitools.model.InlineResponse200151
import org.openapitools.model.InlineResponse200152
import org.openapitools.model.InlineResponse200153
import org.openapitools.model.Model200Success
import org.openapitools.model.TenantsAvailableRoles

class TenantsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addTenant ( InlineObject238 inlineObject238, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/accounts"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject238


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addUserTenant ( Long accountId, InlineObject243 inlineObject243, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/accounts/${accountId}/users"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (accountId == null) {
            throw new RuntimeException("missing required params accountId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject243


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def createTenantSubtenantGroup ( Long accountId, InlineObject240 inlineObject240, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/accounts/${accountId}/groups"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (accountId == null) {
            throw new RuntimeException("missing required params accountId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject240


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse200152.class )

    }

    def getTenant ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/accounts/${id}"

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
                    InlineResponse200150.class )

    }

    def getTenantSubtenantGroup ( Long accountId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/accounts/${accountId}/groups/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (accountId == null) {
            throw new RuntimeException("missing required params accountId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200153.class )

    }

    def listTenantSubtenantGroups ( Long accountId, String phrase, String name, Date lastUpdated, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/accounts/${accountId}/groups"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (accountId == null) {
            throw new RuntimeException("missing required params accountId")
        }

        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (name != null) {
            queryParams.put("name", name)
        }
        if (lastUpdated != null) {
            queryParams.put("lastUpdated", lastUpdated)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listTenants ( Long max, Long offset, String sort, String direction, String phrase, String name, Date lastUpdated, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/accounts"

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
        if (lastUpdated != null) {
            queryParams.put("lastUpdated", lastUpdated)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listTenantsAvailableRoles ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/accounts/available-roles"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    TenantsAvailableRoles.class )

    }

    def removeTenant ( Long id, Boolean removeResources, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/accounts/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (removeResources != null) {
            queryParams.put("removeResources", removeResources)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def removeTenantSubtenantGroup ( Long accountId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/accounts/${accountId}/groups/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (accountId == null) {
            throw new RuntimeException("missing required params accountId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def updateTenant ( Long id, InlineObject239 inlineObject239, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/accounts/${id}"

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
        bodyParams = inlineObject239


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    InlineResponse200151.class )

    }

    def updateTenantSubtenantGroup ( Long accountId, Long id, InlineObject241 inlineObject241, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/accounts/${accountId}/groups/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (accountId == null) {
            throw new RuntimeException("missing required params accountId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject241


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    InlineResponse200152.class )

    }

    def updateTenantSubtenantGroupZones ( Long accountId, Long id, InlineObject242 inlineObject242, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/accounts/${accountId}/groups/${id}/update-zones"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (accountId == null) {
            throw new RuntimeException("missing required params accountId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject242


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

}
