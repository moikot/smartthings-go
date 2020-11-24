# smartthings-go
An automatically generated Go client for Samsung SmartThings API.

# Generate

In order to update models and clients you need to clone the repository and run go-swagger.

```
clone https://github.com/moikot/smartthings-go.git
cd smartthings-go
docker run -it --rm -v $(pwd):/go/src/github.com/moikot/smartthings-go \
   quay.io/goswagger/swagger generate client \
   -f https://swagger.api.smartthings.com/public/st-api.yml \
   --skip-validation -t=/go/src/github.com/moikot/smartthings-go
```
