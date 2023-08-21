package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject196
import org.openapitools.model.InlineObject197
import org.openapitools.model.InlineObject83
import org.openapitools.model.InlineObject84
import org.openapitools.model.InlineObject85
import org.openapitools.model.InlineObject86
import org.openapitools.model.InlineResponse200122
import org.openapitools.model.InlineResponse20052
import org.openapitools.model.InlineResponse20053
import org.openapitools.model.InlineResponse20054
import org.openapitools.model.Model200Success

class ImageBuildsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addBootScript ( InlineObject83 inlineObject83, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/boot-scripts"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject83


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addImageBuild ( InlineObject85 inlineObject85, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/image-builds"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject85


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addPreseedScript ( InlineObject196 inlineObject196, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/preseed-scripts"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject196


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def deleteBootScript ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/boot-scripts/${id}"

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

    def deleteImageBuild ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/image-builds/${id}"

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

    def deletePreseedScript ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/preseed-scripts/${id}"

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

    def executeImageBuild ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/image-builds/${id}/run"

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
                    Model200Success.class )

    }

    def getBootScript ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/boot-scripts/${id}"

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
                    InlineResponse20052.class )

    }

    def getImageBuild ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/image-builds/${id}"

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
                    InlineResponse20053.class )

    }

    def getImageBuildExecutions ( Long id, Long buildNumber, String status, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/image-builds/${id}/list-executions"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (buildNumber != null) {
            queryParams.put("buildNumber", buildNumber)
        }
        if (status != null) {
            queryParams.put("status", status)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getPreseedScript ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/preseed-scripts/${id}"

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
                    InlineResponse200122.class )

    }

    def listBootScripts ( String name, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/boot-scripts"

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
                    Object.class )

    }

    def listImageBuilds ( String name, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/image-builds"

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
                    Object.class )

    }

    def listPreseedScripts ( String name, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/preseed-scripts"

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
                    Object.class )

    }

    def updateBootScript ( Long id, InlineObject84 inlineObject84, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/boot-scripts/${id}"

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
        bodyParams = inlineObject84


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateImageBuild ( Long id, InlineObject86 inlineObject86, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/image-builds/${id}"

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
        bodyParams = inlineObject86


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    InlineResponse20054.class )

    }

    def updatePreseedScript ( Long id, InlineObject197 inlineObject197, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/preseed-scripts/${id}"

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
        bodyParams = inlineObject197


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

}
