# getvero-go
A package to send getvero requests

```console
go get "github.com/StirlingMarketingGroup/getvero-go"
```

## Currently tested with these types
  - Add Tags
  - Remove Tags
  - Event with no data
  - Event with data 

## Examples
  - [User](#user)
    - [Add User](#add-user)
    - [Alias](#alias)
    - [Unsubscribe](#unsubscribe)
    - [Resubscribe](#resubscribe)
    - [Delete](#delete)
  - [Tags](#tags)
    - [Add Tags](#add-tags)
    - [Remove Tags](#remove-tags)
  - [Event](#event)
    - [Track Event](#track-event)
    - [Track Event With Data](#track-event-with-data)

--------
### Basic Example
```go
package main

import (
    "github.com/StirlingMarketingGroup/getvero-go"
    "os"    
)

func main() {
   
    v := getvero.Vero
    v.AuthToken = os.Getenv("YOUR_GETVERO_AUTH_TOKEN")

    v.AddTags(123, "Boring Client")
    // or, the user ID can be a string
    v.AddTags("abc123", "Interesting Client")

}
```

## User
### Add User
```go
package main

import (
    "github.com/StirlingMarketingGroup/getvero-go"
    "os"    
)

func main() {
   
    v := getvero.Vero
    v.AuthToken = os.Getenv("YOUR_GETVERO_AUTH_TOKEN")

    // Add a user with a known email address
    v.IdentifyUserWithEmail(1337, "fakeemail@example.com")

    // Add a user with no known email addresss
    v.IdentifyUserWithoutEmail(897654)

    // Add extra data to user with any format (This can be done in both WithoutEmail and WithEmail)
    d := make(map[string]interface{})
    d["CountryID"] = 235
    d["PhoneNumber"] = "+1-202-555-0127"
    d["pi"] = 3.14159

    v.IdentifyUserWithEmail("123abdf58546", "superfakeemailaddresss@example.com", d)

}
```

### Alias
```go
package main

import (
    "github.com/StirlingMarketingGroup/getvero-go"
    "os"    
)

func main() {
   
    v := getvero.Vero
    v.AuthToken = os.Getenv("YOUR_GETVERO_AUTH_TOKEN")

    // user ID can be either a string or an int for both new and old
    // old is first, new is second
    v.Alias(13548, "786543245")

}
```

### Unsubscribe
```go
package main

import (
    "github.com/StirlingMarketingGroup/getvero-go"
    "os"    
)

func main() {
   
    v := getvero.Vero
    v.AuthToken = os.Getenv("YOUR_GETVERO_AUTH_TOKEN")

    v.Unsubscribe(1337)

}
```

### Resubscribe
```go
package main

import (
    "github.com/StirlingMarketingGroup/getvero-go"
    "os"    
)

func main() {
   
    v := getvero.Vero
    v.AuthToken = os.Getenv("YOUR_GETVERO_AUTH_TOKEN")

    v.Resubscribe(1337)

}
```

### Delete
```go
package main

import (
    "github.com/StirlingMarketingGroup/getvero-go"
    "os"    
)

func main() {
   
    v := getvero.Vero
    v.AuthToken = os.Getenv("YOUR_GETVERO_AUTH_TOKEN")

    v.Delete(1337)

}
```

## Tags
### Add Tags
```go
package main

import (
    "github.com/StirlingMarketingGroup/getvero-go"
    "os"    
)

func main() {
   
    v := getvero.Vero
    v.AuthToken = os.Getenv("YOUR_GETVERO_AUTH_TOKEN")

    // add a single tag through a string
    v.AddTags(1337, "Talkitive Customer")

    // add multiple tags through miltiple string arguments
    v.AddTags(1337, "Talkitive Customer", "t-shirt fanatic", "so many tags")

    // add a single tag in an array
    singleTagArray := make([]string, 1)
    singleTagArray[0] = "Silent" 
    v.AddTags(1337, singleTagArray)

    // add multiple tags in array
    multipleTagArray := make([]string, 2)
    multipleTagArray = append(multipleTagArray, "t-shirt")
    multipleTagArray = append(multipleTagArray, "custom item")
    v.AddTags(1337, multipleTagArray)

}
```

### Remove Tags
```go
package main

import (
    "github.com/StirlingMarketingGroup/getvero-go"
    "os"    
)

func main() {
   
    v := getvero.Vero
    v.AuthToken = os.Getenv("YOUR_GETVERO_AUTH_TOKEN")

    // remove a single tag through a string
    v.RemoveTags(1337, "Talkitive Customer")

    // remove multiple tags through multiple string arguments
    v.RemoveTags(1337, "silent customer", "paid in quarters")

    // remove a single tag in an array
    singleTagArray := make([]string, 1)
    singleTagArray[0] = "Silent" 
    v.RemoveTags(1337, singleTagArray)

    // remove multiple tags in array
    multipleTagArray := make([]string, 2)
    multipleTagArray = append(multipleTagArray, "t-shirt")
    multipleTagArray = append(multipleTagArray, "custom item")
    v.RemoveTags(1337, multipleTagArray)

}
```

## Event
### Track Event
```go
package main

import (
    "github.com/StirlingMarketingGroup/getvero-go"
    "os"    
)

func main() {
   
    v := getvero.Vero
    v.AuthToken = os.Getenv("YOUR_GETVERO_AUTH_TOKEN")

    // used to just add an event occured without any extra data
    v.TrackEvent(1337, "fakeemail@example.com", "item sold")

}
```

### Track Event With Data
```go
package main

import (
    "github.com/StirlingMarketingGroup/getvero-go"
    "os"    
    "time"
)

func main() {
   
    v := getvero.Vero
    v.AuthToken = os.Getenv("YOUR_GETVERO_AUTH_TOKEN")


    // structs can be used as the data sent
    type metaOrder struct {
        FirstOrder bool
        PaymentOption string
    }

    type order struct {
        Item string
        ItemID int
        Qty int
        DateNeeded time.Time
        Meta metaOrder
    }

    d := &order{
        Item: "t-shirt",
        ItemID: 12,
        Qty: 9001,
        DateNeeded: time.Now().Add(time.Hour() * 24 * 3)
        Meta: &metaOrder{
            FirstOrder: true,
            PaymentOption: "credit-card",
        }
    }

    // used to track event and data associated with it
    v.TrackEventWithData(1337, "fakeemail@example.com", "item sold", d)

    // maps can be used as well
    q := make(map[string]interface{})
    q["Qty"] = 9001
    q["Item"] = "t-shirt"

    v.TrackEventWithData(1337, "fakeemail@example.com", "item sold", q)

}
```