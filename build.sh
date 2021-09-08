#! /bin/bash

# Build web UI
cd ~/work/src/github.com/jakecai/WebServer/web
go install
cp ~/work/bin/web ~/work/bin/server_web_ui/web
cp -R ~/work/src/github.com/jakecai/WebServer/templates ~/work/bin/server_web_ui/
