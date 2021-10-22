# Deploy function

    gcloud alpha functions deploy bicycle-data-store \
      --entry-point Bicycle \
      --memory 128MB \
      --env-vars-file envar.yaml \
      --region europe-west1 \
      --runtime go116 \
      --trigger-topic bicycle-data \
      
# Configuration

In a separate `envar.yaml` file :

    PROJECT_ID: "myid"
    PROJECT_REGION: "europe-west1"
    PROJECT_REGISTRY_ID: "myregistryid"
    PROJECT_DEVICE_ID: "mydeviceid"`
    
# Test function

Parameter `-L` to follow the redirection.

    curl -L https://project-url/coffee