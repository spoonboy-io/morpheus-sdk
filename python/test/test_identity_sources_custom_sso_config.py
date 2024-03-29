"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import openapi_client
from openapi_client.model.identity_sources_custom_sso_config_config import IdentitySourcesCustomSSOConfigConfig
from openapi_client.model.identity_sources_ldap_config_default_account_role import IdentitySourcesLDAPConfigDefaultAccountRole
from openapi_client.model.list_deploys200_response_all_of_app_deploys_inner_instance import ListDeploys200ResponseAllOfAppDeploysInnerInstance
globals()['IdentitySourcesCustomSSOConfigConfig'] = IdentitySourcesCustomSSOConfigConfig
globals()['IdentitySourcesLDAPConfigDefaultAccountRole'] = IdentitySourcesLDAPConfigDefaultAccountRole
globals()['ListDeploys200ResponseAllOfAppDeploysInnerInstance'] = ListDeploys200ResponseAllOfAppDeploysInnerInstance
from openapi_client.model.identity_sources_custom_sso_config import IdentitySourcesCustomSSOConfig


class TestIdentitySourcesCustomSSOConfig(unittest.TestCase):
    """IdentitySourcesCustomSSOConfig unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testIdentitySourcesCustomSSOConfig(self):
        """Test IdentitySourcesCustomSSOConfig"""
        # FIXME: construct object with mandatory attributes with example values
        # model = IdentitySourcesCustomSSOConfig()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
