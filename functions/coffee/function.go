package coffee

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"google.golang.org/api/cloudiot/v1"
	"log"
	"net/http"
	"os"
)

const (
	commandID     = "coffee"
	iotBusPattern = "projects/%s/locations/%s/registries/%s/devices/%s"
)

// Coffee launches the coffee making
func Coffee(w http.ResponseWriter, r *http.Request) {

	// read configuration from env
	projectID := os.Getenv("PROJECT_ID")
	if projectID == "" {
		log.Print("Error while reading project ID")
		return
	}

	projectRegion := os.Getenv("PROJECT_REGION")
	if projectRegion == "" {
		log.Print("Error while reading project region")
		return
	}

	projectRegistryID := os.Getenv("PROJECT_REGISTRY_ID")
	if projectRegistryID == "" {
		log.Print("Error while reading project registry ID")
		return
	}

	projectDeviceID := os.Getenv("PROJECT_DEVICE_ID")
	if projectDeviceID == "" {
		log.Print("Error while reading project device ID")
		return
	}

	res, err := sendCommand(projectID, projectRegion, projectRegistryID, projectDeviceID, commandID)

	if err != nil {
		log.Printf("Error while sending coffee %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Error while sending coffee"))
		if err != nil {
			log.Printf("Error while sending error %v", err)
		}
		return
	}

	_, err = fmt.Fprintln(w, "Sent command to device")
	if err != nil {
		log.Printf("Error while sending back answer to caller %v", err)
	}

	log.Printf("Coffee answer %v", res)
}

// sendCommand sends a command to a device listening for commands.
func sendCommand(projectID string, region string, registryID string, deviceID string, sendData string) (*cloudiot.SendCommandToDeviceResponse, error) {
	ctx := context.Background()
	client, err := cloudiot.NewService(ctx)
	if err != nil {
		return nil, err
	}

	req := cloudiot.SendCommandToDeviceRequest{
		BinaryData: b64.StdEncoding.EncodeToString([]byte(sendData)),
	}

	name := fmt.Sprintf(iotBusPattern, projectID, region, registryID, deviceID)

	response, err := client.Projects.Locations.Registries.Devices.SendCommandToDevice(name, &req).Do()
	if err != nil {
		return nil, err
	}

	return response, nil
}
