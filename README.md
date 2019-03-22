# SODA

## Usage

* Set IP to ENV for easy to use
  ```
  export SODA_IP="35.221.180.16"
  ```

* Upload photo
    ```
    $ curl http://${SODA_IP}/photo -X POST -F "file=@/home/arenx/soda/color.png" -H "Content-Type: multipart/form-data" | jq
    % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                    Dload  Upload   Total   Spent    Left  Speed
    100 87414  100   164  100 87250   2733  1420k --:--:-- --:--:-- --:--:-- 1446k
    {
      "data": {
        "ID": "b01a990d-12fb-43f1-90b5-c0217793ac3a",
        "Title": "empty",
        "Description": "empty",
        "File": "photo/b01a990d-12fb-43f1-90b5-c0217793ac3a/photo"
    },
      "status": "ok"
    }
    ```

* Get photo information
    ```
    $ curl http://${SODA_IP}/photo/b01a990d-12fb-43f1-90b5-c0217793ac3a -X GET | jq
    % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                    Dload  Upload   Total   Spent    Left  Speed
    100   164  100   164    0     0   7809      0 --:--:-- --:--:-- --:--:--  7809
    {
      "data": {
        "ID": "b01a990d-12fb-43f1-90b5-c0217793ac3a",
        "Title": "empty",
        "Description": "empty",
        "File": "photo/b01a990d-12fb-43f1-90b5-c0217793ac3a/photo"
    },
      "status": "ok"
    }
    ```

* Change title or description
    ```
    $ curl http://${SODA_IP}/photo/b01a990d-12fb-43f1-90b5-c0217793ac3a -X PUT -d '{ "title": "tttttt", "description": "dddddd" }' | jq
    % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                    Dload  Upload   Total   Spent    Left  Speed
    100   212  100   166  100    46   7904   2190 --:--:-- --:--:-- --:--:-- 10095
    {
      "data": {
        "ID": "b01a990d-12fb-43f1-90b5-c0217793ac3a",
        "Title": "tttttt",
        "Description": "dddddd",
        "File": "photo/b01a990d-12fb-43f1-90b5-c0217793ac3a/photo"
    },
      "status": "ok"
    }
    ```

* List all photos
    ```
    $ curl http://${SODA_IP}/photo -X GET | jq
    % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                    Dload  Upload   Total   Spent    Left  Speed
    100   452  100   452    0     0  21523      0 --:--:-- --:--:-- --:--:-- 22600
    {
        "data": [
            {
                "ID": "20ff196f-2e5b-4298-a6ae-f843af796bf7",
                "Title": "empty",
                "Description": "empty",
                "File": "photo/20ff196f-2e5b-4298-a6ae-f843af796bf7/photo"
            },
            {
                "ID": "b01a990d-12fb-43f1-90b5-c0217793ac3a",
                "Title": "tttttt",
                "Description": "dddddd",
                "File": "photo/b01a990d-12fb-43f1-90b5-c0217793ac3a/photo"
            },
            {
                "ID": "e3110841-4774-41e1-bedd-2e19d00122c0",
                "Title": "empty",
                "Description": "empty",
                "File": "photo/e3110841-4774-41e1-bedd-2e19d00122c0/photo"
            }
    ],
        "status": "ok"
    }
    ```

* Delete a photo
    ```
    $ curl http://${SODA_IP}/photo/b01a990d-12fb-43f1-90b5-c0217793ac3a -X DELETE | jq
    % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                    Dload  Upload   Total   Spent    Left  Speed
    100    15  100    15    0     0    750      0 --:--:-- --:--:-- --:--:--   750
    {
        "status": "ok"
    }
    ```

* Download photo
    ```
    wget http://${SODA_IP}/photo/e3110841-4774-41e1-bedd-2e19d00122c0/photo
    ```

## Build

```

docker build --tag soda .

```

## Local run

```

docker run -it --rm -p 80:8080 soda

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