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
from openapi_client.model.add_security_group_rules_request_rule_destination_group import AddSecurityGroupRulesRequestRuleDestinationGroup
from openapi_client.model.add_security_group_rules_request_rule_destination_tier import AddSecurityGroupRulesRequestRuleDestinationTier
from openapi_client.model.add_security_group_rules_request_rule_source_group import AddSecurityGroupRulesRequestRuleSourceGroup
from openapi_client.model.add_security_group_rules_request_rule_source_tier import AddSecurityGroupRulesRequestRuleSourceTier
globals()['AddSecurityGroupRulesRequestRuleDestinationGroup'] = AddSecurityGroupRulesRequestRuleDestinationGroup
globals()['AddSecurityGroupRulesRequestRuleDestinationTier'] = AddSecurityGroupRulesRequestRuleDestinationTier
globals()['AddSecurityGroupRulesRequestRuleSourceGroup'] = AddSecurityGroupRulesRequestRuleSourceGroup
globals()['AddSecurityGroupRulesRequestRuleSourceTier'] = AddSecurityGroupRulesRequestRuleSourceTier
from openapi_client.model.add_security_group_rules_request_rule import AddSecurityGroupRulesRequestRule


class TestAddSecurityGroupRulesRequestRule(unittest.TestCase):
    """AddSecurityGroupRulesRequestRule unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testAddSecurityGroupRulesRequestRule(self):
        """Test AddSecurityGroupRulesRequestRule"""
        # FIXME: construct object with mandatory attributes with example values
        # model = AddSecurityGroupRulesRequestRule()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
