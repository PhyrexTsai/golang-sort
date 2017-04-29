# golang-sort algorithms

### Algorithm list & test

- Bubble sort
- Insertion sort
- Merge sort
- Quick sort
- Selection sort


### How to build package

```
cd $GOPATH/src/
go install github.com/PhyrexTsai/sort
```

Will generate `sort.a` under `$GOPATH/pkg/darwin_amd64/PhyrexTsai/sort`

### How to run test

```
cd $GOPATH/src/github.com/PhyrexTsai/sort
go test -v
```

### How to import

```
import (
	"github.com/PhyrexTsai/sort"
)
```
