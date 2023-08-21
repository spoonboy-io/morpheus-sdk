<?php
/**
 * InlineResponse20081LoadBalancerProfile
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
 * InlineResponse20081LoadBalancerProfile Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class InlineResponse20081LoadBalancerProfile implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'inline_response_200_81_loadBalancerProfile';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'load_balancer' => '\OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancer',
        'name' => 'string',
        'category' => 'string',
        'service_type' => 'string',
        'service_type_display' => 'string',
        'visibility' => 'string',
        'description' => 'string',
        'internal_id' => 'string',
        'external_id' => 'string',
        'proxy_type' => 'string',
        'redirect_rewrite' => 'string',
        'persistence_type' => 'string',
        'ssl_enabled' => 'string',
        'ssl_cert' => 'string',
        'account_certificate' => 'string',
        'enabled' => 'bool',
        'redirect_url' => 'string',
        'insert_xforwarded_for' => 'bool',
        'persistence_cookie_name' => 'string',
        'persistence_expires_in' => 'string',
        'editable' => 'bool',
        'config' => 'object',
        'created_by' => 'string',
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
        'load_balancer' => null,
        'name' => null,
        'category' => null,
        'service_type' => null,
        'service_type_display' => null,
        'visibility' => null,
        'description' => null,
        'internal_id' => null,
        'external_id' => null,
        'proxy_type' => null,
        'redirect_rewrite' => null,
        'persistence_type' => null,
        'ssl_enabled' => null,
        'ssl_cert' => null,
        'account_certificate' => null,
        'enabled' => null,
        'redirect_url' => null,
        'insert_xforwarded_for' => null,
        'persistence_cookie_name' => null,
        'persistence_expires_in' => null,
        'editable' => null,
        'config' => null,
        'created_by' => null,
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
        'load_balancer' => 'loadBalancer',
        'name' => 'name',
        'category' => 'category',
        'service_type' => 'serviceType',
        'service_type_display' => 'serviceTypeDisplay',
        'visibility' => 'visibility',
        'description' => 'description',
        'internal_id' => 'internalId',
        'external_id' => 'externalId',
        'proxy_type' => 'proxyType',
        'redirect_rewrite' => 'redirectRewrite',
        'persistence_type' => 'persistenceType',
        'ssl_enabled' => 'sslEnabled',
        'ssl_cert' => 'sslCert',
        'account_certificate' => 'accountCertificate',
        'enabled' => 'enabled',
        'redirect_url' => 'redirectUrl',
        'insert_xforwarded_for' => 'insertXforwardedFor',
        'persistence_cookie_name' => 'persistenceCookieName',
        'persistence_expires_in' => 'persistenceExpiresIn',
        'editable' => 'editable',
        'config' => 'config',
        'created_by' => 'createdBy',
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
        'load_balancer' => 'setLoadBalancer',
        'name' => 'setName',
        'category' => 'setCategory',
        'service_type' => 'setServiceType',
        'service_type_display' => 'setServiceTypeDisplay',
        'visibility' => 'setVisibility',
        'description' => 'setDescription',
        'internal_id' => 'setInternalId',
        'external_id' => 'setExternalId',
        'proxy_type' => 'setProxyType',
        'redirect_rewrite' => 'setRedirectRewrite',
        'persistence_type' => 'setPersistenceType',
        'ssl_enabled' => 'setSslEnabled',
        'ssl_cert' => 'setSslCert',
        'account_certificate' => 'setAccountCertificate',
        'enabled' => 'setEnabled',
        'redirect_url' => 'setRedirectUrl',
        'insert_xforwarded_for' => 'setInsertXforwardedFor',
        'persistence_cookie_name' => 'setPersistenceCookieName',
        'persistence_expires_in' => 'setPersistenceExpiresIn',
        'editable' => 'setEditable',
        'config' => 'setConfig',
        'created_by' => 'setCreatedBy',
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
        'load_balancer' => 'getLoadBalancer',
        'name' => 'getName',
        'category' => 'getCategory',
        'service_type' => 'getServiceType',
        'service_type_display' => 'getServiceTypeDisplay',
        'visibility' => 'getVisibility',
        'description' => 'getDescription',
        'internal_id' => 'getInternalId',
        'external_id' => 'getExternalId',
        'proxy_type' => 'getProxyType',
        'redirect_rewrite' => 'getRedirectRewrite',
        'persistence_type' => 'getPersistenceType',
        'ssl_enabled' => 'getSslEnabled',
        'ssl_cert' => 'getSslCert',
        'account_certificate' => 'getAccountCertificate',
        'enabled' => 'getEnabled',
        'redirect_url' => 'getRedirectUrl',
        'insert_xforwarded_for' => 'getInsertXforwardedFor',
        'persistence_cookie_name' => 'getPersistenceCookieName',
        'persistence_expires_in' => 'getPersistenceExpiresIn',
        'editable' => 'getEditable',
        'config' => 'getConfig',
        'created_by' => 'getCreatedBy',
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
        $this->container['load_balancer'] = $data['load_balancer'] ?? null;
        $this->container['name'] = $data['name'] ?? null;
        $this->container['category'] = $data['category'] ?? null;
        $this->container['service_type'] = $data['service_type'] ?? null;
        $this->container['service_type_display'] = $data['service_type_display'] ?? null;
        $this->container['visibility'] = $data['visibility'] ?? null;
        $this->container['description'] = $data['description'] ?? null;
        $this->container['internal_id'] = $data['internal_id'] ?? null;
        $this->container['external_id'] = $data['external_id'] ?? null;
        $this->container['proxy_type'] = $data['proxy_type'] ?? null;
        $this->container['redirect_rewrite'] = $data['redirect_rewrite'] ?? null;
        $this->container['persistence_type'] = $data['persistence_type'] ?? null;
        $this->container['ssl_enabled'] = $data['ssl_enabled'] ?? null;
        $this->container['ssl_cert'] = $data['ssl_cert'] ?? null;
        $this->container['account_certificate'] = $data['account_certificate'] ?? null;
        $this->container['enabled'] = $data['enabled'] ?? null;
        $this->container['redirect_url'] = $data['redirect_url'] ?? null;
        $this->container['insert_xforwarded_for'] = $data['insert_xforwarded_for'] ?? null;
        $this->container['persistence_cookie_name'] = $data['persistence_cookie_name'] ?? null;
        $this->container['persistence_expires_in'] = $data['persistence_expires_in'] ?? null;
        $this->container['editable'] = $data['editable'] ?? null;
        $this->container['config'] = $data['config'] ?? null;
        $this->container['created_by'] = $data['created_by'] ?? null;
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
     * Gets load_balancer
     *
     * @return \OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancer|null
     */
    public function getLoadBalancer()
    {
        return $this->container['load_balancer'];
    }

    /**
     * Sets load_balancer
     *
     * @param \OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancer|null $load_balancer load_balancer
     *
     * @return self
     */
    public function setLoadBalancer($load_balancer)
    {
        $this->container['load_balancer'] = $load_balancer;

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
     * Gets category
     *
     * @return string|null
     */
    public function getCategory()
    {
        return $this->container['category'];
    }

    /**
     * Sets category
     *
     * @param string|null $category category
     *
     * @return self
     */
    public function setCategory($category)
    {
        $this->container['category'] = $category;

        return $this;
    }

    /**
     * Gets service_type
     *
     * @return string|null
     */
    public function getServiceType()
    {
        return $this->container['service_type'];
    }

    /**
     * Sets service_type
     *
     * @param string|null $service_type service_type
     *
     * @return self
     */
    public function setServiceType($service_type)
    {
        $this->container['service_type'] = $service_type;

        return $this;
    }

    /**
     * Gets service_type_display
     *
     * @return string|null
     */
    public function getServiceTypeDisplay()
    {
        return $this->container['service_type_display'];
    }

    /**
     * Sets service_type_display
     *
     * @param string|null $service_type_display service_type_display
     *
     * @return self
     */
    public function setServiceTypeDisplay($service_type_display)
    {
        $this->container['service_type_display'] = $service_type_display;

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
     * @param string|null $visibility visibility
     *
     * @return self
     */
    public function setVisibility($visibility)
    {
        $this->container['visibility'] = $visibility;

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
     * Gets internal_id
     *
     * @return string|null
     */
    public function getInternalId()
    {
        return $this->container['internal_id'];
    }

    /**
     * Sets internal_id
     *
     * @param string|null $internal_id internal_id
     *
     * @return self
     */
    public function setInternalId($internal_id)
    {
        $this->container['internal_id'] = $internal_id;

        return $this;
    }

    /**
     * Gets external_id
     *
     * @return string|null
     */
    public function getExternalId()
    {
        return $this->container['external_id'];
    }

    /**
     * Sets external_id
     *
     * @param string|null $external_id external_id
     *
     * @return self
     */
    public function setExternalId($external_id)
    {
        $this->container['external_id'] = $external_id;

        return $this;
    }

    /**
     * Gets proxy_type
     *
     * @return string|null
     */
    public function getProxyType()
    {
        return $this->container['proxy_type'];
    }

    /**
     * Sets proxy_type
     *
     * @param string|null $proxy_type proxy_type
     *
     * @return self
     */
    public function setProxyType($proxy_type)
    {
        $this->container['proxy_type'] = $proxy_type;

        return $this;
    }

    /**
     * Gets redirect_rewrite
     *
     * @return string|null
     */
    public function getRedirectRewrite()
    {
        return $this->container['redirect_rewrite'];
    }

    /**
     * Sets redirect_rewrite
     *
     * @param string|null $redirect_rewrite redirect_rewrite
     *
     * @return self
     */
    public function setRedirectRewrite($redirect_rewrite)
    {
        $this->container['redirect_rewrite'] = $redirect_rewrite;

        return $this;
    }

    /**
     * Gets persistence_type
     *
     * @return string|null
     */
    public function getPersistenceType()
    {
        return $this->container['persistence_type'];
    }

    /**
     * Sets persistence_type
     *
     * @param string|null $persistence_type persistence_type
     *
     * @return self
     */
    public function setPersistenceType($persistence_type)
    {
        $this->container['persistence_type'] = $persistence_type;

        return $this;
    }

    /**
     * Gets ssl_enabled
     *
     * @return string|null
     */
    public function getSslEnabled()
    {
        return $this->container['ssl_enabled'];
    }

    /**
     * Sets ssl_enabled
     *
     * @param string|null $ssl_enabled ssl_enabled
     *
     * @return self
     */
    public function setSslEnabled($ssl_enabled)
    {
        $this->container['ssl_enabled'] = $ssl_enabled;

        return $this;
    }

    /**
     * Gets ssl_cert
     *
     * @return string|null
     */
    public function getSslCert()
    {
        return $this->container['ssl_cert'];
    }

    /**
     * Sets ssl_cert
     *
     * @param string|null $ssl_cert ssl_cert
     *
     * @return self
     */
    public function setSslCert($ssl_cert)
    {
        $this->container['ssl_cert'] = $ssl_cert;

        return $this;
    }

    /**
     * Gets account_certificate
     *
     * @return string|null
     */
    public function getAccountCertificate()
    {
        return $this->container['account_certificate'];
    }

    /**
     * Sets account_certificate
     *
     * @param string|null $account_certificate account_certificate
     *
     * @return self
     */
    public function setAccountCertificate($account_certificate)
    {
        $this->container['account_certificate'] = $account_certificate;

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
     * @param bool|null $enabled enabled
     *
     * @return self
     */
    public function setEnabled($enabled)
    {
        $this->container['enabled'] = $enabled;

        return $this;
    }

    /**
     * Gets redirect_url
     *
     * @return string|null
     */
    public function getRedirectUrl()
    {
        return $this->container['redirect_url'];
    }

    /**
     * Sets redirect_url
     *
     * @param string|null $redirect_url redirect_url
     *
     * @return self
     */
    public function setRedirectUrl($redirect_url)
    {
        $this->container['redirect_url'] = $redirect_url;

        return $this;
    }

    /**
     * Gets insert_xforwarded_for
     *
     * @return bool|null
     */
    public function getInsertXforwardedFor()
    {
        return $this->container['insert_xforwarded_for'];
    }

    /**
     * Sets insert_xforwarded_for
     *
     * @param bool|null $insert_xforwarded_for insert_xforwarded_for
     *
     * @return self
     */
    public function setInsertXforwardedFor($insert_xforwarded_for)
    {
        $this->container['insert_xforwarded_for'] = $insert_xforwarded_for;

        return $this;
    }

    /**
     * Gets persistence_cookie_name
     *
     * @return string|null
     */
    public function getPersistenceCookieName()
    {
        return $this->container['persistence_cookie_name'];
    }

    /**
     * Sets persistence_cookie_name
     *
     * @param string|null $persistence_cookie_name persistence_cookie_name
     *
     * @return self
     */
    public function setPersistenceCookieName($persistence_cookie_name)
    {
        $this->container['persistence_cookie_name'] = $persistence_cookie_name;

        return $this;
    }

    /**
     * Gets persistence_expires_in
     *
     * @return string|null
     */
    public function getPersistenceExpiresIn()
    {
        return $this->container['persistence_expires_in'];
    }

    /**
     * Sets persistence_expires_in
     *
     * @param string|null $persistence_expires_in persistence_expires_in
     *
     * @return self
     */
    public function setPersistenceExpiresIn($persistence_expires_in)
    {
        $this->container['persistence_expires_in'] = $persistence_expires_in;

        return $this;
    }

    /**
     * Gets editable
     *
     * @return bool|null
     */
    public function getEditable()
    {
        return $this->container['editable'];
    }

    /**
     * Sets editable
     *
     * @param bool|null $editable editable
     *
     * @return self
     */
    public function setEditable($editable)
    {
        $this->container['editable'] = $editable;

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
     * @param object|null $config config
     *
     * @return self
     */
    public function setConfig($config)
    {
        $this->container['config'] = $config;

        return $this;
    }

    /**
     * Gets created_by
     *
     * @return string|null
     */
    public function getCreatedBy()
    {
        return $this->container['created_by'];
    }

    /**
     * Sets created_by
     *
     * @param string|null $created_by created_by
     *
     * @return self
     */
    public function setCreatedBy($created_by)
    {
        $this->container['created_by'] = $created_by;

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


