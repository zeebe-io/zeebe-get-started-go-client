# ⚠️ 📣 This repository is no longer maintained, please visit https://github.com/camunda/camunda-platform-get-started


# Zeebe - Get Started Go Client

You can find the tutorial in the [Zeebe documentation](https://docs.camunda.io/docs/product-manuals/clients/go-client/get-started).)

* [Web Site](https://camunda.com/products/cloud/)
* [Documentation](https://docs.camunda.io/)
* [Issue Tracker](https://github.com/camunda-cloud/zeebe/issues)
* [Slack Channel](https://zeebe-slackin.herokuapp.com/)
* [User Forum](https://forum.camunda.io/)
* [Contribution Guidelines](/CONTRIBUTING.md)

## Camunda Cloud Deployment
Export the connection settings as environment variables:

```
export ZEEBE_ADDRESS='[Zeebe API]'
export ZEEBE_CLIENT_ID='[Client ID]'
export ZEEBE_CLIENT_SECRET='[Client Secret]'
export ZEEBE_AUTHORIZATION_SERVER_URL='[OAuth API]'
```

**Hint:** When you create client credentials in Camunda Cloud you have the option to download a file with above lines filled out for you.

Then run any example in `src`:

```shell
cd src
go run example-4.go
```


## Docker Deployment
First start a broker

```shell
docker run --rm -d -p 26500:26500 --name broker camunda/zeebe:8.0.0
```

Then run any example in `src`:

```shell
cd src
go run example-4.go
```


## Updating this guide

To test and update this guide for a new Zeebe version, go to the [Update Zeebe
Version](https://github.com/zeebe-io/zeebe-get-started-go-client/actions?query=workflow%3A%22Update+the+Zeebe+version%22)
workflow. Click `run workflow`, choose `master`, your version `x.y.z` and
whether or not to push the changes.

Note, that the Zeebe version `x.y.z` must be available as a docker image tagged
as `camunda/zeebe:1.1.0

## Code of Conduct

This project adheres to the Contributor Covenant [Code of
Conduct](/CODE_OF_CONDUCT.md). By participating, you are expected to uphold
this code. Please report unacceptable behavior to code-of-conduct@zeebe.io.

## License

Most Zeebe source files are made available under the [Apache License, Version
2.0](/LICENSE) except for the [broker-core][] component. The [broker-core][]
source files are made available under the terms of the [GNU Affero General
Public License (GNU AGPLv3)][agpl]. See individual source files for
details.

[broker-core]: https://github.com/camunda-cloud/zeebe/tree/master/broker-core
[agpl]: https://github.com/camunda-cloud/zeebe/blob/master/GNU-AGPL-3.0
