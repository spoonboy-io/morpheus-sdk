# ApiSecurityGroupsIdLocationsSecurityGroupLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneId** | **int64** | The ID of the Zone (Cloud) | 
**CustomOptions** | [**AnyOfsecurityGroupLocationAzureCustomOptionssecurityGroupLocationAwsCustomOptionssecurityGroupLocationOpenstackCustomOptions**](anyOf&lt;securityGroupLocationAzureCustomOptions,securityGroupLocationAwsCustomOptions,securityGroupLocationOpenstackCustomOptions&gt;.md) |  | 

## Methods

### NewApiSecurityGroupsIdLocationsSecurityGroupLocation

`func NewApiSecurityGroupsIdLocationsSecurityGroupLocation(zoneId int64, customOptions AnyOfsecurityGroupLocationAzureCustomOptionssecurityGroupLocationAwsCustomOptionssecurityGroupLocationOpenstackCustomOptions, ) *ApiSecurityGroupsIdLocationsSecurityGroupLocation`

NewApiSecurityGroupsIdLocationsSecurityGroupLocation instantiates a new ApiSecurityGroupsIdLocationsSecurityGroupLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSecurityGroupsIdLocationsSecurityGroupLocationWithDefaults

`func NewApiSecurityGroupsIdLocationsSecurityGroupLocationWithDefaults() *ApiSecurityGroupsIdLocationsSecurityGroupLocation`

NewApiSecurityGroupsIdLocationsSecurityGroupLocationWithDefaults instantiates a new ApiSecurityGroupsIdLocationsSecurityGroupLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneId

`func (o *ApiSecurityGroupsIdLocationsSecurityGroupLocation) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ApiSecurityGroupsIdLocationsSecurityGroupLocation) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ApiSecurityGroupsIdLocationsSecurityGroupLocation) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.


### GetCustomOptions

`func (o *ApiSecurityGroupsIdLocationsSecurityGroupLocation) GetCustomOptions() AnyOfsecurityGroupLocationAzureCustomOptionssecurityGroupLocationAwsCustomOptionssecurityGroupLocationOpenstackCustomOptions`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *ApiSecurityGroupsIdLocationsSecurityGroupLocation) GetCustomOptionsOk() (*AnyOfsecurityGroupLocationAzureCustomOptionssecurityGroupLocationAwsCustomOptionssecurityGroupLocationOpenstackCustomOptions, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *ApiSecurityGroupsIdLocationsSecurityGroupLocation) SetCustomOptions(v AnyOfsecurityGroupLocationAzureCustomOptionssecurityGroupLocationAwsCustomOptionssecurityGroupLocationOpenstackCustomOptions)`

SetCustomOptions sets CustomOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


