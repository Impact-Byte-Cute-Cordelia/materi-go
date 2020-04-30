# Struct

---

# Agenda

- Pengenalan Struct
- Struct Method
- Public & Private Property

---

# Struct

- Coba ingat kembali, apa saja tipe data bawaan Go?
- Mungkinkah kita membuat program hanya dengan tipe data tersebut?
- Mungkinkah kita membuat tipe data di luar tipe data yang ada?

---

# Pengertian Struct

- Merupakan _user defined type_ yang mengandung kumpulan
  field/property
- Property tipe data bisa bervariasi
- Proses pendeklarasian type data pada properti wajib di
  dilakukan di awal

---

# Real Implementasi

Jika kalian datang dari object oriented programming (OOP)
mungkin kalian akan mengganggap `struct` merupakan `class`

---

# Manfaat

Memudahkan untuk pembungkusan sebuah informasi menjadi
satu kesatuan unit, sehingga program dapat mudah di
pelihara

---

```go
func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 – x1
    b := y2 – y1
    return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
    l := distance(x1, y1, x1, y2)
    w := distance(x1, y1, x2, y1)
    return l * w
}

func main() {
    var rx1, ry1 float64 = 0, 0
    var rx2, ry2 float64 = 10, 10

    fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
}
```

---

# Struct

- Dengan melihat contoh program tadi, mungkin agak sulit bagi kita untuk memahami apa yang dilakukan program tersebut
- Belum lagi kita harus melacak semua koordinat yang ada
- Bagaimana kalau kita juga menambah ingin menghitung luas lingkaran dengan koordinat yang sama?

---

# Struct

- Cara yang bisa membuat program tadi bisa memiliki kode yang lebih mudah dibaca, adalah dengan menggunakan struct
- Struct adalah tipe data yang mengandung field yang memiliki nama
- `type` memperkenalkan tipe baru, diikuti dengan namanya
- `struct` menandakan bahwa kita sedang mendefinisikan tipe struct
- Daftar field yang ada, ditulis di dalam kurung kurawal

```go
type Circle struct {
    x float64
    y float64
    r float64
}
```

---

# Struct

- Field bisa dibayangkan sebagai sekumpulan variabel
- Setiap field memiliki nama dan tipe, dan disimpan bersebelahan dengan field lain di dalam struct
- Penulisan struct seperti ini juga valid

```
type Circle struct {
  x, y, r float 64
}
```

---

# Inisialisasi Struct

- Struct `Circle` kita kemudian bisa diinisialisasi dengan beberapa cara
- Dengan inisialisasi, kita mengalokasikan memory untuk field-field di dalam struct

```go
var c Circle

// atau
c: = new(Circle)
```

---

# Inisialisasi Struct

- Kita bisa juga menginisialisasi Struct dan memberikan nilai awal untuknya
- Contoh ketiga ini, dilakukan jika kita membutuhkan pointer ke struct

```go
c := Circle{x: 0, y: 0, r: 5}
c := Circle{0, 0, 5}
c := &Circle{0, 0, 5}
```

---

# Field dalam Struct

- Kita bisa mengakses sebuah field dalam struct dengan menggunakan operator `.`

```go
fmt.Println(c.x, c.y, c.r)
c.x = 10
c.y = 5
```

---

# Field dalam Struct

- Setelah diakses, maka kita bisa menggunakan field ini di dalam function

```go
func circleArea(c Circle) float64 {
    return math.Pi * c.r*c.r
}

// main()
c := Circle{0, 0, 5}
fmt.Println(circleArea(c))
```

---

# Menggunakan Struct

- Perlu diingat, karena argumen selalu dikopi di Go, maka jika kita berusaha untuk mengubah salah satu field struct di dalam sebauh fungsi, ia tidak akan mengubah variabel aslinya
- Untuk itu, maka biasanya kita menggunakan pointer, saat kita ingin agar sebuah fungsi bisa mengubah argumen struct yang diberikan kepadanya
- Ini juga meningkatkan efisiensi saat struct besar diberikan atau dikembalikan oleh sebuah fungsi

```go
func setRuas(c *Circle) error {
    c.r = 7
    return nil
}

// main()
c := Circle{0, 0, 5}
fmt.Println(setRuas(&c))
```

---

# Method

- Method adalah sejenis function yang diasosiasikan dengan sebuah struct
- Go memiliki sintaks penulisan method yang cukup singkat dan membantu pembaca kode untuk memahami potongan kode tersebut
- Contoh ini menambahkan method bernama area() untuk Circle dan ini adalah refactor contoh di slide sebelumnya

```go
func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
}
```

---

# Method

- Di antara `func` dan nama function, kita menambahkan _receiver_
- _receiver_ mirip dengan parameter, punya nama dan tipe data
- Dengan menuliskan seperti ini, kita bisa memanggil function dengan operator `.`

```go
fmt.Println(c.area())
```

---

# Method

- Dengan cara ini, kita tidak lagi membutuhkan operator `&`
- Go langsung paham bahwa kita akan memberikan pointer ke circle untuk method ini
- Karena fungsi `area()` ini hanya milik circle saja, kita bisa mengganti nama fungsi menjadi `area` dan bukan `circleArea`

---

```go
type Rectangle struct {
    x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}

// main()
r := Rectangle{0, 0, 10, 10}
fmt.Println(r.area())

```

---

# Embedded Type

- Sejauh ini, kita mendefinisikan hubungan apa-memiliki-apa di dalam struct (misal, lingkaran mempunyai jari-jari)
- Kita juga bisa mendefinisikan hubungan apa-adalah-apa dengan menggunakan struct

---

# Embedded Type

- Contoh, struct person

```go
type Person struct {
    Name string
}

func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}
```

---

# Embedded Type

- Jika kita ingin mendefinsikan struct baru bernama Android, yang memiliki hubungan Android-adalah-Person, kita bisa menggunakan embedded type di dalam Go

```go
type Android struct {
    Person
    Model string
}
```

---

# Embedded Type

- Dengan cara ini, maka method milik Person, bisa langsung diakses oleh Android, karena Android juga dianggap Person

```go
a := new(Android)
a.Talk()
a.Person.Talk() // ini juga valid
```

---

# Private dan Public Property

- Dalam sebuah struct, field yang private, diawali dengan huruf kecil
- Sementara field public, yang bisa diekspor dan digunakan di luar package, diawali dengan huruf besar

```go
type Circle struct {
    x, y   int
    Radius int
}

```

---

# Struct dan OOP

- Karena Go tidak mengenal konsep class, maka sebuah entitas, bisa diwakili dalam bentuk Struct
- Sebuah struct akan memiliki property (field) tertentu dan mempunyai method, untuk menunjukkan apa yang bisa dilakukan entitas tersebut
- Dengan demikian, maka implementasi struktur data seperti dalam OOP, bisa didekati

---

# Constructor untuk Struct

- Jika diinginkan, kita bisa membuat constructor ala OOP untuk diaplikasikan di struct milik Go
- Ini bisa dilakukan untuk mencegah inisialisasi struct dengan cara seperti `var e Employee` yang berisi default value, sehingga sulit untuk digunakan

---

# Constructor untuk Struct

```go
type employee struct {
    firstName   string
    lastName    string
    totalLeaves int
    leavesTaken int
}

func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
    e := employee {firstName, lastName, totalLeave, leavesTaken}
    return e
}
```

---

# Constructor untuk Struct

- Dengan cara tadi, struct bisa dibuat dengan

```go
package main

import "oop/employee"

func main() {
    e := employee.New("Sam", "Adolf", 30, 20)
}
```
