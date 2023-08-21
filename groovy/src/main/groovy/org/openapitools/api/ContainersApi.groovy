package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject62
import org.openapitools.model.InlineObject63
import org.openapitools.model.InlineResponse20034
import org.openapitools.model.InlineResponse20035
import org.openapitools.model.InstancesCloneImage
import org.openapitools.model.Model200Success
import org.openapitools.model.SuccessMessage

class ContainersApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def cloneImageContainerAction ( Long id, InstancesCloneImage instancesCloneImage, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/containers/${id}/clone-image"

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
        bodyParams = instancesCloneImage


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    SuccessMessage.class )

    }

    def containersAttachFloatingIp ( Long id, InlineObject63 inlineObject63, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/containers/${id}/attach-floating-ip"

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
        bodyParams = inlineObject63


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def containersDetachFloatingIp ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/containers/${id}/detach-floating-ip"

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
                    "PUT", "",
                    Model200Success.class )

    }

    def ejectContainerAction ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/containers/${id}/eject"

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
                    "PUT", "",
                    SuccessMessage.class )

    }

    def executeContainerAction ( Long id, String code, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/containers/${id}/action"

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
        if (code == null) {
            throw new RuntimeException("missing required params code")
        }

        if (code != null) {
            queryParams.put("code", code)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    SuccessMessage.class )

    }

    def getContainer ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/containers/${id}"

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
                    InlineResponse20034.class )

    }

    def getContainerActions ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/containers/${id}/actions"

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
                    InlineResponse20035.class )

    }

    def importContainerAction ( Long id, InlineObject62 inlineObject62, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/containers/${id}/import"

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
        bodyParams = inlineObject62


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    SuccessMessage.class )

    }

    def restartContainerAction ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/containers/${id}/restart"

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
                    "PUT", "",
                    SuccessMessage.class )

    }

    def startContainerAction ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/containers/${id}/start"

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
                    "PUT", "",
                    SuccessMessage.class )

    }

    def stopContainerAction ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/containers/${id}/stop"

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
                    "PUT", "",
                    SuccessMessage.class )

    }

    def suspendContainerAction ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/containers/${id}/suspend"

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
                    "PUT", "",
                    SuccessMessage.class )

    }

}
