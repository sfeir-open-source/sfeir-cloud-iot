# Deploy the web app

    gcloud app deploy
    
# Configuration

In a `secrets.yaml` file :

    env_variables:
      GARLAND_ACTION_URL: 'https://mylightcloudfunction.dev?word='    
      
      