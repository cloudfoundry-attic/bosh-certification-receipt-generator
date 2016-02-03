# BOSH Certification Receipt Generator

Generates receipts for certified BOSH release + stemcell combinations.

## Installing pre-built Linux binaries

```
curl ...
chmod +x certify-artifacts
```

## Usage

- To print version information,

  run the following:

  ``` bash
  ./certify-artifacts -v
  ```

- To generate a certification receipt,

  run the following:

  ``` bash
  ./certify-artifacts         \
    --release bosh/250        \
    --release bosh-aws-cpi/42 \
    --stemcell bosh-aws-xen-hvm-ubuntu-trusty-go_agent/3184.1

  # ... where flags are:
  #   --release value (multiple flags allowed)
  #     where value is name/version
  #   --stemcell value (single flag allowed)
  #     where value is name/version
  ```

  resulting in:

  ``` json
  {
    "releases": [
      {
        "name": "bosh",
        "version": "250"
      },
      {
        "name": "bosh-aws-cpi",
        "version": "42"
      }
    ],
    "stemcell": {
      "name": "bosh-aws-xen-hvm-ubuntu-trusty-go_agent",
      "version": "3184.1"
    }
  }
  ```

## Building for Linux

- Install Concourse ([see how](http://concourse.ci/getting-started.html))
- Set up a local pipeline:

  given `path/to/config.yml`, as a YAML configuration file:

  ``` yaml
  ---
  release__bucket:            # An S3 bucket
  release__bucket_access_key: # ... for `release__bucket`
  release__bucket_secret_key: # ... for `release__bucket`
  github__deployment_key: |
    -----BEGIN RSA PRIVATE KEY-----
    GITHUB PRIVATE KEY DATA
    -----END RSA PRIVATE KEY-----
  ```

  run the following:

  ``` bash
  fly set-pipeline
    --pipeline build-certifier \
    --config ci/pipeline.yml   \
    --load-vars-from path/to/config.yml
  ```
