package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject252
import org.openapitools.model.InlineObject255
import org.openapitools.model.InlineObject256
import org.openapitools.model.InlineResponse200157
import org.openapitools.model.InlineResponse200158
import org.openapitools.model.InlineResponse200159
import org.openapitools.model.InlineResponse200167
import org.openapitools.model.Model200Success
import org.openapitools.model.UserSettings
import org.openapitools.model.UserSettingsRegenerateAccessToken
import org.openapitools.model.UsersAvailableRoles

class UsersApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addUser ( Long accountId, InlineObject255 inlineObject255, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/users"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (accountId != null) {
            queryParams.put("accountId", accountId)
        }


        contentType = 'application/json';
        bodyParams = inlineObject255


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def deleteUser ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/users/${id}"

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

    def deleteUserSettingsAccessToken ( Long userId, String clientId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/user-settings/clear-access-token"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (userId != null) {
            queryParams.put("userId", userId)
        }
        if (clientId != null) {
            queryParams.put("clientId", clientId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def deleteUserSettingsAvatar ( Long userId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/user-settings/avatar"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (userId != null) {
            queryParams.put("userId", userId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteUserSettingsDesktopBackground ( Long userId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/user-settings/desktop-background"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (userId != null) {
            queryParams.put("userId", userId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def getUser ( Long id, Boolean includeAccess, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/users/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (includeAccess != null) {
            queryParams.put("includeAccess", includeAccess)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200158.class )

    }

    def getUserPermissions ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/users/${id}/permissions"

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
                    InlineResponse200159.class )

    }

    def getUserSettingsApiClients ( Long userId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/user-settings/api-clients"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (userId != null) {
            queryParams.put("userId", userId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200157.class )

    }

    def listUserSettings ( Long userId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/user-settings"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (userId != null) {
            queryParams.put("userId", userId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    UserSettings.class )

    }

    def listUsers ( Long max, Long offset, String sort, String direction, String phrase, String username, String email, Long roleId, Date lastUpdated, Long accountId, Boolean global, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/users"

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
        if (username != null) {
            queryParams.put("username", username)
        }
        if (email != null) {
            queryParams.put("email", email)
        }
        if (roleId != null) {
            queryParams.put("roleId", roleId)
        }
        if (lastUpdated != null) {
            queryParams.put("lastUpdated", lastUpdated)
        }
        if (accountId != null) {
            queryParams.put("accountId", accountId)
        }
        if (global != null) {
            queryParams.put("global", global)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listUsersAvailableRoles ( Long accountId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/users/available-roles"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (accountId != null) {
            queryParams.put("accountId", accountId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    UsersAvailableRoles.class )

    }

    def updateUserSettings ( Long userId, InlineObject252 inlineObject252, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/user-settings"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (userId != null) {
            queryParams.put("userId", userId)
        }


        contentType = 'application/json';
        bodyParams = inlineObject252


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateUserSettingsAccessToken ( Long userId, String clientId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/user-settings/regenerate-access-token"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (userId != null) {
            queryParams.put("userId", userId)
        }
        if (clientId != null) {
            queryParams.put("clientId", clientId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    UserSettingsRegenerateAccessToken.class )

    }

    def updateUserSettingsAvatar ( Long userId, String userAvatar, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/user-settings/avatar"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (userId != null) {
            queryParams.put("userId", userId)
        }



        contentType = 'multipart/form-data';
        bodyParams = userAvatar

        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def updateUserSettingsDesktopBackground ( Long userId, String userDesktopBackground, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/user-settings/desktop-background"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (userId != null) {
            queryParams.put("userId", userId)
        }



        contentType = 'multipart/form-data';
        bodyParams = userDesktopBackground

        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def updateUsers ( Long id, InlineObject256 inlineObject256, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/users/${id}"

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
        bodyParams = inlineObject256


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def whoami ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/whoami"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200167.class )

    }

}
