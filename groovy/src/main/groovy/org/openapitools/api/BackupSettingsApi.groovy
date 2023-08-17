package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.ListBackupSettings200Response
import org.openapitools.model.Model200Success
import org.openapitools.model.UpdateBackupSettingsRequest

class BackupSettingsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def listBackupSettings ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/backup-settings"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    ListBackupSettings200Response.class )

    }

    def updateBackupSettings ( UpdateBackupSettingsRequest updateBackupSettingsRequest, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/backup-settings"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = updateBackupSettingsRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

}
