# Домашнее задание к занятию «Docker. Часть 2»

## Задача 1

Docker Compose упрощает запуск и управление сразу несколькими контейнерами в различных средах. Можно создавать собственные файлы конфигурации, например, YAML.
Т.е. чтобы запустить приложение, требующее запуск нескольких контейнеров, не нужно поочередно выполнять команды docker run, а выполнить одну docker compose.
В файде конфигурации описано что и в какой последовательности должеон быть запущено.

## задача 2

version: '3'
services:
volumes:
networks:
  my-netology-hw:
    name: volodin-aa-my-netology-hw
    driver: bridge
    ipam:
    config:
      - subnet: 10.5.0.0/16
gateway: 10.5.0.1

## Задача 3

version: '3'
services:
  prometheus:
    image: prom/prometheus:latest
    container_name: volodin-aa-netology-prometheus
    ports:
      - "9090:9090"
    networks:
      volodin-aa-my-netology-hw:
        ipv4_address: 10.5.0.10
    volumes:
      # 1. Проброс файла конфигурации (из папки с проектом в контейнер)
      # :ro (read-only), чтобы не затереть исходны файл
      - ./prometheus.yml:/etc/prometheus/prometheus.yml:ro
      # 2. Проброс именованного тома для данных (чтобы метрики не удалялись)
      - prometheus_data:/prometheus
networks:
  my-netology-hw:
    name: volodin-aa-my-netology-hw
    driver: bridge
    ipam:
    config:
      - subnet: 10.5.0.0/16
gateway: 10.5.0.1
volumes:
  prometheus_data: