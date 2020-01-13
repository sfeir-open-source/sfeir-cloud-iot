#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>
#include <ESP8266httpUpdate.h>

#include <DNSServer.h>           //needed for WiFiManager
#include <ESP8266WebServer.h>    //needed for this program and for WiFiManager
#include <WiFiManager.h>         //https://github.com/tzapu/WiFiManager
#include <ArduinoJson.h>

#include "WiFiClientSecure.h"

#include <CloudIoTCore.h>

#include "Credentials.h"
#include "iot-core.h"

const int GPIO_PIN_4 = 4;
const int GPIO_PIN_5 = 5;
const int GPIO_PIN_12_RELAY = 12;
const String version = "1.0.0";

void on() {
  digitalWrite(GPIO_PIN_12_RELAY, HIGH);
}

void off() {
  digitalWrite(GPIO_PIN_12_RELAY, LOW);
}

void open() {
    digitalWrite(GPIO_PIN_4, LOW);
    delay(300);
    digitalWrite(GPIO_PIN_4, HIGH);
}

void press() {
    digitalWrite(GPIO_PIN_5, HIGH);
    delay(300);
    digitalWrite(GPIO_PIN_5, LOW);
}

void longPress() {
    digitalWrite(GPIO_PIN_5, HIGH);
    delay(5500);
    digitalWrite(GPIO_PIN_5, LOW);
}

void handleCoffee() {
    on();
    delay(1000);
    open();
    delay(1000);
    press();
    delay(120000);
    off();
}

void handleDetartrage() {
    on();
    delay(1000);
    open();
    delay(1000);
    longPress();
    delay(1800000);
    off();
}

void update() {
    t_httpUpdate_return ret = ESPhttpUpdate.update("europe-west1-strangersfeir.cloudfunctions.net", 80, "/download");
    Serial.println(ret);
    switch(ret) {
        case HTTP_UPDATE_FAILED:
            Serial.println("HTTP_UPDATE_FAILD Error code " + String(ESPhttpUpdate.getLastError()));
            Serial.println("HTTP_UPDATE_FAILD Error " + String(ESPhttpUpdate.getLastErrorString().c_str()));
            break;
        case HTTP_UPDATE_NO_UPDATES:
            Serial.println("HTTP_UPDATE_NO_UPDATES");
            break;
        case HTTP_UPDATE_OK:
            Serial.println("HTTP_UPDATE_OK");
            break;
    }
}

void setup ( void ) {
    Serial.begin(115200);
    Serial.println("setup - begin");

    pinMode(GPIO_PIN_4, OUTPUT);
    digitalWrite(GPIO_PIN_4, HIGH); // by default, the open/close switch is in serial with the optocouper, so transistor should be ON by default. Pull-up resistor.

    pinMode(GPIO_PIN_5, OUTPUT);
    digitalWrite(GPIO_PIN_5, LOW);  // by default, the start switch is in parallel with this optocouper, so transistor should be OFF by default. Pull-down resistor.

    pinMode(GPIO_PIN_12_RELAY, OUTPUT);
    digitalWrite(GPIO_PIN_12_RELAY, LOW); // by default, the start switch is in parallel with the relay, so transistor should be OFF by default. Pull-down resistor.

    WiFiManager wifiManager;
    wifiManager.autoConnect("AutoConnectAP");

    // Wait for connection
    while(WiFi.status() != WL_CONNECTED) {
        delay(500);
        Serial.print(".");
    }

    setupCloudIoT();

}

void handleConfig(String payload) {
    Serial.println("New config received : "+payload);
}

void handleCommand(String payload) {
    if(payload == "on") {
        on();
    } else if(payload =="off") {
        off();
    } else if(payload == "coffee") {
        handleCoffee();
    } else if(payload == "open") {
        open();
    } else if(payload == "press") {
        press();
    } else if(payload == "longPress") {
        longPress();
    } else if(payload == "update") {
        update();
    } else if(payload == "detartrage") {
        handleDetartrage();
    }
}

unsigned long lastMillis = 0;

void loop() {
    mqttClient->loop();
    delay(10);  // <- fixes some issues with WiFi stability
    if (!mqttClient->connected()) {
        connect();
    }
    if (millis() - lastMillis > 180000) {
        lastMillis = millis();
        publishTelemetry(version);
    }
}

 
