# Deploy function

    gcloud alpha functions deploy light \
      --entry-point Light \
      --memory 128MB \
      --env-vars-file envar.yaml \
      --region europe-west1 \
      --runtime go113 \
      --allow-unauthenticated \
      --trigger-http

# Configuration

In a separate `envar.yaml` file :

    PROJECT_ID: "myid"
    PROJECT_REGION: "europe-west1"
    PROJECT_REGISTRY_ID: "myregistryid"
    PROJECT_DEVICE_ID: "mydeviceid"`

# Test function

Parameter `-L` to follow the redirection.

    curl -L https://project-url/light?word="test"