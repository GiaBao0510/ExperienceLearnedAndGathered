using System;
using System.Linq;
using System.Net;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Threading;
using System.Threading.Tasks;
using System.IO;
using System.Text;

namespace Buoi35_HttpMessageHandler_CookieContainer
{
    public class MyHttpClientHandler : HttpClientHandler
    {
        public MyHttpClientHandler(CookieContainer _cookieContainer)
        {
            CookieContainer = _cookieContainer;     // Thay thế cho CookieContainer mặc định
            AllowAutoRedirect = false;              // Không cho phép tự động chuyển hướng
            AutomaticDecompression = DecompressionMethods.Deflate | DecompressionMethods.GZip;
            UseCookies = true;
        }
        // Bắt buộc phải nạp chồng
        protected override async Task<HttpResponseMessage> SendAsync(HttpRequestMessage request, CancellationToken _cancellationToken)
        {
            Console.WriteLine("Bat dau ket noi:" + request.RequestUri.ToString());
            // Thực hiện truy vấn đến server
            var response = await base.SendAsync(request, _cancellationToken);
            Console.WriteLine("Hoan thanh tai du lieu");
            return response;
        }
    }

    public class ChangeUri : DelegatingHandler
    {
        public ChangeUri(HttpMessageHandler innerHandler) : base(innerHandler) { }

        protected override Task<HttpResponseMessage> SendAsync(HttpRequestMessage request,
                                                               CancellationToken cancellationToken)
        {
            var host = request.RequestUri.Host.ToLower();
            Console.WriteLine($"Check in ChangeUri - {host}");
            if (host.Contains("google.com"))
            {
                // Đổi địa chỉ truy cập từ google.com sang github
                request.RequestUri = new Uri("https://github.com/");
            }
            // Chuyển truy vấn cho base (thi hành InnerHandler)
            return base.SendAsync(request, cancellationToken);
        }
    }

    public class DenyAccessFacebook : DelegatingHandler
    {
        public DenyAccessFacebook(HttpMessageHandler innerHandler) : base(innerHandler) { }

        protected override async Task<HttpResponseMessage> SendAsync(HttpRequestMessage request,
                                                              CancellationToken cancellationToken)
        {
            var host = request.RequestUri.Host.ToLower();
            Console.WriteLine($"Check in DenyAccessFacebook - {host}");
            if (host.Contains("facebook.com"))
            {
                var response = new HttpResponseMessage(HttpStatusCode.OK);
                response.Content = new ByteArrayContent(Encoding.UTF8.GetBytes("Không được truy cập"));
                return await Task.FromResult<HttpResponseMessage>(response);
            }
            // Chuyển truy vấn cho base (thi hành InnerHandler)
            return await base.SendAsync(request, cancellationToken);
        }
    }

    class Program
    {
        static async Task Main(string[] args)
        {
            //1. Sử dụng SockerHttpHandler
            Console.WriteLine("\n\t======== SockerHttpHandler ======");
            string url = "https://postman-echo.com/post";
            string urlFB = "https://www.facebook.com/";

            //1.1 Tạo cookie controller
            var cookies = new CookieContainer();

            //2. Tạo chuỗi handler()
            var bottomHandler = new MyHttpClientHandler(cookies);
            var changeUrriHandler = new ChangeUri(bottomHandler);
            var denyAccessFB = new DenyAccessFacebook(changeUrriHandler); // Sửa lỗi tại đây

            using var httpClient = new HttpClient(denyAccessFB);
            using var httpRequestMessage = new HttpRequestMessage(HttpMethod.Post, urlFB);
            httpRequestMessage.Headers.Add("User-Agent", "Mozilla/5.0");

            string jsonParameters = @"{ ""Key1"":""5f1A4"", ""Key2"":""5fT71"", ""Key3"":""8ERk4"" }";
            var Content0 = new StringContent(jsonParameters, Encoding.UTF8, "application/json");
            httpRequestMessage.Content = Content0;

            var response0 = await httpClient.SendAsync(httpRequestMessage);
            var html0 = await response0.Content.ReadAsStringAsync();

            // Lấy cookies từ các uri truyền vào
            Console.WriteLine("\nCookies:");
            cookies.GetCookies(new Uri(url)).ToList().ForEach(e =>
            {
                Console.WriteLine($"{e.Name} : {e.Value}");
            });

            Console.WriteLine(html0);

            Console.WriteLine("\n\t======== DelegatingHandler  ======");
        }
    }
}
