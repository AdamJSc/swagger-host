FROM ubuntu:20.04
RUN apt-get update -y && apt-get install -y wget unzip sed
RUN wget -O swagger.zip https://github.com/swagger-api/swagger-ui/archive/refs/tags/v3.49.0.zip \
    && unzip swagger.zip \
    && sed -i 's/https:\/\/petstore.swagger.io\/v2\/swagger.json/doc.yml/g' ./swagger-ui-3.49.0/dist/index.html
CMD ["cp", "-r", "./swagger-ui-3.49.0/dist/", "/swagger/docs"]
