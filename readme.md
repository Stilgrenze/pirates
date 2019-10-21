# Pirate Programming Game
The game is a programming battle game, it can be written in any language which supports websockets.
There is an example implementations in javascript. It's intended for programmers with a bit programming experience.

The goal of the Game is to make as much gold as possible!\
Each Team can control as much ships as they can buy.
The ships are controlled over `Websockets`, the example implementation is ath the bottom of the `shipcontrol/index.html`

The game is written in Go and Javascript with Vue.js, feel free to extend/change

## Teams
Before you can send Pirate Ships out in the Ocean, you must register a Team.
Each Team can buy as much ships as it has money for it.

## Ports
There are Ports out in the Ocean, each Port generates Gold.
As a Pirate you want to raid as much ports as possible.
If a Port is raided, it will be locked some time before a new attack is possible.

## Ships
As a Team you can buy Pirate Ships. Each Ship has 3 attributes.
Each Ship starts with 1 Point on each attribute, but you can buy additional attributes.
You can attack other ships! Each sunken ship will give you 1000 Gold (reward for sinking a pirate ship).

### Speed
How fast the Ship can Travel per Tick. If you want to move the Ship faster than its speed, it will not move!

### Sight
How far you can see with your ships. More sight gives you more information about your surroundings.
Each tick you will get all surrounding tiles you can see.

### Cannons
More cannons means a higher change to win, see fighting.

## Fighting
If a ship attacks a port a random number will be generated from `0 to amount of cannons`\
If the amount of the ship/attacker is greater or equal it will win!\
If a ship attacks another ship the same mechanic applies.\
But only the `defending ship` can sink, because of the suprise effect ;)\
So attack every other ship! Better don't attack your own ships...

## Attacking
For attacking you must add the id of the target in the Attack array.
Attacking is only possible if you are on the same tile.

## Ticks
The Game is turn based, so every ship gets the information and must answer. If all information is given the tick will be calculated.
And the next round is start.

### End of the Game
The Game has no end. Just give a time when you shut down the Server, like 1 hour. Or you can make several rounds.

## Moving
You can move in X and Y direction. Each movement is 1 speed point. So if you move X and Y you need at least a speed of 2.

## Gold
Your Score is based on the gold you made, so you can spend all your gold! It will not reduce your score.

# Setup
Just start the Server with `go run main.go`\
There will be a log with a secret. If you want an overview over the game, copy tha `secret` and paste
it into the index.html file in the _clients/overview folder.

## Overview
After you inserted the `secret` at the start of the script block, you can open it locally in the browser and you should see the map.

## Clients
The Clients should insert the Server IP in their Javascript `(_clients/shipcontrol/index.html)`
After that they can buy Ships.

In the example implementation the page will get all active ships for a team and then create the Ship class.
You can change whatever you want!

![Game field](game.jpg?raw=true "Game filed")

### Licences
Apache 2.0 except the asses in the overview.
These are licensed only in purpose for this games, so if you want to use it somewhere else, please buy the assets.
https://www.gamedevmarket.net/asset/sea-theme-game-assets-3944/
