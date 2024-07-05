# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
# 	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Integration tests for the ACMPCA CertificateAuthority, Certificate, and CertificateAuthorityActivation resources.
"""

import time
import logging
import base64
import pytest

from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s
from acktest.k8s.resource import _get_k8s_api_client
from kubernetes import client
from e2e import service_marker, CRD_GROUP, CRD_VERSION, load_acmpca_resource
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e.tests.helper import ACMPCAValidator
from e2e.fixtures import k8s_secret

CREATE_WAIT_AFTER_SECONDS = 10
UPDATE_WAIT_AFTER_SECONDS = 10
DELETE_WAIT_AFTER_SECONDS = 10

@pytest.fixture(scope="module")
def create_secret(k8s_secret):
    secret = k8s_secret(
        "default",
        random_suffix_name("certificate-secret", 50),
        "certificate",
        "value"
    )
    yield secret

@pytest.fixture(scope="module")
def create_certificate_chain_secret(k8s_secret):
    secret = k8s_secret(
        "default",
        random_suffix_name("certificate-chain-secret", 50),
        "certificateChain",
        "value"
    )
    yield secret

@pytest.fixture(scope="module")
def certificate_authority_hierarchy(create_secret, create_certificate_chain_secret, acmpca_client):
    ca_name = random_suffix_name("certificate-authority", 50)
    cert_name = random_suffix_name("certificate", 30)
    act_name = random_suffix_name("certificate-authority-activation", 50)
    sub_ca_name = random_suffix_name("certificate-authority", 50)
    sub_cert_name = random_suffix_name("certificate", 30)
    sub_act_name = random_suffix_name("certificate-authority-activation", 50)

    secret = create_secret
    certificate_chain_secret = create_certificate_chain_secret
    sub_secret = create_secret
    sub_certificate_chain_secret = create_certificate_chain_secret
    end_entity_secret = create_secret

    replacements = {}
    suffix = random_suffix_name("", 10)
    replacements["NAME"] = ca_name
    replacements["COMMON_NAME"] = "www.example" + suffix + ".com"
    replacements["COUNTRY"] = "US"
    replacements["LOCALITY"] = "Arlington"
    replacements["ORG"] = "Example Organization " + suffix
    replacements["STATE"] = "Virginia"

    # Load CA CR
    ca_resource_data = load_acmpca_resource(
        "certificate_authority",
        additional_replacements=replacements,
    )

    # Create CA resource
    ca_ref = k8s.create_reference(
        CRD_GROUP, CRD_VERSION, "certificateauthorities",
        ca_name, namespace="default",
    )

    replacements["NAME"] = cert_name
    replacements["CA_NAME"] = ca_name
    replacements["CERTIFICATE_SEC_NS"] = secret.ns
    replacements["CERTIFICATE_SEC_NAME"] = secret.name
    replacements["CERTIFICATE_SEC_KEY"] = secret.key
    replacements["TEMPLATE_ARN"] = "arn:aws:acm-pca:::template/RootCACertificate/V1"

    # Load Certificate CR
    cert_resource_data = load_acmpca_resource(
        "certificate_ref",
        additional_replacements=replacements,
    )

    # Create Certificate resource
    cert_ref = k8s.create_reference(
        CRD_GROUP, CRD_VERSION, "certificates",
        cert_name, namespace="default",
    )

    replacements["NAME"] = act_name
    replacements["CERTIFICATE_CHAIN_SEC_NS"] = certificate_chain_secret.ns
    replacements["CERTIFICATE_CHAIN_SEC_NAME"] = certificate_chain_secret.name
    replacements["CERTIFICATE_CHAIN_SEC_KEY"] = certificate_chain_secret.key
    
    # Load CAActivation CR
    act_resource_data = load_acmpca_resource(
        "certificate_authority_activation_ref",
        additional_replacements=replacements,
    )
    
    # Create CAActivation resource
    act_ref = k8s.create_reference(
        CRD_GROUP, CRD_VERSION, "certificateauthorityactivations",
        act_name, namespace="default",
    )

    suffix = random_suffix_name("", 10)
    replacements["NAME"] = sub_ca_name
    replacements["COMMON_NAME"] = "www.example" + suffix + ".com"
    replacements["COUNTRY"] = "US"
    replacements["LOCALITY"] = "Arlington"
    replacements["ORG"] = "Example Organization " + suffix
    replacements["STATE"] = "Virginia"

    # Load CA CR
    sub_ca_resource_data = load_acmpca_resource(
        "subordinate_certificate_authority",
        additional_replacements=replacements,
    )

    # Create CA resource
    sub_ca_ref = k8s.create_reference(
        CRD_GROUP, CRD_VERSION, "certificateauthorities",
        sub_ca_name, namespace="default",
    )

    replacements["NAME"] = sub_cert_name
    replacements["CA_NAME"] = ca_name
    replacements["CSR_CA_NAME"] = sub_ca_name
    replacements["CERTIFICATE_SEC_NS"] = sub_secret.ns
    replacements["CERTIFICATE_SEC_NAME"] = sub_secret.name
    replacements["CERTIFICATE_SEC_KEY"] = sub_secret.key
    replacements["TEMPLATE_ARN"] = "arn:aws:acm-pca:::template/SubordinateCACertificate_PathLen2/V1"

    # Load Certificate CR
    sub_cert_resource_data = load_acmpca_resource(
        "subordinate_certificate_ref",
        additional_replacements=replacements,
    )

    # Create Certificate resource
    sub_cert_ref = k8s.create_reference(
        CRD_GROUP, CRD_VERSION, "certificates",
        sub_cert_name, namespace="default",
    )

    replacements["NAME"] = sub_act_name
    replacements["CA_NAME"] = sub_ca_name
    replacements["CERTIFICATE_CHAIN_SEC_NS"] = certificate_chain_secret.ns
    replacements["CERTIFICATE_CHAIN_SEC_NAME"] = certificate_chain_secret.name
    replacements["CERTIFICATE_CHAIN_SEC_KEY"] = certificate_chain_secret.key
    replacements["CERTIFICATE_SEC_NS"] = sub_secret.ns
    replacements["CERTIFICATE_SEC_NAME"] = sub_secret.name
    replacements["CERTIFICATE_SEC_KEY"] = sub_secret.key
    replacements["COMPLETE_CERTIFICATE_CHAIN_SEC_NS"] = sub_certificate_chain_secret.ns
    replacements["COMPLETE_CERTIFICATE_CHAIN_SEC_NAME"] = sub_certificate_chain_secret.name
    replacements["COMPLETE_CERTIFICATE_CHAIN_SEC_KEY"] = sub_certificate_chain_secret.key
    
    # Load CAActivation CR
    sub_act_resource_data = load_acmpca_resource(
        "subordinate_certificate_authority_activation_ref",
        additional_replacements=replacements,
    )
    
    # Create CAActivation resource
    sub_act_ref = k8s.create_reference(
        CRD_GROUP, CRD_VERSION, "certificateauthorityactivations",
        sub_act_name, namespace="default",
    )

    k8s.create_custom_resource(ca_ref, ca_resource_data)
    k8s.create_custom_resource(cert_ref, cert_resource_data)
    k8s.create_custom_resource(act_ref, act_resource_data)
    k8s.create_custom_resource(sub_ca_ref, sub_ca_resource_data)
    k8s.create_custom_resource(sub_cert_ref, sub_cert_resource_data)
    k8s.create_custom_resource(sub_act_ref, sub_act_resource_data)

    time.sleep(180)

    ca_cr = k8s.wait_resource_consumed_by_controller(ca_ref)
    cert_cr = k8s.wait_resource_consumed_by_controller(cert_ref)
    act_cr = k8s.wait_resource_consumed_by_controller(act_ref)
    sub_ca_cr = k8s.wait_resource_consumed_by_controller(sub_ca_ref)
    sub_cert_cr = k8s.wait_resource_consumed_by_controller(sub_cert_ref)
    sub_act_cr = k8s.wait_resource_consumed_by_controller(sub_act_ref)

    assert ca_cr is not None
    assert k8s.get_resource_exists(ca_ref)
    logging.info(ca_cr)

    assert cert_cr is not None
    assert k8s.get_resource_exists(cert_ref)
    logging.info(cert_cr)

    assert act_cr is not None
    assert k8s.get_resource_exists(act_ref)
    logging.info(act_cr)

    ca_resource_arn =  k8s.get_resource_arn(ca_cr)
    assert ca_resource_arn is not None

    assert sub_ca_cr is not None
    assert k8s.get_resource_exists(sub_ca_ref)
    logging.info(sub_ca_cr)

    assert sub_cert_cr is not None
    assert k8s.get_resource_exists(sub_cert_ref)
    logging.info(sub_cert_cr)

    assert sub_act_cr is not None
    assert k8s.get_resource_exists(sub_act_ref)
    logging.info(sub_act_cr)

    sub_ca_resource_arn =  k8s.get_resource_arn(sub_ca_cr)
    assert sub_ca_resource_arn is not None

    yield (ca_resource_arn, sub_ca_resource_arn)

    #Delete k8s resources
    _, deleted = k8s.delete_custom_resource(sub_act_ref)
    assert deleted is True

    _, deleted = k8s.delete_custom_resource(sub_cert_ref)
    assert deleted is True

    _, deleted = k8s.delete_custom_resource(sub_ca_ref)
    assert deleted is True

    _, deleted = k8s.delete_custom_resource(act_ref)
    assert deleted is True

    _, deleted = k8s.delete_custom_resource(cert_ref)
    assert deleted is True

    _, deleted = k8s.delete_custom_resource(ca_ref)
    assert deleted is True

@service_marker
class TestCertificateAuthorityHierarchy:

    def test_ca_hierarchy(self, acmpca_client, certificate_authority_hierarchy):

        (ca_resource_arn, sub_ca_resource_arn) = certificate_authority_hierarchy

        # Check CA status is ACTIVE
        acmpca_validator = ACMPCAValidator(acmpca_client)
        acmpca_validator.assert_certificate_authority(ca_resource_arn, "ACTIVE")

        acmpca_validator.assert_certificate_authority(sub_ca_resource_arn, "ACTIVE")