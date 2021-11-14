# Simple Echo E-Commerce (Hexagonal)
Simple REST API yang dibuat menggunakan echo web framework.

Implementasi kode terinspirasi pola ```port``` dan ```adapter``` arsitektur ```heksagonal```

![golang clean architecture](Heksagonal.png)

# Fitur REST API
- Registrasi Admin
- Autentikasi Admin dan User
- Create, Read, Update dan Delete Master User
- Create, Read, Update dan Delete Master Produk
- Autentikasi JWT untuk CRUD User dan Produk

# Teknologi dan Framework
- Golang
- Echo Web Framework
- Gorm (ORM)
- PostgreSQL
- Docker Container

# Entity Relationship Diagram (ERD)
![entity relationship diagram](ERD-Echo_E-Commerce.png)

# Use Case Diagram
![UML - Use Case](Use_Case-User-Staff-and-Product.png)

# Activity Diagram - Registrasi User
![UML - Activity Diagram 1](Activity_Diagram-Register_User.png)

# Activity Diagram - Registrasi Master Produk
![UML - Activity Diagram 2](Activity_Diagram-Master_Product.png)

# Menjalankan REST API
Untuk dapat menjalankan layanan ini terdapat 2 cara, pertama melalui lingkungan golang dan kedua melalui sistem container.
Untuk struktur database sudah disiapkan secara langsung ketika service pertama kali dijalankan (sistem migrasi), tetapi terdapat
data yang sifatnya referensi belum tersedia untuk fitur tsb. Untuk memenuhi data tsb kami sudah menyiapkan dummy data pada sumber
```modules/merchant/dummy.sql``` dan ```modules/outlet/dummy.sql```, pastikan Anda berurutan melakukan eksekusi query data dummy.

-   **Menjalankan pada lingkungan golang**<br>
    Sebelum menjalankan perintah dibawah ini, pastikan Anda sudah merubah nama file ```.env_example``` menjadi ```.env```
    dan melakukan penyesuaian konfigurasi sesuai dengan lingkungan yang ada disana.

    Cukup jalankan kode di bawah ini di konsol Anda:
    ```console
    go run .
    ```

    Anda dapat melihat di ```http://localhost:8000``` secara default.

-   **Menjalankan pada lingkungan container**<br>
    Menjalankan service melalui sistem container, dimana service akan dibungkus dalam sebuah sistem container. Semua kebutuhan
    yang berkaitan dengan bahasa golang, pustaka dan perkakas lainnya akan disiapkan dan dibungkus menjadi sebuah image
    container yang siap digunakan. Sebelum ke sana pastikan Anda sudah menginstall [docker](https://docs.docker.com/engine/install/) dan [docker compose](https://docs.docker.com/compose/install/).

    Cukup jalankan kode di bawah ini di konsole Anda:
    ```console
    docker-compose up -d --build
    ```

    Anda dapat melihat di ```http://localhost:8000``` secara default.

    Docker akan melakukan pembuatan image container dan akan membuat instan container yang siap digunakan dibelakang layar.
    Anda tidak perlu repot memikirkan kebutuhan untuk menjalankan service, database dll, docker sudah menyiapkan semua kebutuhan
    itu dan Anda tinggal menunggunakannya saja. Sebagai catatan proses tersebut membutuhkan jaringan internet untuk mendownload
    kebutuhan dalam container dan image yang tidak terdapat pada lokal drive Anda.

    Menon-aktifkan container
    ```console
    docker-compose stop
    ```

    Men-aktifkan container
    ```console
    docker-compose start
    ```

    Menghapus container
    ```console
    docker-compose down
    ```
# Menggunakan REST API
Kami menyiapkan dokumentasi penggunaan REST API dalam bentuk koleksi Postman yang sudah siap digunakan
, sebelum itu pastikan Anda sudah meng-install [Postman](https://www.postman.com/downloads/). Untuk file koleksi
postman dapat di lihat di [postman_collection.json](postman_collection.json).

# Tentang
Saya sangat terbuka untuk menerima kritik dan saran, sungguh suatu kehormatan jika saya bisa berdiskusi dengan Anda lebih lanjut.
Jika ada pertanyaan dan saran bisa langsung di [LinkedIn](https://www.linkedin.com/in/widya-ade-bagus-3a660716b/).