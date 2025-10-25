# andora-wami-ex-main
 bentar....masih beta, males updatenya :v

------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------


# 📶 RAM (Remote Android Modem)

RAM (Remote Android Modem) adalah antarmuka berbasis LuCI yang memungkinkan pengguna menghubungkan dan mengontrol perangkat Android atau device lainnya sebagai modem ADB tethering, lengkap dengan fitur monitoring jaringan, kontrol daya, dan pembacaan pesan SMS secara langsung dari dashboard Luci.

![Tampilan UI](https://github.com/Rama-X2/andora-main/raw/main/screenshot.png)

---

## 🚀 Fitur Unggulan

- 🔌 Deteksi otomatis perangkat Android via ADB
- 🔁 Kontrol daya: Power off, reboot, reboot recovery & bootloader
- 📩 Akses inbox SMS langsung dari LuCI
- 📶 Monitoring jaringan: APN, sinyal, operator, dan jenis jaringan
- 📡 Pengaturan mode pesawat & data seluler dari web
- 🛠️ Terintegrasi langsung di menu **LuCI > Modem > Andora External**
- 🌙 Dukungan theme Argon LuCI yang responsif dan modern

---

## 🧰 SW/FW yang Digunakan

- Rama Server 23.05.x
- LuCI Web Interface
- Shell scripting + ADB
- JSON parsing
- JavaScript & jQuery (dari LuCI core)

---

## ⚙️ Instalasi

### Syarat:
- Devices berbasis Luci
- USB port yang terkoneksi ke perangkat Android
- Android dengan USB Debugging aktif
- ADB terpasang di sistem (bisa melalui `opkg`)

### Langkah Instalasi:

# Clone project ini
```
git clone https://github.com/Rama-X2/andora-main.git
cd andora-main
```
# Jalankan installer
```
sh install.sh
```
Setelah instalasi selesai, buka antarmuka LuCI dan akses melalui:
```
http://{ganti dengan domain/ip luci kalian}/cgi-bin/luci/admin/modem/andora-ex
```

---

---

## 📋 Changelog (v0.70 - Dev)

* ✅ Tambah menu power (Power Off, Reboot, dll)
* ⚙️ Perbaikan parser data SMS
* ✉️ Batas tampilan jumlah pesan di inbox
* 🔃 Perbaikan bug "loading tak berhenti saat refresh inbox"
* 📌 Pindah posisi tombol update
* 🛠️ Andora-EX kini berjalan otomatis saat startup
* 📂 Menu dipindahkan ke `Modem > Andro External`

---

## 💡 Rencana Pengembangan

* Dukungan multi-device ADB
* Pengaturan APN dari LuCI
* Remote via Zerotier/VPN
* Backup & restore pesan

---

## 🧑‍💻 Kontribusi

Pull request sangat diterima! Kamu bisa bantu dalam bentuk:

* Menambahkan fitur baru
* Memperbaiki bug
* Peningkatan UI
* Dokumentasi



Langkah kontribusi:

# Fork dan buat branch baru
```
git checkout -b fitur-baru
```

# Commit perubahan
```
git commit -m "Menambahkan fitur YTTA"
```
# Push dan buat Pull Request

------------------------------------------------------------------------------------------------------------------------

## 📄 Lisensi

Project ini dilisensikan di bawah MIT License. Silakan gunakan, modifikasi, dan distribusikan.



## 🧠 Tentang

Dikembangkan oleh [Rama-X2](https://github.com/Rama-X2) sebagai bagian dari pengembangan sistem Remote Android Modem untuk kebutuhan lokal jaringan rumah / RTRWNet / Server / eksperimen pribadi.
