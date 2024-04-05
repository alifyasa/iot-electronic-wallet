# IoT: Electronic Wallet

Dibuat untuk UTS IF4051 Pengembangan Sistem IoT Semester 2 2023/2024.

## Denah Proyek

```plaintext
.
├── LICENSE
├── README.md
├── cmd
│   ├── logging.go
│   ├── main.go
│   ├── wallet.go
│   └── websockets.go
├── go.mod
├── go.sum
├── microcontroller
│   └── esp32.py
├── run.sh
└── static
    └── index.html
```

Kode backend terdapat di folder `cmd`. Kode mikrokontroller terdapat pada folder `microcontroller`.

## Cara Menjalankan

Untuk menjalankan server, gunakan `run.sh`.

```shell
chmod +x run.sh
./run.sh
```

Untuk program mikrokontroller, gunakan [Thonny](https://thonny.org/).

## Identitas

| Nama | NIM |
| ---- | --- |
| Muhammad Alif Putra Yasa | 13520135 |
