# Motoron-Unit for Carrera-Go motor control

Converts speed control instructions into track voltage.

The conversion is done via Pololu M2T550 Motorons each of which 
can control the car voltage for two tracks.

The Motorons can regulate driving power in a range
from 10 to 22 Volt, compared to the normal input
of 14.5 Volt.

The data is exchanged via the **I2C** bus with the **Raspberry Pi**.

## Benefits

- True 0 - 100 % voltage regulation.
- Input driving voltage from 10 V upto 22 V.
- Controls up to 8 tracks.
- Supports one or two power supplies.
- Pilot individual power curves.
- Pilot individual min and max speed limiting.

Material List
=============

1-4 x Pololu Motoron M2T550 Dual I2C Motor Controller (https://www.amazon.de/dp/B0DJBYSWN5)

1 x Development Board 24 x 18 holes (https://www.amazon.de/dp/B0734XYJPM)

1 x 2.54mm Breakaway 40 Pin Right Angle Single Row Pin Header PCB Connector (https://www.amazon.de/dp/B01MZE0XGZ)

1 x 2.54mm Breakaway 40 Pin Female Header PCB Connector (https://www.amazon.de/dp/B07DBY753C)

1 x 2.54mm Breakaway 40 Pin Connectors (https://www.amazon.de/dp/B0DJ2MYFMB)

4 x Jumper 2.54 mm (https://www.amazon.de/dp/B0CSJDDXML)

1 x Steckersortiment (https://www.amazon.de/dp/B0D7HXG492)

1 x Distance Holder Set (https://www.amazon.de/dp/B08VD1ZCFJ)

1 x Enameled Copper Wire 0.4 mm (https://www.amazon.de/dp/B0DCJQJJCY)

Tool Set
========

[Look here...](../a2z-tools/readme.md)

Step 1
======

Prepare 5 x 4 pin right angel row connectors.

[<img src="images/motoron-unit-step-01-a.jpg" width="400"/>](images/motoron-unit-step-01-a.jpg)

Get Your development 24 x 18 holes board.

[<img src="images/motoron-unit-step-01-b.jpg" width="400"/>](images/motoron-unit-step-01-b.jpg)

Solder the connectors like displayed here:

[<img src="images/motoron-unit-step-01-c.jpg" width="400"/>](images/motoron-unit-step-01-c.jpg)

View from soldering side.

[<img src="images/motoron-unit-step-01-d.jpg" width="400"/>](images/motoron-unit-step-01-d.jpg)

Step 2
======

Prepare 4 x 8 pin, 4 x 6 pin and 4 x 1 pin female headers.

[<img src="images/motoron-unit-step-02-a.jpg" width="400"/>](images/motoron-unit-step-02-a.jpg)

Solder the headers like displayed here:

[<img src="images/motoron-unit-step-02-b.jpg" width="400"/>](images/motoron-unit-step-02-b.jpg)

View from soldering side.

[<img src="images/motoron-unit-step-02-c.jpg" width="400"/>](images/motoron-unit-step-02-c.jpg)

Step 3
======

Prepare 4 x 2 pin straight connectors with jumpers.

[<img src="images/motoron-unit-step-03-a.jpg" width="400"/>](images/motoron-unit-step-03-a.jpg)

Solder the jumpers like displayed here:

[<img src="images/motoron-unit-step-03-b.jpg" width="400"/>](images/motoron-unit-step-03-b.jpg)

Step 4
======

Get two driving power connectors.

[<img src="images/motoron-unit-step-04-a.jpg" width="400"/>](images/motoron-unit-step-04-a.jpg)

Cut the middle pin like this:

[<img src="images/motoron-unit-step-04-b.jpg" width="400"/>](images/motoron-unit-step-04-b.jpg)

Solder them into these positions:

**Remark: The connectors do not exactly fit the raster.**

[<img src="images/motoron-unit-step-04-c.jpg" width="400"/>](images/motoron-unit-step-05-b.jpg)

Step 5
======

Prepare the single power source jumper.

[<img src="images/motoron-unit-step-05-a.jpg" width="400"/>](images/motoron-unit-step-05-a.jpg)

Solder it into this position:

[<img src="images/motoron-unit-step-05-b.jpg" width="400"/>](images/motoron-unit-step-05-b.jpg)

Step 6
======

Soldering the driving power connections.

We start with the GND connect. 

Prepare a thicker wire like this:

[<img src="images/motoron-unit-step-06-a.jpg" width="400"/>](images/motoron-unit-step-06-a.jpg)

Solder it into these positions:

[<img src="images/motoron-unit-step-06-b.jpg" width="400"/>](images/motoron-unit-step-06-b.jpg)

Prepare two thicker VIN connect wires.

[<img src="images/motoron-unit-step-06-c.jpg" width="400"/>](images/motoron-unit-step-06-c.jpg)

Solder them into these positions:

[<img src="images/motoron-unit-step-06-d.jpg" width="400"/>](images/motoron-unit-step-06-d.jpg)

Step 7
======

Connect the Motorons enable I2C address programming jumpers.

The jumpers connect against GND. If pulled to ground,
You can programm the real I2C access address for each unit.

We simply make tin bridges:

[<img src="images/motoron-unit-step-07-a.jpg" width="400"/>](images/motoron-unit-step-07-a.jpg)

The same here:

[<img src="images/motoron-unit-step-07-b.jpg" width="400"/>](images/motoron-unit-step-07-b.jpg)

**Tip: Measure after applying the bridges if they do not produce
a shortcut.**

**Important: Also remember to unplug the jumpers, as they produce
a desired shortcut if active.**

Step 8
======

Connect 5V and GND.

We start with 5V. Prepare a wire like this:

[<img src="images/motoron-unit-step-08-a.jpg" width="400"/>](images/motoron-unit-step-08-a.jpg)

Solder it at the spots marked with red bullets:

[<img src="images/motoron-unit-step-08-b.jpg" width="400"/>](images/motoron-unit-step-08-b.jpg)

Continue with GND. Prepare a wire like this:

[<img src="images/motoron-unit-step-08-c.jpg" width="400"/>](images/motoron-unit-step-08-c.jpg)

Solder it at the spots marked with green bullets:

[<img src="images/motoron-unit-step-08-d.jpg" width="400"/>](images/motoron-unit-step-08-d.jpg)

Step 9
======

Connect I2C SDA and SCL.

Prepare a wire like this:

[<img src="images/motoron-unit-step-09-a.jpg" width="400"/>](images/motoron-unit-step-09-a.jpg)

We start with SDA. Solder it at the spots marked with red bullets:

[<img src="images/motoron-unit-step-09-b.jpg" width="400"/>](images/motoron-unit-step-09-b.jpg)

Continue with SCL. Solder it at the spots marked with green bullets:

[<img src="images/motoron-unit-step-09-c.jpg" width="400"/>](images/motoron-unit-step-09-c.jpg)
