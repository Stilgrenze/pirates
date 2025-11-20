using System;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using Newtonsoft.Json;

public class ShipManager
{
    private static HttpClient httpClient = new HttpClient();
    public List<ShipData>? Ships { get; private set; }
    public PlayerData? Team { get; private set; }
    public string Error { get; private set; }
    public bool Ready { get; private set; }

    public ShipManager() {
    	Error = "";
    	Ready = false;
    }

    public async Task GetShips(bool init)
    {
        try
        {
            var response = await httpClient.GetAsync($"http://{Config.URL}ships/{Config.Name}/{Config.Secret}");
            if (response.IsSuccessStatusCode)
            {
                var responseBody = await response.Content.ReadAsStringAsync();
                Ships = JsonConvert.DeserializeObject<List<ShipData>>(responseBody);
                if (Ships == null) {
                	Error = "could not DeserializeObject GetShips";
                	return;
                }
                if (init)
                {
                    foreach (var ship in Ships)
                    {
                        new ShipControl(ship);
                    }
                }
            }
            else
            {
                Error = await response.Content.ReadAsStringAsync();
            }
        }
        catch (Exception ex)
        {
            Error = ex.Message;
        }
    }

    public async Task GetPlayer()
    {
        try
        {
            var response = await httpClient.GetAsync($"http://{Config.URL}player/{Config.Name}/{Config.Secret}");
            if (response.IsSuccessStatusCode)
            {
                var responseBody = await response.Content.ReadAsStringAsync();
                Team = JsonConvert.DeserializeObject<PlayerData>(responseBody);
                // Update local storage (if needed, you can use a local file or database)
            }
            else
            {
                Error = await response.Content.ReadAsStringAsync();
                Ready = false;
            }
        }
        catch (Exception ex)
        {
            Error = ex.Message;
            Ready = false;
        }
    }

    public async Task BuyShip(string name, int cannons, int sight, int speed)
    {
        try
        {
            var requestBody = new
            {
                ShipName = name,
                Cannons = cannons,
                Sight = sight,
                Speed = speed,
                Player = new { Name = Config.Name, Secret = Config.Secret }
            };
            var content = new StringContent(JsonConvert.SerializeObject(requestBody), Encoding.UTF8, "application/json");
            var response = await httpClient.PostAsync($"http://{Config.URL}buyShip", content);
            if (response.StatusCode == System.Net.HttpStatusCode.Created)
            {
                var responseBody = await response.Content.ReadAsStringAsync();
                var shipData = JsonConvert.DeserializeObject<ShipData>(responseBody);
                if (shipData == null) {
                	Error = "coud not DeserializeObject ShipData";
                	return;
				}
                new ShipControl(shipData);
                await GetShips(false);
            }
            else
            {
                Error = await response.Content.ReadAsStringAsync();
            }
        }
        catch (Exception ex)
        {
            Error = ex.Message;
        }
    }

    public async Task Register(string name, string secret)
    {
        try
        {
            var requestBody = new { Name = name, Secret = secret };
            var content = new StringContent(JsonConvert.SerializeObject(requestBody), Encoding.UTF8, "application/json");
            var response = await httpClient.PostAsync($"http://{Config.URL}registerPlayer", content);
            if (response.StatusCode == System.Net.HttpStatusCode.Created)
            {
                var responseBody = await response.Content.ReadAsStringAsync();
                Team = JsonConvert.DeserializeObject<PlayerData>(responseBody);
                // Save it somewhere if you want it restart ready
                Ready = true;
            }
            else
            {
                Error = await response.Content.ReadAsStringAsync();
            }
        }
        catch (Exception ex)
        {
            Error = ex.Message;
        }
    }
}

public class PlayerData
{
    public int Gold { get; set; }
    public int GoldSpent { get; set; }
}
