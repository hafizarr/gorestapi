# Full Rest API Golang dan Mysql
Full Rest API product dengan bahasa pemrograman Go dan RDMS Mysql

## Quick Start
Semua Endpoint pada aplikasi ini diuji menggukan Aplikasi [Postman](https://www.getpostman.com/).

**1. Mengunduh project**
```bash
# mengunduh project ke GOPATH/src
go get github.com/hafizarr/gorestapi

```

**2. Menjalankan Aplikasi**

#### Linux
```bash
# pindah ke folder GOPATH
cd $GOPATH/github.com/hafizarr/gorestapi

# import table
import file products.sql ke database 

# Memulai aplikasi
make start
atau
go run main.go
```

#### Windows
```bash
# pindah ke folder GOPATH
cd %GOPATH%\src\github.com/hafizarr/gorestapi

# import table
import file products.sql ke database 

# Memulai aplikasi
go run main.go
```

# Endpoints
- API Add product
- API List product dengan sorting

## Products

| Route | Description | params |
| --- | --- | --- |
| `GET  products` | Mengembalikan semua `product` urut berdasarkan product terbaru | None |
| `GET  products?sortBy=lowest-price` | Mengembalikan semua `product` urut berdasarkan harga terendah | Product Price |
| `GET  products?sortBy=highest-price` | Mengembalikan semua `product` urut berdasarkan harga tertinggi | Product Price |
| `GET  products?sortBy=a-z` | Mengembalikan semua `product` urut berdasarkan nama produk a-z | Product Name |
| `GET  products?sortBy=z-a` | Mengembalikan semua `product` urut berdasarkan nama produk z-a | Product Name |
| `POST products/create` | Memasukkan `product` baru ke database | None |
| `PUT products/update{id}` | Meng-update `product` sesuai `product_id` | Product Id |
| `DELETE products/delete/{id}` | Menghapus satu `product` sesuai `product_id` | Product Id |

<br>

Contoh Request Sort By Product Terbaru :
```http
GET     http://localhost:7000/products
```
Respond :
```json
[
    {
        "product_id": 4,
        "product_name": "Xiaomi Black Shark 4 8/128GB",
        "product_price": 8900000,
        "product_description": "Ukuran layar: 6.67 inci, 1080 x 2400 pixels, Super AMOLED, 144Hz",
        "product_quantity": 0
    },
    {
        "product_id": 2,
        "product_name": "Huawei MatePad 11 6/128GB - Matte Grey Free Keyboard + Mouse",
        "product_price": 7299000,
        "product_description": "Ukuran layar: 10.95 inci, 2560 x 1600 pixels, IPS LCD, 120Hz\r\nMemori: RAM 6 GB, ROM 128 GB, MicroSD up to 1TB\r\nSistem operasi: HarmonyOS 2\r\nCPU: Qualcomm Snapdragon 865 Octa-core (1x2.84 GHz Kryo 585 & 3x2.42 GHz Kryo 585 & 4x1.8 GHz Kryo 585)\r\nGPU: Adreno 650\r\nKamera: 13 MP, f/1.8, PDAF, depan 8 MP, f/2.0\r\nWiFi Only\r\nBaterai: Li-Po 7250 mAh, non-removable\r\nBerat: 485 gram\r\nGaransi Resmi.",
        "product_quantity": 100
    },
    {
        "product_id": 1,
        "product_name": "Samsung Galaxy Note20 Ultra 5G 12/512GB - Mystic Bronze Free E Voucher + Samsung Care+*",
        "product_price": 1899000,
        "product_description": "Ukuran layar: 6.9 inci, 1440 x 3088 pixels, Dynamic AMOLED 2X capacitive touchscreen, 16M colors\r\nMemori: RAM 12 GB, ROM 512 GB, MicroSD Slot\r\nSistem operasi: Android 11\r\nCPU: Exynos 990 (7 nm+) Octa-core\r\nGPU: Mali-G77 MP11\r\nKamera: Triple 108 MP, f/1.8; 12 MP, f/3.0; 12 MP, f/2.0, depan 10 MP, f/2.2\r\nSIM: Hybrid Dual SIM (Nano-SIM)\r\nBaterai: Non-removable Li-Ion 4500 mAh\r\nBerat: 208 gram\r\nGaransi Resmi",
        "product_quantity": 100
    }
]

```

Contoh Request Sort By Harga Terendah :
```http
GET     http://localhost:7000/products?shortBy=lowest-price
```
Respond :
```json
[
    {
        "product_id": 1,
        "product_name": "Samsung Galaxy Note20 Ultra 5G 12/512GB - Mystic Bronze Free E Voucher + Samsung Care+*",
        "product_price": 1899000,
        "product_description": "Ukuran layar: 6.9 inci, 1440 x 3088 pixels, Dynamic AMOLED 2X capacitive touchscreen, 16M colors\r\nMemori: RAM 12 GB, ROM 512 GB, MicroSD Slot\r\nSistem operasi: Android 11\r\nCPU: Exynos 990 (7 nm+) Octa-core\r\nGPU: Mali-G77 MP11\r\nKamera: Triple 108 MP, f/1.8; 12 MP, f/3.0; 12 MP, f/2.0, depan 10 MP, f/2.2\r\nSIM: Hybrid Dual SIM (Nano-SIM)\r\nBaterai: Non-removable Li-Ion 4500 mAh\r\nBerat: 208 gram\r\nGaransi Resmi",
        "product_quantity": 100
    },
    {
        "product_id": 2,
        "product_name": "Huawei MatePad 11 6/128GB - Matte Grey Free Keyboard + Mouse",
        "product_price": 7299000,
        "product_description": "Ukuran layar: 10.95 inci, 2560 x 1600 pixels, IPS LCD, 120Hz\r\nMemori: RAM 6 GB, ROM 128 GB, MicroSD up to 1TB\r\nSistem operasi: HarmonyOS 2\r\nCPU: Qualcomm Snapdragon 865 Octa-core (1x2.84 GHz Kryo 585 & 3x2.42 GHz Kryo 585 & 4x1.8 GHz Kryo 585)\r\nGPU: Adreno 650\r\nKamera: 13 MP, f/1.8, PDAF, depan 8 MP, f/2.0\r\nWiFi Only\r\nBaterai: Li-Po 7250 mAh, non-removable\r\nBerat: 485 gram\r\nGaransi Resmi.",
        "product_quantity": 100
    },
    {
        "product_id": 4,
        "product_name": "Xiaomi Black Shark 4 8/128GB",
        "product_price": 8900000,
        "product_description": "Ukuran layar: 6.67 inci, 1080 x 2400 pixels, Super AMOLED, 144Hz",
        "product_quantity": 0
    }
]

```

Contoh Request Sort By Harga Tertinggi :
```http
GET     http://localhost:7000/products?shortBy=lowest-price
```
Respond :
```json
[
    {
        "product_id": 4,
        "product_name": "Xiaomi Black Shark 4 8/128GB",
        "product_price": 8900000,
        "product_description": "Ukuran layar: 6.67 inci, 1080 x 2400 pixels, Super AMOLED, 144Hz",
        "product_quantity": 0
    },
    {
        "product_id": 2,
        "product_name": "Huawei MatePad 11 6/128GB - Matte Grey Free Keyboard + Mouse",
        "product_price": 7299000,
        "product_description": "Ukuran layar: 10.95 inci, 2560 x 1600 pixels, IPS LCD, 120Hz\r\nMemori: RAM 6 GB, ROM 128 GB, MicroSD up to 1TB\r\nSistem operasi: HarmonyOS 2\r\nCPU: Qualcomm Snapdragon 865 Octa-core (1x2.84 GHz Kryo 585 & 3x2.42 GHz Kryo 585 & 4x1.8 GHz Kryo 585)\r\nGPU: Adreno 650\r\nKamera: 13 MP, f/1.8, PDAF, depan 8 MP, f/2.0\r\nWiFi Only\r\nBaterai: Li-Po 7250 mAh, non-removable\r\nBerat: 485 gram\r\nGaransi Resmi.",
        "product_quantity": 100
    },
    {
        "product_id": 1,
        "product_name": "Samsung Galaxy Note20 Ultra 5G 12/512GB - Mystic Bronze Free E Voucher + Samsung Care+*",
        "product_price": 1899000,
        "product_description": "Ukuran layar: 6.9 inci, 1440 x 3088 pixels, Dynamic AMOLED 2X capacitive touchscreen, 16M colors\r\nMemori: RAM 12 GB, ROM 512 GB, MicroSD Slot\r\nSistem operasi: Android 11\r\nCPU: Exynos 990 (7 nm+) Octa-core\r\nGPU: Mali-G77 MP11\r\nKamera: Triple 108 MP, f/1.8; 12 MP, f/3.0; 12 MP, f/2.0, depan 10 MP, f/2.2\r\nSIM: Hybrid Dual SIM (Nano-SIM)\r\nBaterai: Non-removable Li-Ion 4500 mAh\r\nBerat: 208 gram\r\nGaransi Resmi",
        "product_quantity": 100
    }
]

```


Contoh Request Sort By a-z :
```http
GET     http://localhost:7000/products?shortBy=a-z
```
Respond :
```json
[
    {
        "product_id": 2,
        "product_name": "Huawei MatePad 11 6/128GB - Matte Grey Free Keyboard + Mouse",
        "product_price": 7299000,
        "product_description": "Ukuran layar: 10.95 inci, 2560 x 1600 pixels, IPS LCD, 120Hz\r\nMemori: RAM 6 GB, ROM 128 GB, MicroSD up to 1TB\r\nSistem operasi: HarmonyOS 2\r\nCPU: Qualcomm Snapdragon 865 Octa-core (1x2.84 GHz Kryo 585 & 3x2.42 GHz Kryo 585 & 4x1.8 GHz Kryo 585)\r\nGPU: Adreno 650\r\nKamera: 13 MP, f/1.8, PDAF, depan 8 MP, f/2.0\r\nWiFi Only\r\nBaterai: Li-Po 7250 mAh, non-removable\r\nBerat: 485 gram\r\nGaransi Resmi.",
        "product_quantity": 100
    },
    {
        "product_id": 1,
        "product_name": "Samsung Galaxy Note20 Ultra 5G 12/512GB - Mystic Bronze Free E Voucher + Samsung Care+*",
        "product_price": 1899000,
        "product_description": "Ukuran layar: 6.9 inci, 1440 x 3088 pixels, Dynamic AMOLED 2X capacitive touchscreen, 16M colors\r\nMemori: RAM 12 GB, ROM 512 GB, MicroSD Slot\r\nSistem operasi: Android 11\r\nCPU: Exynos 990 (7 nm+) Octa-core\r\nGPU: Mali-G77 MP11\r\nKamera: Triple 108 MP, f/1.8; 12 MP, f/3.0; 12 MP, f/2.0, depan 10 MP, f/2.2\r\nSIM: Hybrid Dual SIM (Nano-SIM)\r\nBaterai: Non-removable Li-Ion 4500 mAh\r\nBerat: 208 gram\r\nGaransi Resmi",
        "product_quantity": 100
    },
    {
        "product_id": 4,
        "product_name": "Xiaomi Black Shark 4 8/128GB",
        "product_price": 8900000,
        "product_description": "Ukuran layar: 6.67 inci, 1080 x 2400 pixels, Super AMOLED, 144Hz",
        "product_quantity": 0
    }
]

```

Contoh Request Sort By z-a :
```http
GET     http://localhost:7000/products?shortBy=z-a
```
Respond :
```json
[
    {
        "product_id": 4,
        "product_name": "Xiaomi Black Shark 4 8/128GB",
        "product_price": 8900000,
        "product_description": "Ukuran layar: 6.67 inci, 1080 x 2400 pixels, Super AMOLED, 144Hz",
        "product_quantity": 0
    },
    {
        "product_id": 1,
        "product_name": "Samsung Galaxy Note20 Ultra 5G 12/512GB - Mystic Bronze Free E Voucher + Samsung Care+*",
        "product_price": 1899000,
        "product_description": "Ukuran layar: 6.9 inci, 1440 x 3088 pixels, Dynamic AMOLED 2X capacitive touchscreen, 16M colors\r\nMemori: RAM 12 GB, ROM 512 GB, MicroSD Slot\r\nSistem operasi: Android 11\r\nCPU: Exynos 990 (7 nm+) Octa-core\r\nGPU: Mali-G77 MP11\r\nKamera: Triple 108 MP, f/1.8; 12 MP, f/3.0; 12 MP, f/2.0, depan 10 MP, f/2.2\r\nSIM: Hybrid Dual SIM (Nano-SIM)\r\nBaterai: Non-removable Li-Ion 4500 mAh\r\nBerat: 208 gram\r\nGaransi Resmi",
        "product_quantity": 100
    },
    {
        "product_id": 2,
        "product_name": "Huawei MatePad 11 6/128GB - Matte Grey Free Keyboard + Mouse",
        "product_price": 7299000,
        "product_description": "Ukuran layar: 10.95 inci, 2560 x 1600 pixels, IPS LCD, 120Hz\r\nMemori: RAM 6 GB, ROM 128 GB, MicroSD up to 1TB\r\nSistem operasi: HarmonyOS 2\r\nCPU: Qualcomm Snapdragon 865 Octa-core (1x2.84 GHz Kryo 585 & 3x2.42 GHz Kryo 585 & 4x1.8 GHz Kryo 585)\r\nGPU: Adreno 650\r\nKamera: 13 MP, f/1.8, PDAF, depan 8 MP, f/2.0\r\nWiFi Only\r\nBaterai: Li-Po 7250 mAh, non-removable\r\nBerat: 485 gram\r\nGaransi Resmi.",
        "product_quantity": 100
    }
]

```

Contoh POST Product :
```http
POST     http://localhost:7000/products/create
```
Contoh :

![source code.](https://github.com/hafizarr/gorestapi/assets/uploads/ss.png)

```
