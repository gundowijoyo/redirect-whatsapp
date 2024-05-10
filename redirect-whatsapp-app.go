package main

import (
    "fmt"
    "os/exec"
)

func main() {
  /*Silahkan atur nomor telepon yang akan dituju*/
    phoneNumber := "62812xxxxxx"

    /* Pesan yang ingin dikirimkan*/
    message := "Halo! Ini adalah pesan otomatis."

    /*Membuat URL WhatsApp dengan nomor telepon dan pesan*/
    whatsappURL := "https://api.whatsapp.com/send?phone=" + phoneNumber + "&text=" + message

    /*Membuka WhatsApp dengan URL yang sudah dibuat*/
    cmd := exec.Command("am", "start", "-a", "android.intent.action.VIEW", "-d", whatsappURL)
    err := cmd.Run()
    if err != nil {
        panic(err)
    }

    // Menampilkan pesan berhasil
  fmt.Println("Yeay, berhasil membuka WhatsApp!")
}
