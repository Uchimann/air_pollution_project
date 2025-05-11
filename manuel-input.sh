#!/bin/bash

if [ $# -ne 4 ]; then
    echo "Kullanım: $0 <latitude> <longitude> <parameter> <value>"
    echo "Geçerli parametre değerleri: PM2.5, PM10, NO2, SO2, O3"
    exit 1
fi

latitude=$1
longitude=$2
parameter=$3
value=$4

valid_params=("PM2.5" "PM10" "NO2" "SO2" "O3")
valid=0

for param in "${valid_params[@]}"; do
    if [ "$parameter" == "$param" ]; then
        valid=1
        break
    fi
done

if [ $valid -eq 0 ]; then
    echo "Hata: Geçersiz parametre. Lütfen aşağıdakilerden birini kullanın:"
    echo "PM2.5, PM10, NO2, SO2, O3"
    exit 1
fi

if ! [[ $latitude =~ ^-?[0-9]+(\.[0-9]+)?$ ]] || ! [[ $longitude =~ ^-?[0-9]+(\.[0-9]+)?$ ]] || ! [[ $value =~ ^[0-9]+(\.[0-9]+)?$ ]]; then
    echo "Hata: Latitude, longitude ve value sayısal değerler olmalıdır."
    exit 1
fi

json_data="{
    \"latitude\": $latitude,
    \"longitude\": $longitude,
    \"pollutant\": \"$parameter\",
    \"value\": $value
}"

echo "Gönderilecek veri: $json_data"

response=$(curl -s -X POST \
    -H "Content-Type: application/json" \
    -d "$json_data" \
    "http://localhost:8080/api/pollution")

if [ $? -eq 0 ]; then
    echo "Başarılı! Sunucu yanıtı: $response"
else
    echo "Hata: API'ye veri gönderilemedi."
fi