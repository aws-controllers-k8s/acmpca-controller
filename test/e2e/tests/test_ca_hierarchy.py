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
from e2e import service_marker, CRD_GROUP, CRD_VERSION, load_all_acmpca_resources
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e.tests.helper import ACMPCAValidator
from e2e.fixtures import k8s_secret


@pytest.fixture(scope="module")
def create_secret(k8s_secret):
    secret = k8s_secret(
        "default",
        random_suffix_name("certificate-secret", 50),
        "certificate",
        "value",
    )
    yield secret


@pytest.fixture(scope="module")
def create_certificate_chain_secret(k8s_secret):
    secret = k8s_secret(
        "default",
        random_suffix_name("certificate-chain-secret", 50),
        "certificateChain",
        "value",
    )
    yield secret


@pytest.fixture(scope="module")
def certificate_authority_hierarchy(
    create_secret, create_certificate_chain_secret, acmpca_client
):
    ca_name = random_suffix_name("certificate-authority", 50)
    cert_name = random_suffix_name("certificate", 30)
    act_name = random_suffix_name("certificate-authority-activation", 50)
    sub_ca_name = random_suffix_name("certificate-authority", 50)
    sub_cert_name = random_suffix_name("certificate", 30)
    sub_act_name = random_suffix_name("certificate-authority-activation", 50)
    end_entity_cert_name = random_suffix_name("certificate", 30)

    secret = create_secret
    certificate_chain_secret = create_certificate_chain_secret
    sub_secret = create_secret
    sub_certificate_chain_secret = create_certificate_chain_secret
    end_entity_secret = create_secret

    replacements = {}
    replacements["CA_NAME"] = ca_name
    replacements["CERT_NAME"] = cert_name
    replacements["CA_ACTIVATION_NAME"] = act_name
    replacements["SUB_CA_NAME"] = sub_ca_name
    replacements["SUB_CERT_NAME"] = sub_cert_name
    replacements["SUB_CA_ACTIVATION_NAME"] = sub_act_name
    replacements["END_ENTITY_CERT_NAME"] = end_entity_cert_name

    replacements["CERTIFICATE_SEC_NS"] = secret.ns
    replacements["CERTIFICATE_SEC_NAME"] = secret.name
    replacements["CERTIFICATE_SEC_KEY"] = secret.key
    replacements["CERTIFICATE_CHAIN_SEC_NS"] = certificate_chain_secret.ns
    replacements["CERTIFICATE_CHAIN_SEC_NAME"] = certificate_chain_secret.name
    replacements["CERTIFICATE_CHAIN_SEC_KEY"] = certificate_chain_secret.key

    replacements["SUB_CERTIFICATE_SEC_NS"] = sub_secret.ns
    replacements["SUB_CERTIFICATE_SEC_NAME"] = sub_secret.name
    replacements["SUB_CERTIFICATE_SEC_KEY"] = sub_secret.key
    replacements["COMPLETE_CERTIFICATE_CHAIN_SEC_NS"] = sub_certificate_chain_secret.ns
    replacements["COMPLETE_CERTIFICATE_CHAIN_SEC_NAME"] = sub_certificate_chain_secret.name
    replacements["COMPLETE_CERTIFICATE_CHAIN_SEC_KEY"] = sub_certificate_chain_secret.key

    replacements["END_CERTIFICATE_SEC_NS"] = end_entity_secret.ns
    replacements["END_CERTIFICATE_SEC_NAME"] = end_entity_secret.name
    replacements["END_CERTIFICATE_SEC_KEY"] = end_entity_secret.key

    # Load CRs
    resource_data = load_all_acmpca_resources(
        "ca_hierarchy",
        additional_replacements=replacements,
    )

    # Create CA reference
    ca_ref = k8s.create_reference(
        CRD_GROUP,
        CRD_VERSION,
        "certificateauthorities",
        ca_name,
        namespace="default",
    )

    # Create Certificate reference
    cert_ref = k8s.create_reference(
        CRD_GROUP, CRD_VERSION, "certificates",
        cert_name, namespace="default",
    )

    # Create CAActivation reference
    act_ref = k8s.create_reference(
        CRD_GROUP,
        CRD_VERSION,
        "certificateauthorityactivations",
        act_name,
        namespace="default",
    )

    # Create Subordinate CA reference
    sub_ca_ref = k8s.create_reference(
        CRD_GROUP,
        CRD_VERSION,
        "certificateauthorities",
        sub_ca_name,
        namespace="default",
    )

    # Create Subordinate Certificate reference
    sub_cert_ref = k8s.create_reference(
        CRD_GROUP,
        CRD_VERSION,
        "certificates",
        sub_cert_name,
        namespace="default",
    )

    # Create Subordinate CAActivation reference
    sub_act_ref = k8s.create_reference(
        CRD_GROUP,
        CRD_VERSION,
        "certificateauthorityactivations",
        sub_act_name,
        namespace="default",
    )

    # Create End-entity Certificate reference
    end_entity_cert_ref = k8s.create_reference(
        CRD_GROUP,
        CRD_VERSION,
        "certificates",
        end_entity_cert_name,
        namespace="default",
    )

    references = [ca_ref, cert_ref, act_ref, sub_ca_ref, sub_cert_ref, sub_act_ref, end_entity_cert_ref]
    i = 0
    for resource in resource_data:
        logging.info(resource)
        k8s.create_custom_resource(references[i], resource)
        assert k8s.wait_on_condition(references[i], "Ready", "True", wait_periods=15)
        i += 1

    time.sleep(180)

    ca_cr = k8s.wait_resource_consumed_by_controller(ca_ref)
    cert_cr = k8s.wait_resource_consumed_by_controller(cert_ref)
    act_cr = k8s.wait_resource_consumed_by_controller(act_ref)
    sub_ca_cr = k8s.wait_resource_consumed_by_controller(sub_ca_ref)
    sub_cert_cr = k8s.wait_resource_consumed_by_controller(sub_cert_ref)
    sub_act_cr = k8s.wait_resource_consumed_by_controller(sub_act_ref)
    end_entity_cert_cr = k8s.wait_resource_consumed_by_controller(end_entity_cert_ref)

    for ref in references:
        assert k8s.wait_on_condition(ref, "Ready", "True", wait_periods=15)

    assert ca_cr is not None
    assert k8s.get_resource_exists(ca_ref)
    logging.info(ca_cr)

    assert cert_cr is not None
    assert k8s.get_resource_exists(cert_ref)
    logging.info(cert_cr)

    assert act_cr is not None
    assert k8s.get_resource_exists(act_ref)
    logging.info(act_cr)

    ca_resource_arn = k8s.get_resource_arn(ca_cr)
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

    sub_ca_resource_arn = k8s.get_resource_arn(sub_ca_cr)
    assert sub_ca_resource_arn is not None

    assert end_entity_cert_cr is not None
    assert k8s.get_resource_exists(end_entity_cert_ref)
    logging.info(end_entity_cert_cr)

    end_entity_cert_arn = k8s.get_resource_arn(end_entity_cert_cr)
    assert end_entity_cert_arn is not None

    yield (ca_resource_arn, sub_ca_resource_arn, end_entity_cert_arn, end_entity_secret)

    # Delete K8s resources
    _, deleted = k8s.delete_custom_resource(end_entity_cert_ref)
    assert deleted is True

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

        (
            ca_resource_arn,
            sub_ca_resource_arn,
            end_entity_cert_arn,
            end_entity_secret,
        ) = certificate_authority_hierarchy

        # Check CA status is ACTIVE
        acmpca_validator = ACMPCAValidator(acmpca_client)
        acmpca_validator.assert_certificate_authority(ca_resource_arn, "ACTIVE")
        acmpca_validator.assert_certificate_authority(sub_ca_resource_arn, "ACTIVE")

        # Get End-entity certificate
        end_entity_cert = acmpca_validator.get_certificate(
            ca_arn=sub_ca_resource_arn, cert_arn=end_entity_cert_arn
        )
        assert end_entity_cert is not None

        # Check certificate is in secret
        _api_client = _get_k8s_api_client()
        api_response = (
            client.CoreV1Api(_api_client)
            .read_namespaced_secret(end_entity_secret.name, end_entity_secret.ns)
            .data
        )

        assert end_entity_secret.key in api_response
        assert (
            base64.b64decode(api_response[end_entity_secret.key]).decode("ascii")
            == end_entity_cert
        )
