# Dasar-dasar Go

---

# Go

- System programming language yang dikembangkan oleh Google
- Pertama kali muncul di tahun 2009
- Multi-paradigm, static typed dan compiled
- Secara fungsi mirip C/C++ namun dengan sintaks yang lebih sederhana

---

# Go

- Dalam prakteknya, Go banyak digunakan di sisi server, baik sebagai back end untuk sebuah aplikasi, hingga untuk melakukan manajemen terhadap infrastruktur, terutama untuk yang berbasis cloud
- Kini, banyak startup yang menggunakan Go sebagai backendnya, hingga memicu permintaan pasar yang tinggi
- Salah satu yang membuatnya menarik adalah kemudahan dalam melakukan pemrograman paralel

---

# RunTime

![runtime](slides/04-dasar-dasar-go/images/runtime.png)

---

# CompileTime

![runtime](slides/04-dasar-dasar-go/images/compiletime.png)

---

# Memulai Menggunakan Go

- Text editor: VS Code dengan plugin Golang
- Terminal
- Environment variable: 

---

# GOPATH dan Workspace

- `GOPATH` adalah sebuah variable yang digunakan golang sebagai rujukan lokasi dimana semua folder project disimpan.
- `workspace` adalah area lokasi folder yang akan dirujuk oleh `GOPATH` yang dijadikan tempat menyimpan sebuah project Go.
- Biasanya `workspace` ini berada di `~/go`
- `workspace` tidak lagi wajib untuk ada, per Go 1.11+

---

# GOPATH variable

Buat direktori go di home kalian

```
$ mkdir -p $HOME/go
```

Tambahkan dua buah variable di dalam bash kalian bisa di `.bash_profile` atau `.zshrc`

```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin 
```

lalu jalankan
```
$ env
```
di terminal

---

# Struktur Folder ~/go

- bin
- pkg
  - darwin_amd64
  - dep
- src
  - github.com
  - github.com
    - {username}/{project} 
  - golang.org
  - gopkg.in

---

# Memulai Menggunakan Go

- Go adalah nama bahasa pemrograman dan juga toolkit yang digunakan untuk mem-_build_ dan menjalankan program Go
- Instal Go dari websitenya [http://www.golang.org](http://www.golang.org)
- Jalankan instalasi
- Verifikasi instalasi dengan perintah `go version`
- Untuk melihat apa saja yang bisa dilakukan Go toolkit, jalankan `go help`

---

# Hello World Go

- Buat direktori baru `~/go/src/github.com/<username>/<project>`
- Ketik kode dan simpan dalam direktori itu dengan nama file `main.go`

```
package main

import "fmt"

// menulis Hello, World

func main() {
    fmt.Println("Hello, World")
}
```

---

# Hello World Go

- Compile program tadi dengan menjalankan

```go
cd ~/go/src/github.com/<username>/<project>
go run main.go
```

---

# Analisa Program

```
package main
```

- Dikenal dengan deklarasi package
- Setiap program Go membutuhkan ini
- Metode untuk mengorganisasi dan menggunakan ulang kode

---

# Analisa Program

Dalam Go, sebuah program bisa berupa  _executables_ ataupun _libraries_

- executables berarti program bisa langsung dijalankan dari command line
- libraries artinya ia bisa dikumpulkan dalam program lain dan dijalankan dari situ

---

# Analisa Program

```
import "fmt"
```

- Ini berarti kita memasukkan kode dari program lain ke dalam kode kita
- Package `fmt` mengimplementasikan formatting untuk input dan output
- Ketika diimpor, package yang diimpor ditulis sebagai _string_

---

# Analisa Program

- Komentar (comment) dalam Go, bisa ditulis baik dengan `//` maupun `/* */`

---

# Analisa Program

```
func main() {
    fmt.Println("Hello, World")
}
```

- Dalam Go, function adalah blok pembangun sebuah program
- Ia memiliki input, output dan pernyataan yang dijalankan berurutan dari atas ke bawah
- Sebuah fungsi dalam Go ditulis dengan keyword `func` dan ia bisa memiliki parameter atau argumen, return value serta body (tubuh) function

---

# Analisa Program

```go
fmt.Println("Hello, World")
```

- Kita mengakses sebuah fungsi di dalam package `fmt` bernama `Println` yang berarti print (tampilkan sebuah teks) dalam 1 line (baris)
- Kita membuat string baru yang mengandung frase `Hello, World`
- Kita memanggil fungsi tersebut dengan string tadi sebagai satu-satunya argumen

---

# Membuat Binary

Aplikasi yang kita buat dalam bahasa Go, sangat dengan mudah dijadikan file binary.

```
$ go build main.go
```

---

# Build dan Install

```
$ go install github.com/<username>/<project>
```

---

# Melihat Dokumentasi Program

- Untuk melihat dokumentasi fungsi `Println` kita bisa menjalankan

  ```
  go doc fmt Println
  ```

- Coba periksa, apa yang dilakukan fungsi tersebut?

