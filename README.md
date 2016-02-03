# BOSH Certification Receipt Generator

Generates receipts for certified BOSH release + stemcell combinations.

## Installing pre-built Linux binaries

``` bash
URL=https://s3.amazonaws.com/bosh-certification-generator-releases/certify-artifacts-linux-amd64
curl -f -L -o certify-artifacts $URL
chmod +x certify-artifacts
```

## Usage

- To print version information,

  run the following:

  ``` bash
  ./certify-artifacts -v
  ```

- To generate a certification receipt,

  given the following flags:

  - `--release <value>` (multiple flags allowed)<br>
    where value is `name/version`
  - `--stemcell <value>` (single flag allowed)<br>
    where value is `name/version`

  run (something like) the following:

  ``` bash
  ./certify-artifacts         \
    --release bosh/250        \
    --release bosh-aws-cpi/42 \
    --stemcell bosh-aws-xen-hvm-ubuntu-trusty-go_agent/3184.1
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

  open http://192.168.100.4:8080/pipelines/build-certifier
  ```

  ...and start the build.
