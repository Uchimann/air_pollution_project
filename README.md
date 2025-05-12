# Hava Kirliliği İzleme Sistemi

Bu proje, dünya genelinde hava kirlilik verilerini toplayan, analiz eden ve görselleştiren web tabanlı bir platform sunmaktadır. Sistem, kirlilik verilerini işleyerek belirli eşik değerlerini aşan kirlilik seviyelerini tespit eder ve kullanıcıları gerçek zamanlı olarak uyarır.


## Proje Yapısı

Bu proje, modern bir microservice mimarisi kullanarak geliştirilmiştir ve üç ana servisten oluşmaktadır:

1. **Data Collector**: Kirlilik verilerini toplayan ve veritabanına kaydeden servis
2. **Analyzer**: Kirlilik verilerini analiz eden ve anomalileri tespit eden servis
3. **Notifier**: Tespit edilen anomalileri kullanıcılara bildiren servis
4. **Frontend**: Kullanıcı arayüzü

## Teknoloji Seçimleri

### Backend
- **Dil**: Go (Golang)
- **Mimari**: Microservice (3 servis + frontend)
- **API Framework**: Fiber (Data Collector)
- **Veritabanı**: TimescaleDB (PostgreSQL tabanlı zaman serisi veritabanı)
- **Kuyruklama Sistemi**: RabbitMQ
- **ORM**: GORM

### Frontend
- **Framework**: React.js
- **Harita Kütüphanesi**: Leaflet.js + Leaflet.Heat
- **Grafik Kütüphanesi**: Chart.js + React-Chartjs-2
- **SSE (Server-Sent Events)**: Gerçek zamanlı bildirimler için

### Containerization
- Docker ve Docker Compose

## Sistem Mimarisi



1. **Veri Akışı**:
   - Kirlilik verileri manuel olarak ya da test scripti ile sisteme eklenir
   - Data Collector servisi verileri alır, veritabanına kaydeder ve RabbitMQ'ya gönderir
   - Analyzer servisi kuyruktaki verileri işler ve anomalileri tespit eder
   - Anomali tespit edildiğinde, Analyzer sonuçları RabbitMQ üzerinden Notifier servisine gönderir
   - Notifier servisi, SSE ile frontend'e gerçek zamanlı bildirim gönderir
   - Frontend harita ve uyarı panelini günceller
![Screenshot_49](https://github.com/user-attachments/assets/8af87acb-bfed-40a1-80d8-6c20bf57a36a)

[Excalidraw](https://excalidraw.com/#json=iAGSXPdfYvJ6wuQgzg_C4,ywCkswDoKXX247J4h0vN8w)


## Kurulum

### Ön Koşullar
- Docker ve Docker Compose kurulu olmalıdır
- Git kurulu olmalıdır

### Kurulum Adımları

1. Repoyu klonlayın:
```bash
git clone https://github.com/uchimann/air_pollution_project.git
cd air_pollution_project
```

2. Docker Compose ile tüm servisleri başlatın:
```bash
docker-compose up --build
```

3. Sistemin başlaması için birkaç saniye bekleyin. Tüm servisler hazır olduğunda, frontend'e erişebilirsiniz:
```
http://localhost:3000
```

## Kullanım Rehberi

### Manuel Veri Girişi
Manuel olarak kirlilik verisi girmek için manuel-input.sh scriptini kullanabilirsiniz:

```bash
./manuel-input.sh <latitude> <longitude> <parameter> <value>
```

Örnek:
```bash
./manuel-input.sh 41.0082 28.9784 "PM2.5" 35.4
```

Geçerli parametre değerleri: PM2.5, PM10, NO2, SO2, O3

### Otomatik Test

Sistemi test etmek için auto-test.sh scriptini kullanabilirsiniz:

```bash
./auto-test.sh [options]
```

Opsiyonlar:
- `--duration=<seconds>`: Script'in çalışma süresi (saniye)
- `--rate=<requests_per_second>`: Saniyede kaç istek atılacağı
- `--anomaly-chance=<percentage>`: Anomali oluşturma olasılığı (%)

Örnek:
```bash
./auto-test.sh --duration=300 --rate=2 --anomaly-chance=15
```

## API Dokümantasyonu

### Data Collector API

#### Kirlilik Verisi Eklemek

**Endpoint**: `POST /api/pollution`

**Payload**:
```json
{
  "latitude": 41.0082,
  "longitude": 28.9784,
  "pollutant": "PM2.5",
  "value": 35.4
}
```

**Cevap**:
```json
{
  "data": {
    "timestamp": "2023-08-15T10:30:45Z",
    "latitude": 41.0082,
    "longitude": 28.9784,
    "pollutant": "PM2.5",
    "value": 35.4
  }
}
```

## SSE (Server-Sent Events)

Frontend, Notifier servisinden gerçek zamanlı güncellemeleri almak için SSE kullanır.

**SSE Endpoint**: `http://localhost:8081/events`

**Format**:
```json
data: {
  "id": 1,
  "pollution_data_id": 123,
  "timestamp": "2023-08-15T10:30:45Z",
  "latitude": 41.0082,
  "longitude": 28.9784,
  "pollutant": "PM2.5",
  "value": 35.4,
  "threshold_value": 12.0,
  "is_anomalous": true,
  "anomaly_level": "Medium",
  "health_risk_level": "Unhealthy"
}
```

## Anomali Tespit Kriterleri

Sistem, aşağıdaki kriterlere göre hava kalitesi anomalilerini tespit eder:

1. **Threshold değerlendirmesi**:
   - Her kirletici için WHO standartlarına dayalı eşik değerleri (PM2.5, PM10, NO2, SO2, O3 için)
   - Üç seviyeli anomali tespiti: Düşük, Orta, Yüksek
   - Üç seviyeli sağlık riski: Orta, Sağlıksız, Tehlikeli

2. **Anomali seviyeleri**:
   - Düşük: Orta seviye eşiğini aşan değerler
   - Orta: Sağlıksız seviye eşiğini aşan değerler
   - Yüksek: Tehlikeli seviye eşiğini aşan değerler

## Servis Detayları

### Data Collector
- Port: 8080
- Kirlilik verilerini HTTP API üzerinden kabul eder
- Verileri PostgreSQL/TimescaleDB'ye kaydeder
- Verileri RabbitMQ'ya yayınlar

### Analyzer
- Kirlilik verilerini işler ve anomalileri tespit eder
- Anomali tespit edildiğinde sonuçları RabbitMQ üzerinden Notifier servisine gönderir

### Notifier
- Port: 8081
- RabbitMQ'dan anomali bildirimleri alır
- SSE yoluyla frontend'e gerçek zamanlı bildirimler gönderir

### Frontend
- Port: 3000
- Harita üzerinde kirlilik yoğunluğunu görselleştirir
- Zaman içindeki kirlilik seviyelerini gösteren grafikler sunar
- Gerçek zamanlı anomali uyarıları gösterir

## Sorun Giderme (Troubleshooting)

### Servisler başlamıyor
```bash
# Tüm container'ları durdurup yeniden başlatın
docker-compose down
docker-compose up -d

# Container'ların durumunu kontrol edin
docker-compose ps

# Hata loglarını kontrol edin
docker-compose logs -f data-collector
docker-compose logs -f analyzer
docker-compose logs -f notifier
docker-compose logs -f frontend
```

### Frontend haritasında veri görünmüyor
1. Notifier servisinin çalışıp çalışmadığını kontrol edin: `docker-compose logs -f notifier`
2. Tarayıcının geliştirici konsolunda SSE bağlantı hatalarını kontrol edin
3. auto-test.sh ile test verisi gönderin

### RabbitMQ ile ilgili sorunlar
RabbitMQ yönetim arayüzüne erişerek kuyrukları kontrol edin:
```
http://localhost:15672
Kullanıcı adı: guest
Şifre: guest
```

### Frontend SSE bağlantı hataları
CORS ayarlarını kontrol edin. Notifier servisinin header'larında `Access-Control-Allow-Origin` olduğundan emin olun.

## Geliştirme

Geliştirme için her servisin kendi klasöründe çalışabilirsiniz. Docker Compose kullanmadan geliştirmek için:

1. PostgreSQL ve RabbitMQ için Docker kullanın:
```bash
docker-compose up -d postgres rabbitmq
```

2. Her servisi ayrı terminallerde çalıştırın:
```bash
# Data Collector
cd data-collector
go run ./cmd/app/main.go

# Analyzer
cd analyzer
go run ./cmd/app/main.go

# Notifier
cd notifier
go run ./cmd/app/main.go

# Frontend
cd frontend
npm start
```

