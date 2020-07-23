# DIY IoT Thermostat

This is a package that enables users to make use of cheap temperature (and humidity, barometric pressure, and volatile organic compound gas) sensors and read sensor-generated data into their Home Assistant instance to automate functions around their home.

This repo contains temperature sensor circuit designs, code that runs on a microcomputer to poll the sensor and serve data, and YAML configurations that are loaded into Home Assistant.

Version 1: DS18B20  
Version 2: BME680 [Upcoming]

<p align="center">
  <img src="https://i.ibb.co/tLtbWQL/IMG-20200719-163154.jpg" height="600" alt="Version 1 Hardware"/>
  <img src="https://i.ibb.co/gtjVvfm/thermostat-panel.png" height="600" alt="Version 1 Hardware"/>
</p>
<p align="center">
  <img src="https://i.ibb.co/41WF6sR/thermostat.png" alt="Version 1 Hardware"/>
</p>

### List of Equipment

DS18B20
LED of choice + associated resistor (has to be driven from a 3.3V digital signal)
Raspberry Pi
Logitech Harmony Hub
Seperate server (optional)
