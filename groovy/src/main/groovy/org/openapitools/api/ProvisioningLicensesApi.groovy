package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject202
import org.openapitools.model.InlineObject203
import org.openapitools.model.InlineResponse200126
import org.openapitools.model.InlineResponse200127
import org.openapitools.model.Model200Success

class ProvisioningLicensesApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addProvisioningLicense ( InlineObject202 inlineObject202, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/provisioning-licenses"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject202


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def getProvisioningLicense ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/provisioning-licenses/${id}"

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
                    InlineResponse200126.class )

    }

    def getProvisioningLicenseReservations ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/provisioning-licenses/${id}/reservations"

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
                    InlineResponse200127.class )

    }

    def listProvisioningLicenses ( Long max, Long offset, String sort, String direction, String phrase, String name, String licenseType, String licenseVersion, String orgName, String fullName, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/provisioning-licenses"

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
        if (name != null) {
            queryParams.put("name", name)
        }
        if (licenseType != null) {
            queryParams.put("licenseType", licenseType)
        }
        if (licenseVersion != null) {
            queryParams.put("licenseVersion", licenseVersion)
        }
        if (orgName != null) {
            queryParams.put("orgName", orgName)
        }
        if (fullName != null) {
            queryParams.put("fullName", fullName)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def removeProvisioningLicense ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/provisioning-licenses/${id}"

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

    def updateProvisioningLicense ( Long id, InlineObject203 inlineObject203, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/provisioning-licenses/${id}"

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
        bodyParams = inlineObject203


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

}
