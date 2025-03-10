# Ambient Sensors Module

This module collects ambient sensor data like
temperature, humidity, air-pressure, air-quality
and illumination and RGB-color of the ambient light.

## Benefits

- Temperature in Celsius
- Humidity in Percent
- Air-Pressure in hPa
- Air-Quality in Percent
- Illumination in Lux
- Ambient light color in RGB

Material List
=============

1 x TCS34725 I2C RGB Illumination Sensor (https://www.amazon.de/dp/B07DK4NHBY)

1 x BMP280 + AHT20 I2C Humidity / Temperature / Pressure (https://www.amazon.de/dp/B0B76Y29T7)

1 x SGP40 eCO2 Kohlendioxid Raumluftmonitor (https://www.amazon.de/dp/B0CJY3SM8S)

1 x Development Board 24 x 18 holes (https://www.amazon.de/dp/B0734XYJPM)

1 x 2.54mm Breakaway 40 Pin Right Angle Single Row Pin Header PCB Connector (https://www.amazon.de/dp/B01MZE0XGZ)

1 x 2.54mm Breakaway 40 Pin Connectors (https://www.amazon.de/dp/B0DJ2MYFMB)

1 x Distance Holder Set (https://www.amazon.de/dp/B08VD1ZCFJ)

1 x Enameled Copper Wire 0.4 mm (https://www.amazon.de/dp/B0DCJQJJCY)

Tool Set
========

[Look here...](../a2z-tools/readme.md)

Step 1
======

Solder the right angle connectors to sensors.

**Notice: The SGP30 sensor You see here has been replaced with
a SGP40 sensor, since the SGP30 did not work at all. The SGP40
has the same form factor.**

[<img src="images/ambient-module-step-01-a.jpg" width="400"/>](images/ambient-module-step-01-a.jpg)

Soldering the connector to the combined BMP280 + AHT20 I2C is tricky, since
You have to solder on top of the contacts.

[<img src="images/ambient-module-step-01-b.jpg" width="400"/>](images/ambient-module-step-01-b.jpg)

Top-side view.

[<img src="images/ambient-module-step-01-c.jpg" width="400"/>](images/ambient-module-step-01-c.jpg)

Step 2
======

Prepare and solder the sensor connectors.

You will need 2 x 4 pin and 1 x 7 pin cuts.

[<img src="images/ambient-module-step-02-a.jpg" width="400"/>](images/ambient-module-step-02-a.jpg)

Bend the pins to a right angle.

[<img src="images/ambient-module-step-02-b.jpg" width="400"/>](images/ambient-module-step-02-b.jpg)

Get the development board now.

[<img src="images/ambient-module-step-02-c.jpg" width="400"/>](images/ambient-module-step-02-c.jpg)

Attach the connectors like You see on this image.

[<img src="images/ambient-module-step-02-d.jpg" width="400"/>](images/ambient-module-step-02-d.jpg)

View from back-side after soldering.

[<img src="images/ambient-module-step-02-e.jpg" width="400"/>](images/ambient-module-step-02-e.jpg)

Step 3
======

Drill holes for the sensor distance holders and attachments.

You will need a 2 - 2.5 mm drill.

[<img src="images/ambient-module-step-03-a.jpg" width="400"/>](images/ambient-module-step-03-a.jpg)

View after drilling.

[<img src="images/ambient-module-step-03-b.jpg" width="400"/>](images/ambient-module-step-03-b.jpg)

Step 4
======

Mount the development board corners distance holders.

[<img src="images/ambient-module-step-04-a.jpg" width="400"/>](images/ambient-module-step-04-a.jpg)

[<img src="images/ambient-module-step-04-b.jpg" width="400"/>](images/ambient-module-step-04-b.jpg)

[<img src="images/ambient-module-step-04-c.jpg" width="400"/>](images/ambient-module-step-04-c.jpg)

Step 5
======

Mount the sensors distance holders and attachments.

[<img src="images/ambient-module-step-05-a.jpg" width="400"/>](images/ambient-module-step-05-a.jpg)

The longer holdes go left.

[<img src="images/ambient-module-step-05-b.jpg" width="400"/>](images/ambient-module-step-05-b.jpg)

[<img src="images/ambient-module-step-05-c.jpg" width="400"/>](images/ambient-module-step-05-c.jpg)

View with sensors attached.

[<img src="images/ambient-module-step-05-d.jpg" width="400"/>](images/ambient-module-step-05-d.jpg)

Step 6
======

Solder the 4 pin right angle I2C bus connector.

[<img src="images/ambient-module-step-06-a.jpg" width="400"/>](images/ambient-module-step-06-a.jpg)

View from soldering side.

[<img src="images/ambient-module-step-06-b.jpg" width="400"/>](images/ambient-module-step-06-b.jpg)

Step 7
======

Solder the wiring.

Prepare 4 wires like this.

[<img src="images/ambient-module-step-07-a.jpg" width="400"/>](images/ambient-module-step-07-a.jpg)

The 5V connections.

[<img src="images/ambient-module-step-07-b.jpg" width="400"/>](images/ambient-module-step-07-b.jpg)

The GND connections.

[<img src="images/ambient-module-step-07-c.jpg" width="400"/>](images/ambient-module-step-07-c.jpg)

The I2C-SCL connections.

[<img src="images/ambient-module-step-07-d.jpg" width="400"/>](images/ambient-module-step-07-d.jpg)

The I2C-SDA connections.

[<img src="images/ambient-module-step-07-e.jpg" width="400"/>](images/ambient-module-step-07-e.jpg)

Good job!
