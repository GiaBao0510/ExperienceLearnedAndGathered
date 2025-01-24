using System;
using System.Linq;
using System.Net;
using System.Net.NetworkInformation;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Threading;
using System.Threading.Tasks;
using System.IO;
using System.Text;

namespace Buoi34_SuDungHttpClient
{
	class Program
	{
		//Hiển thị thông tin header sau khi nhận phản hồi
		static void ShowHeaders(HttpResponseHeaders tieude)
		{
			Console.WriteLine("\n\t=== Cac Headers: ===");
			foreach(var item in tieude)
			{
				Console.WriteLine($"{item.Key} : {item.Value}");
			}
		}

		// Phương thức này dùng để trả nội dung HTML của trang web
		public static async Task<string> GetWebContent(string url)
		{
			// Khởi tạo HttpClient - Thêm từ khóa using để đối tượng này tự hủy khi thoát ra khỏi phương thức
			using var httpclient = new HttpClient();

			try
			{
				// Thiết lập các Header nếu cần
				httpclient.DefaultRequestHeaders.Add("Accept", "text/html,application/xhtml+xml+json");

				// Thực hiện truy vấn GET
				HttpResponseMessage response = await httpclient.GetAsync(url);

				//Hiển thị các headers
				ShowHeaders(response.Headers);

				// Đọc nội dung từ response
				string html = await response.Content.ReadAsStringAsync();

				return html;
			}
			catch(Exception e)
			{
				Console.WriteLine($"Loi trong phuong thuc lay noi dung tu url: {e.Message}");
				return "404";
			}
		}

		//Đọc HttpResponseMessage bằng ReadAsByteArrayAsync
		public static async Task<byte[]> DownloadDataBytes(string url)
		{
			using var httpClient = new HttpClient();
			try
			{
				HttpResponseMessage response = await httpClient.GetAsync(url);
				var bytes = await response.Content.ReadAsByteArrayAsync();
				return bytes;

			}
			catch (Exception e)
			{
				Console.WriteLine($"Loi: {e.Message}");
				return null;
			}
		}

		//Đọc HttpResponseMessage bằng ReadAsStreamAsync
		//Dùng để đọc từ byte 1, từng mảng bute một hoặc từng khối byte 1.
		public static async Task DownloadStream(string url, string filename)
		{
			HttpClient httpClient = new HttpClient();
			try
			{
				var httpResponseMessage = await httpClient.GetAsync(url);
				using var stream = await httpResponseMessage.Content.ReadAsStreamAsync();

				//Tạo vùng đệm byte - dùng để đọc các byte từ stream trên
				int SIZEBUFFER = 500;
				var buffer = new byte[SIZEBUFFER];
				
				//Tạo vòng lạp để đọc từ đầu đến cuối stream
				bool endread = false;

					//Khi đọc được các byte thì nó lộ ra cá stream khác nên ta tạo các stream mới
				using var streamwrite = File.OpenWrite(filename);
				do
				{
					//ReadAsync: mảng byte đọc được sẽ ghi vào vùng đệm buffer
					int numberbyte = await stream.ReadAsync(buffer, 0, SIZEBUFFER);

					if (numberbyte == 0)
					{
						endread = true;
					}
					else
					{   //Đổ các byte từ chỉ số 0 đến chỉ số numberbyte
						streamwrite.WriteAsync(buffer, 0, numberbyte);
					}
				} while (!endread);
			
			
			}catch(Exception e)
			{
				Console.WriteLine($"Loi: {e.Message}");
			}
		}


		static async Task Main(string[] args)
		{
				// 1. Ví dụ hiển thị các giá trị thuộc tính Uri
			Console.WriteLine("\n\tGia Tri cac thuoc tinh Uri");
			string ImageFile = "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQZjMoYxBIgTnNkvxhHJ-PrCtrEjiOrt1-47Q&s";
			string ImageWindow = "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSt_H_F63j7JXjHQo5eiqbjRilzWQIdI9kbNQ&s";
			string url = "https://xuanthulab.net/lap-trinh/csharp/?page=3#acff";
			var uri = new Uri(url);         // class uri
			var uritype = uri.GetType();    // Sử dụng GetType() để lấy kiểu của đối tượng
			uritype.GetProperties().ToList().ForEach(
				property =>
				{
					Console.WriteLine($"{property.Name,15} {property.GetValue(uri)}");
				}
			);
			Console.WriteLine($"Segment: {string.Join(", ", uri.Segments)}");

				// 2. Lớp tính DNS và lớp IPHostEntry
			Console.WriteLine("\n\tClass DNS and Class IPHostEntry");
			string uri2 = "https://www.bootstrapcdn.com/";
			uri = new Uri(uri2);
			var hostEntry = Dns.GetHostEntry(uri.Host);  // Phân giải host hoặc IP thành đối tượng IPHostEntry
			Console.WriteLine($"Local host name: {Dns.GetHostName()}");
			Console.WriteLine($"Host: {uri.Host} co cac IP.");
			hostEntry.AddressList.ToList().ForEach(ip => Console.WriteLine(ip));

				// 3. Class Ping
			Console.WriteLine("\n\tClass Ping");
			var ping = new Ping();
			var pingReply = ping.Send("google.com.vn"); // Gửi đến tên miền

			Console.WriteLine(pingReply.Status);        // Lấy trạng thái trả về
			if (pingReply.Status == IPStatus.Success)
			{
				Console.WriteLine("Thoi gian phan hoi: " + pingReply.RoundtripTime);
				Console.WriteLine("Dia chi phan hoi: " + pingReply.Address);
			}

				// Tạo truy vấn GET bất đồng bộ với HttpClient
			Console.WriteLine("\n\tTruy van bat dong bo:");
			var html = await GetWebContent(url);
			Console.WriteLine(html);

				//Đọc httpresponseMessage bằng ReadAsByteArrayAsync
			Console.WriteLine("\n\tReadAsByteArrayAsync");
			var bytes = await DownloadDataBytes(ImageFile);
			string filename = "viralhit1.png";

			using var streams = new FileStream(filename, FileMode.Create, FileAccess.Write, FileShare.None );

				//Đổ dữ liệu byte vào
			Console.WriteLine("\n\tReadAsStreamAsync:");
			await DownloadStream(ImageWindow, "window.png");

				//Phương thức SendAsync .Mạnh hơn và tổng quát hơn
			Console.WriteLine("\n\tSendAsync");
			using var httpClient = new HttpClient();

			var httpMessageRequest = new HttpRequestMessage();
			httpMessageRequest.Method = HttpMethod.Get;     //Thiết lập phương thức sẽ thực hiện
			httpMessageRequest.RequestUri = new Uri("https://mangaplus.shueisha.co.jp/updates");      //THiết lập uri
			httpMessageRequest.Headers.Add("Accept", "text/html,application/xhtml+xml+json"); //Thiết lập header

				//Gửi thông diệp yêu cầu đến server
			var httpResponseMessage = await httpClient.SendAsync(httpMessageRequest);

			var html1 = await httpResponseMessage.Content.ReadAsStringAsync();
			Console.WriteLine(html1);

				//post dữ liệu form HTML lên server với FormUrlEncodedContent
				//Thực hiện thao tác post dữ liệu giống như submit data
			var httpMessageRequest1 = new HttpRequestMessage(HttpMethod.Post, "https://postman-echo.com/post");

			httpMessageRequest1.Headers.Add("Accept", "text/html,application/xhtml+xml+json");

				//Thiết lập nội dung html bằng lớp FormUrlEncodedContent()
				//Tham số khởi tạo cho đối tượng này là KeyValuePair<string,string>
			Console.WriteLine("\n\tFormUrlEncodedContent:");
			var parameters = new List<KeyValuePair<string, string>>();
			parameters.Add(new KeyValuePair<string, string> ( "Key1", " 5sdgrr1SA6t"));
			parameters.Add(new KeyValuePair<string, string> ( "Key2", "F54dfg210dfgd0" ));
			parameters.Add(new KeyValuePair<string, string> ( "Key3", "F54fdfg8r210dfgd0" ));

			var Content = new FormUrlEncodedContent(parameters);
			httpMessageRequest1.Content = Content;

			var phanHoiHttp1 = await httpClient.SendAsync(httpMessageRequest1);

			ShowHeaders(phanHoiHttp1.Headers);
			var html2 = await phanHoiHttp1.Content.ReadAsStringAsync();
			Console.WriteLine(html2);

	            //Post dữ liệu nội dung json với StringContent
            Console.WriteLine("\n\t Post json data with StringContent:");
			var httpMessageRequest2 = new HttpRequestMessage(HttpMethod.Post, "https://postman-echo.com/post");
			string jsonData = @"
				""Key101"":""8dfg2rKD5f"",
				""Key102"":""8d8rTrKD5f"",
				""Key103"":""8df75dhfT5f"",
			";
			var Content2 = new StringContent(jsonData, Encoding.UTF8, "application/json");
			httpMessageRequest2.Content = Content2;

			var phanHoiHttp2 = await httpClient.SendAsync(httpMessageRequest2);
			ShowHeaders(phanHoiHttp1.Headers);
			var html3 = await phanHoiHttp2.Content.ReadAsStringAsync();
			Console.WriteLine(html3);

			//Upload dữ liệu phức tạp (upload file) với MultipartFormDataContent, StreamContent
			Console.WriteLine("\n\t MultipartFormDataContent:");
			var httpMessageRequest3 = new HttpRequestMessage(HttpMethod.Post, "https://postman-echo.com/post");

			var Content3 = new MultipartFormDataContent();

			//Uploadfile
			Stream FileStream = File.OpenRead("Text.txt");//Đọc file
			var fileUpload = new StreamContent(FileStream); //Chứa nội dung đính kèm

			Content3.Add(fileUpload,"fileupload","abc.html");
			Content3.Add(new StringContent("value"),"key115");

			httpMessageRequest3.Content = Content3;

			var PhanHoiHttp3 = await httpClient.SendAsync(httpMessageRequest3);
			ShowHeaders(PhanHoiHttp3.Headers);

			var html4 = await PhanHoiHttp3.Content.ReadAsStringAsync();
			Console.WriteLine(html4);

        }
	}
}
