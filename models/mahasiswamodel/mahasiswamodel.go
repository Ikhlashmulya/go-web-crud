package mahasiswamodel

import (
  "go-web-crud/entities"
  "go-web-crud/config"
  "log"
)

//function mengambil semua data dari database
func GetAllData() []entities.Mahasiswa {
  //memanggil koneksi database
  db := config.GetConnection()
  //menutup database saat sudah selesai
  defer db.Close()
  
  //melakukan query ke database
  rows, err := db.Query(`SELECT id, nama, nim, jurusan FROM mahasiswa`)
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()
  
  
  //mengiterasi data dari database ke slice of struct result
  
  //membuat struct untuk return func
  var result []entities.Mahasiswa
  
  //perulangan akan terus berjalan sampai data habis
  for rows.Next() {
    //deklarasi struct tempat penyimpanan sementara dari database
    var mhs entities.Mahasiswa
    //rows.Scan mengisi data dari query database ke struct mhs
    rows.Scan(&mhs.Id, &mhs.Nama, &mhs.Nim, &mhs.Jurusan)
    //var mhs akan di-append atau ditambahkan ke struct result
    result = append(result, mhs)
  }
  
  //return hasil iterasi
  return result
}

//function untuk insert data ke database
func CreateData(mhs entities.Mahasiswa) bool {
  //memanggil koneksi database
  db := config.GetConnection() 
  //menutup database saat selesai
  defer db.Close()
  
  //proses eksekusi SQL
  res, err := db.Exec(
    `INSERT INTO mahasiswa(nama, nim, jurusan) VALUES(?,?,?)`,
    mhs.Nama, 
    mhs.Nim,
    mhs.Jurusan,
  )
  if err != nil {
    log.Fatal(err)
  }
  
  //return true jika last insert id lebih dari nol
  //lastinsertid didapat dari hasi eksekusi SQL
  id, _ := res.LastInsertId() 
  return id > 0
}

//function untuk mengambil data berdasarkan id
func GetDataById(id int) entities.Mahasiswa {
  //memanggil koneksi database
  db := config.GetConnection() 
  //menutup database saat sudah selesai
  defer db.Close() 
  
  //deklarasi struct yang akan dijadikan sebagai return value
  var mhs entities.Mahasiswa
  
  //proses query dari database, query ini hanya mengambil satu row data (single record)
  err := db.QueryRow(`SELECT id, nama, nim, jurusan FROM mahasiswa WHERE id = ?`, id).Scan(&mhs.Id, &mhs.Nama, &mhs.Nim, &mhs.Jurusan)
  if err != nil {
    log.Fatal(err)
  }
  
  //return struct mhs
  return mhs
}

//function untuk mengupdate data 
func UpdateData(mhs entities.Mahasiswa) bool {
  //memanggil koneksi database
  db := config.GetConnection() 
  //menutup database saat sudah selesai
  defer db.Close() 
  
  //proses eksekusi SQL
  res, err := db.Exec(
    `UPDATE mahasiswa SET nama = ?, nim = ?, jurusan = ? WHERE id = ?`,
    mhs.Nama,
    mhs.Nim,
    mhs.Jurusan,
    mhs.Id,
  )
  if err != nil {
    log.Fatal(err)
  }
  
  //return true jika rowsaffected lebih dari 0
  //rowsaffected didapat dari hasil eksekusi SQL
  num, _ := res.RowsAffected()
  return num > 0
}

//function untuk delete data
func DeleteData(id int) bool {
  //memanggil koneksi database
  db := config.GetConnection() 
  //menutup database saat sudah selesai
  defer db.Close() 
  
  //proses eksekusi SQL
  res, err := db.Exec(`DELETE FROM mahasiswa WHERE id = ?`, id)
  if err != nil {
    log.Fatal(err)
  }
  
  //return true jika rowsaffected lebih dari 0
  //rowsaffected didapat dari hasil eksekusi SQL
  num, _ := res.RowsAffected()
  return num > 0
}