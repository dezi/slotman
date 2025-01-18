# Raspberry Pi Hat with I2C and SPI interfaces.

bla bla 

## Benefits

- Hassle-free breakout - no modification of Pi required.
- 4 pin breakout I2C-1 bus connector with 5V power supply.
- 8 pin breakout SPI-0 connector with 3.3V power supply and 2 chip selects.
- 8 pin breakout SPI-1 connector with 3.3V power supply and 3 chip selects.

Material List
=============

1 x Development board 24 x 10 holes (https://www.amazon.de/dp/B0734XYJPM)

2 x Female header 2x20 pins 2.54 mm pitch extra high (https://www.amazon.de/dp/B07YDKX8SR)

1 x 2.54mm Breakaway 40 Pin Right Angle Single Row Pin Header PCB Connector (https://www.amazon.de/dp/B01MZE0XGZ)

1 x Enameled copper wire 0.4 mm (https://www.amazon.de/dp/B0DCJQJJCY)

Tool Set
========

[Look here...](../a2z-tools/readme.md)

Step 1
======

Mount the female header and solder it.

[<img src="images/raspberry-pi-hat-step-01-a.jpg" width="400"/>](images/raspberry-pi-hat-step-01-a.jpg)

This will be the position.

[<img src="images/raspberry-pi-hat-step-01-b.jpg" width="400"/>](images/raspberry-pi-hat-step-01-b.jpg)

After soldering.

[<img src="images/raspberry-pi-hat-step-01-c.jpg" width="400"/>](images/raspberry-pi-hat-step-01-c.jpg)

Cut the pins down to the soldering with a nail-scissor.

[<img src="images/raspberry-pi-hat-step-01-d.jpg" width="400"/>](images/raspberry-pi-hat-step-01-d.jpg)

Step 2
======

Now we mount the right angle single pin headers for I2C-1, SPI-0, SPI-1 and
an additional two pin 5V and GND connector for Raspberry Pi's fan if desired.

[<img src="images/raspberry-pi-hat-step-02-a.jpg" width="400"/>](images/raspberry-pi-hat-step-02-a.jpg)

This will be the positions.

**Important: it is correct that these pins go to soldering side of board.**

**Tip: Solder a single pin of each connector first and then make
sure they are standing straight.**

[<img src="images/raspberry-pi-hat-step-02-b.jpg" width="400"/>](images/raspberry-pi-hat-step-02-b.jpg)

Step 3
======

Now we start with the additional 5V pin header because it is simple.

Prepare Your connecting wires.

[<img src="images/raspberry-pi-hat-step-03-a.jpg" width="400"/>](images/raspberry-pi-hat-step-03-a.jpg)

Solder each wire on the connectors pin and thread through the holes as displayed.

[<img src="images/raspberry-pi-hat-step-03-b.jpg" width="400"/>](images/raspberry-pi-hat-step-03-b.jpg)

Pull the wires tight with a tweezer or plunge.

[<img src="images/raspberry-pi-hat-step-03-c.jpg" width="400"/>](images/raspberry-pi-hat-step-03-c.jpg)

Unconnected wire after pulling them tight.

[<img src="images/raspberry-pi-hat-step-03-d.jpg" width="400"/>](images/raspberry-pi-hat-step-03-d.jpg)

Solder them to 5V pin and GND pin as displayed in the image.

[<img src="images/raspberry-pi-hat-step-03-e.jpg" width="400"/>](images/raspberry-pi-hat-step-03-e.jpg)

[<img src="images/raspberry-pi-hat-gpio-pins.jpg" width="400"/>](images/raspberry-pi-hat-gpio-pins.jpg)

Step 4
======

We continue with the 4 pin I2C connector.

Prepare Your connecting wires.

[<img src="images/raspberry-pi-hat-step-04-a.jpg" width="400"/>](images/raspberry-pi-hat-step-04-a.jpg)

Pull the wires tight with a tweezer or plunge.

[<img src="images/raspberry-pi-hat-step-04-b.jpg" width="400"/>](images/raspberry-pi-hat-step-04-b.jpg)

Unconnected wires after pulling them tight.

[<img src="images/raspberry-pi-hat-step-04-c.jpg" width="400"/>](images/raspberry-pi-hat-step-04-c.jpg)

Solder them as displayed in the image.

[<img src="images/raspberry-pi-hat-step-04-d.jpg" width="400"/>](images/raspberry-pi-hat-step-04-d.jpg)

The I2C connector pins from left to right:

- 5V
- GND
- SCL (GPIO 3 Serial Clock I2C)
- SDA (GPIO 2 Serial Data I2C)

[<img src="images/raspberry-pi-hat-gpio-pins.jpg" width="400"/>](images/raspberry-pi-hat-gpio-pins.jpg)

Step 5
======

We continue with the 8 pin SPI-0 connector.

Prepare Your 7 connecting wires.

[<img src="images/raspberry-pi-hat-step-05-a.jpg" width="400"/>](images/raspberry-pi-hat-step-05-a.jpg)

Unconnected wires after pulling them tight.

[<img src="images/raspberry-pi-hat-step-05-b.jpg" width="400"/>](images/raspberry-pi-hat-step-05-b.jpg)

Unconnected wires before soldering.

[<img src="images/raspberry-pi-hat-step-05-c.jpg" width="400"/>](images/raspberry-pi-hat-step-05-c.jpg)

Connected wires after soldering.

[<img src="images/raspberry-pi-hat-step-05-d.jpg" width="400"/>](images/raspberry-pi-hat-step-05-d.jpg)

The SPI-0 connector pins from left to right:

- 3.3V
- GND
- SCLK (GPIO 11 SCXL SPI 0)
- MISO (GPIO 9 MISO SPI 0)
- MOSI (GPIO 10 MOSI SPI 0)
- CE0 (GPIO 8 Chip Enabled-CE0 SPI 0)
- CE1 (GPIO 7 Chip Enabled-CE1 SPI 0)

[<img src="images/raspberry-pi-hat-gpio-pins.jpg" width="400"/>](images/raspberry-pi-hat-gpio-pins.jpg)

