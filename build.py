#!/usr/bin/env python3
# 2021.1.15
# -- Imports --------------------------------------------------------------------------

from subprocess import call
from pathlib import Path
from platform import system

# -------------------------------------------------------------------------- Imports --

# -- Build project --------------------------------------------------------------------------

PROJECT_DIR = Path(__file__).parent
LINUX_BIN = PROJECT_DIR.joinpath('el-logserver-linux')
WIN_BIN = PROJECT_DIR.joinpath('el-logserver-win')
MAC_BIN = PROJECT_DIR.joinpath('el-logserver-mac')
OS = system()

# compile for linux
call(
    f"GOOS=linux GOARCH=amd64;"
    f" go build -o \"{LINUX_BIN}\" "
    f"\"{PROJECT_DIR.joinpath('main.go')}\"",
    shell=True
)
# compile for win
call(
    f"GOOS=windows GOARCH=amd64;"
    f" go build -o \"{WIN_BIN}\" "
    f"\"{PROJECT_DIR.joinpath('main.go')}\"",
    shell=True
)
# compile for mac
call(
    f"GOOS=darwin GOARCH=amd64;"
    f" go build -o \"{MAC_BIN}\" "
    f"\"{PROJECT_DIR.joinpath('main.go')}\"",
    shell=True
)

# export routers info
if OS == "Linux":
    call(f"\"{LINUX_BIN}\" -export-route-list=true", shell=True)
elif OS == "Windows":
    call(f"\"{WIN_BIN}\" -export-route-list=true", shell=True)
elif OS == "Darwin":
    call(f"\"{MAC_BIN}\" -export-route-list=true", shell=True)

# -------------------------------------------------------------------------- Build project --
