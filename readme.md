#Aplikasi Booking

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
