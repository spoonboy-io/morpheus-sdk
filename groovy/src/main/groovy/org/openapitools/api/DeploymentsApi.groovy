package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject67
import org.openapitools.model.InlineObject68
import org.openapitools.model.InlineObject69
import org.openapitools.model.InlineObject70
import org.openapitools.model.InlineResponse20038
import org.openapitools.model.InlineResponse20039
import org.openapitools.model.Model200Success

class DeploymentsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addDeploymentFile ( Long deploymentId, Long id, String filepath, File file, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/deployments/${deploymentId}/versions/${id}/files${filepath}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (deploymentId == null) {
            throw new RuntimeException("missing required params deploymentId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }
        // verify required params are set
        if (filepath == null) {
            throw new RuntimeException("missing required params filepath")
        }




        contentType = 'multipart/form-data';
        bodyParams = file

        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def addDeploymentVersion ( Long deploymentId, InlineObject69 inlineObject69, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/deployments/${deploymentId}/versions"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (deploymentId == null) {
            throw new RuntimeException("missing required params deploymentId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject69


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addDeployments ( InlineObject67 inlineObject67, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/deployments"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject67


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def deleteDeployment ( Long deploymentId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/deployments/${deploymentId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (deploymentId == null) {
            throw new RuntimeException("missing required params deploymentId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteDeploymentFile ( Long deploymentId, Long id, String filepath, String force, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/deployments/${deploymentId}/versions/${id}/files${filepath}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (deploymentId == null) {
            throw new RuntimeException("missing required params deploymentId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }
        // verify required params are set
        if (filepath == null) {
            throw new RuntimeException("missing required params filepath")
        }

        if (force != null) {
            queryParams.put("force", force)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteDeploymentVersion ( Long deploymentId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/deployments/${deploymentId}/versions/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (deploymentId == null) {
            throw new RuntimeException("missing required params deploymentId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def getDeployment ( Long deploymentId, Long maxVersions, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/deployments/${deploymentId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (deploymentId == null) {
            throw new RuntimeException("missing required params deploymentId")
        }

        if (maxVersions != null) {
            queryParams.put("maxVersions", maxVersions)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20038.class )

    }

    def getDeploymentVersion ( Long deploymentId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/deployments/${deploymentId}/versions/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (deploymentId == null) {
            throw new RuntimeException("missing required params deploymentId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20039.class )

    }

    def listDeploymentFiles ( Long deploymentId, Long id, String filepath, Long max, Long offset, String phrase, Long version, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/deployments/${deploymentId}/versions/${id}/files${filepath}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (deploymentId == null) {
            throw new RuntimeException("missing required params deploymentId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }
        // verify required params are set
        if (filepath == null) {
            throw new RuntimeException("missing required params filepath")
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
        if (version != null) {
            queryParams.put("version", version)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listDeploymentVersions ( Long deploymentId, Long max, Long offset, String phrase, Long version, String type, String dateCreated, Date lastUpdated, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/deployments/${deploymentId}/versions"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (deploymentId == null) {
            throw new RuntimeException("missing required params deploymentId")
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
        if (version != null) {
            queryParams.put("version", version)
        }
        if (type != null) {
            queryParams.put("type", type)
        }
        if (dateCreated != null) {
            queryParams.put("dateCreated", dateCreated)
        }
        if (lastUpdated != null) {
            queryParams.put("lastUpdated", lastUpdated)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listDeployments ( Long max, Long offset, String phrase, String name, String description, String dateCreated, Date lastUpdated, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/deployments"

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
        if (description != null) {
            queryParams.put("description", description)
        }
        if (dateCreated != null) {
            queryParams.put("dateCreated", dateCreated)
        }
        if (lastUpdated != null) {
            queryParams.put("lastUpdated", lastUpdated)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def updateDeployment ( Long deploymentId, InlineObject68 inlineObject68, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/deployments/${deploymentId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (deploymentId == null) {
            throw new RuntimeException("missing required params deploymentId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject68


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateDeploymentVersion ( Long deploymentId, Long id, InlineObject70 inlineObject70, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/deployments/${deploymentId}/versions/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (deploymentId == null) {
            throw new RuntimeException("missing required params deploymentId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject70


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

}
