FROM denouche/docker-arduino

ENV ARDUINO_ESP8266_VERSION 2.4.0-rc1

RUN mkdir -p /root/.arduino15/ \
    && echo "boardsmanager.additional.urls=https://github.com/esp8266/Arduino/releases/download/$ARDUINO_ESP8266_VERSION/package_esp8266com_index.json" > /root/.arduino15/preferences.txt

RUN arduino --install-boards "esp8266:esp8266"
RUN arduino --install-library "WiFiManager,ArduinoJson:5.13.5,MQTT,Google Cloud IoT Core JWT"

WORKDIR /usr/src/app

