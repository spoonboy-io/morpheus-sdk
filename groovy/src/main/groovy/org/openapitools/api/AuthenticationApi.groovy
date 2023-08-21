package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.AccessToken
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject10
import org.openapitools.model.InlineObject11
import org.openapitools.model.InlineResponse200167
import org.openapitools.model.InlineResponse2006
import org.openapitools.model.InlineResponse2007

class AuthenticationApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def forgotPassword ( InlineObject11 inlineObject11, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/forgot/send-email"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject11


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse2007.class )

    }

    def getAccessToken ( String clientId, String grantType, String scope, String username, String password, Object refreshToken, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/oauth/token"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clientId == null) {
            throw new RuntimeException("missing required params clientId")
        }
        // verify required params are set
        if (grantType == null) {
            throw new RuntimeException("missing required params grantType")
        }
        // verify required params are set
        if (scope == null) {
            throw new RuntimeException("missing required params scope")
        }

        if (clientId != null) {
            queryParams.put("client_id", clientId)
        }
        if (grantType != null) {
            queryParams.put("grant_type", grantType)
        }
        if (scope != null) {
            queryParams.put("scope", scope)
        }



        contentType = 'application/x-www-form-urlencoded';
        bodyParams = [:]
        bodyParams.put("username", username)
        bodyParams.put("password", password)
        bodyParams.put("refresh_token", refreshToken)

        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    AccessToken.class )

    }

    def resetPassword ( InlineObject10 inlineObject10, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/forgot/reset-password"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject10


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse2006.class )

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
