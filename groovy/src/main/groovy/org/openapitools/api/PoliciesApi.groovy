package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject184
import org.openapitools.model.InlineObject185
import org.openapitools.model.InlineObject186
import org.openapitools.model.InlineObject187
import org.openapitools.model.InlineObject188
import org.openapitools.model.InlineObject189
import org.openapitools.model.Model200Success

class PoliciesApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addPolicies ( InlineObject184 inlineObject184, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/policies"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject184


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addPoliciesCloud ( Long cloudId, InlineObject188 inlineObject188, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${cloudId}/policies"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (cloudId == null) {
            throw new RuntimeException("missing required params cloudId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject188


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addPoliciesGroup ( Long groupId, InlineObject186 inlineObject186, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/groups/${groupId}/policies"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (groupId == null) {
            throw new RuntimeException("missing required params groupId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject186


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def getPolicies ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/policies/${id}"

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
                    Object.class )

    }

    def getPoliciesCloud ( Long cloudId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${cloudId}/policies/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (cloudId == null) {
            throw new RuntimeException("missing required params cloudId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getPoliciesGroup ( Long groupId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/groups/${groupId}/policies/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (groupId == null) {
            throw new RuntimeException("missing required params groupId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listPolicies ( Long max, Long offset, String sort, String direction, String phrase, String name, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/policies"

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

    def listPoliciesCloud ( Long cloudId, Long max, Long offset, String sort, String direction, String phrase, String name, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${cloudId}/policies"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (cloudId == null) {
            throw new RuntimeException("missing required params cloudId")
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

    def listPoliciesGroup ( Long groupId, Long max, Long offset, String sort, String direction, String phrase, String name, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/groups/${groupId}/policies"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (groupId == null) {
            throw new RuntimeException("missing required params groupId")
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

    def listPolicyTypes ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/policy-types"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def removePolicies ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/policies/${id}"

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

    def removePoliciesCloud ( Long cloudId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${cloudId}/policies/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (cloudId == null) {
            throw new RuntimeException("missing required params cloudId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def removePoliciesGroup ( Long groupId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/groups/${groupId}/policies/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (groupId == null) {
            throw new RuntimeException("missing required params groupId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def updatePolicies ( Long id, InlineObject185 inlineObject185, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/policies/${id}"

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
        bodyParams = inlineObject185


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updatePoliciesCloud ( Long cloudId, Long id, InlineObject189 inlineObject189, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${cloudId}/policies/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (cloudId == null) {
            throw new RuntimeException("missing required params cloudId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject189


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updatePoliciesGroup ( Long groupId, Long id, InlineObject187 inlineObject187, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/groups/${groupId}/policies/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (groupId == null) {
            throw new RuntimeException("missing required params groupId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject187


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

}
