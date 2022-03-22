# FuckCoolapkTokenV2

Generate random request token for coolapk v11+

## Usage

```go
package main

import (
    "fmt"
    "github.com/XiaoMengXinX/FuckCoolapkTokenV2"
)

func main() {
    deviceCode, reqToken := token.GetToken()
    fmt.Printf("Random devicde code: %s\nRandom token: %s\n", deviceCode, reqToken)
}
```

output:

```aidl
Random devicde code: IDMw4SNwEDMyIjLBFTUTByOhRDIsVGepBFI7UGbn92bHByOlx2Zv92RgszM4oTO5oTZ4ojYlpTO0oDNwAyOgsDI7MjM0cjRCFENFREOGNzNwQUM1AzN4YkQwgTOwMzMDNTM
Random token: v2JDJ5JDEwJE1UWTBOemt6TVRJME5BLzA4NWI0M3U1NzdKdDFMOUVFTm9Mb3I2WUwwbWFXQ1Q2bHJtN2NP
```

###Enjoy it~

##Thanks to

[QQ little ice](https://github.com/qqlittleice233)