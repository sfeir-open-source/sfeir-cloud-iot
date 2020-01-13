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

#include <Adafruit_NeoPixel.h> 

#define PIN 4
#define NUMPIXELS 26

Adafruit_NeoPixel pixels = Adafruit_NeoPixel(NUMPIXELS, PIN);

const char* software_version = "1.1.9";
unsigned long standbyMillis = 0;
String url = "";
int port = 0;
String path = "";

void update() {
    t_httpUpdate_return ret = ESPhttpUpdate.update(url, port, path);
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

    pixels.begin();
    pixels.clear();
    pixels.show();
    WiFiManager wifiManager;
    wifiManager.autoConnect("AutoConnectAP");

    Serial.println("");
    // Wait for connection
    while(WiFi.status() != WL_CONNECTED) {
        delay(500);
        Serial.print(".");
    }

    Serial.println("\nConnected!\n");

    setupCloudIoT();

    Serial.println("\n Connected to MQTT !\n");

    pixels.setBrightness(127);
}

void handleConfig(String payload) {
    Serial.println("New config received : "+payload);
    DynamicJsonBuffer jsonBuffer;

    JsonObject& root = jsonBuffer.parseObject(payload);
    String autoUpdate = root["autoUpdate"];

    url = root["server"]["url"].as<String>();
    port = root["server"]["port"].as<int>();
    path = root["server"]["path"].as<String>();

    if(autoUpdate == "true" && url.length()>0 && port!=0 && path.length()>0) {
        Serial.println("Updating device !");
        update();
    }

}

String message="";

void handleCommand(String payload) {
    Serial.println("New message received : "+payload);
    message = payload;
}

struct colorCode {
    char letter;
    unsigned int red;
    unsigned int green;
    unsigned int blue;
};

struct colorCode alphabet[26] = {
    {'a', 255, 255, 225}, // A 0 white
    {'b', 0, 0, 225}, // B 1 dark blue
    {'c', 210, 30, 210}, // C 2 magenta
    {'d', 103, 228, 255}, // D 3 light blue
    {'e', 0, 0, 225}, // E 4 dark blue
    {'f', 255, 255, 0}, // F 5 yellow
    {'g', 255, 0, 0}, // G 6 red
    {'h', 0, 255, 100}, // H 7 green

    /* second row, reverse order */
    {'q', 210, 30, 210}, // Q 8 magenta
    {'p', 103, 228, 255}, // P 9 icy blue
    {'o', 210, 30, 210}, // O 10 magenta
    {'n', 255, 0, 0}, // N 11 red
    {'m', 255, 255, 0}, // M 12 yellow
    {'l', 103, 228, 255}, // L 13 light blue
    {'k', 0, 0, 225}, // K 14 dark blue
    {'j', 255, 0, 0}, // J 15 red
    {'i', 0, 255, 100}, // I 16 green

    /* new row, normal order */
    {'r', 103, 228, 255}, // R 17 light blue
    {'s', 255, 255, 225}, // S 18 white
    {'t', 255, 255, 0}, // T 19 yellow
    {'u', 0, 0, 225}, // U 20 dark blue
    {'v', 255, 0, 0}, // V 21 red
    {'w', 0, 255, 100}, // W 22 green
    {'x', 255, 255, 0}, // X 23 yellow
    {'y', 210, 30, 210}, // Y 24 magenta
    {'z', 255, 0, 0}, // Z 25 red
};

/* Helper function that gets the index of a given letter 
from the alphabet RGB data structure */
int getIndexOfLetter(char letter) {
    for (int i=0; i<NUMPIXELS; i++) {
        if(alphabet[i].letter == letter) {
            return i;
        }
    }
    return -1;
}

void displayLetter(int index, int displayTime, int blankTime) {
    int red = alphabet[index].red;
    int blue = alphabet[index].blue;
    int green = alphabet[index].green;
    
    pixels.setPixelColor(index, pixels.Color(red, green, blue));
    pixels.show();
        
    // after 1.5 seconds
    delay(displayTime);
    // clear pixels,
    pixels.clear();
    pixels.show();
    // then wait a little to see clear pixels
    delay(blankTime);
}

void loop() {
    mqttClient->loop();
    delay(10);  // <- fixes some issues with WiFi stability
    if (!mqttClient->connected()) {
        connect();
    }

    // At start, or if no message for 15 minutes, start random mode
    if (millis() - standbyMillis > 900000 || standbyMillis == 0) {
        int randomIndex = int(random(0,NUMPIXELS));
        displayLetter(randomIndex, 500, 2000);
    }

    if (message.length() >0) {
        pixels.clear();
        pixels.show();
        delay(2000);

        pixels.setBrightness(255);
        // Reset stand by timer
        standbyMillis = millis();

        Serial.println("Message to display : "+message);
        for(int i=0; i<message.length();i++) {
            int j = getIndexOfLetter(message.charAt(i));
            if (0 <= j && j < NUMPIXELS) {
                displayLetter(j, 1500, 500);
            }
        }
        pixels.clear();
        pixels.show();
        message="";
        pixels.setBrightness(127);
    }

}

 
