package mahasiswacontroller

import (
  "net/http"
  "html/template"
  "go-web-crud/models/mahasiswamodel"
  "go-web-crud/entities"
  "log"
  "strconv"
)

//handle function untuk "/mahasiswa"
func Index(w http.ResponseWriter, r *http.Request) {
  //membuat template html
  temp, err := template.ParseFiles("views/mahasiswa/index.html")
  if err != nil {
    log.Fatal(err)
  }
  
  //mengambil data dari model yang akan ditampilkan di views
  data := mahasiswamodel.GetAllData()
  
  //mengeksekusi template dan mengirimkan data
  err = temp.Execute(w, data)
  if err != nil {
    log.Fatal(err)
  }
}

//handle function untuk "/mahasiswa/add" (menambah data)
func Add(w http.ResponseWriter, r *http.Request) {
  //jika method request get maka akan menampilkan form untuk menambahkan data
  if r.Method == "GET" {
    temp, err := template.ParseFiles("views/mahasiswa/add.html")
    if err != nil {
      log.Fatal(err)
    }
    
    err = temp.Execute(w, nil) 
    if err != nil {
      log.Fatal(err)
    }
  }
  
  //method request post untuk menerima data dari form melalui method post yang akan masuk ke database
  if r.Method == "POST" {
    //memasukan data dari form ke struct mhs; mhs data yang akan dimasukan kedalam function di model
    //FormValue untuk mengambil data dari method post
    var mhs entities.Mahasiswa
    mhs.Nama = r.FormValue("nama")
    mhs.Nim = r.FormValue("nim")
    mhs.Jurusan = r.FormValue("jurusan")
    
    //mengecek jika return dari create data false berarti ada yang error (data bisa jadi tidak masuk)
    if ok := mahasiswamodel.CreateData(mhs); !ok {
      http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
    }
    
    //setelah selesai akan langsung diredirect ke "/mahasiswa"
    http.Redirect(w, r, "/mahasiswa", http.StatusFound)
  }
}

//handle function untuk "/mahasiswa/add" (mengedit data)
func Edit(w http.ResponseWriter, r *http.Request) {
  //jika method request get maka akan menampilkan data yang akan diedit ke form 
  if r.Method == "GET" {
    //membuat template html, file yang akan ditampilkan
    temp, err := template.ParseFiles("views/mahasiswa/edit.html")
    if err != nil {
      log.Fatal(err)
    }
    
    //mengambil id dari method get dan mengconvert ke int
    //harus diconvert terlebih dahulu karena data yang dikirimkan melalui get berupa string 
    idStr := r.URL.Query().Get("id")
    id, _ := strconv.Atoi(idStr)
    data := mahasiswamodel.GetDataById(id)
    
    //ekseskusi template dan mengirim data
    err = temp.Execute(w, data)
    if err != nil {
      log.Fatal(err)
    }
  }
  
  //method request post untuk menerima data dari form melalui method post yang akan masuk ke database
  if r.Method == "POST" {
    //mengconvert id ke int
    idStr := r.FormValue("id")
    id, _ := strconv.Atoi(idStr)
    
    //memasukan data ke struct dari form yang dikirimkan melalui method post
    var mhs entities.Mahasiswa
    mhs.Nama = r.FormValue("nama")
    mhs.Nim = r.FormValue("nim")
    mhs.Jurusan = r.FormValue("jurusan")
    mhs.Id = uint(id)
    
    //mengecek jika return dari create data false berarti ada yang error (data bisa jadi tidak mberubah) 
    if ok := mahasiswamodel.UpdateData(mhs); !ok {
      http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
    }
    
    //setelah selesai akan langsung diredirect ke "/mahasiswa" 
    http.Redirect(w, r, "/mahasiswa", http.StatusFound)
  }
}

//handle function untuk "/mahasiswa/delete" (menghapus data)
func Delete(w http.ResponseWriter, r *http.Request) {
  
  //mengambil id dari methid get dan mengconvert ke int
  //harus diconvert terlebih dahulu karena data yang dikirimkan melalui get berupa string 
  idStr := r.URL.Query().Get("id")
  id, _ := strconv.Atoi(idStr)
  
  //mengecek jika return dari create data false berarti ada yang error (data bisa jadi tidak terhapus) 
  if ok := mahasiswamodel.DeleteData(id); !ok {
    http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
  }
  
  //setelah selesai akan langsung diredirect ke "/mahasiswa" 
  http.Redirect(w, r, "/mahasiswa", http.StatusFound)
}