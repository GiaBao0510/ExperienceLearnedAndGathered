// Online C# Editor for free
// Write, Edit and Run your C# code using C# Online Compiler

using System;
using System.Collections;
using System.Collections.Generic;
using System.Numerics;
using System.Security.Cryptography;
using System.Text;
using Google.Apis.Auth.OAuth2;
using Google.Apis.Drive.v3;
using Google.Apis.Services;
using Google.Apis.Upload;
using Google.Apis.Util.Store;
using System.IO;
using System.Threading;

public class HelloWorld
{
    static string[] Scopes = { DriveService.Scope.DriveFile };
    static string ApplicationName = "VoteSecure";

    public static async Task Main(string[] args)
    {
         UserCredential credential;

        // Đường dẫn tệp credentials.json
        using (var stream = new FileStream("credentials.json", FileMode.Open, FileAccess.Read))
        {
            // Tạo hoặc lấy token.json
            string credPath = "token.json";
            credential = await GoogleWebAuthorizationBroker.AuthorizeAsync(
                GoogleClientSecrets.FromStream(stream).Secrets,
                Scopes,
                "user",
                CancellationToken.None,
                new FileDataStore(credPath, true));
            
            Console.WriteLine("Credential file saved to: " + credPath);
        }

        // Khởi tạo dịch vụ Drive
        var service = new DriveService(new BaseClientService.Initializer()
        {
            HttpClientInitializer = credential,
            ApplicationName = ApplicationName,
        });

        // Đoạn mã này sẽ in ra thông báo xác nhận token đã được tạo và lưu
        Console.WriteLine("Google Drive API Service is created successfully.");
    }
}