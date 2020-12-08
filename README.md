# Zeebe get start with Go client


You can find the tutorial in the [Zeebe documentation](https://docs.zeebe.io/clients/go-client/get-started)

First start a broker

```shell
docker run --rm -d -p 26500:26500 --name broker camunda/zeebe:0.25.3
```

Then run any example in `src`:

```shell
cd src
go run example-4.go
```

* [Web Site](https://zeebe.io)
* [Documentation](https://docs.zeebe.io)
* [Issue Tracker](https://github.com/zeebe-io/zeebe/issues)
* [Slack Channel](https://zeebe-slackin.herokuapp.com/)
* [User Forum](https://forum.zeebe.io)
* [Contribution Guidelines](/CONTRIBUTING.md)

## Updating this guide

To test and update this guide for a new Zeebe version, go to the [Update Zeebe
Version](https://github.com/zeebe-io/zeebe-get-started-go-client/actions?query=workflow%3A%22Update+the+Zeebe+version%22)
workflow. Click `run workflow`, choose `master`, your version `x.y.z` and
whether or not to push the changes.

Note, that the Zeebe version `x.y.z` must be available as a docker image tagged
as `camunda/zeebe:0.25.3`

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

[broker-core]: https://github.com/zeebe-io/zeebe/tree/master/broker-core
[agpl]: https://github.com/zeebe-io/zeebe/blob/master/GNU-AGPL-3.0
