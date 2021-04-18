# WanMei OCR

WanMei OCR是一个基于 `WmCode.dll` 完美验证码识别库的go语言调用库封装。`example` 里包含一个gRPC服务示例。

### 编译

#### cmd

```bash
SET GOOS=windows 
SET GOARCH=386
go build -o wmOCRSrv.exe
```
#### powershell

```powershell
$env:GOOS="windows"
$env:GOARCH="386"
go build -o wmOCRSrv.exe
```