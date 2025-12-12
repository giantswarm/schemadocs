[![CircleCI](https://dl.circleci.com/status-badge/img/gh/giantswarm/schemadocs/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/giantswarm/schemadocs/tree/main)

# schemadocs - the JSON Schema docs generator

schemadocs helps you generate documentation from your JSON schema and store it in README and other documentation files. It is one possible way to generate documentation for Helm chart configuration options.

## Features

- Generate Markdown documentation from a JSON schema file and inject it into an output text file
- Validate generated Markdown documentation in a text file and check whether it is aligned with documentation generated from the supplied JSON schema file

## Example

Given the following JSON schema:

<details>
<summary>Schow JSON schema</summary>

```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "type": "object",
  "properties": {
    "clusterName": {
      "type": "string",
      "title": "Cluster name",
      "description": "Unique identifier for the cluster"
    },
    "replicas": {
      "type": "integer",
      "title": "Replicas",
      "default": 3,
      "description": "Number of replicas to run"
    },
    "enabled": {
      "type": "boolean",
      "title": "Enable feature",
      "default": true
    },
    "database": {
      "type": "object",
      "title": "Database settings",
      "properties": {
        "host": {
          "type": "string",
          "title": "Host",
          "description": "Database host address"
        },
        "port": {
          "type": "integer",
          "title": "Port",
          "default": 5432
        }
      }
    }
  }
}
```

</details>

The generated documentation will look like this using the default (tabular) output format:

---

### Database settings
Properties within the `.database` top-level object

| **Property** | **Description** | **More Details** |
| :----------- | :-------------- | :--------------- |
| `database.host` | **Host** - Database host address|**Type:** `string`<br/>|
| `database.port` | **Port**|**Type:** `integer`<br/>**Default:** `5432`|

### Other

| **Property** | **Description** | **More Details** |
| :----------- | :-------------- | :--------------- |
| `clusterName` | **Cluster name** - Unique identifier for the cluster|**Type:** `string`<br/>|
| `enabled` | **Enable feature**|**Type:** `boolean`<br/>**Default:** `true`|
| `replicas` | **Replicas** - Number of replicas to run|**Type:** `integer`<br/>**Default:** `3`|

---

## Installation

```nohighlight
go install github.com/giantswarm/schemadocs@latest
```

## Usage

### Generating documentation

Executing `schemadocs generate` without any options will generate Markdown documentation from a JSON schema file and store it in a `README.md` file in the current working directory.

It is required that the README file contains exactly one pair of placeholder strings, which mark the start and end of the documentation.
By default, the placeholders are `<!-- DOCS_START -->` and `<!-- DOCS_END -->`.

If the README file does not exist or in case it does not contain one pair of valid placeholder strings, the execution will fail.

```nohighlight
$ schemadocs generate schema.json

Generating documentation
Schema: schema.json
Destination file: README.md
Using default placeholders '<!-- DOCS_START -->' and '<!-- DOCS_END -->'

[SUCCESS] Documentation generated successfully!
```

To generate the documentation to a custom file, apply the `--output-path` option.
To use a pair of custom placeholder strings, apply the `--doc-placeholder-start` and `--doc-placeholder-end`.

```nohighlight
schemadocs generate schema.json \
  --output-path README.md \
  --doc-placeholder-start <!-- DOCS_START --> \
  --doc-placeholder-end <!-- DOCS_END -->
```

**Note:** The `--doc-placeholder-start` and `--doc-placeholder-end` need to be provided as a pair. So, if one of them is specified, the other one needs to be specified, too.

Use `--help` to learn about more options.

### Validation

Validate documentation in a text file by comparing it to documentation generated from a JSON schema file specified in the required `--schema` option.
The validation passes in case both documentations match exactly.

It is required that the text file contains exactly one pair of placeholder strings, which mark the start and end of the documentation.
By default, the placeholders are `<!-- DOCS_START -->` and `<!-- DOCS_END -->`.

If the text file or the schema file do not exist or in case the text file does not contain one pair of valid placeholder strings, the execution will fail.

```nohighlight
$ schemadocs validate README.md --schema schema.json

Validating documentation
Source file: README.md
Schema file: schema.json
Using default placeholders '<!-- DOCS_START -->' and '<!-- DOCS_END -->'

[SUCCESS] Documentation is valid!
```

To use a pair of custom placeholder strings, apply the `--doc-placeholder-start` and `--doc-placeholder-end`.

```nohighlight
schemadocs validate README.md --schema schema.json --doc-placeholder-start [DOCS_START] --doc-placeholder-end [DOCS_END]
```

**Note:** The `--doc-placeholder-start` and `--doc-placeholder-end` need to be provided as a pair. So, if one of them is specified, the other one needs to be specified, too.

Use `--help` to learn about more options.
