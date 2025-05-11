#!/bin/bash

DURATION=60
RATE=1
ANOMALY_CHANCE=20

for i in "$@"; do
  case $i in
    --duration=*)
      DURATION="${i#*=}"
      ;;
    --rate=*)
      RATE="${i#*=}"
      ;;
    --anomaly-chance=*)
      ANOMALY_CHANCE="${i#*=}"
      ;;
    *)
      echo "Kullanım: $0 [options]"
      echo "Opsiyonlar:"
      echo "  --duration=<seconds>: Script'in çalışma süresi"
      echo "  --rate=<requests_per_second>: Saniyede kaç istek atılacağı"
      echo "  --anomaly-chance=<percentage>: Anomali oluşturma olasılığı"
      exit 1
      ;;
  esac
done

echo "Test başlıyor..."
echo "Süre: $DURATION saniye"
echo "Hız: $RATE istek/saniye"
echo "Anomali olasılığı: %$ANOMALY_CHANCE"

START_TIME=$(date +%s)
END_TIME=$((START_TIME + DURATION))
REQUEST_COUNT=0
ANOMALY_COUNT=0

POLLUTANTS=("PM2.5" "PM10" "NO2" "SO2" "O3")

while [ $(date +%s) -lt $END_TIME ]; do
  for ((i=0; i<RATE; i++)); do
    LAT=$(awk -v r=$RANDOM 'BEGIN {printf "%.6f", -90 + (r/32767) * 180}')  # -90 ile 90 arası
    LON=$(awk -v r=$RANDOM 'BEGIN {printf "%.6f", -180 + (r/32767) * 360}') # -180 ile 180 arası
    
    POLLUTANT=${POLLUTANTS[$((RANDOM % 5))]}
    
    VALUE=$((RANDOM % 50 + 10))  # 10-60 arası normal değerler
    
    ANOMALY_TEXT=""
    if [ $((RANDOM % 100)) -lt $ANOMALY_CHANCE ]; then
      MULTIPLIER=$(awk -v r=$RANDOM 'BEGIN {printf "%.1f", 3 + (r/32767) * 2}')
      VALUE=$(awk -v v=$VALUE -v m=$MULTIPLIER 'BEGIN {printf "%.1f", v * m}')
      ANOMALY_TEXT=" - ANOMALİ!"
      ANOMALY_COUNT=$((ANOMALY_COUNT + 1))
    fi
    
    JSON_DATA="{\"latitude\": $LAT, \"longitude\": $LON, \"pollutant\": \"$POLLUTANT\", \"value\": $VALUE}"
    
    curl -s -X POST \
      -H "Content-Type: application/json" \
      -d "$JSON_DATA" \
      http://localhost:8080/api/pollution > /dev/null
    
    REQUEST_COUNT=$((REQUEST_COUNT + 1))
    echo "İstek #$REQUEST_COUNT: $POLLUTANT @ ($LAT, $LON) = $VALUE$ANOMALY_TEXT"
  done
  
  sleep 1
done

echo "----------------------------------------"
echo "Test tamamlandı!"
echo "Toplam istek sayısı: $REQUEST_COUNT"
echo "Oluşturulan anomali sayısı: $ANOMALY_COUNT"
echo "Toplam süre: $(($(date +%s) - START_TIME)) saniye"