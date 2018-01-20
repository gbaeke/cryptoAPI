# cryptoAPI
Simple example with goa plus build and run on Kubernetes

* Install latest Go (had issue with 1.7; went away with 1.9)
* Install goa with go get -u github.com/goadesign/goa/...
* Modify the design/design.go file using the goa dsl
* Code generation with: goagen bootstrap -d crypto\design
  * This creates a bunch of code plus a main.go & currencies.go
* Modify currencies.go with your implementation
  * See currencies.go for a simple retrieval of cryptocurrency info from coinmarketcap.com
* build with **go build**
  * static build for linux with **CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' .**
* build Docker image with **docker build -t yourtag .**
* push container to ACR
  * az acr login -name name-of-Azure-container-registry
  * docker push name-of-your-image
    * tag should be DNS-name-of-ACR/image-name
 * create Kubernetes deployment
   * kubectl apply -f crypto.yaml
 * expose service
   * kubectl expose deployment crypto --type LoadBalancer
