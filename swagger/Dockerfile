FROM ubuntu:20.04
RUN apt-get update -y && apt-get install -y wget unzip sed
RUN wget -O swagger.zip https://github.com/swagger-api/swagger-ui/archive/refs/tags/v3.49.0.zip \
    && unzip swagger.zip \
    && rm -rf ./swagger-ui-3.49.0/docs \
    && mv ./swagger-ui-3.49.0/dist ./swagger-ui-3.49.0/docs \
    && sed -i 's/https:\/\/petstore.swagger.io\/v2\/swagger.json/definition.yml/g' ./swagger-ui-3.49.0/docs/index.html \
    && wget -O ./swagger-ui-3.49.0/docs/definition.yml https://idratherbewriting.com/learnapidoc/docs/rest_api_specifications/openapi_openweathermap.yml
CMD ["cp", "-r", "./swagger-ui-3.49.0/docs", "/swagger"]
