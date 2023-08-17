package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.AddAppInstanceRequest
import org.openapitools.model.AddApps200Response
import org.openapitools.model.AppCreate
import org.openapitools.model.AppUpdate
import org.openapitools.model.ApplyAppStateRequest
import org.openapitools.model.DefaultError
import org.openapitools.model.GetApp200Response
import org.openapitools.model.GetAppSecurityGroups200Response
import org.openapitools.model.GetAppState200Response
import org.openapitools.model.ListApps200Response
import org.openapitools.model.Model200Success
import org.openapitools.model.PrepareAppApply200Response
import org.openapitools.model.RemoveAppInstanceRequest
import org.openapitools.model.SetAppSecurityGroups200Response
import org.openapitools.model.SetAppSecurityGroupsRequest
import org.openapitools.model.ValidateAppState200Response

class AppsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addAppInstance ( Long id, AddAppInstanceRequest addAppInstanceRequest, Closure onSuccess, Closure onFailure)  {
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
        bodyParams = addAppInstanceRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    GetApp200Response.class )

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
                    GetApp200Response.class )

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
                    AddApps200Response.class )

    }

    def applyAppState ( Long id, ApplyAppStateRequest applyAppStateRequest, Closure onSuccess, Closure onFailure)  {
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
        bodyParams = applyAppStateRequest


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
                    GetApp200Response.class )

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
                    GetAppSecurityGroups200Response.class )

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
                    GetAppState200Response.class )

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
                    ListApps200Response.class )

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
                    PrepareAppApply200Response.class )

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

    def removeAppInstance ( Long id, RemoveAppInstanceRequest removeAppInstanceRequest, Closure onSuccess, Closure onFailure)  {
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
        bodyParams = removeAppInstanceRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    GetApp200Response.class )

    }

    def setAppSecurityGroups ( Long id, SetAppSecurityGroupsRequest setAppSecurityGroupsRequest, Closure onSuccess, Closure onFailure)  {
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
        bodyParams = setAppSecurityGroupsRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SetAppSecurityGroups200Response.class )

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
                    GetApp200Response.class )

    }

    def validateAppState ( Long id, ApplyAppStateRequest applyAppStateRequest, Closure onSuccess, Closure onFailure)  {
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
        bodyParams = applyAppStateRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    ValidateAppState200Response.class )

    }

}
