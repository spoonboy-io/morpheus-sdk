package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject230
import org.openapitools.model.InlineObject231
import org.openapitools.model.InlineResponse200143
import org.openapitools.model.InlineResponse200144
import org.openapitools.model.Model200Success

class SslCertificatesApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addCertificate ( InlineObject230 inlineObject230, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/certificates"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject230


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def deleteCertificate ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/certificates/${id}"

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

    def getCertificate ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/certificates/${id}"

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
                    InlineResponse200144.class )

    }

    def listCertificates ( String name, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/certificates"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (name != null) {
            queryParams.put("name", name)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200143.class )

    }

    def updateCertificate ( Long id, InlineObject231 inlineObject231, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/certificates/${id}"

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
        bodyParams = inlineObject231


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    InlineResponse200144.class )

    }

}
