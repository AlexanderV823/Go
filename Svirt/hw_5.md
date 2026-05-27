# Домашнее задание к занятию «Kubernetes. Часть 1»

## Задача 1

1. Запустите Kubernetes локально, используя k3s или minikube на свой выбор.

* Установка Minikube
cd ~ && wget -O minikube.deb https://github.com/kubernetes/minikube/releases/latest/download/minikube-linux-amd64

cd ~ && sudo install minikube.deb /usr/local/bin/minikube

* Установка kubectl

sudo snap install kubectl --classic

* Запуск локального кластера

minikube start

* Проверка работы

kubectl get nodes

![Скрин 1](./hw5-1_1.jpg)

* Остановка и удаление

minikube stop
minikube delete

2. Добейтесь стабильной работы всех системных контейнеров.

* Создаем группу docker (если она еще не создана)

sudo groupadd docker

* Добавляем вашего текущего пользователя в эту группу

sudo usermod -aG docker $USER

* Применяем изменения прав немедленно без перезагрузки системы

newgrp docker

* Первый чистый запуск Minikube

minikube delete
minikube start --driver=docker --addons=ingress,dashboard

* Проверка стабильности системных контейнеров

kubectl get pods -n kube-system

![Скрин 2](./hw5-1_2.jpg)

## Задача 2

1. Измените [файл](redis-app-base.yaml) с учётом условий:

* redis должен запускаться без пароля;
* создайте Service, который будет направлять трафик на этот Deployment;
* версия образа redis должна быть зафиксирована на 6.0.13.

[Измененный файл](redis-app.yaml)

2. Запустите Deployment в своём кластере и добейтесь его стабильной работы.

* Локальное скачивание образа через Docker

docker pull redis:6.0.13

* Перенос образа внутрь Minikube

minikube image load redis:6.0.13

* Запуск деплоймента

kubectl apply -f redis-app.yaml

В консоль должно вывестись:

deployment.apps/redis created
service/redis-service created

![Скрин 3](./hw5-1_3.jpg)

* Контроль стабильности и проверка работы

kubectl get pods -l app=redis -w

Дождитесь, пока статус изменится на Running, а в столбце READY появится 1/1. Чтобы выйти из режима мониторинга, нажмите Ctrl + C.

![Скрин 4](./hw5-1_4.jpg)

Убедитесь, что служба создана и имеет внутренний IP-адрес.

kubectl get svc redis-service

![Скрин 5](./hw5-1_5.jpg)

* Проверка логов контейнера (для подтверждения беспарольного режима)

kubectl logs deployment/redis

В логах вы увидите приветственное сообщение Redis и строку, подтверждающую готовность принимать соединения на порту 6379.

![Скрин 6](./hw5-1_6.jpg)