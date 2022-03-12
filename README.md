# References
https://iraspa.org/blog/visual-studio-code-c-cpp-fortran-with-multiple-source-files/

# Extensions
* C/C++
* C/C++ Extension Pack
* CMake
* CMake Tools

# Overview
Platform specifics goes to settings.json:
terminal.integrated.env.windows/linux/osx

tasks.json and c_cpp_properties.json references specifics as e.g:
`${config:terminal.integrated.env.windows.app_name}`

# Toolchain setup
TBD

# Usage
## Setup/Change configuration
`ctrl-shift-p c++ select a configuration` choose `msys`/`Mac`

## Build
`ctrl-shift-p cmake:build`
or `ctrl-shift-p tasks run task -> cmakebuild`

## Run
`ctrl-shift-p tasks run task` -> `runappmsys`/`runappmac`
run should dependsOn cmakebuild, but currently not working

## Debug
Set a breakpoint in source file,
select Run and Debug on the Debug start view,
select `(gdb) LaunchMsys`/`(gdb) LaunchMac`, press F5
