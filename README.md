<br />

![alt welcome](./frontend/src/assets/home/welcome.png)


<div align="center">
  <h2 align="center">Tugas Besar 2 IF2211 Strategi Algoritma</h2>

  <p> Dibuat oleh :
    <p>13522062 - Salsabiila | 13522085 - Zahira Dina Amalia | 13522087 - Shulha
    <br />
    <br />
  </p>
</div>


<!-- Daftar Isi -->
<details>
  <summary>Daftar Isi</summary>
  <ol>
    <li>
      <a href="#penjelasan-singkat">Penjelasan Singkat</a>
      <ul>
        <li><a href="#dependansi">Dependansi</a></li>
      </ul>
    </li>
    <li>
      <a href="#cara-penggunaan">Cara Penggunaan</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#instalasi">Instalasi</a></li>
      </ul>
    </li>
    <li><a href="#struktru-program">Struktur Program</a></li>
    <li><a href="#contoh-pemakaian">Contoh Pemakaian</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>


## Penjelasan Singkat

Program ini dibuat berdasarkan spesifikasi dari tugas besar 2 dari mata kuliah IF2211 Strategi Algoritma yang terdapat pada tautan berikut, <a href= https://informatika.stei.itb.ac.id/~rinaldi.munir/Stmik/2023-2024/Tubes2-Stima-2024.pdf> Spesifikasi Tugas Besar 2</a>.

WikiRace atau Wiki Game adalah sebuah permainan yang menggunakan Wikipedia, sebuah situs ensiklopedia daring yang dikelola secara sukarela oleh berbagai orang di seluruh dunia. Dalam permainan ini, pemain dimulai dari satu artikel Wikipedia dan berusaha untuk menavigasi melalui berbagai artikel lainnya dengan mengikuti tautan di dalam setiap artikel, dengan tujuan mencapai artikel tertentu dalam waktu sesingkat mungkin atau dengan jumlah klik yang terbatas.


### Dependansi
Program ini menggunakan beberapa dependansi utama dalam pengembangannya

* <a href=https://nodejs.org/en> Nodejs </a>
* <a href=https://reactjs.org/> React </a>
* <a href=https://go.dev/> Go Language </a>

<p align="right"><a href="#readme-top">kembali</a></p>


## Struktur Program
```
│
├── frontend
│   ├── node_modules
│   |   ├── ...
│   │   └── ...
│   │
│   ├── public
│   |   ├── favicon.ico
│   |   ├── index.html
│   |   ├── logo192.png
│   |   ├── logo512.png
│   |   ├── manifest.json
│   │   └── robots.txt
│   │
│   ├── src
│   |   ├── assets
|   |   |   ├── about
|   |   |   ├── home
|   |   |   ├── navbar
|   |   |   └── race.
|   |   |   
│   |   ├── components
|   |   |   ├── about.js
|   |   |   ├── home.js
|   |   |   ├── navbar.js
|   |   |   └── race.js
|   |   |  
│   |   ├── App.js
│   |   ├── index,js
│   |   ├── styles.css
│   │   └── robots.txt
│   │
│   ├── .gitignore
│   ├── package-lock.json
│   ├── package.json
│   ├── README.md
│   └── yarn.lock
│
├── logic
│   ├── internal
│   │   ├── entities
|   |   |   └── nodes.go
│   │   ├── getPath
|   |   |   ├── bfs.go
|   |   |   └── ids.go
|   |   └── tools
|   |       └── scraping.go
|   |
│   ├── go.mod
│   ├── go.sum
│   └── main.go
│
├── node_modules
│   ├── ...
│   ├── ...
│   └── ...
│
├── image-2.png
├── image-3.png
├── image-4.png
|
├── package-lock.json
│  
├── package.json
|
└── README.md

```

## Cara Penggunaan

Program dapat digunakan dengan melakukan penginstalan dengan mengikuti langkah-langkah di bawah ini.


### Prerequisites

Sebelum melakukan penginstallan, pastikan npm merupakaan versi terbaru dengan mengetikkan hal berikut pada terminal.
* npm
  ```sh
  npm install npm@latest -g
  ```

### Instalasi

Instalasi dapat dilakukan dengan melakukan langkah-langkah berikut.

1. Lakukan `git clone` pada repository ini
   ```sh
   git clone https://github.com/hiirrs/Tubes2_JB-racing.git
   ```
2. Install NPM packages
   ```sh
   npm install
   ```
   
3. Untuk menggunakan program, ketikkan berikut pada terminal folder repository
   ```js
   npm run dev
   ```
<p align="right"><a href="#readme-top">kembali</a></p>


## Contoh Pemakaian
### Landing Page
![alt website](image-2.png)
 - Setelah melakukan `npm run dev`, website akan menampilkan landing page yang memiliki dua buah tombol. 
 - Tombol bernama `Start Race` akan mengarahkan pada halaman race yang merupakan tempat terdapatnya fitur wikirace,  laman race juga dapat diakses dengan menekan tombol `RACE` yang terdapat pada navigation bar.
-  Tombol `Know More` akan mengarahkan ke halaman about us dari website berikut. Laman about us juga dapat diakses melalui tombol `ABOUT US` pada navigation bar.

<br />

### Race Page
![alt website](./image-3.png)
 - Sesuai perintah pada tampilan, pada halaman ini dapat dimasukkan dua buah hal yang memiliki link wikipedia pada en.wikipedia.org/.
 - Halaman ini juga menjadi tempat pemilihan jenis algoritma yang ingin digunakan antara `BFS` dan `IDS`.
 - Setelah melakukan pemilihan jenis algoritma, laman kemudian akan menampilkan tombol `GO` untuk memulai proses pencarian.

 Berikut salah satu contoh penggunaan wikiracing,
 ![alt text](image-4.png)

<br />
Adapun halaman about us berisikan informasi singkat mengenai pembuat program ini.

<p align="right"><a href="#readme-top">kembali</a></p>



## Contributors
|   NIM    |                  Nama                  |
| :------: | :------------------------------------: |
| 13522062 |              Salsabiila                |
| 13522085 |          Zahira Dina Amalia            |
| 13522087 |                Shulha                  |

<p align="right"><a href="#readme-top">kembali</a></p>