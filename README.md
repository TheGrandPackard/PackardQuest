# PackardQuest

PackardQuest is a personal project that imitates the [MagiQuest](https://www.greatwolf.com/southern-california/waterpark-attractions/attractions/magiquest) game. PackardQuest provides software to facilitate a multiplayer experience using the official or homebrew wands. Hardware schematics and firmware are also available.

## Overview

### Hardware

The PackardQuest game is meant to imitate the full capabilities of MagiQuest, which consists of multiple variations of interactions with wands to make progress on the quest.

#### Wands

The official wands can be used and the unique infrared identifier can be extracted from the wands to track player progress. Homebrew wands are also compatible so long as they transmit MagiQuest encoded data like the official wands.

#### Interaction Nodes

The interaction nodes are what the wands will interact with. The nodes are designed to be low cost and simple components. The nodes are currently built around the Arduino platform and communicate with a PC, but can just as easily be ported to a Raspberry Pi or ESP32 type device for a standalone node that communicates with a centralized server. The interaction nodes can be programmed to do many things, the proof of concept included leveraging IFTT to turn on an IoT light.

### Software

The PackardQuest software can be deployed as standalone node that also contains the player and game database, or it can be deployed with a centralized server and many interaction nodes.

#### Node

The node is the interface between the interaction node hardware and the rest of the game. For simplicity, the Arduino node only handles processing the infrared and passes along all data to the software to further process the interfaction. For other platforms such as Raspberry Pi and ESP32, all functionality can be handled in a single piece of software or firmware.


#### Server

When using many nodes, a centralized server can be used to track players and their progress in the game. The nodes will communicate over TCP/IP to the server to access player and game data and update the player's progress in the game.

## Getting Started

The easiest way to get started is to set up a standalone interaction node using an Arduino attached to a PC. You will need the following:

* Arduino Uno or similar
* Infrared receiver
* Breadboard and wiring to hook up infrared receiver to the Arduino
* USB cable to connect Arduino to PC
* PC with the following software
    * [PackardQuest Node software](https://github.com/TheGrandPackard/PackardQuest/releases)
    * [Arduino IDE](https://www.arduino.cc/en/software)
* A MagiQuest wand
* A free [IFTTT](https://ifttt.com/) account 

### Arduino Sketch

To configure the interaction node, install the Arduino IDE and flash the InteractNode sketch. Refer to the schematic below for how to wire up the infared receiver. After the sketch has been flashed, interaction from a wand can be verified through the serial console in the Arduino IDE. After verification, the Arduino IDE can be closed.

![](https://cdn-learn.adafruit.com/assets/assets/000/000/555/original/light_arduinopna4602.gif?1447976120)

### Node Software

The software also expects a player database file to exist in the same directory - you can rename the example file from `players.example.json` to `players.json` and update the player entry. The wand ID is the integer value of your wand (check the serial console in the Arduino IDE).


The software runs in a console and expects arguments to be provided, for example:

```./client -players-file=players.json -serial-port=COM6 -ifttt-api-key=<YOUR_KEY_GOES_HERE>```

With the correct arguments, the interaction node should be detected and process wand interactions and forward the events to IFTT.