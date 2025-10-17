# rvmbuilder

rvmbuilder outputs a macro to export Aveva Review files and the Navisworks file from the json format options.

## Installation

If you're using Go:

```bash
go install github.com/k-awata/rvmbuilder@latest
```

Otherwise you can download a binary from [Releases](https://github.com/k-awata/rvmbuilder/releases).

## Usage

- Output a sample JSON

```bash
rvmbuilder -s > sample.json
```

- Output a macro to run in E3D

```bash
rvmbuilder sample.json > export.mac
```

## JSON Editor

[RvmBuilder JSON Editor](https://k-awata.github.io/rvmbuilder/)

## License

[MIT License](LICENSE)
