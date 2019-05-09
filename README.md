```
___                        ___                 
|  |                       |  |            ____   
|  |___  ____  ___ _______ |  |___ _______ |  |___
|   _  | |  |_|  | |   __| |  __ / |  \__| |   ___|
|______/ |_______| |_____| |__|\_\ |_____| |____/
```

Bucket is a file server storage based on REST.

### HOW-TO GET STARTED

1. Pull bucket docker image
```
$ docker pull madebyais/bucket
```

2. Create a folder to store the config file, e.g.
```
$ mkdir -p /etc/bucket
```

3. Copy following sample config and save it to `bucket.yaml`
```
$ touch bucket.yaml
```
bucket.yaml
```
local:
    folder: /data/bucket
    bucket:
    - name: sample
      token: 8151325dcdbae9e0ff95f9f9658432dbedfdb209
```

4. Create a folder to store the bucket and its images
```
$ mkdir -p /data/bucket
```

5. Start the server
```
$ docker run -p 8700:8700 -v ${local_folder}:/data/bucket -v ${local_folder_for_config}:/etc/bucket madebyais/bucket
```
Notes
```
${local_folder} replace this with your host folder to store the bucket and images
```
```
${local_folder_for_config} replace this with your host folder to store bucket config file
```

6. You should see something like this
```
BUCKET LOGO
http server started on [::]:8700
```

7. Try to curl the API to upload a file
```
curl -X POST \
  http://localhost:8700/local/sample \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Postman-Token: f7a7a1ea-66b9-496c-88ba-438d738cc01d' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -H 'x-bucket-token: 8151325dcdbae9e0ff95f9f9658432dbedfdb209' \
  -F 'file=@sample.png'
```

8. You will receive following response
```
{
    "data": {
        "url": "/local/sample/Screen-Shot-2019-05-08-at-3.02.20-PM.png"
    },
    "status": "success"
}
```

9. Then you can access the file using this curl as an example
```
curl -X GET \
  'http://localhost:8700/local/sample/sample.png?token=8151325dcdbae9e0ff95f9f9658432dbedfdb209' \
  -H 'Cache-Control: no-cache' \
  -H 'Postman-Token: 246ce5df-cff0-4ab5-9ed7-c3b57ecf8aff'
```

10. or you can open in your browser by copy pasting this url
```
http://localhost:8700/local/sample/sample.png?token=8151325dcdbae9e0ff95f9f9658432dbedfdb209
```