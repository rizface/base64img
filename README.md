## Go BASE64 IMAGE

convert base64 DataURLs to image

### Install

```shell
go get github.com/rizface/base64img
```

### Usage

```go
package main
import "github.com/rizface/base64img"
func main() {
        jpg := new(Jpg)
        jpg.Filename = "test"
        jpg.DataUrl = <DataURLs>
}
```
