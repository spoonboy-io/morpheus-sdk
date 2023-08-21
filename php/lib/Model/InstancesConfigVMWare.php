<?php
/**
 * InstancesConfigVMWare
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
 * InstancesConfigVMWare Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class InstancesConfigVMWare implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'instancesConfigVMWare';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'no_agent' => 'bool',
        'resource_pool_id' => 'string',
        'host_id' => 'string',
        'smbios_asset_tag' => 'string',
        'nested_virtualization' => 'string',
        'vmware_folder_id' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'no_agent' => null,
        'resource_pool_id' => null,
        'host_id' => null,
        'smbios_asset_tag' => null,
        'nested_virtualization' => null,
        'vmware_folder_id' => null
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
        'no_agent' => 'noAgent',
        'resource_pool_id' => 'resourcePoolId',
        'host_id' => 'hostId',
        'smbios_asset_tag' => 'smbiosAssetTag',
        'nested_virtualization' => 'nestedVirtualization',
        'vmware_folder_id' => 'vmwareFolderId'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'no_agent' => 'setNoAgent',
        'resource_pool_id' => 'setResourcePoolId',
        'host_id' => 'setHostId',
        'smbios_asset_tag' => 'setSmbiosAssetTag',
        'nested_virtualization' => 'setNestedVirtualization',
        'vmware_folder_id' => 'setVmwareFolderId'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'no_agent' => 'getNoAgent',
        'resource_pool_id' => 'getResourcePoolId',
        'host_id' => 'getHostId',
        'smbios_asset_tag' => 'getSmbiosAssetTag',
        'nested_virtualization' => 'getNestedVirtualization',
        'vmware_folder_id' => 'getVmwareFolderId'
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

    const NESTED_VIRTUALIZATION_ON = 'on';
    const NESTED_VIRTUALIZATION_OFF = 'off';
    

    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getNestedVirtualizationAllowableValues()
    {
        return [
            self::NESTED_VIRTUALIZATION_ON,
            self::NESTED_VIRTUALIZATION_OFF,
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
        $this->container['no_agent'] = $data['no_agent'] ?? false;
        $this->container['resource_pool_id'] = $data['resource_pool_id'] ?? null;
        $this->container['host_id'] = $data['host_id'] ?? null;
        $this->container['smbios_asset_tag'] = $data['smbios_asset_tag'] ?? null;
        $this->container['nested_virtualization'] = $data['nested_virtualization'] ?? 'off';
        $this->container['vmware_folder_id'] = $data['vmware_folder_id'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        $allowedValues = $this->getNestedVirtualizationAllowableValues();
        if (!is_null($this->container['nested_virtualization']) && !in_array($this->container['nested_virtualization'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'nested_virtualization', must be one of '%s'",
                $this->container['nested_virtualization'],
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
     * Gets no_agent
     *
     * @return bool|null
     */
    public function getNoAgent()
    {
        return $this->container['no_agent'];
    }

    /**
     * Sets no_agent
     *
     * @param bool|null $no_agent Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected.
     *
     * @return self
     */
    public function setNoAgent($no_agent)
    {
        $this->container['no_agent'] = $no_agent;

        return $this;
    }

    /**
     * Gets resource_pool_id
     *
     * @return string|null
     */
    public function getResourcePoolId()
    {
        return $this->container['resource_pool_id'];
    }

    /**
     * Sets resource_pool_id
     *
     * @param string|null $resource_pool_id id of the resource group to be used, can be prefixed with `pool-`. A resource pool group can be specified instead by prefixing its ID with `poolGroup-`.
     *
     * @return self
     */
    public function setResourcePoolId($resource_pool_id)
    {
        $this->container['resource_pool_id'] = $resource_pool_id;

        return $this;
    }

    /**
     * Gets host_id
     *
     * @return string|null
     */
    public function getHostId()
    {
        return $this->container['host_id'];
    }

    /**
     * Sets host_id
     *
     * @param string|null $host_id Specific host to deploy to if so desired.
     *
     * @return self
     */
    public function setHostId($host_id)
    {
        $this->container['host_id'] = $host_id;

        return $this;
    }

    /**
     * Gets smbios_asset_tag
     *
     * @return string|null
     */
    public function getSmbiosAssetTag()
    {
        return $this->container['smbios_asset_tag'];
    }

    /**
     * Sets smbios_asset_tag
     *
     * @param string|null $smbios_asset_tag Sets the asset tag on the SMBIOS for use by the guest operating system. If left blank, the virtual machine name will be used.
     *
     * @return self
     */
    public function setSmbiosAssetTag($smbios_asset_tag)
    {
        $this->container['smbios_asset_tag'] = $smbios_asset_tag;

        return $this;
    }

    /**
     * Gets nested_virtualization
     *
     * @return string|null
     */
    public function getNestedVirtualization()
    {
        return $this->container['nested_virtualization'];
    }

    /**
     * Sets nested_virtualization
     *
     * @param string|null $nested_virtualization Enable Nested Virtualization
     *
     * @return self
     */
    public function setNestedVirtualization($nested_virtualization)
    {
        $allowedValues = $this->getNestedVirtualizationAllowableValues();
        if (!is_null($nested_virtualization) && !in_array($nested_virtualization, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'nested_virtualization', must be one of '%s'",
                    $nested_virtualization,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['nested_virtualization'] = $nested_virtualization;

        return $this;
    }

    /**
     * Gets vmware_folder_id
     *
     * @return string|null
     */
    public function getVmwareFolderId()
    {
        return $this->container['vmware_folder_id'];
    }

    /**
     * Sets vmware_folder_id
     *
     * @param string|null $vmware_folder_id VMWare Folder External ID (as a String) or ID (as an Integer or String)
     *
     * @return self
     */
    public function setVmwareFolderId($vmware_folder_id)
    {
        $this->container['vmware_folder_id'] = $vmware_folder_id;

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


