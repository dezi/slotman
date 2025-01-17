# Analog-Digital Converters for Carrera-Go Speed Controllers

Converts the resistor values from the original **Carrera-Go Speed 
Controllers** into digital value for further processing.

The conversion is done via dual **ADS1115** analog/digital with 4
analog inputs each yielding in 8 tracks.

The data is sent via the **I2C** bus to the **Raspberry Pi**.

## Benefits

- Enables **full throttle speed** with **Carrera-Go** Controllers w/o hardware modification
- Can be used with or without **Carrera-Go** turbo-switch.
- **Very fine** speed control also at **lower speeds**.
- Pilots can configure **individual response curves**.
- Pilots can configure **minimum low and top speed** individually.
- Controls up to **6 pilots/tracks** plus 2 spare inputs.

Material List
=============

  2 x I2C ADS1115 16Bit ADC 4 channel Module Gain Amplifier (https://www.amazon.de/dp/B07RJT3GHC)

  1 x Development board 24 x 18 holes (https://www.amazon.de/dp/B0734XYJPM)

16 x 1k Ohm resistor (https://www.amazon.de/dp/B0BMXBZCTF) 

  1 x 2.54mm Breakaway 40 Pin Right Angle Single Row Pin Header PCB Connector (https://www.amazon.de/dp/B01MZE0XGZ)

  1 x 2.54mm Breakaway 40 Pin Female Header PCB Connector (https://www.amazon.de/dp/B01MZE0XGZ)

  1 x Distance holder set (https://www.amazon.de/dp/B08VD1ZCFJ)

  1 x Enameled copper wire 0.4 mm (https://www.amazon.de/dp/B0DCJQJJCY)

Step 1
======

Mount the 1k Ohm voltage divider resistors. Pull them tight
with a plunge so that they line up nicely on the front side.

**Do no soldering yet!**

[<img src="images/ad-converter-step-01-a.jpg" width="400"/>](images/ad-converter-step-01-a.jpg)

[<img src="images/ad-converter-step-01-b.jpg" width="400"/>](images/ad-converter-step-01-b.jpg)

[<img src="images/ad-converter-step-01-c.jpg" width="400"/>](images/ad-converter-step-01-c.jpg)

Step 2
======

Now solder the spots at the bottom of the picture. 

After soldering use a nail scissor to cut the remainder of the resistor's wire. 

Bent and cut it that way, so it just connects to the next contact hole. This is required later on.

[<img src="images/ad-converter-step-02-a.jpg" width="400"/>](images/ad-converter-step-02-a.jpg)

Now we want to accomplish a connections between all ends of the lower resistor array.

Bent and cut the wire a shown in the picture and solder the spots.

[<img src="images/ad-converter-step-02-b.jpg" width="400"/>](images/ad-converter-step-02-b.jpg)

Now we want to accomplish a connections between all ends of the upper resistor array the same way.

[<img src="images/ad-converter-step-02-c.jpg" width="400"/>](images/ad-converter-step-02-c.jpg)

**Tip: The resistor wires are sometimes reluctant to soldering.**

**Use a multi-meter to make sure they are all connected.**

**This will save You from debugging later on.**

Step 3
======

Mount the analog input right angle connectors.

[<img src="images/ad-converter-step-03-a.jpg" width="400"/>](images/ad-converter-step-03-a.jpg)

Connect and solder the ends of the resistors together with the connector pins.

[<img src="images/ad-converter-step-03-b.jpg" width="400"/>](images/ad-converter-step-03-b.jpg)

**Tip: Again, the resistor wires are sometimes reluctant to soldering.**

**Use a multi-meter to make sure they are all connected.**

**This will save You from debugging later on.**

Step 4
======

Mount the I2C input connector and the ADS1115 inline sockets.

[<img src="images/ad-converter-step-04-a.jpg" width="400"/>](images/ad-converter-step-04-a.jpg)

Solder only one pin of the inline sockets first, then make sure before soldering the rest,
they are standing nicely upright.

[<img src="images/ad-converter-step-04-b.jpg" width="400"/>](images/ad-converter-step-04-b.jpg)
