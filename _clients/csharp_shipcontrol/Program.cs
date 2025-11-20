using System;
using System.Threading.Tasks;

public static class Config
{
    public static string URL = "162.55.58.253:1337/";
    public static string Name = "CSharp_Team";
    public static string Secret = "Test123";
}

class Program
{
    private static ShipManager manager = new ShipManager();

    static async Task Main(string[] args)
    {
    	await manager.Register(Config.Name, Config.Secret);

    	await manager.BuyShip("Demoship", 1, 1, 1);

  		await manager.GetShips(init: true);

        Console.WriteLine("Press any key to exit...");
        Console.ReadKey();
    }
}