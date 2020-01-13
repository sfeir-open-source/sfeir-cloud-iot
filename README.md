# SFEIR IoT projects

## Projects

### Google Home coffee maker

An Arduino IoT Core connected project to make Tassimo coffee with help from Google Home.

### Google Home stranger thing garland

An Arduino IoT Core connected project to display words letter by letter on a LED garland
like in the Stranger Thing Netflix TV show.

## Code base

### Cloud functions

Code base is built in separate cloud functions

#### Download

Download function is meant to do OTA device firmware updates from firmwares in the cloud storage.
It is HTTP triggered from the devices wanted to be updated.

#### Coffee

Coffee function is also HTTP triggered from IFTTT through Google Home
to launch a coffee making on the connected coffee machine with an Arduino.


#### Words

Words function is HTTP triggered from IFTTT through Google Home
to display a word or a list of words on the connected garland with an Arduino.

### Firmwares

The list of firmwares

#### Coffee

Coffee machine firmware.

#### Words

Stanger Things garlang firmware. 

### Web App Engine

A simple App Engine web page to launch the display of words on the garland.