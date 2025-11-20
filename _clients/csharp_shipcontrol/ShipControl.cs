using System;
using System.Net.WebSockets;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using Newtonsoft.Json;

public class ShipControl
{
    public string ShipName { get; set; }
    public string Id { get; set; }
    public int Cannons { get; set; }
    public int Sight { get; set; }
    public int Speed { get; set; }

    private ClientWebSocket _socket;
    private Uri _uri;
    private CancellationTokenSource _cts;

    public ShipControl(ShipData shipData)
    {
        ShipName = shipData.ShipName;
        Id = shipData.Id;
        Cannons = shipData.Cannons;
        Sight = shipData.Sight;
        Speed = shipData.Speed;

        _socket = new ClientWebSocket();
        _cts = new CancellationTokenSource();

        _uri = new Uri($"ws://{Config.URL}shipControl/id/{Config.Name}/{Config.Secret}");

        // Start WebSocket connection
        _ = Connect();
    }

    private async Task Connect()
    {
        while (!_cts.Token.IsCancellationRequested)
        {
            try
            {
                await _socket.ConnectAsync(_uri, _cts.Token);
                Console.WriteLine($"Ship {Id} ready");

                // Start receiving messages
                await ReceiveLoop();
            }
            catch (Exception ex)
            {
                Console.Error.WriteLine($"Ship {Id} connection error: {ex.Message}. Reconnecting in 5s...");
                await Task.Delay(5000);
                _socket.Dispose();
                _socket = new ClientWebSocket();
            }
        }
    }

    private async Task ReceiveLoop()
    {
        var buffer = new byte[8192];

        while (_socket.State == WebSocketState.Open)
        {
            var result = await _socket.ReceiveAsync(new ArraySegment<byte>(buffer), _cts.Token);

            if (result.MessageType == WebSocketMessageType.Close)
            {
                Console.WriteLine($"Ship {Id} disconnected. Attempting reconnect...");
                await _socket.CloseAsync(WebSocketCloseStatus.NormalClosure, string.Empty, _cts.Token);
                break;
            }

            var message = Encoding.UTF8.GetString(buffer, 0, result.Count);
            await OnMessage(message);
        }
    }

    private async Task OnMessage(string json)
    {
        var info = JsonConvert.DeserializeObject(json);
        Console.WriteLine(info);

        // TODO: implement your game logic here (within 100 ms)

        var command = new
        {
            MoveX = -1,
            MoveY = 0,
            Attack = new string[] { "PORT_wQkGBSYrSQPECsJJ" }
        };

        await SendCommand(command);
    }

    private async Task SendCommand(object command)
    {
        if (_socket.State != WebSocketState.Open) return;

        var json = JsonConvert.SerializeObject(command);
        var bytes = Encoding.UTF8.GetBytes(json);
        var segment = new ArraySegment<byte>(bytes);

        await _socket.SendAsync(segment, WebSocketMessageType.Text, true, _cts.Token);
    }

    public async Task Close()
    {
        _cts.Cancel();
        if (_socket.State == WebSocketState.Open)
        {
            await _socket.CloseAsync(WebSocketCloseStatus.NormalClosure, "Closing", CancellationToken.None);
        }
        _socket.Dispose();
    }
}
