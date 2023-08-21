<?php
/**
 * IdentitySourcesOneLoginConfig
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
 * IdentitySourcesOneLoginConfig Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class IdentitySourcesOneLoginConfig implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'identitySourcesOneLoginConfig';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'name' => 'string',
        'description' => 'string',
        'code' => 'string',
        'type' => 'string',
        'active' => 'bool',
        'deleted' => 'bool',
        'auto_sync_on_login' => 'bool',
        'external_login' => 'bool',
        'allow_custom_mappings' => 'bool',
        'manual_role_assignment' => 'bool',
        'account' => '\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance',
        'default_account_role' => '\OpenAPI\Client\Model\IdentitySourcesLDAPConfigDefaultAccountRole',
        'config' => '\OpenAPI\Client\Model\IdentitySourcesOneLoginConfigConfig',
        'role_mappings' => '\OpenAPI\Client\Model\IdentitySourcesJumpCloudConfigRoleMappings[]',
        'subdomain' => 'string',
        'login_url' => 'string',
        'provider_settings' => 'object',
        'date_created' => '\DateTime',
        'last_updated' => '\DateTime'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'id' => 'int64',
        'name' => null,
        'description' => null,
        'code' => null,
        'type' => null,
        'active' => null,
        'deleted' => null,
        'auto_sync_on_login' => null,
        'external_login' => null,
        'allow_custom_mappings' => null,
        'manual_role_assignment' => null,
        'account' => null,
        'default_account_role' => null,
        'config' => null,
        'role_mappings' => null,
        'subdomain' => null,
        'login_url' => null,
        'provider_settings' => null,
        'date_created' => 'date-time',
        'last_updated' => 'date-time'
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
        'id' => 'id',
        'name' => 'name',
        'description' => 'description',
        'code' => 'code',
        'type' => 'type',
        'active' => 'active',
        'deleted' => 'deleted',
        'auto_sync_on_login' => 'autoSyncOnLogin',
        'external_login' => 'externalLogin',
        'allow_custom_mappings' => 'allowCustomMappings',
        'manual_role_assignment' => 'manualRoleAssignment',
        'account' => 'account',
        'default_account_role' => 'defaultAccountRole',
        'config' => 'config',
        'role_mappings' => 'roleMappings',
        'subdomain' => 'subdomain',
        'login_url' => 'loginURL',
        'provider_settings' => 'providerSettings',
        'date_created' => 'dateCreated',
        'last_updated' => 'lastUpdated'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'name' => 'setName',
        'description' => 'setDescription',
        'code' => 'setCode',
        'type' => 'setType',
        'active' => 'setActive',
        'deleted' => 'setDeleted',
        'auto_sync_on_login' => 'setAutoSyncOnLogin',
        'external_login' => 'setExternalLogin',
        'allow_custom_mappings' => 'setAllowCustomMappings',
        'manual_role_assignment' => 'setManualRoleAssignment',
        'account' => 'setAccount',
        'default_account_role' => 'setDefaultAccountRole',
        'config' => 'setConfig',
        'role_mappings' => 'setRoleMappings',
        'subdomain' => 'setSubdomain',
        'login_url' => 'setLoginUrl',
        'provider_settings' => 'setProviderSettings',
        'date_created' => 'setDateCreated',
        'last_updated' => 'setLastUpdated'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'name' => 'getName',
        'description' => 'getDescription',
        'code' => 'getCode',
        'type' => 'getType',
        'active' => 'getActive',
        'deleted' => 'getDeleted',
        'auto_sync_on_login' => 'getAutoSyncOnLogin',
        'external_login' => 'getExternalLogin',
        'allow_custom_mappings' => 'getAllowCustomMappings',
        'manual_role_assignment' => 'getManualRoleAssignment',
        'account' => 'getAccount',
        'default_account_role' => 'getDefaultAccountRole',
        'config' => 'getConfig',
        'role_mappings' => 'getRoleMappings',
        'subdomain' => 'getSubdomain',
        'login_url' => 'getLoginUrl',
        'provider_settings' => 'getProviderSettings',
        'date_created' => 'getDateCreated',
        'last_updated' => 'getLastUpdated'
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
        $this->container['id'] = $data['id'] ?? null;
        $this->container['name'] = $data['name'] ?? null;
        $this->container['description'] = $data['description'] ?? null;
        $this->container['code'] = $data['code'] ?? null;
        $this->container['type'] = $data['type'] ?? null;
        $this->container['active'] = $data['active'] ?? null;
        $this->container['deleted'] = $data['deleted'] ?? null;
        $this->container['auto_sync_on_login'] = $data['auto_sync_on_login'] ?? null;
        $this->container['external_login'] = $data['external_login'] ?? null;
        $this->container['allow_custom_mappings'] = $data['allow_custom_mappings'] ?? null;
        $this->container['manual_role_assignment'] = $data['manual_role_assignment'] ?? null;
        $this->container['account'] = $data['account'] ?? null;
        $this->container['default_account_role'] = $data['default_account_role'] ?? null;
        $this->container['config'] = $data['config'] ?? null;
        $this->container['role_mappings'] = $data['role_mappings'] ?? null;
        $this->container['subdomain'] = $data['subdomain'] ?? null;
        $this->container['login_url'] = $data['login_url'] ?? null;
        $this->container['provider_settings'] = $data['provider_settings'] ?? null;
        $this->container['date_created'] = $data['date_created'] ?? null;
        $this->container['last_updated'] = $data['last_updated'] ?? null;
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
     * Gets id
     *
     * @return int|null
     */
    public function getId()
    {
        return $this->container['id'];
    }

    /**
     * Sets id
     *
     * @param int|null $id id
     *
     * @return self
     */
    public function setId($id)
    {
        $this->container['id'] = $id;

        return $this;
    }

    /**
     * Gets name
     *
     * @return string|null
     */
    public function getName()
    {
        return $this->container['name'];
    }

    /**
     * Sets name
     *
     * @param string|null $name name
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
     * @param string|null $description description
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
     * @param string|null $code code
     *
     * @return self
     */
    public function setCode($code)
    {
        $this->container['code'] = $code;

        return $this;
    }

    /**
     * Gets type
     *
     * @return string|null
     */
    public function getType()
    {
        return $this->container['type'];
    }

    /**
     * Sets type
     *
     * @param string|null $type type
     *
     * @return self
     */
    public function setType($type)
    {
        $this->container['type'] = $type;

        return $this;
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
     * @param bool|null $active active
     *
     * @return self
     */
    public function setActive($active)
    {
        $this->container['active'] = $active;

        return $this;
    }

    /**
     * Gets deleted
     *
     * @return bool|null
     */
    public function getDeleted()
    {
        return $this->container['deleted'];
    }

    /**
     * Sets deleted
     *
     * @param bool|null $deleted deleted
     *
     * @return self
     */
    public function setDeleted($deleted)
    {
        $this->container['deleted'] = $deleted;

        return $this;
    }

    /**
     * Gets auto_sync_on_login
     *
     * @return bool|null
     */
    public function getAutoSyncOnLogin()
    {
        return $this->container['auto_sync_on_login'];
    }

    /**
     * Sets auto_sync_on_login
     *
     * @param bool|null $auto_sync_on_login auto_sync_on_login
     *
     * @return self
     */
    public function setAutoSyncOnLogin($auto_sync_on_login)
    {
        $this->container['auto_sync_on_login'] = $auto_sync_on_login;

        return $this;
    }

    /**
     * Gets external_login
     *
     * @return bool|null
     */
    public function getExternalLogin()
    {
        return $this->container['external_login'];
    }

    /**
     * Sets external_login
     *
     * @param bool|null $external_login external_login
     *
     * @return self
     */
    public function setExternalLogin($external_login)
    {
        $this->container['external_login'] = $external_login;

        return $this;
    }

    /**
     * Gets allow_custom_mappings
     *
     * @return bool|null
     */
    public function getAllowCustomMappings()
    {
        return $this->container['allow_custom_mappings'];
    }

    /**
     * Sets allow_custom_mappings
     *
     * @param bool|null $allow_custom_mappings allow_custom_mappings
     *
     * @return self
     */
    public function setAllowCustomMappings($allow_custom_mappings)
    {
        $this->container['allow_custom_mappings'] = $allow_custom_mappings;

        return $this;
    }

    /**
     * Gets manual_role_assignment
     *
     * @return bool|null
     */
    public function getManualRoleAssignment()
    {
        return $this->container['manual_role_assignment'];
    }

    /**
     * Sets manual_role_assignment
     *
     * @param bool|null $manual_role_assignment manual_role_assignment
     *
     * @return self
     */
    public function setManualRoleAssignment($manual_role_assignment)
    {
        $this->container['manual_role_assignment'] = $manual_role_assignment;

        return $this;
    }

    /**
     * Gets account
     *
     * @return \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null
     */
    public function getAccount()
    {
        return $this->container['account'];
    }

    /**
     * Sets account
     *
     * @param \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null $account account
     *
     * @return self
     */
    public function setAccount($account)
    {
        $this->container['account'] = $account;

        return $this;
    }

    /**
     * Gets default_account_role
     *
     * @return \OpenAPI\Client\Model\IdentitySourcesLDAPConfigDefaultAccountRole|null
     */
    public function getDefaultAccountRole()
    {
        return $this->container['default_account_role'];
    }

    /**
     * Sets default_account_role
     *
     * @param \OpenAPI\Client\Model\IdentitySourcesLDAPConfigDefaultAccountRole|null $default_account_role default_account_role
     *
     * @return self
     */
    public function setDefaultAccountRole($default_account_role)
    {
        $this->container['default_account_role'] = $default_account_role;

        return $this;
    }

    /**
     * Gets config
     *
     * @return \OpenAPI\Client\Model\IdentitySourcesOneLoginConfigConfig|null
     */
    public function getConfig()
    {
        return $this->container['config'];
    }

    /**
     * Sets config
     *
     * @param \OpenAPI\Client\Model\IdentitySourcesOneLoginConfigConfig|null $config config
     *
     * @return self
     */
    public function setConfig($config)
    {
        $this->container['config'] = $config;

        return $this;
    }

    /**
     * Gets role_mappings
     *
     * @return \OpenAPI\Client\Model\IdentitySourcesJumpCloudConfigRoleMappings[]|null
     */
    public function getRoleMappings()
    {
        return $this->container['role_mappings'];
    }

    /**
     * Sets role_mappings
     *
     * @param \OpenAPI\Client\Model\IdentitySourcesJumpCloudConfigRoleMappings[]|null $role_mappings role_mappings
     *
     * @return self
     */
    public function setRoleMappings($role_mappings)
    {
        $this->container['role_mappings'] = $role_mappings;

        return $this;
    }

    /**
     * Gets subdomain
     *
     * @return string|null
     */
    public function getSubdomain()
    {
        return $this->container['subdomain'];
    }

    /**
     * Sets subdomain
     *
     * @param string|null $subdomain subdomain
     *
     * @return self
     */
    public function setSubdomain($subdomain)
    {
        $this->container['subdomain'] = $subdomain;

        return $this;
    }

    /**
     * Gets login_url
     *
     * @return string|null
     */
    public function getLoginUrl()
    {
        return $this->container['login_url'];
    }

    /**
     * Sets login_url
     *
     * @param string|null $login_url login_url
     *
     * @return self
     */
    public function setLoginUrl($login_url)
    {
        $this->container['login_url'] = $login_url;

        return $this;
    }

    /**
     * Gets provider_settings
     *
     * @return object|null
     */
    public function getProviderSettings()
    {
        return $this->container['provider_settings'];
    }

    /**
     * Sets provider_settings
     *
     * @param object|null $provider_settings provider_settings
     *
     * @return self
     */
    public function setProviderSettings($provider_settings)
    {
        $this->container['provider_settings'] = $provider_settings;

        return $this;
    }

    /**
     * Gets date_created
     *
     * @return \DateTime|null
     */
    public function getDateCreated()
    {
        return $this->container['date_created'];
    }

    /**
     * Sets date_created
     *
     * @param \DateTime|null $date_created date_created
     *
     * @return self
     */
    public function setDateCreated($date_created)
    {
        $this->container['date_created'] = $date_created;

        return $this;
    }

    /**
     * Gets last_updated
     *
     * @return \DateTime|null
     */
    public function getLastUpdated()
    {
        return $this->container['last_updated'];
    }

    /**
     * Sets last_updated
     *
     * @param \DateTime|null $last_updated last_updated
     *
     * @return self
     */
    public function setLastUpdated($last_updated)
    {
        $this->container['last_updated'] = $last_updated;

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


