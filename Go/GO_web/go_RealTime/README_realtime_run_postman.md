# Huong dan sua loi `http.Handler`, chay realtime server va test bang Postman

## 1. Nguyen nhan loi trong `main.go`

Loi:

```text
cannot use handler.NewSSEHandler() (value of type *handler.SSEHandler) as http.Handler value in argument to mux.Handle:
*handler.SSEHandler does not implement http.Handler (missing method ServeHTTP)

cannot use handler.NewWSHanlder(wsHub) (value of type *handler.WSHanlder) as http.Handler value in argument to mux.Handle:
*handler.WSHanlder does not implement http.Handler (missing method ServeHTTP)
```

Xuat hien vi `mux.Handle(...)` chi nhan gia tri implement interface `http.Handler`.

Trong Go, interface `http.Handler` duoc dinh nghia nhu sau:

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

Nghia la struct handler cua ban bat buoc phai co method ten chinh xac la:

```go
ServeHTTP
```

Trong code hien tai, ca hai handler dang viet nham thanh:

```go
ServerHTTP
```

Khac nhau chi mot chu `r`, nhung Go coi day la method hoan toan khac. Vi vay `*SSEHandler` va `*WSHanlder` chua implement `http.Handler`.

## 2. Cach sua loi `ServeHTTP`

### Sua SSE handler

Mo file:

```text
internal/handler/sse_handler.go
```

Tim dong:

```go
func (h *SSEHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
```

Doi thanh:

```go
func (h *SSEHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
```

### Sua WebSocket handler

Mo file:

```text
internal/handler/ws_handler.go
```

Tim dong:

```go
func (h *WSHanlder) ServerHTTP(w http.ResponseWriter, r *http.Request) {
```

Doi thanh:

```go
func (h *WSHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
```

Sau khi sua hai dong nay, `mux.Handle("/sse", handler.NewSSEHandler())` va `mux.Handle("/ws", handler.NewWSHanlder(wsHub))` moi hop le.

## 3. Nen sua them loi logic SSE

Trong `internal/handler/sse_handler.go`, code hien tai co dong:

```go
fmt.Fprintf(w, "event status-update\n")
```

Dong nay co hai van de:

1. Format SSE dung phai la `event: ten_event`, co dau `:`.
2. File `static/sse_demo.html` dang lang nghe event ten `stats-update`, khong phai `status-update`.

Nen doi thanh:

```go
fmt.Fprintf(w, "event: stats-update\n")
```

Neu khong sua dong nay, server co the van chay, nhung trang SSE demo se khong nhan dung event cap nhat.

## 4. Nen doi ten `WSHanlder` cho de doc

Hien tai code dang dat ten:

```go
type WSHanlder struct
func NewWSHanlder(...)
```

Day la loi chinh ta: `Hanlder` thay vi `Handler`. Neu ban muon code sach hon, co the doi thanh:

```go
type WSHandler struct {
	hub *hub.Hub
}

func NewWSHandler(hub *hub.Hub) *WSHandler {
	return &WSHandler{hub: hub}
}
```

Khi doi ten nay thi phai doi dong trong `cmd/server/main.go`:

```go
mux.Handle("/ws", handler.NewWSHandler(wsHub))
```

Luu y: viec doi ten nay khong bat buoc de het loi compile neu ban van dung nhat quan `WSHanlder`. Loi bat buoc can sua la `ServerHTTP` -> `ServeHTTP`.

## 5. Kiem tra sau khi sua

Chay cac lenh sau tai thu muc goc du an:

```powershell
go test ./...
```

Neu khong co test, ket qua mong doi se gan nhu:

```text
?   	github.com/GiaBao0510/Go-Realtime/cmd/server	[no test files]
?   	github.com/GiaBao0510/Go-Realtime/internal/handler	[no test files]
?   	github.com/GiaBao0510/Go-Realtime/internal/hub	[no test files]
?   	github.com/GiaBao0510/Go-Realtime/internal/model	[no test files]
```

Neu lenh tren con loi, doc loi dau tien Go in ra. Sau khi sua `ServeHTTP`, cac loi co the gap tiep thuong la:

1. Sai ten function neu ban doi `NewWSHanlder` thanh `NewWSHandler` nhung quen sua `main.go`.
2. Sai event SSE neu van de `event status-update` chua sua.
3. Loi port `8080` da duoc tien trinh khac su dung.

## 6. Cach chay chuong trinh

Tai thu muc goc du an, chay:

```powershell
go run ./cmd/server
```

Server dang cau hinh chay tai:

```text
http://localhost:8080
```

Voi `main.go` hien tai:

```go
mux.Handle("/", http.FileServer(http.Dir("./static")))
```

File trong thu muc `static` duoc serve truc tiep tu root URL. Vi vay URL dung la:

```text
http://localhost:8080/sse_demo.html
http://localhost:8080/ws_demo.html
```

Khong phai:

```text
http://localhost:8080/static/sse_demo.html
http://localhost:8080/static/ws_demo.html
```

Neu ban muon URL co `/static/...`, can doi route static trong `main.go` thanh:

```go
mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
```

Neu khong co nhu cau do, khong can sua.

## 7. Test SSE bang Postman

SSE la ket noi mot chieu: server day data xuong client, client khong gui message nguoc lai trong cung connection SSE.

Trong Postman:

1. Tao request moi.
2. Chon method `GET`.
3. Nhap URL:

```text
http://localhost:8080/sse
```

4. Trong tab Headers, co the them:

```text
Accept: text/event-stream
```

5. Bam `Send`.

Ket qua mong doi:

```text
event: connected
data: {"message": "...", "time": "..."}

id: 1
event: stats-update
data: {"online_users":...,"server_time":"...","cpu_load":...,"memory_usage":...}
```

Luu y: SSE la long-running request, nen Postman co the hien request dang loading lien tuc. Day la binh thuong.

Neu Postman cua ban khong hien stream tot, co the test nhanh bang PowerShell/curl:

```powershell
curl.exe -N http://localhost:8080/sse
```

## 8. Test WebSocket bang Postman

WebSocket la ket noi hai chieu: client gui message len server, server broadcast message lai cho cac client dang ket noi.

Trong Postman:

1. Chon `New`.
2. Chon `WebSocket`.
3. Nhap URL:

```text
ws://localhost:8080/ws?username=GiaBao
```

4. Bam `Connect`.
5. Gui message JSON:

```json
{"content":"Xin chao realtime"}
```

De thay ro broadcast, nen mo 2 ket noi WebSocket trong Postman:

Ket noi 1:

```text
ws://localhost:8080/ws?username=GiaBao
```

Ket noi 2:

```text
ws://localhost:8080/ws?username=TestUser
```

Khi gui message tu ket noi 1, ket noi 2 cung se nhan duoc message neu Hub va WebSocket handler chay dung.

## 9. Co can sua cau lenh Postman khong?

Khong can sua code rieng cho Postman neu ban giu endpoint hien tai:

```text
GET http://localhost:8080/sse
WS  ws://localhost:8080/ws?username=TenCuaBan
```

Chi can dam bao:

1. Server da chay bang `go run ./cmd/server`.
2. Da sua `ServerHTTP` thanh `ServeHTTP`.
3. Da sua SSE event thanh `event: stats-update` neu muon test dung voi demo HTML.
4. Port `8080` khong bi ung dung khac chiem.

## 10. Checklist tu thuc hien

1. Sua `ServerHTTP` -> `ServeHTTP` trong `sse_handler.go`.
2. Sua `ServerHTTP` -> `ServeHTTP` trong `ws_handler.go`.
3. Sua `event status-update` -> `event: stats-update`.
4. Chay `go test ./...`.
5. Chay `go run ./cmd/server`.
6. Mo browser test:

```text
http://localhost:8080/sse_demo.html
http://localhost:8080/ws_demo.html
```

7. Test Postman:

```text
GET http://localhost:8080/sse
WS  ws://localhost:8080/ws?username=GiaBao
```
