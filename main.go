package main

import(
  "net/http"
  "log" 
  "go-web-crud/controllers/homecontroller"
  "go-web-crud/controllers/mahasiswacontroller"
)

func main() {
  
  //mendaftarkan serve mux (router)
  mux := http.NewServeMux()
    
  //router 
  //1. Halaman Home
  mux.HandleFunc("/", homecontroller.Welcome)
  //2. Halaman Mahasiswa
  mux.HandleFunc("/mahasiswa", mahasiswacontroller.Index)
  mux.HandleFunc("/mahasiswa/add", mahasiswacontroller.Add)
  mux.HandleFunc("/mahasiswa/edit", mahasiswacontroller.Edit)
  mux.HandleFunc("/mahasiswa/delete", mahasiswacontroller.Delete)
  
  //membuat web server
  server := http.Server{
      Addr: ":8080", 
      Handler: mux, 
  }
    
  //menyalakan Server
  log.Println("Starting web on port 8080")
  err := server.ListenAndServe()
  log.Fatal(err)
}