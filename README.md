# gobcache

gobcache is a simple way to store data temporarily. 

[![GoDoc](https://godoc.org/github.com/ilk/gobcache?status.svg)](https://godoc.org/github.com/ilk/gobcache)

Main usage is for development where you have to deal with data which takes long to fetch.

## Example usage
```go
  import (
    "fmt"

    "github.com/ilk/gobcache"
  )

  func main() {
    cache := gobcache.NewCache(gobcache.Config{})
    data := ""
    exampleData := `{"data":"I'm a huge json'"}`

    // get data from cache if exists
    if err := cache.GetData("123", &data); err != nil {
      // instead of assign exampleData, fetch your "fetch" function or whatever
      data = exampleData

      if err := cache.SaveData("123", data); err != nil {
        // handle error
        fmt.Printf("ooops %v\n", err)
      }
    }
    //
    fmt.Println(data)
  }
```
