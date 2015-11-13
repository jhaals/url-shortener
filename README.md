# url-shortener
Simple URL shortener backed by sqlite.


    docker run -dv /local/data/path:/data \
    	-p 1337:1337 \
    	-e BASE_URL=http://mydomain.com \
    	-e DB_PATH=/data \
    	jhaals/url-shortener

