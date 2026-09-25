using System;
using Buoi41_TaoModel_EntityFramwork.Model;
namespace Buoi41_TaoModel_EntityFramwork;

class Program
{
    static async Task Main(string[] args)
    {
        ShopContext context = new ShopContext();
        await context.CreateDB();

    }
}
