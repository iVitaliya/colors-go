# colors-go Documentation
### Install & Importing
`go get github.com/iVitaliya/colors-go`
```go
import "github.com/iVitaliya/colors-go"
```

### Current Allowed Color Types
* [Bright Colors](https://github.com/iVitaliya/colors-go/docs/BrightColors.md)
* [Dark Colors](https://github.com/iVitaliya/colors-go/docs/DarkColors.md)

### Current Allowed Background Color Types
* [Bright Background Colors](https://github.com/iVitaliya/colors-go/docs/BrightBGColors.md)
* [Dark Background Colors](https://github.com/iVitaliya/colors-go/docs/DarkBGColors.md)

### Current Allowed Util Color Types
* [Util Colors](https://github.com/iVitaliya/colors-go/docs/UtilColors.md)

# Correct Use
```go
package main

import (
  "fmt"
  "github.com/iVitaliya/colors-go"
)

func main() {
  // Bright
  fmt.Println(colors.BrightGreen("This is testing the Bright Green color"))
  
  // Dark
  fmt.Println(colors.Red("This is testing the Red color"))
  
  // Bright Background Color
  fmt.Println(colors.BgBrightYellow("This is testing the Bright Background Yellow color"))
  
  // Dark Background Color
  fmt.Println(colors.BgBlue("This is testing the Blue Background color"))
  
  // Utils + Color
  fmt.Println(colors.Dim(colors.BrightYellow("This is bright Yellow")))
}
```
