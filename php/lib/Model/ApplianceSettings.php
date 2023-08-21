<?php
/**
 * ApplianceSettings
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
 * ApplianceSettings Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ApplianceSettings implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'applianceSettings';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'appliance_url' => 'string',
        'internal_appliance_url' => 'string',
        'cors_allowed' => 'string',
        'registration_enabled' => 'bool',
        'default_role_id' => 'string',
        'default_user_role_id' => 'string',
        'docker_privileged_mode' => 'bool',
        'expire_pwd_days' => 'string',
        'disable_after_attempts' => 'string',
        'disable_after_days_inactive' => 'string',
        'warn_user_days_before' => 'string',
        'smtp_mail_from' => 'string',
        'smtp_server' => 'string',
        'smtp_port' => 'string',
        'smtp_ssl' => 'bool',
        'smtp_tls' => 'bool',
        'smtp_user' => 'string',
        'smtp_password' => 'string',
        'smtp_password_hash' => 'string',
        'proxy_host' => 'string',
        'proxy_port' => 'string',
        'proxy_user' => 'string',
        'proxy_password' => 'string',
        'proxy_password_hash' => 'string',
        'proxy_domain' => 'string',
        'proxy_workstation' => 'string',
        'currency_provider' => 'string',
        'currency_key' => 'string',
        'enabled_zone_types' => '\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance[]',
        'stats_retainment_period' => 'int'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'appliance_url' => null,
        'internal_appliance_url' => null,
        'cors_allowed' => null,
        'registration_enabled' => null,
        'default_role_id' => null,
        'default_user_role_id' => null,
        'docker_privileged_mode' => null,
        'expire_pwd_days' => null,
        'disable_after_attempts' => null,
        'disable_after_days_inactive' => null,
        'warn_user_days_before' => null,
        'smtp_mail_from' => null,
        'smtp_server' => null,
        'smtp_port' => null,
        'smtp_ssl' => null,
        'smtp_tls' => null,
        'smtp_user' => null,
        'smtp_password' => null,
        'smtp_password_hash' => null,
        'proxy_host' => null,
        'proxy_port' => null,
        'proxy_user' => null,
        'proxy_password' => null,
        'proxy_password_hash' => null,
        'proxy_domain' => null,
        'proxy_workstation' => null,
        'currency_provider' => null,
        'currency_key' => null,
        'enabled_zone_types' => null,
        'stats_retainment_period' => 'int64'
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
        'appliance_url' => 'applianceUrl',
        'internal_appliance_url' => 'internalApplianceUrl',
        'cors_allowed' => 'corsAllowed',
        'registration_enabled' => 'registrationEnabled',
        'default_role_id' => 'defaultRoleId',
        'default_user_role_id' => 'defaultUserRoleId',
        'docker_privileged_mode' => 'dockerPrivilegedMode',
        'expire_pwd_days' => 'expirePwdDays',
        'disable_after_attempts' => 'disableAfterAttempts',
        'disable_after_days_inactive' => 'disableAfterDaysInactive',
        'warn_user_days_before' => 'warnUserDaysBefore',
        'smtp_mail_from' => 'smtpMailFrom',
        'smtp_server' => 'smtpServer',
        'smtp_port' => 'smtpPort',
        'smtp_ssl' => 'smtpSSL',
        'smtp_tls' => 'smtpTLS',
        'smtp_user' => 'smtpUser',
        'smtp_password' => 'smtpPassword',
        'smtp_password_hash' => 'smtpPasswordHash',
        'proxy_host' => 'proxyHost',
        'proxy_port' => 'proxyPort',
        'proxy_user' => 'proxyUser',
        'proxy_password' => 'proxyPassword',
        'proxy_password_hash' => 'proxyPasswordHash',
        'proxy_domain' => 'proxyDomain',
        'proxy_workstation' => 'proxyWorkstation',
        'currency_provider' => 'currencyProvider',
        'currency_key' => 'currencyKey',
        'enabled_zone_types' => 'enabledZoneTypes',
        'stats_retainment_period' => 'statsRetainmentPeriod'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'appliance_url' => 'setApplianceUrl',
        'internal_appliance_url' => 'setInternalApplianceUrl',
        'cors_allowed' => 'setCorsAllowed',
        'registration_enabled' => 'setRegistrationEnabled',
        'default_role_id' => 'setDefaultRoleId',
        'default_user_role_id' => 'setDefaultUserRoleId',
        'docker_privileged_mode' => 'setDockerPrivilegedMode',
        'expire_pwd_days' => 'setExpirePwdDays',
        'disable_after_attempts' => 'setDisableAfterAttempts',
        'disable_after_days_inactive' => 'setDisableAfterDaysInactive',
        'warn_user_days_before' => 'setWarnUserDaysBefore',
        'smtp_mail_from' => 'setSmtpMailFrom',
        'smtp_server' => 'setSmtpServer',
        'smtp_port' => 'setSmtpPort',
        'smtp_ssl' => 'setSmtpSsl',
        'smtp_tls' => 'setSmtpTls',
        'smtp_user' => 'setSmtpUser',
        'smtp_password' => 'setSmtpPassword',
        'smtp_password_hash' => 'setSmtpPasswordHash',
        'proxy_host' => 'setProxyHost',
        'proxy_port' => 'setProxyPort',
        'proxy_user' => 'setProxyUser',
        'proxy_password' => 'setProxyPassword',
        'proxy_password_hash' => 'setProxyPasswordHash',
        'proxy_domain' => 'setProxyDomain',
        'proxy_workstation' => 'setProxyWorkstation',
        'currency_provider' => 'setCurrencyProvider',
        'currency_key' => 'setCurrencyKey',
        'enabled_zone_types' => 'setEnabledZoneTypes',
        'stats_retainment_period' => 'setStatsRetainmentPeriod'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'appliance_url' => 'getApplianceUrl',
        'internal_appliance_url' => 'getInternalApplianceUrl',
        'cors_allowed' => 'getCorsAllowed',
        'registration_enabled' => 'getRegistrationEnabled',
        'default_role_id' => 'getDefaultRoleId',
        'default_user_role_id' => 'getDefaultUserRoleId',
        'docker_privileged_mode' => 'getDockerPrivilegedMode',
        'expire_pwd_days' => 'getExpirePwdDays',
        'disable_after_attempts' => 'getDisableAfterAttempts',
        'disable_after_days_inactive' => 'getDisableAfterDaysInactive',
        'warn_user_days_before' => 'getWarnUserDaysBefore',
        'smtp_mail_from' => 'getSmtpMailFrom',
        'smtp_server' => 'getSmtpServer',
        'smtp_port' => 'getSmtpPort',
        'smtp_ssl' => 'getSmtpSsl',
        'smtp_tls' => 'getSmtpTls',
        'smtp_user' => 'getSmtpUser',
        'smtp_password' => 'getSmtpPassword',
        'smtp_password_hash' => 'getSmtpPasswordHash',
        'proxy_host' => 'getProxyHost',
        'proxy_port' => 'getProxyPort',
        'proxy_user' => 'getProxyUser',
        'proxy_password' => 'getProxyPassword',
        'proxy_password_hash' => 'getProxyPasswordHash',
        'proxy_domain' => 'getProxyDomain',
        'proxy_workstation' => 'getProxyWorkstation',
        'currency_provider' => 'getCurrencyProvider',
        'currency_key' => 'getCurrencyKey',
        'enabled_zone_types' => 'getEnabledZoneTypes',
        'stats_retainment_period' => 'getStatsRetainmentPeriod'
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
        $this->container['appliance_url'] = $data['appliance_url'] ?? null;
        $this->container['internal_appliance_url'] = $data['internal_appliance_url'] ?? null;
        $this->container['cors_allowed'] = $data['cors_allowed'] ?? null;
        $this->container['registration_enabled'] = $data['registration_enabled'] ?? null;
        $this->container['default_role_id'] = $data['default_role_id'] ?? null;
        $this->container['default_user_role_id'] = $data['default_user_role_id'] ?? null;
        $this->container['docker_privileged_mode'] = $data['docker_privileged_mode'] ?? null;
        $this->container['expire_pwd_days'] = $data['expire_pwd_days'] ?? null;
        $this->container['disable_after_attempts'] = $data['disable_after_attempts'] ?? null;
        $this->container['disable_after_days_inactive'] = $data['disable_after_days_inactive'] ?? null;
        $this->container['warn_user_days_before'] = $data['warn_user_days_before'] ?? null;
        $this->container['smtp_mail_from'] = $data['smtp_mail_from'] ?? null;
        $this->container['smtp_server'] = $data['smtp_server'] ?? null;
        $this->container['smtp_port'] = $data['smtp_port'] ?? null;
        $this->container['smtp_ssl'] = $data['smtp_ssl'] ?? null;
        $this->container['smtp_tls'] = $data['smtp_tls'] ?? null;
        $this->container['smtp_user'] = $data['smtp_user'] ?? null;
        $this->container['smtp_password'] = $data['smtp_password'] ?? null;
        $this->container['smtp_password_hash'] = $data['smtp_password_hash'] ?? null;
        $this->container['proxy_host'] = $data['proxy_host'] ?? null;
        $this->container['proxy_port'] = $data['proxy_port'] ?? null;
        $this->container['proxy_user'] = $data['proxy_user'] ?? null;
        $this->container['proxy_password'] = $data['proxy_password'] ?? null;
        $this->container['proxy_password_hash'] = $data['proxy_password_hash'] ?? null;
        $this->container['proxy_domain'] = $data['proxy_domain'] ?? null;
        $this->container['proxy_workstation'] = $data['proxy_workstation'] ?? null;
        $this->container['currency_provider'] = $data['currency_provider'] ?? null;
        $this->container['currency_key'] = $data['currency_key'] ?? null;
        $this->container['enabled_zone_types'] = $data['enabled_zone_types'] ?? null;
        $this->container['stats_retainment_period'] = $data['stats_retainment_period'] ?? null;
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
     * Gets appliance_url
     *
     * @return string|null
     */
    public function getApplianceUrl()
    {
        return $this->container['appliance_url'];
    }

    /**
     * Sets appliance_url
     *
     * @param string|null $appliance_url appliance_url
     *
     * @return self
     */
    public function setApplianceUrl($appliance_url)
    {
        $this->container['appliance_url'] = $appliance_url;

        return $this;
    }

    /**
     * Gets internal_appliance_url
     *
     * @return string|null
     */
    public function getInternalApplianceUrl()
    {
        return $this->container['internal_appliance_url'];
    }

    /**
     * Sets internal_appliance_url
     *
     * @param string|null $internal_appliance_url internal_appliance_url
     *
     * @return self
     */
    public function setInternalApplianceUrl($internal_appliance_url)
    {
        $this->container['internal_appliance_url'] = $internal_appliance_url;

        return $this;
    }

    /**
     * Gets cors_allowed
     *
     * @return string|null
     */
    public function getCorsAllowed()
    {
        return $this->container['cors_allowed'];
    }

    /**
     * Sets cors_allowed
     *
     * @param string|null $cors_allowed cors_allowed
     *
     * @return self
     */
    public function setCorsAllowed($cors_allowed)
    {
        $this->container['cors_allowed'] = $cors_allowed;

        return $this;
    }

    /**
     * Gets registration_enabled
     *
     * @return bool|null
     */
    public function getRegistrationEnabled()
    {
        return $this->container['registration_enabled'];
    }

    /**
     * Sets registration_enabled
     *
     * @param bool|null $registration_enabled registration_enabled
     *
     * @return self
     */
    public function setRegistrationEnabled($registration_enabled)
    {
        $this->container['registration_enabled'] = $registration_enabled;

        return $this;
    }

    /**
     * Gets default_role_id
     *
     * @return string|null
     */
    public function getDefaultRoleId()
    {
        return $this->container['default_role_id'];
    }

    /**
     * Sets default_role_id
     *
     * @param string|null $default_role_id default_role_id
     *
     * @return self
     */
    public function setDefaultRoleId($default_role_id)
    {
        $this->container['default_role_id'] = $default_role_id;

        return $this;
    }

    /**
     * Gets default_user_role_id
     *
     * @return string|null
     */
    public function getDefaultUserRoleId()
    {
        return $this->container['default_user_role_id'];
    }

    /**
     * Sets default_user_role_id
     *
     * @param string|null $default_user_role_id default_user_role_id
     *
     * @return self
     */
    public function setDefaultUserRoleId($default_user_role_id)
    {
        $this->container['default_user_role_id'] = $default_user_role_id;

        return $this;
    }

    /**
     * Gets docker_privileged_mode
     *
     * @return bool|null
     */
    public function getDockerPrivilegedMode()
    {
        return $this->container['docker_privileged_mode'];
    }

    /**
     * Sets docker_privileged_mode
     *
     * @param bool|null $docker_privileged_mode docker_privileged_mode
     *
     * @return self
     */
    public function setDockerPrivilegedMode($docker_privileged_mode)
    {
        $this->container['docker_privileged_mode'] = $docker_privileged_mode;

        return $this;
    }

    /**
     * Gets expire_pwd_days
     *
     * @return string|null
     */
    public function getExpirePwdDays()
    {
        return $this->container['expire_pwd_days'];
    }

    /**
     * Sets expire_pwd_days
     *
     * @param string|null $expire_pwd_days expire_pwd_days
     *
     * @return self
     */
    public function setExpirePwdDays($expire_pwd_days)
    {
        $this->container['expire_pwd_days'] = $expire_pwd_days;

        return $this;
    }

    /**
     * Gets disable_after_attempts
     *
     * @return string|null
     */
    public function getDisableAfterAttempts()
    {
        return $this->container['disable_after_attempts'];
    }

    /**
     * Sets disable_after_attempts
     *
     * @param string|null $disable_after_attempts disable_after_attempts
     *
     * @return self
     */
    public function setDisableAfterAttempts($disable_after_attempts)
    {
        $this->container['disable_after_attempts'] = $disable_after_attempts;

        return $this;
    }

    /**
     * Gets disable_after_days_inactive
     *
     * @return string|null
     */
    public function getDisableAfterDaysInactive()
    {
        return $this->container['disable_after_days_inactive'];
    }

    /**
     * Sets disable_after_days_inactive
     *
     * @param string|null $disable_after_days_inactive disable_after_days_inactive
     *
     * @return self
     */
    public function setDisableAfterDaysInactive($disable_after_days_inactive)
    {
        $this->container['disable_after_days_inactive'] = $disable_after_days_inactive;

        return $this;
    }

    /**
     * Gets warn_user_days_before
     *
     * @return string|null
     */
    public function getWarnUserDaysBefore()
    {
        return $this->container['warn_user_days_before'];
    }

    /**
     * Sets warn_user_days_before
     *
     * @param string|null $warn_user_days_before warn_user_days_before
     *
     * @return self
     */
    public function setWarnUserDaysBefore($warn_user_days_before)
    {
        $this->container['warn_user_days_before'] = $warn_user_days_before;

        return $this;
    }

    /**
     * Gets smtp_mail_from
     *
     * @return string|null
     */
    public function getSmtpMailFrom()
    {
        return $this->container['smtp_mail_from'];
    }

    /**
     * Sets smtp_mail_from
     *
     * @param string|null $smtp_mail_from smtp_mail_from
     *
     * @return self
     */
    public function setSmtpMailFrom($smtp_mail_from)
    {
        $this->container['smtp_mail_from'] = $smtp_mail_from;

        return $this;
    }

    /**
     * Gets smtp_server
     *
     * @return string|null
     */
    public function getSmtpServer()
    {
        return $this->container['smtp_server'];
    }

    /**
     * Sets smtp_server
     *
     * @param string|null $smtp_server smtp_server
     *
     * @return self
     */
    public function setSmtpServer($smtp_server)
    {
        $this->container['smtp_server'] = $smtp_server;

        return $this;
    }

    /**
     * Gets smtp_port
     *
     * @return string|null
     */
    public function getSmtpPort()
    {
        return $this->container['smtp_port'];
    }

    /**
     * Sets smtp_port
     *
     * @param string|null $smtp_port smtp_port
     *
     * @return self
     */
    public function setSmtpPort($smtp_port)
    {
        $this->container['smtp_port'] = $smtp_port;

        return $this;
    }

    /**
     * Gets smtp_ssl
     *
     * @return bool|null
     */
    public function getSmtpSsl()
    {
        return $this->container['smtp_ssl'];
    }

    /**
     * Sets smtp_ssl
     *
     * @param bool|null $smtp_ssl smtp_ssl
     *
     * @return self
     */
    public function setSmtpSsl($smtp_ssl)
    {
        $this->container['smtp_ssl'] = $smtp_ssl;

        return $this;
    }

    /**
     * Gets smtp_tls
     *
     * @return bool|null
     */
    public function getSmtpTls()
    {
        return $this->container['smtp_tls'];
    }

    /**
     * Sets smtp_tls
     *
     * @param bool|null $smtp_tls smtp_tls
     *
     * @return self
     */
    public function setSmtpTls($smtp_tls)
    {
        $this->container['smtp_tls'] = $smtp_tls;

        return $this;
    }

    /**
     * Gets smtp_user
     *
     * @return string|null
     */
    public function getSmtpUser()
    {
        return $this->container['smtp_user'];
    }

    /**
     * Sets smtp_user
     *
     * @param string|null $smtp_user smtp_user
     *
     * @return self
     */
    public function setSmtpUser($smtp_user)
    {
        $this->container['smtp_user'] = $smtp_user;

        return $this;
    }

    /**
     * Gets smtp_password
     *
     * @return string|null
     */
    public function getSmtpPassword()
    {
        return $this->container['smtp_password'];
    }

    /**
     * Sets smtp_password
     *
     * @param string|null $smtp_password smtp_password
     *
     * @return self
     */
    public function setSmtpPassword($smtp_password)
    {
        $this->container['smtp_password'] = $smtp_password;

        return $this;
    }

    /**
     * Gets smtp_password_hash
     *
     * @return string|null
     */
    public function getSmtpPasswordHash()
    {
        return $this->container['smtp_password_hash'];
    }

    /**
     * Sets smtp_password_hash
     *
     * @param string|null $smtp_password_hash smtp_password_hash
     *
     * @return self
     */
    public function setSmtpPasswordHash($smtp_password_hash)
    {
        $this->container['smtp_password_hash'] = $smtp_password_hash;

        return $this;
    }

    /**
     * Gets proxy_host
     *
     * @return string|null
     */
    public function getProxyHost()
    {
        return $this->container['proxy_host'];
    }

    /**
     * Sets proxy_host
     *
     * @param string|null $proxy_host proxy_host
     *
     * @return self
     */
    public function setProxyHost($proxy_host)
    {
        $this->container['proxy_host'] = $proxy_host;

        return $this;
    }

    /**
     * Gets proxy_port
     *
     * @return string|null
     */
    public function getProxyPort()
    {
        return $this->container['proxy_port'];
    }

    /**
     * Sets proxy_port
     *
     * @param string|null $proxy_port proxy_port
     *
     * @return self
     */
    public function setProxyPort($proxy_port)
    {
        $this->container['proxy_port'] = $proxy_port;

        return $this;
    }

    /**
     * Gets proxy_user
     *
     * @return string|null
     */
    public function getProxyUser()
    {
        return $this->container['proxy_user'];
    }

    /**
     * Sets proxy_user
     *
     * @param string|null $proxy_user proxy_user
     *
     * @return self
     */
    public function setProxyUser($proxy_user)
    {
        $this->container['proxy_user'] = $proxy_user;

        return $this;
    }

    /**
     * Gets proxy_password
     *
     * @return string|null
     */
    public function getProxyPassword()
    {
        return $this->container['proxy_password'];
    }

    /**
     * Sets proxy_password
     *
     * @param string|null $proxy_password proxy_password
     *
     * @return self
     */
    public function setProxyPassword($proxy_password)
    {
        $this->container['proxy_password'] = $proxy_password;

        return $this;
    }

    /**
     * Gets proxy_password_hash
     *
     * @return string|null
     */
    public function getProxyPasswordHash()
    {
        return $this->container['proxy_password_hash'];
    }

    /**
     * Sets proxy_password_hash
     *
     * @param string|null $proxy_password_hash proxy_password_hash
     *
     * @return self
     */
    public function setProxyPasswordHash($proxy_password_hash)
    {
        $this->container['proxy_password_hash'] = $proxy_password_hash;

        return $this;
    }

    /**
     * Gets proxy_domain
     *
     * @return string|null
     */
    public function getProxyDomain()
    {
        return $this->container['proxy_domain'];
    }

    /**
     * Sets proxy_domain
     *
     * @param string|null $proxy_domain proxy_domain
     *
     * @return self
     */
    public function setProxyDomain($proxy_domain)
    {
        $this->container['proxy_domain'] = $proxy_domain;

        return $this;
    }

    /**
     * Gets proxy_workstation
     *
     * @return string|null
     */
    public function getProxyWorkstation()
    {
        return $this->container['proxy_workstation'];
    }

    /**
     * Sets proxy_workstation
     *
     * @param string|null $proxy_workstation proxy_workstation
     *
     * @return self
     */
    public function setProxyWorkstation($proxy_workstation)
    {
        $this->container['proxy_workstation'] = $proxy_workstation;

        return $this;
    }

    /**
     * Gets currency_provider
     *
     * @return string|null
     */
    public function getCurrencyProvider()
    {
        return $this->container['currency_provider'];
    }

    /**
     * Sets currency_provider
     *
     * @param string|null $currency_provider currency_provider
     *
     * @return self
     */
    public function setCurrencyProvider($currency_provider)
    {
        $this->container['currency_provider'] = $currency_provider;

        return $this;
    }

    /**
     * Gets currency_key
     *
     * @return string|null
     */
    public function getCurrencyKey()
    {
        return $this->container['currency_key'];
    }

    /**
     * Sets currency_key
     *
     * @param string|null $currency_key currency_key
     *
     * @return self
     */
    public function setCurrencyKey($currency_key)
    {
        $this->container['currency_key'] = $currency_key;

        return $this;
    }

    /**
     * Gets enabled_zone_types
     *
     * @return \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance[]|null
     */
    public function getEnabledZoneTypes()
    {
        return $this->container['enabled_zone_types'];
    }

    /**
     * Sets enabled_zone_types
     *
     * @param \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance[]|null $enabled_zone_types enabled_zone_types
     *
     * @return self
     */
    public function setEnabledZoneTypes($enabled_zone_types)
    {
        $this->container['enabled_zone_types'] = $enabled_zone_types;

        return $this;
    }

    /**
     * Gets stats_retainment_period
     *
     * @return int|null
     */
    public function getStatsRetainmentPeriod()
    {
        return $this->container['stats_retainment_period'];
    }

    /**
     * Sets stats_retainment_period
     *
     * @param int|null $stats_retainment_period stats_retainment_period
     *
     * @return self
     */
    public function setStatsRetainmentPeriod($stats_retainment_period)
    {
        $this->container['stats_retainment_period'] = $stats_retainment_period;

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


