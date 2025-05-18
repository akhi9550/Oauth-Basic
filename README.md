## Getting Started

1. Clone this repository

   ```bash
   git clone https://github.com/akhi9550/OauthBasic.git
   ```

2. Download the requirements

   ```bash
   go mod download
   go mod tidy
   ```

3. Get Client ID and Client Secret from the Oauth provider. I am using Google for example. Go to the [Google cloud credential page](https://console.cloud.google.com/apis/credentials) for getting client id and secret.

4. Add the Client ID and secret as an environment variable and run the application

   ```bash
   CLIENT_ID="clientid" CLIENT_SECRET_ID="secret" 
   ```
5. Run the Code 

```bash
go run .
```