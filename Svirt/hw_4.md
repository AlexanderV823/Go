# Домашнее задание к занятию «Docker. Часть 2»

## Задача 1

Docker Compose упрощает запуск и управление сразу несколькими контейнерами в различных средах. Можно создавать собственные файлы конфигурации, например, YAML.
Т.е. чтобы запустить приложение, требующее запуск нескольких контейнеров, не нужно поочередно выполнять команды docker run, а выполнить одну docker compose.
В файде конфигурации описано что и в какой последовательности должеон быть запущено.

## Задача 2

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

## Задача 4

pushgateway:
    image: prom/pushgateway:latest
    container_name: volodin-aa-netology-pushgateway
    ports:
    - "9091:9091"
    networks:
    volodin-aa-my-netology-hw:
        ipv4_address: 10.5.0.11
    restart: always

## Задача 5

grafana:
    image: grafana/grafana:latest
    container_name: volodin-aa-netology-grafana
    ports:
      - "80:3000"
    environment:
      - GF_PATHS_CONFIG=/etc/grafana/custom.ini
    volumes:
      - grafana_data:/var/lib/grafana
      - ./custom.ini:/etc/grafana/custom.ini:ro
    networks:
      volodin-aa-my-netology-hw:
        ipv4_address: 10.5.0.12
    restart: always

## Задача 6

Добавлены разделы depends_on: для pushgateway и grafana. Раздел restart: always был сразу добавлен.