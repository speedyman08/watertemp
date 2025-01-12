# Build
```fyne package -os (darwin/windows/linux) -icon resources/thermometer.png```

using "go build ." is also possible however most operating systems will not allow for
notifications to be sent from apps that aren't packaged, e.g. windows

# Requirements
- Go toolchain
- Fyne helper tool (```go install fyne.io/fyne/v2/cmd/fyne@latest```)
- C Compiler (gcc, clang)

# Information
It'll take a while to build the first time due to fyne, and it's C stuff needing to be compiled.
Will not take as long for later builds 