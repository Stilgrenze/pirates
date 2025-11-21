using System;
using WebSocketSharp;

public class ShipControl
{
	public string name { get; set; }
    public string id { get; set; }
    public int cannons { get; set; }
    public int sight { get; set; }
    public int speed { get; set; }
    private WebSocket websocket;

    public ShipControl(ShipData shipData)
    {
    	this.name = shipData.ShipName;
        this.id = shipData.Id;
        this.cannons = shipData.Cannons;
        this.sight = shipData.Sight;
        this.speed = shipData.Speed;

        // WebSocket connection
        this.websocket = new WebSocket("ws://" + Config.URL + "shipControl/" + this.id + "/" + Config.Name + "/" + Config.Secret);
        this.websocket.OnOpen += (sender, e) => this.OnOpen(sender, e);
        this.websocket.OnClose += (sender, e) => this.OnClose(sender, e);
        this.websocket.OnMessage += (sender, e) => this.OnMessage(sender, e);
        this.websocket.OnError += (sender, e) => this.OnError(sender, e);
        this.websocket.Connect();
    }

    private void OnOpen(object sender, EventArgs e)
    {
        Console.WriteLine("Ship " + this.id + " ready");
    }

    private void OnClose(object sender, CloseEventArgs e)
    {
        // https://en.wikipedia.org/wiki/IP_over_Avian_Carriers
        Console.WriteLine("Ship " + this.id + " no more pigeons to send");
    }

    private void OnError(object sender, WebSocketSharp.ErrorEventArgs e)
    {
        // https://en.wikipedia.org/wiki/IP_over_Avian_Carriers
        Console.Error.WriteLine("Ship " + this.id + " pigeons died: " + e.Message);
    }

    private void OnMessage(object sender, MessageEventArgs e)
    {
        var info = Newtonsoft.Json.JsonConvert.DeserializeObject(e.Data);
        Console.WriteLine(info);

        // TODO Implement your code for your ships here!
        // You have only 100ms time to react! Your Ship will disappear after 60 Second Idle
        this.websocket.Send(
            Newtonsoft.Json.JsonConvert.SerializeObject(new
            {
                // TODO Movement here
                MoveX = -1,
                MoveY = 0,
                // TODO Insert your attack here
                Attack = new string[] { "PORT_wQkGBSYrSQPECsJJ" }
            })
        );
    }
}
