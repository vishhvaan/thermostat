# DIY IoT Thermostat

This is a package that enables users to make use of cheap temperature (and humidity, barometric pressure, and volatile organic compound gas) sensors and read sensor-generated data into their Home Assistant instance to automate functions around their home.

This repo contains temperature sensor circuit designs, code that runs on a microcomputer to poll the sensor and serve data, and YAML configurations that are loaded into Home Assistant.

Version 1: DS18B20  
Version 2: BME680 [Upcoming]

<p align="center">
  <img src="https://i.ibb.co/tLtbWQL/IMG-20200719-163154.jpg" height="400" alt="Version 1 Hardware"/>
  <img src="https://i.ibb.co/gtjVvfm/thermostat-panel.png" height="400" alt="Version 1 Hardware"/>
</p>
<p align="center">
  <img src="https://i.ibb.co/41WF6sR/thermostat.png" alt="Version 1 Hardware"/>
</p>

## Hardware
### List of Equipment

* DS18B20 + 4k7 pullup resistor
* LED of choice + associated resistor (has to be driven from a 3.3V digital signal)
* Raspberry Pi
* Logitech Harmony Hub
* Seperate server (optional)

```
	+---------------+           +---------------+  Webhook +----------------+
	|               |  1-Wire   |               +<---------+                |
	|  Temperature  +---------->+   Raspberry   |   WiFi   |  Server with   |
	|    Sensor     |  or I2C   |      Pi       |          | Home Assistant |
	|               |           |               +--------->+                |
	+---------------+           +---------------+  Stdout  ++--------------++
								|              ^
								|  Harmony API |
								|     WiFi     |
								v              |
				    +---------------+    IR    ++--------------++
				    |     Cooler    +<---------+                |
				    +---------------+          |    Logitech    |
				    +---------------+          |   Harmony Hub  |
				    |     Heater    +<---------+                |
				    +---------------+    IR    +----------------+
```
The server with home assistant could reside on the Raspberry Pi and could run the thermostat program directly instead of through a webhook.

## Software

Thermostat data can be exposed to the network in two ways. The preferred way is through webhooks. This project makes use of the very versatile [webhook][wh] tool. Control of the LED can also be accomplished through webhooks.

Releases of the thermostat software are packaged with webhook binaries that have been tested and pre-configured to work with the thermostat program.

Alternatively, users may choose to make use of the Pi's in-built SSH server to execute the themostat program. This method relies of the generation of SSH keys and ensuring the Home Assistant instance has access to the key. If Home Assistant is run from Docker, the SSH server has to be trusted every time a new container is created (after an upgrade).



[wh]: https://github.com/adnanh/webhook
