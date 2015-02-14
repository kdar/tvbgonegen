TV-B-Gone Gen
=============

A package that generates the compressed structures for v1.2 of [TV-B-Gone](https://learn.adafruit.com/tv-b-gone-kit).

For now, this only works on codes generated from the library [Arduino-IRremote](https://github.com/shirriff/Arduino-IRremote). I wasn't able to get codes generated from [Raw-IR-decoder](https://github.com/adafruit/Raw-IR-decoder-for-Arduino) from Adafruit to work. It seems some of the timings are off.

### Usage

* Install [Arduino-IRremote](https://github.com/shirriff/). You can also refer to his [blog post](http://www.righto.com/2009/08/multi-protocol-infrared-remote-library.html).
* Use the [IRrecvDump](https://github.com/shirriff/Arduino-IRremote/tree/master/examples/IRrecvDump) example from [Arduino-IRremote](https://github.com/shirriff/).
* Upload the sketch to your Arduino with the IR receiver.
* Open up the Arduino serial monitor.
* Point remote at receive and click a button.
* Copy the codes printed in the monitor. Should look like

```
Decoded NEC: 1CE348B7 (32 bits)
Raw (68): -31588 7600 -3600 550 -350 600 -350 600 -350 600 -1250 600 -1250 600 -1250 550 -400 550 -350 650 -1200 650 -1200 650 -1200 600 -350 600 -350 550 -350 650 -1200 650 -1200 600 -350 650 -1200 600 -350 550 -400 600 -1250 550 -350 600 -350 600 -350 600 -1250 550 -400 600 -1250 600 -1250 550 -350 650 -1200 650 -1200 650 -1200 600 
```

* Run the tvbgonegen command on this input.
