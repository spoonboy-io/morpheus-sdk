package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject267
import org.openapitools.model.InlineObject268
import org.openapitools.model.InlineObject269
import org.openapitools.model.InlineObject270
import org.openapitools.model.InlineObject271
import org.openapitools.model.InlineObject272
import org.openapitools.model.InlineObject273
import org.openapitools.model.InlineObject274
import org.openapitools.model.InlineResponse200168
import org.openapitools.model.InlineResponse200169
import org.openapitools.model.Model200Success

class WikiApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addWiki ( InlineObject272 inlineObject272, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/wiki/pages"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject272


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def getWiki ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/wiki/pages/${id}"

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
                    InlineResponse200168.class )

    }

    def getWikiApp ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/apps/${id}/wiki"

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
                    InlineResponse200168.class )

    }

    def getWikiCategories ( String phrase, String pagePhrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/wiki/categories"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (pagePhrase != null) {
            queryParams.put("pagePhrase", pagePhrase)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200169.class )

    }

    def getWikiCloud ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${id}/wiki"

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
                    InlineResponse200168.class )

    }

    def getWikiCluster ( Integer clusterId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/wiki"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200168.class )

    }

    def getWikiGroup ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/groups/${id}/wiki"

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
                    InlineResponse200168.class )

    }

    def getWikiInstance ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/wiki"

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
                    InlineResponse200168.class )

    }

    def getWikiServer ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/${id}/wiki"

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
                    InlineResponse200168.class )

    }

    def listWiki ( String name, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/wiki/pages"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (name != null) {
            queryParams.put("name", name)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200168.class )

    }

    def removeWiki ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/wiki/pages/${id}"

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

    def updateWiki ( Long id, InlineObject273 inlineObject273, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/wiki/pages/${id}"

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
        bodyParams = inlineObject273


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateWikiApp ( Long id, InlineObject267 inlineObject267, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/apps/${id}/wiki"

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
        bodyParams = inlineObject267


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateWikiCloud ( Long id, InlineObject274 inlineObject274, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${id}/wiki"

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
        bodyParams = inlineObject274


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateWikiCluster ( Integer clusterId, InlineObject268 inlineObject268, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/wiki"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject268


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateWikiGroup ( Long id, InlineObject269 inlineObject269, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/groups/${id}/wiki"

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
        bodyParams = inlineObject269


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateWikiInstance ( Long id, InlineObject270 inlineObject270, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/wiki"

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
        bodyParams = inlineObject270


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateWikiServer ( Long id, InlineObject271 inlineObject271, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/${id}/wiki"

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
        bodyParams = inlineObject271


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

}
