# Deploy function

    gcloud alpha functions deploy download \
      --entry-point Download \
      --memory 128MB \
      --env-vars-file envar.yaml \
      --region europe-west1 \
      --runtime go111 \
      --allow-unauthenticated \
      --trigger-http

# Configuration

In a separate `envar.yaml` file :

    FIRMWARE_BUCKET: "bucketname"
    COFFEE_FIRMWARE: "firmwarenameinbucketforcoffee.bin"
    LIGHT_FIRMWARE: "firmwarenameinbucketforlight.bin"
   
# Test function

Parameter `-L` to follow the redirection.

    curl -L https://project-url/download?mac=azerrtyuiiop&type=light