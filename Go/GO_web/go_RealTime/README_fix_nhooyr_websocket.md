# Huong dan khac phuc loi `main module does not need module nhooyr.io/websocket`

## Tinh trang da kiem tra

Khi chay:

```powershell
go mod why -m nhooyr.io/websocket
```

Go tra ve:

```text
# nhooyr.io/websocket
(main module does not need module nhooyr.io/websocket)
```

Thong bao nay khong co nghia la thu vien `nhooyr.io/websocket` chac chan sai. No co nghia la Go khong thay package nao trong main module dang import module do.

Trong du an hien tai, nguyen nhan chinh nam o file:

```text
internal/hub/hub.go
```

File nay dang duoc luu bang encoding `UTF-16 BE`. Khi Go doc file `.go`, cac byte NUL trong file lam Go bao loi:

```text
read internal\hub\hub.go: unexpected NUL in input
```

Vi Go khong parse duoc `hub.go`, no cung khong nhin thay dong import:

```go
import "nhooyr.io/websocket"
```

Ket qua la `go mod why` ket luan nham theo goc nhin cua Go tool: main module khong can `nhooyr.io/websocket`.

## Nguyen nhan phu trong `go.mod`

`go.mod` hien co:

```go
require nhooyr.io/websocket v1.8.11
replace nhooyr.io/websocket => ./vendor/nhooyr.io/websocket
```

Nhung trong du an hien tai khong co thu muc:

```text
vendor/nhooyr.io/websocket
```

Sau khi sua encoding cua `hub.go`, Go se doc duoc import `nhooyr.io/websocket`, nhung co the tiep tuc loi vi `replace` dang tro toi mot thu muc khong ton tai.

## Cach khac phuc de xuat

### 1. Luu lai `internal/hub/hub.go` thanh UTF-8

Dung editor nhu VS Code:

1. Mo `internal/hub/hub.go`.
2. Chon encoding o goc duoi ben phai.
3. Chon `Save with Encoding`.
4. Chon `UTF-8`.
5. Luu file.

Co the kiem tra file nao dang co byte NUL bang PowerShell:

```powershell
Get-ChildItem -Recurse -File -Include *.go | ForEach-Object {
  $bytes = [System.IO.File]::ReadAllBytes($_.FullName)
  if ($bytes -contains 0) { $_.FullName }
}
```

Sau khi sua xong, lenh tren khong nen in ra file `.go` nao.

### 2. Xu ly dong `replace` trong `go.mod`

Neu ban khong co y dinh tu quan ly source cua `nhooyr.io/websocket` trong thu muc local, nen xoa dong:

```go
replace nhooyr.io/websocket => ./vendor/nhooyr.io/websocket
```

Sau do chay:

```powershell
go mod tidy
```

Lenh nay se giu `require nhooyr.io/websocket v1.8.11` neu code that su import thu vien nay.

Neu ban muon dung `vendor`, cach thong dung hon la:

```powershell
go mod tidy
go mod vendor
```

Khong nen dung `replace ... => ./vendor/...` tru khi `vendor/nhooyr.io/websocket` la mot module local day du, co `go.mod` rieng va ban that su muon thay the module goc bang ban local do.

### 3. Kiem tra lai module graph

Sau khi sua encoding va `go.mod`, chay:

```powershell
go mod why -m nhooyr.io/websocket
```

Ket qua dung nen hien duong import dai loai:

```text
# nhooyr.io/websocket
github.com/GiaBao0510/Go-Realtime/internal/hub
nhooyr.io/websocket
```

Neu van hien `(main module does not need module nhooyr.io/websocket)`, hay kiem tra lai xem package nao con import `nhooyr.io/websocket`:

```powershell
rg 'nhooyr.io/websocket|websocket' -n
```

## Loi khac co the phat sinh sau khi sua

Sau khi Go doc duoc `hub.go`, du an co the tiep tuc lo ra cac loi khac. Cac diem can ra soat:

1. `go.mod` dang replace toi `./vendor/nhooyr.io/websocket`, nhung thu muc nay khong ton tai. Cach xu ly la xoa dong `replace` hoac tao dung module local ma dong `replace` tro toi.
2. Trong `internal/hub/hub.go` co ten constant `websocket.StatusNornamClosure`. Ten dung cua nhooyr la `websocket.StatusNormalClosure`. Neu chua sua, build co the loi undefined identifier.
3. `cmd/server/main.go` hien chi co `package server`, khong phai `package main` va chua co ham `main()`. Neu muc tieu la chay server bang `go run ./cmd/server`, can bo sung entrypoint rieng.
4. `internal/handler/ws_handler.go` hien chi co `package handler`, nen route `/ws` trong file demo `static/ws_demo.html` chua co handler WebSocket hoan chinh.
5. `internal/handler/sse_handler.go` co phan `ServeHTTP` dang bi comment/dang do. Neu build server day du, can hoan thien handler nay.

## Lenh kiem tra sau khi khac phuc

Chay lan luot:

```powershell
go mod tidy
go list ./...
go test ./...
```

Neu dung vendor:

```powershell
go mod vendor
go list -mod=vendor ./...
go test -mod=vendor ./...
```

Neu cac lenh tren qua het, loi module `nhooyr.io/websocket` da duoc xu ly on dinh. Neu tiep tuc fail, hay doc loi dau tien ma Go in ra; kha nang cao se la loi code chua hoan thien nhu cac muc da liet ke o tren, khong con la loi dependency thua trong `go.mod`.
