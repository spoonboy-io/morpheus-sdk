package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.AppCreate
import org.openapitools.model.AppPrepareApply
import org.openapitools.model.AppState
import org.openapitools.model.AppUpdate
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject267
import org.openapitools.model.InlineObject3
import org.openapitools.model.InlineObject4
import org.openapitools.model.InlineObject5
import org.openapitools.model.InlineObject6
import org.openapitools.model.InlineResponse200168
import org.openapitools.model.InlineResponse2002
import org.openapitools.model.InlineResponse2003
import org.openapitools.model.Model200Success
import org.openapitools.model.UNKNOWN_BASE_TYPE

class AppsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addAppInstance ( Long id, InlineObject3 inlineObject3, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/apps/${id}/add-instance"

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
        bodyParams = inlineObject3


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse2003.class )

    }

    def addAppUndoDelete ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/apps/${id}/cancel-removal"

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
                    InlineResponse2003.class )

    }

    def addApps ( AppCreate appCreate, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/apps"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = appCreate


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse2002.class )

    }

    def applyAppState ( Long id, InlineObject4 inlineObject4, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/apps/${id}/apply"

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
        bodyParams = inlineObject4


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def deleteApp ( Long id, String removeInstances, String preserveVolumes, String keepBackups, String releaseFloatingIps, String releaseEIPs, String force, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/apps/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (removeInstances != null) {
            queryParams.put("removeInstances", removeInstances)
        }
        if (preserveVolumes != null) {
            queryParams.put("preserveVolumes", preserveVolumes)
        }
        if (keepBackups != null) {
            queryParams.put("keepBackups", keepBackups)
        }
        if (releaseFloatingIps != null) {
            queryParams.put("releaseFloatingIps", releaseFloatingIps)
        }
        if (releaseEIPs != null) {
            queryParams.put("releaseEIPs", releaseEIPs)
        }
        if (force != null) {
            queryParams.put("force", force)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def getApp ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/apps/${id}"

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
                    InlineResponse2003.class )

    }

    def getAppSecurityGroups ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/apps/${id}/security-groups"

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

    def getAppState ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/apps/${id}/state"

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
                    AppState.class )

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

    def listApps ( Long max, Long offset, String name, String phrase, Long createdBy, Boolean showDeleted, String labels, String allLabels, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/apps"

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
        if (name != null) {
            queryParams.put("name", name)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (createdBy != null) {
            queryParams.put("createdBy", createdBy)
        }
        if (showDeleted != null) {
            queryParams.put("showDeleted", showDeleted)
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

    def prepareAppApply ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/apps/${id}/prepare-apply"

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
                    AppPrepareApply.class )

    }

    def refreshAppState ( Long id, Object body, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/apps/${id}/refresh"

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
        bodyParams = body


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def removeAppInstance ( Long id, InlineObject5 inlineObject5, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/apps/${id}/remove-instance"

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
        bodyParams = inlineObject5


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse2003.class )

    }

    def setAppSecurityGroups ( Long id, UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/apps/${id}/security-groups"

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
        bodyParams = UNKNOWN_BASE_TYPE


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def updateApp ( Long id, AppUpdate appUpdate, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/apps/${id}"

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
        bodyParams = appUpdate


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    InlineResponse2003.class )

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

    def validateAppState ( Long id, InlineObject6 inlineObject6, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/apps/${id}/validate-apply"

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
        bodyParams = inlineObject6


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

}
