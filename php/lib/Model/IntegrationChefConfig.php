<?php
/**
 * IntegrationChefConfig
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
 * IntegrationChefConfig Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class IntegrationChefConfig implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'integrationChef_config';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'databags' => '\OpenAPI\Client\Model\IntegrationChefConfigDatabags[]',
        'endpoint' => 'string',
        'org' => 'string',
        'chef_user' => 'string',
        'user_key' => 'string',
        'org_key' => 'string',
        'version' => 'string',
        'chef_use_fqdn' => 'bool',
        'windows_version' => 'string',
        'windows_install_url' => 'string',
        'user_key_hash' => 'string',
        'org_key_hash' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'databags' => null,
        'endpoint' => null,
        'org' => null,
        'chef_user' => null,
        'user_key' => null,
        'org_key' => null,
        'version' => null,
        'chef_use_fqdn' => null,
        'windows_version' => null,
        'windows_install_url' => null,
        'user_key_hash' => null,
        'org_key_hash' => null
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
        'databags' => 'databags',
        'endpoint' => 'endpoint',
        'org' => 'org',
        'chef_user' => 'chefUser',
        'user_key' => 'userKey',
        'org_key' => 'orgKey',
        'version' => 'version',
        'chef_use_fqdn' => 'chefUseFqdn',
        'windows_version' => 'windowsVersion',
        'windows_install_url' => 'windowsInstallUrl',
        'user_key_hash' => 'userKeyHash',
        'org_key_hash' => 'orgKeyHash'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'databags' => 'setDatabags',
        'endpoint' => 'setEndpoint',
        'org' => 'setOrg',
        'chef_user' => 'setChefUser',
        'user_key' => 'setUserKey',
        'org_key' => 'setOrgKey',
        'version' => 'setVersion',
        'chef_use_fqdn' => 'setChefUseFqdn',
        'windows_version' => 'setWindowsVersion',
        'windows_install_url' => 'setWindowsInstallUrl',
        'user_key_hash' => 'setUserKeyHash',
        'org_key_hash' => 'setOrgKeyHash'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'databags' => 'getDatabags',
        'endpoint' => 'getEndpoint',
        'org' => 'getOrg',
        'chef_user' => 'getChefUser',
        'user_key' => 'getUserKey',
        'org_key' => 'getOrgKey',
        'version' => 'getVersion',
        'chef_use_fqdn' => 'getChefUseFqdn',
        'windows_version' => 'getWindowsVersion',
        'windows_install_url' => 'getWindowsInstallUrl',
        'user_key_hash' => 'getUserKeyHash',
        'org_key_hash' => 'getOrgKeyHash'
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
        $this->container['databags'] = $data['databags'] ?? null;
        $this->container['endpoint'] = $data['endpoint'] ?? null;
        $this->container['org'] = $data['org'] ?? null;
        $this->container['chef_user'] = $data['chef_user'] ?? null;
        $this->container['user_key'] = $data['user_key'] ?? null;
        $this->container['org_key'] = $data['org_key'] ?? null;
        $this->container['version'] = $data['version'] ?? null;
        $this->container['chef_use_fqdn'] = $data['chef_use_fqdn'] ?? null;
        $this->container['windows_version'] = $data['windows_version'] ?? null;
        $this->container['windows_install_url'] = $data['windows_install_url'] ?? null;
        $this->container['user_key_hash'] = $data['user_key_hash'] ?? null;
        $this->container['org_key_hash'] = $data['org_key_hash'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

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
     * Gets databags
     *
     * @return \OpenAPI\Client\Model\IntegrationChefConfigDatabags[]|null
     */
    public function getDatabags()
    {
        return $this->container['databags'];
    }

    /**
     * Sets databags
     *
     * @param \OpenAPI\Client\Model\IntegrationChefConfigDatabags[]|null $databags databags
     *
     * @return self
     */
    public function setDatabags($databags)
    {
        $this->container['databags'] = $databags;

        return $this;
    }

    /**
     * Gets endpoint
     *
     * @return string|null
     */
    public function getEndpoint()
    {
        return $this->container['endpoint'];
    }

    /**
     * Sets endpoint
     *
     * @param string|null $endpoint endpoint
     *
     * @return self
     */
    public function setEndpoint($endpoint)
    {
        $this->container['endpoint'] = $endpoint;

        return $this;
    }

    /**
     * Gets org
     *
     * @return string|null
     */
    public function getOrg()
    {
        return $this->container['org'];
    }

    /**
     * Sets org
     *
     * @param string|null $org org
     *
     * @return self
     */
    public function setOrg($org)
    {
        $this->container['org'] = $org;

        return $this;
    }

    /**
     * Gets chef_user
     *
     * @return string|null
     */
    public function getChefUser()
    {
        return $this->container['chef_user'];
    }

    /**
     * Sets chef_user
     *
     * @param string|null $chef_user chef_user
     *
     * @return self
     */
    public function setChefUser($chef_user)
    {
        $this->container['chef_user'] = $chef_user;

        return $this;
    }

    /**
     * Gets user_key
     *
     * @return string|null
     */
    public function getUserKey()
    {
        return $this->container['user_key'];
    }

    /**
     * Sets user_key
     *
     * @param string|null $user_key user_key
     *
     * @return self
     */
    public function setUserKey($user_key)
    {
        $this->container['user_key'] = $user_key;

        return $this;
    }

    /**
     * Gets org_key
     *
     * @return string|null
     */
    public function getOrgKey()
    {
        return $this->container['org_key'];
    }

    /**
     * Sets org_key
     *
     * @param string|null $org_key org_key
     *
     * @return self
     */
    public function setOrgKey($org_key)
    {
        $this->container['org_key'] = $org_key;

        return $this;
    }

    /**
     * Gets version
     *
     * @return string|null
     */
    public function getVersion()
    {
        return $this->container['version'];
    }

    /**
     * Sets version
     *
     * @param string|null $version version
     *
     * @return self
     */
    public function setVersion($version)
    {
        $this->container['version'] = $version;

        return $this;
    }

    /**
     * Gets chef_use_fqdn
     *
     * @return bool|null
     */
    public function getChefUseFqdn()
    {
        return $this->container['chef_use_fqdn'];
    }

    /**
     * Sets chef_use_fqdn
     *
     * @param bool|null $chef_use_fqdn chef_use_fqdn
     *
     * @return self
     */
    public function setChefUseFqdn($chef_use_fqdn)
    {
        $this->container['chef_use_fqdn'] = $chef_use_fqdn;

        return $this;
    }

    /**
     * Gets windows_version
     *
     * @return string|null
     */
    public function getWindowsVersion()
    {
        return $this->container['windows_version'];
    }

    /**
     * Sets windows_version
     *
     * @param string|null $windows_version windows_version
     *
     * @return self
     */
    public function setWindowsVersion($windows_version)
    {
        $this->container['windows_version'] = $windows_version;

        return $this;
    }

    /**
     * Gets windows_install_url
     *
     * @return string|null
     */
    public function getWindowsInstallUrl()
    {
        return $this->container['windows_install_url'];
    }

    /**
     * Sets windows_install_url
     *
     * @param string|null $windows_install_url windows_install_url
     *
     * @return self
     */
    public function setWindowsInstallUrl($windows_install_url)
    {
        $this->container['windows_install_url'] = $windows_install_url;

        return $this;
    }

    /**
     * Gets user_key_hash
     *
     * @return string|null
     */
    public function getUserKeyHash()
    {
        return $this->container['user_key_hash'];
    }

    /**
     * Sets user_key_hash
     *
     * @param string|null $user_key_hash user_key_hash
     *
     * @return self
     */
    public function setUserKeyHash($user_key_hash)
    {
        $this->container['user_key_hash'] = $user_key_hash;

        return $this;
    }

    /**
     * Gets org_key_hash
     *
     * @return string|null
     */
    public function getOrgKeyHash()
    {
        return $this->container['org_key_hash'];
    }

    /**
     * Sets org_key_hash
     *
     * @param string|null $org_key_hash org_key_hash
     *
     * @return self
     */
    public function setOrgKeyHash($org_key_hash)
    {
        $this->container['org_key_hash'] = $org_key_hash;

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


