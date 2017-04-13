# golang_microservice
use Go-kit write a microservice; armed with godepth &amp; docker

  
* clone project
~$ mkdir gloang_microservice  
~$ cd gloang_microservice  
~/gloang_microservice$ git clone git@github.com:billiepander/golang_microservice.git  
~/gloang_microservice$ cd golang_microservice  
* create docker image:  
~/gloang_microservice/golang_microservice$ docker build -t `gomicro_image` .  
* create & run container based on upper image:  
~/gloang_microservice/golang_microservice$ docker run --publish 6666:8080 --name `gomicro_container` --rm `gomicro_image`  
  
* finally, following http POST command will return what you want, you can change 小米 to other key word  
~$ curl -XPOST -d'{"s":"小米"}' localhost:6666/getcompany
