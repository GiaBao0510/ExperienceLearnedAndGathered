## **Lớp HttpMessageHandler:**
- Lớp **HttpMessageHandler** là một lớp trừu tượng được thư viện .NET Core triển khai các lớp như: _Delegatinghandler, HttpMessageHandler, HttpClientHandler,..._ các lớp triển khai này(hoặc nếu tự xây dụng lớp triển khai HttpMessage Handler) thì phải nạp chồng phương thức SendAsync.
```
protected Task<HttpResponseMessage> SendAsync (HttpRequestMessage request, CancellationToken cancellationToken);
```
- Các lớp triển khai HttpMessageHandler dùng để khởi tạo HttpClient, lúc này HttpClient thực hiện gửi truy vấn (SendAsync) thì SendAsync của handler sẽ thực thi.
- **HttpClientHandler** là một lớp triển khai từ **HttpMessageHandler**, nó thực hiện cuối cùng trong chuỗi các Handler nếu có thể để thực sự gửi truy vấn HTTP.
- Một số thuộc tính trong **HttpClientHandler**:

| **Thuộc tính**         | Mô tả                                                                                                                                                                                                                                                                                                                                                                                                                    |
| ---------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| AllowAutoRedirect      | Thuộc tính, mặc định là `true`, để thiết lập tự động chuyển hướng. Ví dụ truy vấn đến URI có chuyển hướng đến đích mới (301) thì - HttpClient sẽ tự động chuyển hướng truy vấn đến đó.                                                                                                                                                                                                                                   |
| AutomaticDecompression | Thuộc tính thuộc tính để handler tự động giải nén / nén nội dung HTTP, nó thuộc kiểu enum `DecompressionMethods` gồm có:<br><br>- `DecompressionMethods.None` không sử dụng nén<br>- `DecompressionMethods.GZip` dùng thuật toán gZip<br>- `DecompressionMethods.Deflate` dùng thuật toán nén deflate<br><br>Ví dụ có thể gán:  ```AutomaticDecompression = DecompressionMethods.Deflate \| DecompressionMethods.GZip``` |
| UseCookies             | Mặc định là `true`: cho phép sử dụng thuộc tính CookieContainer để lưu các Cookie của server khi respone trả về, cũng như tự động gửi Cookie khi gửi truy vấn.                                                                                                                                                                                                                                                           |
| CookieContainer        | Thuộc tính thuộc lớp `CookieContainer`, nó lưu các cookie.                                                                                                                                                                                                                                                                                                                                                               |
- ## **Khuyến khích nên sử dụng SockerHttpHandler ,Vì nó có thể chạy được đa nền tảng**. 

## **DelegatingHandler cho HttpClient**
- **DelegatingHandler** (cũng triển khải từ HttpMessageHandler) là một handler đặc biệt, nó như một MiddleWare để tạo ra một pipeline (chuỗi các handler). Mỗi đối tượng DelegatingHandler có một thuộc tính `InnerHandler` (kiểu HttpMessageHandler), **phải được gán** bằng một đối tượng `SocketsHttpHandler`, `HttpClientHandler` hoặc `DelegatingHandler`...
- Thiết lập `InnerHandler` qua phương thức khởi tạo lớp DelegatingHandler. Khi thực hiện truy vấn `SendAsync` thì nó tiếp tục gọi `SendAsync` trong InnerHandler, cứ như vậy nó sẽ tạo thành chuỗi.
- Nếu InnerHandler không phải là một DelegatingHandler khác thì InnerHandler đó là handler dưới cùng của chuỗi handler. Request - respone sẽ đi qua chuỗi handler từ trên cùng xuống dưới khi truy vấn và ngược lại khi trả về.
![[Pasted image 20240810224447.png]]

