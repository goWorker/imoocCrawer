#! /bin/bash

#Build web UI
cd ~/gofile/src/video_server/web
go install
cp ~/gofile/bin/web ~/gofile/bin/video_server_web_ui/web
cp -R ~/gofile/src/video_server/templates ~/gofile/bin/video_server_web_ui/