<?php
/**
 * ApiZonesZoneIdDataStoresIdDatastore
 *
 * PHP version 7.2
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Morpheus API
 *
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.0.0
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace OpenAPI\Client\Model;

use \ArrayAccess;
use \OpenAPI\Client\ObjectSerializer;

/**
 * ApiZonesZoneIdDataStoresIdDatastore Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ApiZonesZoneIdDataStoresIdDatastore implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = '_api_zones__zoneId__data_stores__id__datastore';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'active' => 'bool',
        'visibility' => 'string',
        'tenant_permissions' => '\OpenAPI\Client\Model\ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions[]',
        'resource_permissions' => '\OpenAPI\Client\Model\ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'active' => null,
        'visibility' => null,
        'tenant_permissions' => null,
        'resource_permissions' => null
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'active' => 'active',
        'visibility' => 'visibility',
        'tenant_permissions' => 'tenantPermissions',
        'resource_permissions' => 'resourcePermissions'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'active' => 'setActive',
        'visibility' => 'setVisibility',
        'tenant_permissions' => 'setTenantPermissions',
        'resource_permissions' => 'setResourcePermissions'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'active' => 'getActive',
        'visibility' => 'getVisibility',
        'tenant_permissions' => 'getTenantPermissions',
        'resource_permissions' => 'getResourcePermissions'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }

    const VISIBILITY__PUBLIC = 'public';
    const VISIBILITY__PRIVATE = 'private';
    

    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getVisibilityAllowableValues()
    {
        return [
            self::VISIBILITY__PUBLIC,
            self::VISIBILITY__PRIVATE,
        ];
    }
    

    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->container['active'] = $data['active'] ?? null;
        $this->container['visibility'] = $data['visibility'] ?? 'private';
        $this->container['tenant_permissions'] = $data['tenant_permissions'] ?? null;
        $this->container['resource_permissions'] = $data['resource_permissions'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        $allowedValues = $this->getVisibilityAllowableValues();
        if (!is_null($this->container['visibility']) && !in_array($this->container['visibility'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'visibility', must be one of '%s'",
                $this->container['visibility'],
                implode("', '", $allowedValues)
            );
        }

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets active
     *
     * @return bool|null
     */
    public function getActive()
    {
        return $this->container['active'];
    }

    /**
     * Sets active
     *
     * @param bool|null $active Activate `true` or disable `false` the datastore
     *
     * @return self
     */
    public function setActive($active)
    {
        $this->container['active'] = $active;

        return $this;
    }

    /**
     * Gets visibility
     *
     * @return string|null
     */
    public function getVisibility()
    {
        return $this->container['visibility'];
    }

    /**
     * Sets visibility
     *
     * @param string|null $visibility Setting `private` or `public`
     *
     * @return self
     */
    public function setVisibility($visibility)
    {
        $allowedValues = $this->getVisibilityAllowableValues();
        if (!is_null($visibility) && !in_array($visibility, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'visibility', must be one of '%s'",
                    $visibility,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['visibility'] = $visibility;

        return $this;
    }

    /**
     * Gets tenant_permissions
     *
     * @return \OpenAPI\Client\Model\ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions[]|null
     */
    public function getTenantPermissions()
    {
        return $this->container['tenant_permissions'];
    }

    /**
     * Sets tenant_permissions
     *
     * @param \OpenAPI\Client\Model\ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions[]|null $tenant_permissions tenant_permissions
     *
     * @return self
     */
    public function setTenantPermissions($tenant_permissions)
    {
        $this->container['tenant_permissions'] = $tenant_permissions;

        return $this;
    }

    /**
     * Gets resource_permissions
     *
     * @return \OpenAPI\Client\Model\ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions|null
     */
    public function getResourcePermissions()
    {
        return $this->container['resource_permissions'];
    }

    /**
     * Sets resource_permissions
     *
     * @param \OpenAPI\Client\Model\ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions|null $resource_permissions resource_permissions
     *
     * @return self
     */
    public function setResourcePermissions($resource_permissions)
    {
        $this->container['resource_permissions'] = $resource_permissions;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}

