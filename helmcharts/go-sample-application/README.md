# Пример развертывания простого приложения
### Новая установка приложения используя helm

#### Установка приложения со значениями переменных из файла values.yaml
```shell
helm upgrade --install --create-namespace --namespace go-sample-app go-sample-app helmcharts/go-sample-application
```
#### Установка приложения со значениями переменных из файла values-dev.yaml
```shell
helm upgrade --install --create-namespace --namespace go-sample-app go-sample-app helmcharts/go-sample-application \
--values values-dev.yaml
```
#### Установка приложения со значениями переменных из файла values.yaml и параметром командной строи --set  
```shell
helm upgrade --install --create-namespace --namespace go-sample-app go-sample-app helmcharts/go-sample-application \
--set backend.env.banner="SIMPLE APPLICATION for DEVELOPMENT"
```



## Структура приложения:
- BACKEND Простое приложение backend_server.py выводит информацию о переменных окружения среды исполнения 
  в формате JSON. 
  Приложение выполняется в контейнере из образа python.
  Приложение доступно по порту 8070.
  API:
  hostname:8070/       - {"ok": "Ok"}           
  hostname:8070/info   - Переменные среды ОС
  hostname:8070/health - {"health": "healthy"}     
  hostname:8070/about  - {"/": "Return Ok", "/info": "Return OS environment", "/health": "Return healthy", "/about": "Return this info"}   
- WEBSERVER Простое приложение выводит информацию о переменных окружения среды исполнения
  в формате HTML. Принимает запросы на порту 8080 и перенаправляет их на порт 8070 BACKEND  

