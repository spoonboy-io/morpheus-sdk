# AddCertificateRequestCertificate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A unique name scoped to your account for the key | [optional] 
**CertFile** | **String** | The contents of the certificate file | [optional] 
**KeyFile** | **String** | The contents of the key file | [optional] 
**DomainName** | **String** | The domain name this certificate is tied to | [optional] 
**Wildcard** | **Boolean** | Wether or not this certificate is a wildcard cert | [optional] [default to $false]

## Examples

- Prepare the resource
```powershell
$AddCertificateRequestCertificate = Initialize-PSOpenAPIToolsAddCertificateRequestCertificate  -Name null `
 -CertFile null `
 -KeyFile null `
 -DomainName null `
 -Wildcard null
```

- Convert the resource to JSON
```powershell
$AddCertificateRequestCertificate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

