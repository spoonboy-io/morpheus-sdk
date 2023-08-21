package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.Error
import org.openapitools.model.InlineObject208
import org.openapitools.model.InlineObject209
import org.openapitools.model.InlineObject210
import org.openapitools.model.Model200Success
import org.openapitools.model.OneOfrolePermissionBlueprintrolePermissionBlueprintAll
import org.openapitools.model.OneOfrolePermissionCatalogItemTyperolePermissionCatalogItemTypeAll
import org.openapitools.model.OneOfrolePermissionCloudrolePermissionCloudAll
import org.openapitools.model.OneOfrolePermissionFeaturerolePermissionDefaultGrouprolePermissionDefaultCloudrolePermissionDefaultInstanceTyperolePermissionDefaultBlueprintrolePermissionDefaultReportTyperolePermissionDefaultPersonarolePermissionDefaultCatalogItemTyperolePermissionDefaultVDIPoolrolePermissionDefaultTaskSetrolePermissionDefaultTask
import org.openapitools.model.OneOfrolePermissionGrouprolePermissionGroupAll
import org.openapitools.model.OneOfrolePermissionInstanceTyperolePermissionInstanceTypeAll
import org.openapitools.model.OneOfrolePermissionPersonarolePermissionPersonaAll
import org.openapitools.model.OneOfrolePermissionReportTyperolePermissionReportTypeAll
import org.openapitools.model.OneOfrolePermissionTaskSetrolePermissionTaskSetAll
import org.openapitools.model.OneOfrolePermissionTaskrolePermissionTaskAll
import org.openapitools.model.Role
import org.openapitools.model.UNKNOWN_BASE_TYPE

class RolesApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addRoles ( Boolean includeDefaultAccess, InlineObject208 inlineObject208, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/roles"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (includeDefaultAccess != null) {
            queryParams.put("includeDefaultAccess", includeDefaultAccess)
        }


        contentType = 'application/json';
        bodyParams = inlineObject208


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Role.class )

    }

    def deleteRole ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/roles/${id}"

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

    def getRole ( Long id, Boolean includeDefaultAccess, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/roles/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (includeDefaultAccess != null) {
            queryParams.put("includeDefaultAccess", includeDefaultAccess)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Role.class )

    }

    def listRoles ( String phrase, Long max, Long offset, String sort, String direction, String authority, String roleType, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/roles"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (phrase != null) {
            queryParams.put("phrase", phrase)
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
        if (authority != null) {
            queryParams.put("authority", authority)
        }
        if (roleType != null) {
            queryParams.put("roleType", roleType)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def updateRole ( Long id, Boolean includeDefaultAccess, InlineObject209 inlineObject209, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/roles/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (includeDefaultAccess != null) {
            queryParams.put("includeDefaultAccess", includeDefaultAccess)
        }


        contentType = 'application/json';
        bodyParams = inlineObject209


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Role.class )

    }

    def updateRoleBlueprintAccess ( Long id, UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/roles/${id}/update-blueprint"

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
                    "PUT", "",
                    Object.class )

    }

    def updateRoleCatalogItemTypeAccess ( Long id, UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/roles/${id}/update-catalog-item-type"

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
                    "PUT", "",
                    Object.class )

    }

    def updateRoleCloudAccess ( Long id, UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/roles/${id}/update-cloud"

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
                    "PUT", "",
                    Object.class )

    }

    def updateRoleGroupAccess ( Long id, UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/roles/${id}/update-group"

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
                    "PUT", "",
                    Object.class )

    }

    def updateRoleInstanceTypeAccess ( Long id, UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/roles/${id}/update-instance-type"

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
                    "PUT", "",
                    Object.class )

    }

    def updateRolePermission ( Long id, UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/roles/${id}/update-permission"

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
                    "PUT", "",
                    Object.class )

    }

    def updateRolePersonaAccess ( Long id, UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/roles/${id}/update-persona"

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
                    "PUT", "",
                    Object.class )

    }

    def updateRoleReportTypeAccess ( Long id, UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/roles/${id}/update-report-type"

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
                    "PUT", "",
                    Object.class )

    }

    def updateRoleTaskAccess ( Long id, UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/roles/${id}/update-task"

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
                    "PUT", "",
                    Object.class )

    }

    def updateRoleVDIPoolAccess ( Long id, InlineObject210 inlineObject210, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/roles/${id}/update-vdi-pool"

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
        bodyParams = inlineObject210


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateRoleWorkflowAccess ( Long id, UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/roles/${id}/update-task-set"

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
                    "PUT", "",
                    Object.class )

    }

}
