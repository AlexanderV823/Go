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