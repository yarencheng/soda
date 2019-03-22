# SODA

## Build

```

docker build --tag soda .

```

## Run

```

docker run -it --rm soda

```

## Debug

```

docker build --tag soda . && docker run -it --rm -p 8080:8080 soda


curl http://localhost:3000/photo -X POST -F "file=@/home/arenx/soda/color.png" -H "Content-Type: multipart/form-data"
export ID="faba9abb-c879-4d3e-9810-0afb72040abd"
curl http://localhost:3000/photo/${ID}/photo -X GET
curl http://localhost:3000/photo/${ID} -X GET
curl http://localhost:3000/photo/${ID} -X PUT -d '{ "title": "tttttt", "description": "dddddd" }'
curl http://localhost:3000/photo -X GET
curl http://localhost:3000/photo/${ID} -X DELETE
curl http://localhost:3000/photo -X GET


```