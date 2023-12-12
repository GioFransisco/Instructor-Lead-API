# Untuk mengambil base image
FROM alpine
# Untuk membuat working direktori
WORKDIR /app
# Untuk menjalankan perintah / command
# RUN go mod tidy
# RUN go build -o be-enigma-laundry
# Untuk mengcopy/menyalin file dari local ke container
COPY .env /app
COPY instructor-led-apps /app
# Untuk mengeksekusi program
ENTRYPOINT [ "/app/instructor-led-apps" ]