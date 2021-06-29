# Project infomation

# Cách chạy project Tfs2

-Bước 1: truy cập forder backend->main.go, tại file main.go sửa đường dẫn mySql thành của bạn
example:
```bash
var mysqlCredentials = fmt.Sprintf(
	"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	"root",
	"password",
	"tcp",
	"localhost",
	"3306",
	"tfs2",
)
```
<div>-Bước 2: Tạo module bằng lệnh "go mod init ....", tải các pakage bằng "go get .."
<div>-Bước 3: tạo database tfs2 trong mySql</div>
<div>-Bước 4: Run file main.go bằng lệnh "go run main.go"</div>
<div>-Bước 5: Truy cập forder frontend->html, chạy file index.html</div>

Good luck!
