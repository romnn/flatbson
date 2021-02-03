## flatbson

[![Build Status](https://github.com/romnn/flatbson/workflows/test/badge.svg)](https://github.com/romnn/flatbson/actions)
[![GitHub](https://img.shields.io/github/license/romnn/flatbson)](https://github.com/romnn/flatbson)
[![GoDoc](https://godoc.org/github.com/romnn/flatbson?status.svg)](https://godoc.org/github.com/romnn/flatbson)
[![Test Coverage](https://codecov.io/gh/romnn/flatbson/branch/master/graph/badge.svg)](https://codecov.io/gh/romnn/flatbson)
[![Release](https://img.shields.io/github/release/romnn/flatbson)](https://github.com/romnn/flatbson/releases/latest)

Go package for flattening bson documents.



#### Usage as a library

```golang
import "github.com/romnn/flatbson"
```

Example:
```golang
import (
	"fmt"
	"github.com/romnn/flatbson"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
    input := bson.D{
		{"user", bson.D{{"email", "test"}}},
		{"metadata", bson.D{{"city", bson.D{{"name", "Berlin"}}}}},
	}
	// Flatten the document
	flattened, err := flatbson.Flattened(input, ".")
	if err != nil {
		panic(err)
	}
    fmt.Println(flattened)
}
```

For more examples, see `examples/`.


#### Development

######  Prerequisites

Before you get started, make sure you have installed the following tools::

    $ python3 -m pip install -U cookiecutter>=1.4.0
    $ python3 -m pip install pre-commit bump2version invoke ruamel.yaml halo
    $ go get -u golang.org/x/tools/cmd/goimports
    $ go get -u golang.org/x/lint/golint
    $ go get -u github.com/fzipp/gocyclo
    $ go get -u github.com/mitchellh/gox  # if you want to test building on different architectures

**Remember**: To be able to excecute the tools downloaded with `go get`, 
make sure to include `$GOPATH/bin` in your `$PATH`.
If `echo $GOPATH` does not give you a path make sure to run
(`export GOPATH="$HOME/go"` to set it). In order for your changes to persist, 
do not forget to add these to your shells `.bashrc`.

With the tools in place, it is strongly advised to install the git commit hooks to make sure checks are passing in CI:
```bash
invoke install-hooks
```

You can check if all checks pass at any time:
```bash
invoke pre-commit
```

Note for Maintainers: After merging changes, tag your commits with a new version and push to GitHub to create a release:
```bash
bump2version (major | minor | patch)
git push --follow-tags
```

#### Note

This project is still in the alpha stage and should not be considered production ready.
