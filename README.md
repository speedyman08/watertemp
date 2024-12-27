# Building this tool
## Windows
- ```fyne package -os windows -icon resources/thermometer.png```
## Linux and Mac
- ```fyne package -os linux -icon resources/thermometer.png```
- ```fyne package -os darwin -icon resources/thermometer.png```

You could just also use ```go build . ``` however there is a high change that notifications may not work because most 
OS's will silence notifications from apps that aren't packaged properly or apps that do not have an icon

# Requirements
- Go tooling
- Fyne helper tool (```go install fyne.io/fyne/v2/cmd/fyne@latest```)
- A C Compiler (Gcc, Clang), for Fyne
- System graphics driver

# Information
This project uses Fyne for managing the system tray which is a pretty large dependency resulting in a big build time. 
However, it only takes a long time for the first build, subsequent builds will not take as long due to incremental recomp