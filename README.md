# NDC example with CSV

The connector is ported from [NDC Reference Connector](https://github.com/hasura/ndc-spec/tree/main/ndc-reference) that read CSV files into memory. It is intended to illustrate the concepts involved, and should be complete, in the sense that every specification feature is covered. It is not optimized and is not intended for production use, but might be useful for testing.

## Getting Started

```bash
go run . serve
```

## Using the reference connector

The reference connector runs on http://localhost:8080:

```sh
curl http://localhost:8080/schema | jq .
```

## Steps to deploy a connector

1. Clone reference connector from github [NDC Reference Connector](https://github.com/hasura/ndc-spec/tree/main/ndc-reference)
2. Push connector's code to your own public github repository.
3. Install Hasura CLI
4. Install connector plugin, follow (https://hasura.io/docs/latest/hasura-cli/connector-plugin/)
5. Create connector

```bash
hasura3 connector create <CONNECTOR_NAME> --github-repo-url <CONNECTOR_GITHUB_URL>
```

6. List connector to get the url of deployed connector.

```bash
hasura3 connector list
```

#### Use connector in hasura project:
7. create folder <FOLDER_NAME> in subgraphs/default/dataconnectors 
8. Create file in that folder with <CONNECTOR>.hml format and add following configuration in it.

```bash
kind: DataConnector
version: v2
definition:
  name: elastic
  url:
    singleUrl:
      value: <CONNECTOR_URL>
```

9. Specify connector in build-profile.yaml

```bash
version: 2
spec:
  environment: default
  mode: replace
  supergraph:
    resources:
    - supergraph/*
  subgraphs:
  - name: default
    resources:
    - subgraphs/default/**/*.hml
    connectors:
      <CONNECTOR_NAME>:
        path: subgraphs/default/dataconnectors/<FOLDER_NAME>
        connectorConfigFile: <CONNECTOR>.hml
```

10. Build hasura project

```bash
hasura3 watch --dir .
```

11. If getting err that versions not found than rename version to versions in capability response and run fillowing command:

```bash
hasura3 build create -d "<BUILD_NAME>"
```