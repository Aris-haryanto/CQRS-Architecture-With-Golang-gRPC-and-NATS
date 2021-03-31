## Database Setup
untuk merubah database connection silahkan buka file `database/connection.go` pada fungsi `init()`
```go
conn, err := gorm.Open("mysql", "username-kamu@(localhost)/db-kamu?charset=utf8&parseTime=True&loc=Local")
```

## Run Nats
Buka 1 terminal untuk menjalankan service nats. service ini yang akan dipakai oleh service push notif nantinya
```bash
$ nats-streaming-server
```

## Run service
Buka 5 terminal ketikan masing-masing terminal

##### Untuk run service server
bagian ini untuk menerima request dari services client menggunakan grpc
```bash
$ cd services
$ go run main.go
```

##### Untuk run service client
bagian ini untuk mengirim request ke services server menggunakan grpc
```bash
$ cd client
$ go run main.go
```

##### Untuk run service create deposit
bagian ini untuk kita mengcreate deposit yang masuk yang dilisten dari nats streaming
```bash
$ cd deposit-create
$ go run main.go
```

##### Untuk run service approve deposit
bagian ini untuk kita mengapprove deposit yang sudah ada yang dilisten dari nats streaming
```bash
$ cd deposit-approve
$ go run main.go
```

##### Untuk run push notification
bagian ini untuk melihat proses push notifikasi dari services server ketika ada request dari client
```bash
$ cd pushnotif
$ go run main.go
```
## Untuk mengcompile protobuff
```bash
$ protoc -I proto/ proto/*.proto --go_out=plugins=grpc:proto
```