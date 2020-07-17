package download

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/binary"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	macParam         = "macParam"
	deviceTypeParam  = "type"
	deviceTypeCoffee = "coffee"
	deviceTypeLight  = "light"
)

// Download returns the lastest arduino firmware for a device if ok
func Download(w http.ResponseWriter, r *http.Request) {

	// check macParam
	// short list of devices for security
	deviceMac := r.URL.Query().Get(macParam)
	log.Printf("Connected device MAC address is : %s", deviceMac)

	// check type
	deviceType := r.URL.Query().Get(deviceTypeParam)

	firmwareBucket := os.Getenv("FIRMWARE_BUCKET")
	if firmwareBucket == "" {
		log.Print("Error while reading firmware bucket name")
		return
	}

	coffeeFirmware := os.Getenv("COFFEE_FIRMWARE")
	if coffeeFirmware == "" {
		log.Print("Error while reading coffee firmware name")
		return
	}

	lightFirmware := os.Getenv("LIGHT_FIRMWARE")
	if lightFirmware == "" {
		log.Print("Error while reading light firmware name")
		return
	}

	// connect to storage
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var firmware []byte
	switch deviceType {
	case deviceTypeLight:
		// type == light ==> light firmware
		// gs://firmwareBucket/lightFirmware
		firmware, err = read(client, firmwareBucket, lightFirmware)
		if err != nil {
			log.Printf("Cannot read lightning firmware: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte("Error while reading firmware"))
			if err != nil {
				log.Printf("Error while sending error %v", err)
			}
			return
		}
		log.Printf("Device type is : %s", deviceType)
	default:
		// no type or coffee ==> coffee machine firmware
		// gs://firmwareBucket/coffeeFirmware
		deviceType = deviceTypeCoffee
		firmware, err = read(client, firmwareBucket, coffeeFirmware)
		if err != nil {
			log.Printf("Cannot read coffee firmware: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte("Error while reading firmware"))
			if err != nil {
				log.Printf("Error while sending error %v", err)
			}
			return
		}
		log.Printf("Device type is : %s", deviceType)
	}

	log.Printf("Object content retrieved: %v\n", binary.Size(firmware))

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+deviceType+".bin")
	w.Header().Set("Content-Length", strconv.Itoa(binary.Size(firmware)))
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(firmware)
	if err != nil {
		log.Printf("Error sending firmware: %v", err)
	}
}

func read(client *storage.Client, bucket, object string) ([]byte, error) {
	ctx := context.Background()
	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return data, nil
}
