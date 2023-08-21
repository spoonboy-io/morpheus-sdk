package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject125
import org.openapitools.model.InlineObject126
import org.openapitools.model.License
import org.openapitools.model.Model200Success

class LicenseApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def getLicense ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/license"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    License.class )

    }

    def installLicense ( InlineObject125 inlineObject125, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/license"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject125


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    License.class )

    }

    def testLicense ( InlineObject126 inlineObject126, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/license/test"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject126


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    License.class )

    }

    def uninstallLicense ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/license"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

}
