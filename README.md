[![Build Status](https://travis-ci.com/avdienko/gomem.svg?branch=master)](https://travis-ci.com/avdienko/gomem)
[![Go Report Card](https://goreportcard.com/badge/github.com/avdienko/gomem)](https://goreportcard.com/report/github.com/avdienko/gomem)

# gomem
gomem this is a simple thread-safe in-memory key / value store

## Requirement
Go 1.8 or higher

## Installation 
```
go get github.com/avdienko/gomem
```

## Usage
```go
package main

import (
	"fmt"
	"time"

	"github.com/avdienko/gomem"
)

func main() {
	// Set default expiration time and GC interval
	cache := gomem.New(5*time.Second, 1*time.Second)
	// If the expiration time is 0 then the default value will be used
	cache.Save("key", "value", 2*time.Minute)
	// If the value is not then the found will be false
	result, found := cache.Get("key")
	if found {
		fmt.Println(result) // output: value
	}
}

```
## Supported methods
* Get    - get value by key  
* Save   - save value by key  
* Delete - delete by key  
* Rename - rename key name   
* Count  - get the number of items  
* Exist  - check the existence of the key

## Contributing
Contributors are welcome! Feel free to improve this project or share your ideas via pull request.

## License
Distributed under MIT License, please see license file within the code for more details.
