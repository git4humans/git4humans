set "version=0.0.1"

goreleaser release --snapshot --skip-publish --rm-dist

rmdir /s /q release\%version%\bin
Xcopy dist release\%version%\bin /E/H/C/I
rmdir /s /q dist

cd release\%version%\bin 

ren git4humans-linux-v0.0.0-32bit.zip git4humans-linux-%version%-32bit.zip
ren git4humans-linux-v0.0.0-64bit.zip git4humans-linux-%version%-64bit.zip
ren git4humans-linux-v0.0.0-arm64.zip git4humans-linux-%version%-arm64.zip
ren git4humans-macOS-v0.0.0-64bit.zip git4humans-macOS-%version%-64bit.zip
ren git4humans-macOS-v0.0.0-arm64.zip git4humans-macOS-%version%-arm64.zip
ren git4humans-win-v0.0.0-32bit.zip git4humans-win-%version%-32bit.zip
ren git4humans-win-v0.0.0-64bit.zip git4humans-win-%version%-64bit.zip
ren git4humans-win-v0.0.0-arm64.zip git4humans-win-%version%-arm64.zip

cd ../../..