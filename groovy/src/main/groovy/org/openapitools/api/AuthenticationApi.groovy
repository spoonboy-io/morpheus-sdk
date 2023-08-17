package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.ForgotPassword200Response
import org.openapitools.model.ForgotPasswordRequest
import org.openapitools.model.ResetPassword200Response
import org.openapitools.model.ResetPasswordRequest

class AuthenticationApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def forgotPassword ( ForgotPasswordRequest forgotPasswordRequest, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/forgot/send-email"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = forgotPasswordRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    ForgotPassword200Response.class )

    }

    def resetPassword ( ResetPasswordRequest resetPasswordRequest, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/forgot/reset-password"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = resetPasswordRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    ResetPassword200Response.class )

    }

}
