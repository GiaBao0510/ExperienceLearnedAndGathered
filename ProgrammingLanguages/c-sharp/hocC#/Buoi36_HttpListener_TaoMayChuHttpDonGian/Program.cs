using System.Net;
using System.Text;
using System.Text.Json;
using System.Text.Json.Serialization;
using Newtonsoft.Json;

namespace Buoi36_HttpListener_TaoMayChuHttpDonGian;

class Program
{
    //Tạo lớp HttpServer
    class MyHttpServer
    {
        private HttpListener _listener;

        //Tạo ra phương thức khởi tạo, có tham số là mảng chuỗi
        public MyHttpServer(string[] prifixes)
        {
                //Đầu tiên là kiểm tra xem nó có hỗi trọ HttpListener hay không
            if (!HttpListener.IsSupported)
            {
                throw new Exception("HttpListener is not supported");
            }

            _listener = new HttpListener();

            //Duyệt qua các prifixes. Để thêm tường đường dẫn với các cổng khác nhau
            foreach (string prefix in prifixes) _listener.Prefixes.Add(prefix);
        }

        //Phương thức dùng để khởi chạy
        public async Task Start()
        {
            _listener.Start();
            Console.WriteLine("Http Server Started");
            do
            {
                try
                {
                    Console.WriteLine(DateTime.Now.ToLongTimeString() + " waiting a client connect...");
                    
                    //Một client kết nối đến
                    var conText = await _listener.GetContextAsync();

                    //Xử lý nội dung trả về cho client
                    await ProcessRequest(conText);

                    Console.WriteLine(DateTime.Now.ToLongTimeString() + " Client connected.");
                }
                catch (Exception e)
                {
                    Console.WriteLine("Error: "+e.Message);
                }
                Console.WriteLine("...");
            } while (_listener.IsListening);
        }

        //Chuyên xử lý để trả về phản hồi khi người dùng kết nối được
        async Task ProcessRequest(HttpListenerContext conetext)
        {
            HttpListenerRequest req = conetext.Request;
            HttpListenerResponse res = conetext.Response;

            Console.WriteLine($"{req.HttpMethod} {req.RawUrl} {req.Url.AbsolutePath}");

            //Xử lý nội dung trả về cho client
            var outputStream = res.OutputStream;

            switch (req.Url.AbsolutePath)
            {
                case "/requestinfomation":        //Trả về thông tin máy chủ
                    {
                        //Gửi thông tin về cho client
                        res.Headers.Add("content-Type", "text-html");
                        res.StatusCode = (int)HttpStatusCode.OK;

                        string responseString = this.GenerateHTML(req);
                        var buffer = Encoding.UTF8.GetBytes(responseString);
                        res.ContentLength64 = buffer.Length;
                        await outputStream.WriteAsync(buffer, 0, buffer.Length);
                    }
                    break;
                case "/":       //Trả về trang chủ
                    {
                        var buffer = Encoding.UTF8.GetBytes("\n == Welcome to my website ==\n");
                        res.ContentLength64 = buffer.Length;
                        await outputStream.WriteAsync(buffer, 0, buffer.Length);
                    }
                    break;
                case "/json":   //Trả về dữ liệu json
                    {
                        res.Headers.Add("Content-Type", "application/json");
                        var product = new
                        {
                            Name = "Mat kinh",
                            Price = 150000,
                        };

                        //Chuyển đối tượng thành chuỗi json
                        var json = JsonConvert.SerializeObject(product);
                        var buffer = Encoding.UTF8.GetBytes(json);
                        res.ContentLength64 = buffer.Length;
                        await outputStream.WriteAsync(buffer, 0, buffer.Length);
                    }
                    break;
                case "/anhtusi":    //Trả về dữ liệu ảnh
                    {
                        res.Headers.Add("Content-Type", "image/jpg");
                        var buffer = await File.ReadAllBytesAsync("TuSi.jpg");
                        res.ContentLength64 = buffer.Length;
                        await outputStream.WriteAsync(buffer, 0, buffer.Length);
                    }
                    break;
                case "/stop":   //dừng hoạt động
                    {
                        _listener.Stop();
                        Console.WriteLine("\n\t Stop server");
                    }
                    break;
                default:    //không tìm thấy 404
                    {
                        var buffer = Encoding.UTF8.GetBytes("\n Not Found 404.\n");
                        res.ContentLength64 = buffer.Length;
                        await outputStream.WriteAsync(buffer, 0, buffer.Length);
                    }
                    break;
            }
            outputStream.Close();
        }

        //Tạo nội dung hmtl trả về cho client (HTML chứa thông tin Request)
        public string GenerateHTML(HttpListenerRequest res)
        {
            string format = @"<!DOCTYPE html>
                            <html lang=""en""> 
                                <head>
                                    <meta charset=""UTF-8"">
                                    {0}
                                 </head> 
                                <body>
                                    {1}
                                </body> 
                            </html>";
            string head = "<title>Test Server</title>";
            var body = new StringBuilder();
            body.Append("<h1>Request Info</h1>");
            body.Append("<h2>Request Header</h2>");

            //Header Infor
            var headers = from Key in res.Headers.AllKeys
                         select $"<div>{Key}: {string.Join(", ", res.Headers.GetValues(Key))} </div>";
            body.Append(string.Join("", headers));

            //Extract request properties
            body.Append("<h2>Request properties</h2>");
            var properties = res.GetType().GetProperties();
            foreach (var property in properties) {
                var name_property = property.Name;
                string values_property;
                try
                {
                    values_property = property.GetValue(res).ToString();
                }
                catch (Exception ex) {
                    values_property = ex.Message;
                }
                body.Append($"<div>{name_property} : {values_property}</div>");
            }

            string html = string.Format(format, head, body.ToString());
            return html;

        }

    }

    static async Task Main(string[] args)
    {
        /*  == Ví dụ 1: ==
        
        //1.Kiểm tra máy có hỗ trợ HttpListener hay không
        if (HttpListener.IsSupported)
        {
            Console.WriteLine("Suported HttpListener");
        }
        else
        {
            Console.WriteLine("Not suported HttpListener");
            throw new Exception("Not suported HttpListener.");
        }

        var server = new HttpListener();
        server.Prefixes.Add("http://*:8081/");  //Chấp nhận mọi kết nối từ cổng 8080. Mặc định lắng nghe trên cổng 80

        //Bắt đầu cho phép nhận các yêu cầu
        server.Start();

        Console.WriteLine("\n\t== Server HTTP start ==");

        //Nhận thông tin gửi đến từ client .Nếu không gọi thêm phương thức GetContextAsync khác thì nó 
        //sẽ tự động kết thúc việc lắng nghe. Vì vậy ta cần đưa nó vào vòng lặp Do
        do
        {
            Console.WriteLine("Client connected");

            var conetext_1 = await server.GetContextAsync();
            var response_1 = conetext_1.Response; //Thông điệp httpResponse trả về
            var outputStream_1 = response_1.OutputStream; //Viết dữ liệu ra và client có thể nhận được

            //Trả về các header của response
            response_1.Headers.Add("content-type", "text/html");

            var html = "<h1>Hello World</h1>";
            var bytes = Encoding.UTF8.GetBytes(html);
            await outputStream_1.WriteAsync(bytes, 0, bytes.Length);
            outputStream_1.Close(); //Đóng

        } while (server.IsListening);

        */
            // == ví dụ 2 ==

        var server = new MyHttpServer(new string[] {"http://*:8080/"});
        await server.Start();


    }
}
