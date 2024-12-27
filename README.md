# Building this tool
## Windows
- Use the script provided in ```scripts/build-windows.bat```
## Linux and Mac
- Just use ```go build```

# Requirements
- Go tooling
- A C Compiler (Gcc, Clang), for Fyne
- System graphics driver

# Information
This project uses Fyne for managing the system tray which is a pretty large dependency resulting in a big build time. 
However, it only takes a long time for the first build, subsequent builds will not take as long due to incremental recomp.