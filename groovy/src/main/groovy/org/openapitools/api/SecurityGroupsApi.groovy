package org.openapitools.api;

import org.openapitools.api.ApiUtils
import java.math.BigDecimal
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject213
import org.openapitools.model.InlineObject214
import org.openapitools.model.InlineObject215
import org.openapitools.model.InlineObject216
import org.openapitools.model.InlineObject217
import org.openapitools.model.InlineResponse200133
import org.openapitools.model.InlineResponse200134
import org.openapitools.model.InlineResponse200135
import org.openapitools.model.Model200Success

class SecurityGroupsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addSecurityGroupLocations ( Long id, InlineObject215 inlineObject215, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/security-groups/${id}/locations"

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
        bodyParams = inlineObject215


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addSecurityGroupRules ( Long id, InlineObject216 inlineObject216, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/security-groups/${id}/rules"

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
        bodyParams = inlineObject216


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addSecurityGroups ( InlineObject213 inlineObject213, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/security-groups"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject213


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse200133.class )

    }

    def getSecurityGroupRules ( Long id, BigDecimal sgId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/security-groups/${id}/rules/${sgId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }
        // verify required params are set
        if (sgId == null) {
            throw new RuntimeException("missing required params sgId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200135.class )

    }

    def getSecurityGroups ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/security-groups/${id}"

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
                    InlineResponse200134.class )

    }

    def listSecurityGroupRules ( Long id, Long max, Long offset, String sort, String direction, String phrase, String name, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/security-groups/${id}/rules"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
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




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listSecurityGroups ( Long max, Long offset, String sort, String direction, String phrase, String name, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/security-groups"

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




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def removeSecurityGroupLocations ( Long id, BigDecimal locationId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/security-groups/${id}/locations/${locationId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }
        // verify required params are set
        if (locationId == null) {
            throw new RuntimeException("missing required params locationId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def removeSecurityGroupRules ( Long id, BigDecimal sgId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/security-groups/${id}/rules/${sgId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }
        // verify required params are set
        if (sgId == null) {
            throw new RuntimeException("missing required params sgId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def removeSecurityGroups ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/security-groups/${id}"

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

    def updateSecurityGroupRules ( Long id, BigDecimal sgId, InlineObject217 inlineObject217, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/security-groups/${id}/rules/${sgId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }
        // verify required params are set
        if (sgId == null) {
            throw new RuntimeException("missing required params sgId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject217


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateSecurityGroups ( Long id, InlineObject214 inlineObject214, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/security-groups/${id}"

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
        bodyParams = inlineObject214


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    InlineResponse200133.class )

    }

}
