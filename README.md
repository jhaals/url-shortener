# url-shortener
Simple URL shortener backed by sqlite.

Using the API

        $ curl -X POST http://mydomain.com/save -d '{"url": "http://google.com"}'
        {"error":"","id":"M","url":"http://mydomain.com/M"}

There's also a simple web ui available

#### Run in docker:

    docker run -dv /local/data/path:/data \
    	-p 1337:1337 \
    	-e BASE_URL=http://mydomain.com \
    	-e DB_PATH=/data \
    	jhaals/url-shortener

or use `docker-compose`  

    docker-compose up -d
