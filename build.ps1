# 删除 target 目录及其内容
Remove-Item -Recurse -Force ./target -ErrorAction SilentlyContinue

# 构建 Go 程序并输出到 target 目录
go build -ldflags="-s -w" -o ./target/QuarkShell.exe .
