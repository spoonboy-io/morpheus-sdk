<?php
/**
 * UserSourceCreateActiveDirectory
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
 * UserSourceCreateActiveDirectory Class Doc Comment
 *
 * @category Class
 * @description Active Directory Configuration
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class UserSourceCreateActiveDirectory implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'userSourceCreateActiveDirectory';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'url' => 'string',
        'domain' => 'string',
        'use_ssl' => 'bool',
        'binding_username' => 'string',
        'binding_password' => 'string',
        'required_group' => 'string',
        'search_member_groups' => 'bool'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'url' => null,
        'domain' => null,
        'use_ssl' => null,
        'binding_username' => null,
        'binding_password' => null,
        'required_group' => null,
        'search_member_groups' => null
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
        'url' => 'url',
        'domain' => 'domain',
        'use_ssl' => 'useSSL',
        'binding_username' => 'bindingUsername',
        'binding_password' => 'bindingPassword',
        'required_group' => 'requiredGroup',
        'search_member_groups' => 'searchMemberGroups'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'url' => 'setUrl',
        'domain' => 'setDomain',
        'use_ssl' => 'setUseSsl',
        'binding_username' => 'setBindingUsername',
        'binding_password' => 'setBindingPassword',
        'required_group' => 'setRequiredGroup',
        'search_member_groups' => 'setSearchMemberGroups'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'url' => 'getUrl',
        'domain' => 'getDomain',
        'use_ssl' => 'getUseSsl',
        'binding_username' => 'getBindingUsername',
        'binding_password' => 'getBindingPassword',
        'required_group' => 'getRequiredGroup',
        'search_member_groups' => 'getSearchMemberGroups'
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
        $this->container['url'] = $data['url'] ?? null;
        $this->container['domain'] = $data['domain'] ?? null;
        $this->container['use_ssl'] = $data['use_ssl'] ?? false;
        $this->container['binding_username'] = $data['binding_username'] ?? null;
        $this->container['binding_password'] = $data['binding_password'] ?? null;
        $this->container['required_group'] = $data['required_group'] ?? null;
        $this->container['search_member_groups'] = $data['search_member_groups'] ?? false;
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
     * Gets url
     *
     * @return string|null
     */
    public function getUrl()
    {
        return $this->container['url'];
    }

    /**
     * Sets url
     *
     * @param string|null $url AD Server IP/FQDN
     *
     * @return self
     */
    public function setUrl($url)
    {
        $this->container['url'] = $url;

        return $this;
    }

    /**
     * Gets domain
     *
     * @return string|null
     */
    public function getDomain()
    {
        return $this->container['domain'];
    }

    /**
     * Sets domain
     *
     * @param string|null $domain Domain
     *
     * @return self
     */
    public function setDomain($domain)
    {
        $this->container['domain'] = $domain;

        return $this;
    }

    /**
     * Gets use_ssl
     *
     * @return bool|null
     */
    public function getUseSsl()
    {
        return $this->container['use_ssl'];
    }

    /**
     * Sets use_ssl
     *
     * @param bool|null $use_ssl Use SSL
     *
     * @return self
     */
    public function setUseSsl($use_ssl)
    {
        $this->container['use_ssl'] = $use_ssl;

        return $this;
    }

    /**
     * Gets binding_username
     *
     * @return string|null
     */
    public function getBindingUsername()
    {
        return $this->container['binding_username'];
    }

    /**
     * Sets binding_username
     *
     * @param string|null $binding_username Binding Username
     *
     * @return self
     */
    public function setBindingUsername($binding_username)
    {
        $this->container['binding_username'] = $binding_username;

        return $this;
    }

    /**
     * Gets binding_password
     *
     * @return string|null
     */
    public function getBindingPassword()
    {
        return $this->container['binding_password'];
    }

    /**
     * Sets binding_password
     *
     * @param string|null $binding_password Binding Password
     *
     * @return self
     */
    public function setBindingPassword($binding_password)
    {
        $this->container['binding_password'] = $binding_password;

        return $this;
    }

    /**
     * Gets required_group
     *
     * @return string|null
     */
    public function getRequiredGroup()
    {
        return $this->container['required_group'];
    }

    /**
     * Sets required_group
     *
     * @param string|null $required_group Required Group
     *
     * @return self
     */
    public function setRequiredGroup($required_group)
    {
        $this->container['required_group'] = $required_group;

        return $this;
    }

    /**
     * Gets search_member_groups
     *
     * @return bool|null
     */
    public function getSearchMemberGroups()
    {
        return $this->container['search_member_groups'];
    }

    /**
     * Sets search_member_groups
     *
     * @param bool|null $search_member_groups Include Member Groups
     *
     * @return self
     */
    public function setSearchMemberGroups($search_member_groups)
    {
        $this->container['search_member_groups'] = $search_member_groups;

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

