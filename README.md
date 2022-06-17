## banimport

This tool detects prohibited imports. You can control "import dependencies" among go packages, for example, in your monorepo.

## Install

```shell
go install github.com/nrnrk/banimport@latest
```

## How to use

1. Install command

```shell
go install github.com/nrnrk/banimport@latest
```

2. Prepare config file
    * Edit a config file to prohibit "import dependencies"

```shell
cd <your repository root>
vi .banimport.json
```

example: .banimport.json
```json:
{
  "pattern": {
    "github.com/example/monorepo/svca": [
        "github.com/example/monorepo/svcb",
        "github.com/example/monorepo/svcc",
    ],
    "github.com/example/monorepo/svcb": [
        "github.com/example/monorepo/svcc",
        "github.com/example/monorepo/svca",
    ],
    "github.com/example/monorepo/svcc": [
        "github.com/example/monorepo/svca",
        "github.com/example/monorepo/svcb",
    ]
  }
}
```

3. Execute by `go vet`

```shell
go vet vettool=$(which banimport) -banimport.config=$(cat .banimport.json) ./...
```

## Usecase

```text
├── libA
├── libB
├── svcA
├── svcB
└── svcC
```

* `libA` and`libB` are intended to be used by other packages.
* `svcA`, `svcB`, and `svcC` sholud not be used by other packaes even in the module.

In the above case, you can prohibit "import dependencies" on `svcA`, `svcB`, and `svcC` to use the following config file.

.banimport.json
```json
{
  "pattern": {
    "github.com/nrnrk/monorepo/libA": [
        "github.com/nrnrk/monorepo/svcA",
        "github.com/nrnrk/monorepo/svcB",
        "github.com/nrnrk/monorepo/svcC"
    ],
    "github.com/nrnrk/monorepo/libB": [
        "github.com/nrnrk/monorepo/svcA",
        "github.com/nrnrk/monorepo/svcB",
        "github.com/nrnrk/monorepo/svcC"
    ],
    "github.com/nrnrk/monorepo/svcA": [
        "github.com/nrnrk/monorepo/svcB",
        "github.com/nrnrk/monorepo/svcC"
    ],
    "github.com/nrnrk/monorepo/svcB": [
        "github.com/nrnrk/monorepo/svcC",
        "github.com/nrnrk/monorepo/svcA"
    ],
    "github.com/nrnrk/monorepo/svcC": [
        "github.com/nrnrk/monorepo/svcA",
        "github.com/nrnrk/monorepo/svcB"
    ]
  }
}
```


## Hints

* If your service packages include some packages used by other external packages, you might want to use `internal` package. This checker is useful when your service packages do not have any packages supposed to be used by external packages.
    * You can check the details about `internal` package [here](https://go.dev/doc/go1.4#internalpackages).


This repository is originally created using [skeleton](https://github.com/gostaticanalysis/skeleton).
