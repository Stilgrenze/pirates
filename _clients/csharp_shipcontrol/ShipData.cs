public class ShipData
{
	public string ShipName { get; set; }
    public string Id { get; set; }
    public int Cannons { get; set; }
    public int Sight { get; set; }
    public int Speed { get; set; }

    public ShipData() {
    	ShipName = "";
    	Id = "";
    	Cannons = 0;
    	Sight = 0;
    	Speed = 0;
    }
}
