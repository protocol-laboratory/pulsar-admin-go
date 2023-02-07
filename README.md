# pulsar-admin-go
This project is a Go client library for the Apache Pulsar Admin API.<br/>
## usage

```go
package main

import (
	"fmt"
	"github.com/protocol-laboratory/pulsar-admin-go/padmin"
)

func main() {
	admin, err := padmin.NewDefaultPulsarAdmin()
	if err != nil {
		panic(err)
	}
	// get namespace topic list
	topics, err := admin.PersistentTopics.ListNamespaceTopics("tenant", "namespace")
	if err != nil {
		panic(err)
	}
	fmt.Println(topics)
}
```

## how to set up the pulsar broker in unit-tests
The testcases will automatically start a pulsar broker in the background using testcontainers.<br/>
Also by detect if pulsar broker is already running, the testcases will use the existing broker. That means you can use the already exists local pulsar broker for the testcases.<br/>
You can start a pulsar broker using commands below:<br/>
```bash
docker run -it --rm -p 6650:6650 -p 8080:8080 apachepulsar/pulsar:latest /pulsar/bin/pulsar-daemon start standalone --no-functions-worker --no-stream-storage
```
