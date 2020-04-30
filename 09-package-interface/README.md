# Package

- Merupakan cara kita melakukan namespacing project dalam go
- Sebuah namespace biasanya ditulis dengan mengunakan huruf kecil dan nama folder biasanya 
  merpresentasikan nama package yang akan kita buat

--- 

- Mari kita coba ulang dari awal dan ayo kita ulang kembali

```
$ mkdir awesome
$ cd awesome
$ go mod init github.com/<username>/awesome
$ touch main.go
```

---

Klo kita print dari current directory kita maka akan muncul beberapa file berikut

```sh
- ./awesome 
   - main.go
   - go.mod
```

FYI, package main dan func main itu merupakan entry point untuk menjalankan go kita.

---

Mari kita buah sebuah package dengan nama `model` sudah barang tentu kita harus membuat direktori
baru pada project tadi. Nama foldernya musti `model` karena kita menggunakan nama folder ini sebagai 
nama package

```
$ mkdir model
```

---

Tentunya saat ini kita memiliki folder baru dengan nama model bukan? Mari kita print directorynya

```sh
- ./awesome 
   - model/
   - main.go
   - go.mod
```

---

Mari kita buat sebuah file di dalam folder model tersebut. Kita buah `user.go`, sehingga struktur
folder kita akan menjadi seperti ini

```sh
- ./awesome 
   - model/
      - user.go
   - main.go
   - go.mod
```

---

Mari kita buka file tersebut` dan kita coba isi dengan code sebagai berikut:

```go
package model

type User struct {
   ID    int
   Name  string
}

func (u User) Hello() string {
   return "Hello " + user.Name
}
```

---

# Import package ke package yang lain

Karena kita menggunakan gomod, maka untuk melakukan call package yang sudah kita buat di 
program main kita, bisa dilakukan sebagai berikut:

```go
package main

import (
   "fmt"
   "github.com/<username>/awesome/model"
)

func main() {
   user := model.User{ID: 1, Name: "Foobar"}
   fmt.Println(user.Hello())
}

```

---

# Ketentuan Package

- Library yang kita tulis dalam package yang sama, secara otomatis akan di load oleh go
- Perhatikan naming method dan struct haruslah unique


```sh
- ./awesome 
   - model/
      - user.go
      - project.go
   - main.go
   - go.mod
```

---

Berikut adalah sebuah struct project yang berada pada package model

```go
// project.go
package model

type Project struct {
   ID    int
   Name  string
   User  User
}

// Kita sudah melakukan define Create function, apabila ini kita buat juga pada user 
// akan menjadi error
func Create(name, user User) (Project, error) {
   // Implementasi body function
}

// Ini masih bisa digunakan walaupun nama methodnya sama, tapi karena fungsi ini 
// terdapat pada struct project maka masih di anggap OK
func (p Project) Hello() string {
   return "Hello " + p.Name
}
```

---

# Interface

- Merupakan user type definition yang di deklarasikan
- Di dalam body nya hanya terdiri dari definisi fungsi tanpa ada implementasinya
- Fungsi - fungsi yang terdapat pada body interface harus di deklarasikan pada kelas implementasinya
- Seluruh fungsi yang terdapat pada body interface harus bersifat public

---

```go
type Mamalia interface{
   Move() string
}

type Person struct{}

func(p Person) Move() string {
   return "Walking"
}
```

---

- Interface merupakan kontrak dalam konsep OOP
- Ketika sebuah interface yang di masukkan ke dalam sebuah argument ke sebuah pembuatan
  function, maka ketika function tersebut kita panggil terhadap argumen yang kita passing
  haruslah merupakan implementasi dari interface tersebut

---

```go
// Kita tidak secara explisit menggunakan struct Person, 
// alih menggunakan interface yang kita passing
func Work(Mamalia) string {
   return Mamalia.Move()
}

func main() {
   // Karena `Person` memiliki method Move, maka struct yang kita call sudah di anggap 
   // sebagai implementasi interface tersebut
   p := Person{}
   fmt.Println(Work(p))
}

```

--- 

Dalam go tidak perlu melakukan penjelasan secara explisit apakah struct yang kita buat itu
merupakan implementasi dari sebuah interface tertentu, asalkah sebuah struct memiliki method 
dan signature yang sama dengan method interface, maka itu sudah dikatakan implementasi dari
interface yang kita buat 

---

# Interface merupakan tipe data

```go
type Mamalia interface{
   Move() string
}

type Animal struct {
   Mamalia Mamalia
}
```

---

# Interface Kosong

- Merupakan special data untuk menyimpan segala jenis type data
- Sangat cocok untuk digunakan untuk menampung result dari hasil sebuah operasi

---

# Contoh

```go
type QueryResult struct {
   Result interface{}
}

```

Result bisa merupakan array bisa juga merupakan single object

---

# Contoh

```go
result := QueryResult{
   Result: domain.Farm{Name: "Farm Keluarga"},
}

fmt.Println(result.Result)
```

*Notes*

Perlu diingat sebuah informasi yang dimasukkan kedalam interface kosong kan ditampung dalam bentuk string, sehingga bisa di print menggunakan `fmt.Println`

---

# Casting

Untuk informasi yang sudah dimasukkan ke dalam interface agar di dapat di akses secara object harus dilakukan casting terhadap informasi tersebut

```go
result := QueryResult{
   Result: domain.Farm{Name: "Farm Keluarga"},
}

fmt.Println(result.Result)

farm := result.Result.(domain.Farm)
fmt.Println(farm.Name)
```