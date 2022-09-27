# Aplikasi Booking

Tools yang saya gunakan untuk membuat aplikasi ini adalah

- go version 1.19.1
- Chi [Router](https://github.com/go-chi/chi)
- Alex edwards [SCS session management](https://github.com/alexedwards/scs)
- Use [surf](https://github.com/justinas/nosurf)
- Use validator [validator](https://github.com/asaskevich/govalidator)
<!-- jika sudah selesai buat dengan template ini -->
- Alternatife [Golang Template](https://github.com/CloudyKit/jet)

<!-- Testing -->

go test -v // all past
go test -cover 76%

<!-- karena ada yang tidak dipakai saya sengaja karena sebagai pembelajaran -->

go test -coverprofile=coverage.out && go tool cover -html=coverage.out

pada command line (masuk ke root folder aplikasi disini booking)

test dengan perintah
go test -v ./...
artinya test semua sub folder

<!-- Soda bagian dari framework buffalo -->
<!-- Cek Apakah sudah ada jika belum install -->

=$GOPATH/bin/soda

<!-- Buat migrasi dengan nama migrasi jika sukses maka akan muncul folder baru dengan nama migrations  -->

=$GOPATH/bin/soda generate fizz CreateTableUsers

<!-- buat database pada table up bisa menggunakan sql dan fuzz -->
<!-- cara migratenya dengan menggunakan perintah $GOPATH/bin/soda migrate -->
