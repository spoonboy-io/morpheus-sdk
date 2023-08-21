package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject72
import org.openapitools.model.InlineObject73
import org.openapitools.model.InlineObject92
import org.openapitools.model.InlineResponse20040
import org.openapitools.model.Model200Success

class DeploysApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addInstanceDeploy ( Long id, InlineObject92 inlineObject92, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/deploys"

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
        bodyParams = inlineObject92


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse20040.class )

    }

    def deletedeploy ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/deploys/${id}"

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

    def getInstanceDeploys ( Long id, Long max, Long offset, String phrase, String name, Long deploymentId, String instanceName, Long instanceId, Long version, Long versionId, Long createdById, String deployType, String dateCreated, Date lastUpdated, String deployDate, String status, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/instances/${id}/deploys"

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
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (name != null) {
            queryParams.put("name", name)
        }
        if (deploymentId != null) {
            queryParams.put("deploymentId", deploymentId)
        }
        if (instanceName != null) {
            queryParams.put("instanceName", instanceName)
        }
        if (instanceId != null) {
            queryParams.put("instanceId", instanceId)
        }
        if (version != null) {
            queryParams.put("version", version)
        }
        if (versionId != null) {
            queryParams.put("versionId", versionId)
        }
        if (createdById != null) {
            queryParams.put("createdById", createdById)
        }
        if (deployType != null) {
            queryParams.put("deployType", deployType)
        }
        if (dateCreated != null) {
            queryParams.put("dateCreated", dateCreated)
        }
        if (lastUpdated != null) {
            queryParams.put("lastUpdated", lastUpdated)
        }
        if (deployDate != null) {
            queryParams.put("deployDate", deployDate)
        }
        if (status != null) {
            queryParams.put("status", status)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listDeploys ( Long max, Long offset, String phrase, String name, Long deploymentId, String instanceName, Long instanceId, Long version, Long versionId, Long createdById, String deployType, String dateCreated, Date lastUpdated, String deployDate, String status, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/deploys"

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
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (name != null) {
            queryParams.put("name", name)
        }
        if (deploymentId != null) {
            queryParams.put("deploymentId", deploymentId)
        }
        if (instanceName != null) {
            queryParams.put("instanceName", instanceName)
        }
        if (instanceId != null) {
            queryParams.put("instanceId", instanceId)
        }
        if (version != null) {
            queryParams.put("version", version)
        }
        if (versionId != null) {
            queryParams.put("versionId", versionId)
        }
        if (createdById != null) {
            queryParams.put("createdById", createdById)
        }
        if (deployType != null) {
            queryParams.put("deployType", deployType)
        }
        if (dateCreated != null) {
            queryParams.put("dateCreated", dateCreated)
        }
        if (lastUpdated != null) {
            queryParams.put("lastUpdated", lastUpdated)
        }
        if (deployDate != null) {
            queryParams.put("deployDate", deployDate)
        }
        if (status != null) {
            queryParams.put("status", status)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def runDeploy ( Long id, InlineObject73 inlineObject73, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/deploys/${id}/deploy"

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
        bodyParams = inlineObject73


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse20040.class )

    }

    def updateDeploy ( Long id, InlineObject72 inlineObject72, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/deploys/${id}"

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
        bodyParams = inlineObject72


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    InlineResponse20040.class )

    }

}
