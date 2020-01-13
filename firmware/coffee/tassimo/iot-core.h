/******************************************************************************
 * Copyright 2018 Google
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *****************************************************************************/
// This file contains static methods for API requests using Wifi / MQTT
#ifndef __ESP8266_MQTT_H__
#define __ESP8266_MQTT_H__
#include <ESP8266WiFi.h>
#include <time.h>
#include <WiFiManager.h>         //https://github.com/tzapu/WiFiManager


#include <MQTT.h>

#include <CloudIoTCore.h>
#include <CloudIoTCoreMqtt.h>
#include "iot-config.h" // Iot core configuration here

void on();
void off();
void handleCoffee();
void open();
void press();

void handleConfig(String payload);
void handleCommand(String payload);

void messageReceived(String &topic, String &payload) {
  if(topic == "/devices/"+String(device_id)+"/config") {
    handleConfig(payload);
  } else if(topic == "/devices/"+String(device_id)+"/commands") {
    handleCommand(payload);
  }
}
// Initialize WiFi and MQTT for this board
MQTTClient *mqttClient;
Client *netClient;
CloudIoTCoreDevice *device;
CloudIoTCoreMqtt *mqtt;
unsigned long iss = 0;
String jwt;
WiFiManager wifiManager;

///////////////////////////////
// Helpers specific to this board
///////////////////////////////
String getDefaultSensor() {
  return  "Wifi: " + String(WiFi.RSSI()) + "db";
}

String getJwt() {
  // Disable software watchdog as these operations can take a while.
  ESP.wdtDisable();
  iss = time(nullptr);
  Serial.println("Refreshing JWT");
  jwt = device->createJWT(iss, jwt_exp_secs);
  Serial.println("JWT Refreshed");
  ESP.wdtEnable(0);
  return jwt;
}

///////////////////////////////
// Orchestrates various methods from preceeding code.
///////////////////////////////
void publishTelemetry(String data) {
  mqtt->publishTelemetry(data);
}

void publishTelemetry(const char* data, int length) {
  mqtt->publishTelemetry(data, length);
}

void publishTelemetry(String subfolder, String data) {
  mqtt->publishTelemetry(subfolder, data);
}

void publishTelemetry(String subfolder, const char* data, int length) {
  mqtt->publishTelemetry(subfolder, data, length);
}

void connect() {
  mqtt->mqttConnect();
  //mqttClient->subscribe("/devices/esp8266-tanguy/commands/#", 1);
}

// TODO: fix globals
void setupCloudIoT() {
    configTime(0, 0, ntp_primary, ntp_secondary);
    Serial.println("Waiting on time sync...");
    while (time(nullptr) < 1510644967) {
        delay(10);
    }
    Serial.println("Setting up cloud iot core");

      netClient = new WiFiClientSecure();

  // Create the device
  device = new CloudIoTCoreDevice(
      project_id, location, registry_id, device_id,
      private_key_str);

  mqttClient = new MQTTClient(512);
  mqttClient->setOptions(180, true, 1000); // keepAlive, cleanSession, timeout
  mqtt = new CloudIoTCoreMqtt(mqttClient, netClient, device);
  mqtt->startMQTT(); // Opens connection
}

#endif //__ESP8266_MQTT_H__