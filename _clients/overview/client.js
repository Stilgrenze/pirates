// TODO Insert URL of the Server here
const URL = 'localhost:1337/';

function InitShip(shipData) {
    // TODO you can change your ship controls here if you want
    new Ship(shipData);
}

// Examle Ship Implementation
class Ship {
    constructor(shipData) {
        this.id = shipData.id;
        this.cannons = shipData.cannons;
        this.sight = shipData.sight;
        this.speed = shipData.speed;

        // Websocket connection
        this.websocket = new WebSocket('ws://'+URL+'shipControl/'+this.id+'/'+TEAM+'/'+SECRET);
        this.websocket.onopen = (evt) => { this.onOpen(evt) };
        this.websocket.onclose = (evt) => { this.onClose(evt) };
        this.websocket.onmessage = (evt) => { this.onMessage(evt) };
        this.websocket.onerror = (evt) => { this.onError(evt) };
    }

    onOpen(evt) {
        console.log('Ship ' + this.id + ' ready');
    }

    onClose(evt) {
        // https://en.wikipedia.org/wiki/IP_over_Avian_Carriers
        console.log('Ship ' + this.id + ' no more pigeons to send');
    }

    onError(evt) {
        // https://en.wikipedia.org/wiki/IP_over_Avian_Carriers
        console.error('Ship ' + this.id + ' pigeons died', evt);
    }

    onMessage(evt) {
        let info = JSON.parse(evt.data);
        console.log(info);

        // TODO Implement your code for your ships here!

        // You have only 100ms time to react! Your Ship will dissapear after 60 Second Idle
        this.websocket.send(
            JSON.stringify({
                // TODO Movement here
                MoveX: 1,
                MoveY: 0,
                // TODO Insert your attack here
                Attack: ["PORT_wQkGBSYrSQPECsJJ"]
            })
        );
    }
}