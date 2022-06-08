--- consola local
--construir imagen
docker build --tag=img-lab-go-back .

--login en cuenta de dockerhub
docker login

--creacion de etiqueta para dockerhub
docker tag img-lab-go-back rmmdiferentt/img-lab-go-back:v1

--push en dockerhub
docker push rmmdiferentt/img-lab-go-back:v1


--- consola server con docker
--crear e iniciar docker con la imagen de dockerhub, expone puerto 5000 del server a los clientes
docker run -p 5000:3000 -d rmmdiferentt/img-lab-go-back:v1

--iniciar bash en el contenedor de docker
docker exec -it NOMBRE_CONTENEDOR bash

