# Analog-Digital Converters for Carrera-Go Speed Controllers

Converts the resistor values from the original **Carrera-Go Speed 
Controllers** into digital value for further processing.

The conversion is done via dual **ADS1115** analog/digital with 4
analog inputs each yielding in 8 tracks.

The data is sent via the **I2C** bus to the **Raspberry Pi**.

## Benefits

- Enables **full throttle speed** with **Carrera-Go** Controllers w/o hardware modification
- Can be used with or without **Carrera-Go** turbo-switch.
- Very fine speed control also at lower speeds.
- Pilots can configure individual response curves.
- Pilots can configure minimum low and top speed individually.
- Up to **6 pilots/tracks** plus 2 spare inputs.

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

**Tip: The resistor wire are sometimes reluctant to soldering. Use a multi-meter to make
sure they are all connected. This will save You from debugging later on.**