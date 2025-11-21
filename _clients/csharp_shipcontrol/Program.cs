using System;
using System.Threading.Tasks;

public static class Config
{
    public static string URL = "";
    public static string Name = "C_Sharp_Team";
    public static string Secret = "supersecret";
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