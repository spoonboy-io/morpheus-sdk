<?php
/**
 * ZoneCreate
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
 * ZoneCreate Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ZoneCreate implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'zoneCreate';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'name' => 'string',
        'description' => 'string',
        'code' => 'string',
        'location' => 'string',
        'visibility' => 'string',
        'zone_type' => 'AnyOfObjectObject',
        'group_id' => 'int',
        'account_id' => 'int',
        'enabled' => 'bool',
        'auto_recover_power_state' => 'bool',
        'scale_priority' => 'int',
        'linked_account_id' => 'int',
        'config' => 'object',
        'security_mode' => 'string',
        'credential' => '\OpenAPI\Client\Model\ZoneCreateCredential'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'name' => null,
        'description' => null,
        'code' => null,
        'location' => null,
        'visibility' => null,
        'zone_type' => null,
        'group_id' => 'int64',
        'account_id' => 'int64',
        'enabled' => null,
        'auto_recover_power_state' => null,
        'scale_priority' => 'int64',
        'linked_account_id' => 'int64',
        'config' => null,
        'security_mode' => null,
        'credential' => null
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
        'name' => 'name',
        'description' => 'description',
        'code' => 'code',
        'location' => 'location',
        'visibility' => 'visibility',
        'zone_type' => 'zoneType',
        'group_id' => 'groupId',
        'account_id' => 'accountId',
        'enabled' => 'enabled',
        'auto_recover_power_state' => 'autoRecoverPowerState',
        'scale_priority' => 'scalePriority',
        'linked_account_id' => 'linkedAccountId',
        'config' => 'config',
        'security_mode' => 'securityMode',
        'credential' => 'credential'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'name' => 'setName',
        'description' => 'setDescription',
        'code' => 'setCode',
        'location' => 'setLocation',
        'visibility' => 'setVisibility',
        'zone_type' => 'setZoneType',
        'group_id' => 'setGroupId',
        'account_id' => 'setAccountId',
        'enabled' => 'setEnabled',
        'auto_recover_power_state' => 'setAutoRecoverPowerState',
        'scale_priority' => 'setScalePriority',
        'linked_account_id' => 'setLinkedAccountId',
        'config' => 'setConfig',
        'security_mode' => 'setSecurityMode',
        'credential' => 'setCredential'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'name' => 'getName',
        'description' => 'getDescription',
        'code' => 'getCode',
        'location' => 'getLocation',
        'visibility' => 'getVisibility',
        'zone_type' => 'getZoneType',
        'group_id' => 'getGroupId',
        'account_id' => 'getAccountId',
        'enabled' => 'getEnabled',
        'auto_recover_power_state' => 'getAutoRecoverPowerState',
        'scale_priority' => 'getScalePriority',
        'linked_account_id' => 'getLinkedAccountId',
        'config' => 'getConfig',
        'security_mode' => 'getSecurityMode',
        'credential' => 'getCredential'
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

    const VISIBILITY__PRIVATE = 'private';
    const VISIBILITY__PUBLIC = 'public';
    

    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getVisibilityAllowableValues()
    {
        return [
            self::VISIBILITY__PRIVATE,
            self::VISIBILITY__PUBLIC,
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
        $this->container['name'] = $data['name'] ?? null;
        $this->container['description'] = $data['description'] ?? null;
        $this->container['code'] = $data['code'] ?? null;
        $this->container['location'] = $data['location'] ?? null;
        $this->container['visibility'] = $data['visibility'] ?? 'private';
        $this->container['zone_type'] = $data['zone_type'] ?? null;
        $this->container['group_id'] = $data['group_id'] ?? null;
        $this->container['account_id'] = $data['account_id'] ?? null;
        $this->container['enabled'] = $data['enabled'] ?? true;
        $this->container['auto_recover_power_state'] = $data['auto_recover_power_state'] ?? false;
        $this->container['scale_priority'] = $data['scale_priority'] ?? 1;
        $this->container['linked_account_id'] = $data['linked_account_id'] ?? null;
        $this->container['config'] = $data['config'] ?? null;
        $this->container['security_mode'] = $data['security_mode'] ?? 'off';
        $this->container['credential'] = $data['credential'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if ($this->container['name'] === null) {
            $invalidProperties[] = "'name' can't be null";
        }
        $allowedValues = $this->getVisibilityAllowableValues();
        if (!is_null($this->container['visibility']) && !in_array($this->container['visibility'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'visibility', must be one of '%s'",
                $this->container['visibility'],
                implode("', '", $allowedValues)
            );
        }

        if ($this->container['zone_type'] === null) {
            $invalidProperties[] = "'zone_type' can't be null";
        }
        if ($this->container['group_id'] === null) {
            $invalidProperties[] = "'group_id' can't be null";
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
     * Gets name
     *
     * @return string
     */
    public function getName()
    {
        return $this->container['name'];
    }

    /**
     * Sets name
     *
     * @param string $name A unique name scoped to your account for the cloud
     *
     * @return self
     */
    public function setName($name)
    {
        $this->container['name'] = $name;

        return $this;
    }

    /**
     * Gets description
     *
     * @return string|null
     */
    public function getDescription()
    {
        return $this->container['description'];
    }

    /**
     * Sets description
     *
     * @param string|null $description Optional description field if you want to put more info there
     *
     * @return self
     */
    public function setDescription($description)
    {
        $this->container['description'] = $description;

        return $this;
    }

    /**
     * Gets code
     *
     * @return string|null
     */
    public function getCode()
    {
        return $this->container['code'];
    }

    /**
     * Sets code
     *
     * @param string|null $code Optional code for use with policies
     *
     * @return self
     */
    public function setCode($code)
    {
        $this->container['code'] = $code;

        return $this;
    }

    /**
     * Gets location
     *
     * @return string|null
     */
    public function getLocation()
    {
        return $this->container['location'];
    }

    /**
     * Sets location
     *
     * @param string|null $location Optional location for your cloud
     *
     * @return self
     */
    public function setLocation($location)
    {
        $this->container['location'] = $location;

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
     * @param string|null $visibility private or public
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
     * Gets zone_type
     *
     * @return AnyOfObjectObject
     */
    public function getZoneType()
    {
        return $this->container['zone_type'];
    }

    /**
     * Sets zone_type
     *
     * @param AnyOfObjectObject $zone_type zone_type
     *
     * @return self
     */
    public function setZoneType($zone_type)
    {
        $this->container['zone_type'] = $zone_type;

        return $this;
    }

    /**
     * Gets group_id
     *
     * @return int
     */
    public function getGroupId()
    {
        return $this->container['group_id'];
    }

    /**
     * Sets group_id
     *
     * @param int $group_id Specifies which Server group this cloud should be assigned to
     *
     * @return self
     */
    public function setGroupId($group_id)
    {
        $this->container['group_id'] = $group_id;

        return $this;
    }

    /**
     * Gets account_id
     *
     * @return int|null
     */
    public function getAccountId()
    {
        return $this->container['account_id'];
    }

    /**
     * Sets account_id
     *
     * @param int|null $account_id Specifies which Tenant this cloud should be assigned to
     *
     * @return self
     */
    public function setAccountId($account_id)
    {
        $this->container['account_id'] = $account_id;

        return $this;
    }

    /**
     * Gets enabled
     *
     * @return bool|null
     */
    public function getEnabled()
    {
        return $this->container['enabled'];
    }

    /**
     * Sets enabled
     *
     * @param bool|null $enabled Can be used to disable the cloud
     *
     * @return self
     */
    public function setEnabled($enabled)
    {
        $this->container['enabled'] = $enabled;

        return $this;
    }

    /**
     * Gets auto_recover_power_state
     *
     * @return bool|null
     */
    public function getAutoRecoverPowerState()
    {
        return $this->container['auto_recover_power_state'];
    }

    /**
     * Sets auto_recover_power_state
     *
     * @param bool|null $auto_recover_power_state Automatically Power on VMs
     *
     * @return self
     */
    public function setAutoRecoverPowerState($auto_recover_power_state)
    {
        $this->container['auto_recover_power_state'] = $auto_recover_power_state;

        return $this;
    }

    /**
     * Gets scale_priority
     *
     * @return int|null
     */
    public function getScalePriority()
    {
        return $this->container['scale_priority'];
    }

    /**
     * Sets scale_priority
     *
     * @param int|null $scale_priority Scale Priority
     *
     * @return self
     */
    public function setScalePriority($scale_priority)
    {
        $this->container['scale_priority'] = $scale_priority;

        return $this;
    }

    /**
     * Gets linked_account_id
     *
     * @return int|null
     */
    public function getLinkedAccountId()
    {
        return $this->container['linked_account_id'];
    }

    /**
     * Sets linked_account_id
     *
     * @param int|null $linked_account_id Linked Account ID (enter commercial ID to get costing for AWS Govcloud)
     *
     * @return self
     */
    public function setLinkedAccountId($linked_account_id)
    {
        $this->container['linked_account_id'] = $linked_account_id;

        return $this;
    }

    /**
     * Gets config
     *
     * @return object|null
     */
    public function getConfig()
    {
        return $this->container['config'];
    }

    /**
     * Sets config
     *
     * @param object|null $config Map containing zone configuration settings. See the section on specific zone types for details.
     *
     * @return self
     */
    public function setConfig($config)
    {
        $this->container['config'] = $config;

        return $this;
    }

    /**
     * Gets security_mode
     *
     * @return string|null
     */
    public function getSecurityMode()
    {
        return $this->container['security_mode'];
    }

    /**
     * Sets security_mode
     *
     * @param string|null $security_mode host firewall. `off` or `internal`. a.k.a. \"local firewall\"
     *
     * @return self
     */
    public function setSecurityMode($security_mode)
    {
        $this->container['security_mode'] = $security_mode;

        return $this;
    }

    /**
     * Gets credential
     *
     * @return \OpenAPI\Client\Model\ZoneCreateCredential|null
     */
    public function getCredential()
    {
        return $this->container['credential'];
    }

    /**
     * Sets credential
     *
     * @param \OpenAPI\Client\Model\ZoneCreateCredential|null $credential credential
     *
     * @return self
     */
    public function setCredential($credential)
    {
        $this->container['credential'] = $credential;

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


