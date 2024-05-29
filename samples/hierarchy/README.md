# CA Hierarchy Sample

This sample demonstrates how to set up a CA hierarchy using the AWS Controllers for Kubernetes (ACK) service controller for AWS Private Certificate Authority.

## Create a root CA

To create and activate your root CA, apply the manifest, `root_ca.yaml`, to your Kubernetes cluster using the following command:

```
kubectl apply -f root_ca.yaml
```

Applying this manifest should create a `CertificateAuthority` resource, `Certificate` resource, `CertificateAuthority` resource, and 2 `Secret` resources, storing the CA certificate and certificate chain. 

The following command will describe the `CertificateAuthority` resource, named `root-ca`. Once the resource's status is `ACTIVE`, the subordinate CA can be created.
```
kubectl describe certificateauthority/root-ca
```

## Create a subordinate CA

To create and activate your subordinate CA, apply the manifest, `sub_ca.yaml`, to your Kubernetes cluster using the following command:

```
kubectl apply -f sub_ca.yaml
```

Applying this manifest should create a `CertificateAuthority` resource, `Certificate` resource, `CertificateAuthority` resource, and 2 `Secret` resources, storing the CA certificate and certificate chain.

The following command will describe the `CertificateAuthority` resource, named `sub-ca`. Once the resource's status is `ACTIVE`, the end entity certificate can be issued.
```
kubectl describe certificateauthority/sub-ca
```

## Issue an end entity certificate

To issue the end entity certificate, modify the placeholder values in `end_entity_certificate.yaml` and apply the manifest to your Kubernetes cluster using the following command:

```
kubectl apply -f end_entity.yaml
```

Applying this manifest should create a `Certificate` resource and `Secret` resource, storing the end entity certificate. 
