# Raspberry Pi Hat with I2C and SPI interfaces.

bla bla 

## Benefits

- Hassle-free breakout - no modification of Raspberry Pi required.
- 4 pin breakout I2C-1 bus connector with 5V power supply.
- 8 pin breakout SPI-0 connector with 3.3V power supply and 2 chip selects.
- 8 pin breakout SPI-1 connector with 3.3V power supply and 3 chip selects.

Material List
=============

1 x Development Board 24 x 10 holes (https://www.amazon.de/dp/B0734XYJPM)

2 x Female Header 2x20 pins 2.54 mm Pitch extra high (https://www.amazon.de/dp/B07YDKX8SR)

1 x 2.54mm Breakaway 40 Pin Right Angle Single Row Pin Header PCB Connector (https://www.amazon.de/dp/B01MZE0XGZ)

1 x Enameled Copper Wire 0.4 mm (https://www.amazon.de/dp/B0DCJQJJCY)

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

Prepare Your 7 connecting wires, 5 short wires and 3 longer wires.


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
- SCLK (GPIO 11 SCLK SPI 0)
- MISO (GPIO 9 MISO SPI 0)
- MOSI (GPIO 10 MOSI SPI 0)
- CE0 (GPIO 8 Chip Enabled-CE0 SPI 0)
- CE1 (GPIO 7 Chip Enabled-CE1 SPI 0)

[<img src="images/raspberry-pi-hat-gpio-pins.jpg" width="400"/>](images/raspberry-pi-hat-gpio-pins.jpg)

Step 6
======

We continue with the 8 pin SPI-1 connector.

Prepare now 8 connecting wires with different lengths. 

The SPI-1 connector pins from left to right:

- 3.3V
- GND
- SCLK (GPIO 21 SCLK SPI 1)
- MISO (GPIO 19 MISO SPI 1)
- MOSI (GPIO 20 MOSI SPI 1)
- CE0 (GPIO 18 Chip Enabled-CE0 SPI 1)
- CE1 (GPIO 17 Chip Enabled-CE1 SPI 1)
- CE2 (GPIO 16 Chip Enabled-CE2 SPI 1)

[<img src="images/raspberry-pi-hat-gpio-pins.jpg" width="400"/>](images/raspberry-pi-hat-gpio-pins.jpg)

[<img src="images/raspberry-pi-hat-step-06-a.jpg" width="400"/>](images/raspberry-pi-hat-step-06-a.jpg)

Unconnected wires of backside.

[<img src="images/raspberry-pi-hat-step-06-b.jpg" width="400"/>](images/raspberry-pi-hat-step-06-b.jpg)

Unconnected wires of topside.

[<img src="images/raspberry-pi-hat-step-06-c.jpg" width="400"/>](images/raspberry-pi-hat-step-06-c.jpg)

Connected wires after soldering.

[<img src="images/raspberry-pi-hat-step-06-d.jpg" width="400"/>](images/raspberry-pi-hat-step-06-d.jpg)

Step 7
======

Mounting the hat to the Raspberry Pi.

[<img src="images/raspberry-pi-hat-step-07-a.jpg" width="400"/>](images/raspberry-pi-hat-step-07-a.jpg)

This is my old and rugged Raspberry Pi 3B in a GeekPi housing.

As You can see, the hat needs an extension to connect.

[<img src="images/raspberry-pi-hat-step-07-b.jpg" width="400"/>](images/raspberry-pi-hat-step-07-b.jpg)

Hat mounted with extension.

[<img src="images/raspberry-pi-hat-step-07-c.jpg" width="400"/>](images/raspberry-pi-hat-step-07-c.jpg)

Hat mounted to Pi and fan power attached.

[<img src="images/raspberry-pi-hat-step-07-d.jpg" width="400"/>](images/raspberry-pi-hat-step-07-d.jpg)

Step 8
======

Testing the I2C-1 breakout.

I will use a ADS1115 chip which will show up at address 0x48 in the I2C scan.

[<img src="images/raspberry-pi-hat-step-08-a.jpg" width="400"/>](images/raspberry-pi-hat-step-08-a.jpg)

4 line connecting cable.

[<img src="images/raspberry-pi-hat-step-08-b.jpg" width="400"/>](images/raspberry-pi-hat-step-08-b.jpg)

4 line connecting cable with ADS1115 attached.

**Double-check the polarity!**

[<img src="images/raspberry-pi-hat-step-08-c.jpg" width="400"/>](images/raspberry-pi-hat-step-08-c.jpg)

Device attached to I2C-1 breakout.

**Double-check the polarity!**

[<img src="images/raspberry-pi-hat-step-08-d.jpg" width="400"/>](images/raspberry-pi-hat-step-08-d.jpg)

Output of the i2cdetect tool on Raspberry Pi terminal.

[<img src="images/raspberry-pi-hat-step-08-e.png" width="400"/>](images/raspberry-pi-hat-step-08-e.png)

Good job!!!

Step 9
======

Addendum:

We want to use the SPI-0 breakout for a nice display
device GC9A01. (https://www.amazon.de/dp/B0CFXVD9HX)

These devices need another control pin named DC. 
So we need to add another wire to the board.

[<img src="images/raspberry-pi-hat-step-09-a.jpg" width="400"/>](images/raspberry-pi-hat-step-09-a.jpg)

The wire goes to the unconnected last pin of the SPI-0 connector.

[<img src="images/raspberry-pi-hat-step-09-b.jpg" width="400"/>](images/raspberry-pi-hat-step-09-b.jpg)

The extra wire is now attached to GPIO pin 25.

[<img src="images/raspberry-pi-hat-step-09-c.jpg" width="400"/>](images/raspberry-pi-hat-step-09-c.jpg)

The SPI-0 connector pins from left to right:

- 3.3V
- GND
- SCLK (GPIO 11 SCLK SPI 0)
- MISO (GPIO 9 MISO SPI 0)
- MOSI (GPIO 10 MOSI SPI 0)
- CE0 (GPIO 8 Chip Enabled-CE0 SPI 0)
- CE1 (GPIO 7 Chip Enabled-CE1 SPI 0)
- DC (GPIO 25 )

[<img src="images/raspberry-pi-hat-gpio-pins.jpg" width="400"/>](images/raspberry-pi-hat-gpio-pins.jpg)
