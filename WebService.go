package main

import (
	"encoding/json"
    "fmt"
    "log"
    "net/http"
    "database/sql"
    _ "mysql-master"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to Web Service Penduduk!\nWeb Service yang tersedia (alamat url) : \n\nSemua Penduduk : /Penduduk\nSearch Penduduk berdasarkan NIK : /Penduduk/NIK/NIK=?\nSearch Penduduk berdasarkan nama lengkap : /Penduduk/nama/nama=?\nPenduduk Perempuan : Penduduk/Perempuan\nPenduduk Laki-laki : Penduduk/Laki-laki\nPenduduk Kewarganegaraan Indonesia : Penduduk/Indonesia\nPenduduk Beragama Islam : Penduduk/Islam\nPenduduk Beragama Kristen : Penduduk/Kristen\nPenduduk Beragama Katholik : Penduduk/Katholik\nPenduduk Beragama Hindu : Penduduk/Hindu\nPenduduk Beragama Buddha : Penduduk/Buddha\n\nSemua Working Place : /WorkingPlace\nSearch Working Place berdasarkan nik : /WorkingPlace/NIK/NIK=?\nSearch Working Place berdasarkan nama lengkap : /WorkingPlace/nama/nama=?")
    fmt.Println("Endpoint Hit: homePage")
}

type Perusahaan struct {
	NIK   string    `json:"nik"`
	Nama_Perusahaan string `json:"nama_perusahaan"`
	Alamat_Perusahaan string `json:"alamat_perusahaan"`
}

type Penduduk struct {
	NIK   string    `json:"nik"`
	Nama string `json:"nama"`
	Tanggal_Lahir string `json:"tanggal_lahir"`
	Jenis_Kelamin string `json:"jenis_kelamin"`
	Alamat string `json:"alamat"`
	Agama string `json:"agama"`
	Status_Perkawinan string `json:"status_perkawinan"`
	Pekerjaan string `json:"pekerjaan"`
	Kewarganegaraan string `json:"kewarganegaraan"`
}


func main() {
	handleRequests()
}


func AllPenduduk(w http.ResponseWriter, r *http.Request){
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/penduduk_bandung")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	

	results, err := db.Query("SELECT nik,nama,tanggal_lahir,jenis_kelamin,alamat,agama,status_perkawinan,pekerjaan,kewarganegaraan FROM penduduk")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
   for results.Next() {
		var penduduk Penduduk
		// for each row, scan the result into our tag composite object
		
		err = results.Scan(&penduduk.NIK,&penduduk.Nama,&penduduk.Tanggal_Lahir,&penduduk.Jenis_Kelamin,&penduduk.Alamat,&penduduk.Agama,&penduduk.Status_Perkawinan,&penduduk.Pekerjaan,&penduduk.Kewarganegaraan)
                // and then print out the tag's Name attribute
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		 fmt.Println("Endpoint Hit: returnAllPenduduk")
    
   		 json.NewEncoder(w).Encode(penduduk)

	}
   
}
func Indonesia(w http.ResponseWriter, r *http.Request){
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/penduduk_bandung")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	

	results, err := db.Query("SELECT nik,nama,tanggal_lahir,jenis_kelamin,alamat,agama,status_perkawinan,pekerjaan,kewarganegaraan FROM penduduk where kewarganegaraan like '%Indonesia%'")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
   for results.Next() {
		var penduduk Penduduk
		// for each row, scan the result into our tag composite object
		
		err = results.Scan(&penduduk.NIK,&penduduk.Nama,&penduduk.Tanggal_Lahir,&penduduk.Jenis_Kelamin,&penduduk.Alamat,&penduduk.Agama,&penduduk.Status_Perkawinan,&penduduk.Pekerjaan,&penduduk.Kewarganegaraan)
                // and then print out the tag's Name attribute
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		 fmt.Println("Endpoint Hit: returnPendudukIndonesia")
    
   		 json.NewEncoder(w).Encode(penduduk)

	}
   
}
func Perempuan(w http.ResponseWriter, r *http.Request){
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/penduduk_bandung")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	

	results, err := db.Query("SELECT nik,nama,tanggal_lahir,jenis_kelamin,alamat,agama,status_perkawinan,pekerjaan,kewarganegaraan FROM penduduk where jenis_kelamin like '%P%'")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
   for results.Next() {
		var penduduk Penduduk
		// for each row, scan the result into our tag composite object
		
		err = results.Scan(&penduduk.NIK,&penduduk.Nama,&penduduk.Tanggal_Lahir,&penduduk.Jenis_Kelamin,&penduduk.Alamat,&penduduk.Agama,&penduduk.Status_Perkawinan,&penduduk.Pekerjaan,&penduduk.Kewarganegaraan)
                // and then print out the tag's Name attribute
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		 fmt.Println("Endpoint Hit: returnAllPendudukPerempuan")
    
   		 json.NewEncoder(w).Encode(penduduk)

	}
   
}
func LakiLaki(w http.ResponseWriter, r *http.Request){
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/penduduk_bandung")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	

	results, err := db.Query("SELECT nik,nama,tanggal_lahir,jenis_kelamin,alamat,agama,status_perkawinan,pekerjaan,kewarganegaraan FROM penduduk where jenis_kelamin like '%L%'")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
   for results.Next() {
		var penduduk Penduduk
		// for each row, scan the result into our tag composite object
		
		err = results.Scan(&penduduk.NIK,&penduduk.Nama,&penduduk.Tanggal_Lahir,&penduduk.Jenis_Kelamin,&penduduk.Alamat,&penduduk.Agama,&penduduk.Status_Perkawinan,&penduduk.Pekerjaan,&penduduk.Kewarganegaraan)
                // and then print out the tag's Name attribute
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		 fmt.Println("Endpoint Hit: returnAllPendudukLakiLaki")
    
   		 json.NewEncoder(w).Encode(penduduk)

	}
   
}
func Islam(w http.ResponseWriter, r *http.Request){
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/penduduk_bandung")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	

	results, err := db.Query("SELECT nik,nama,tanggal_lahir,jenis_kelamin,alamat,agama,status_perkawinan,pekerjaan,kewarganegaraan FROM penduduk where agama like '%Islam%'")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
   for results.Next() {
		var penduduk Penduduk
		// for each row, scan the result into our tag composite object
		
		err = results.Scan(&penduduk.NIK,&penduduk.Nama,&penduduk.Tanggal_Lahir,&penduduk.Jenis_Kelamin,&penduduk.Alamat,&penduduk.Agama,&penduduk.Status_Perkawinan,&penduduk.Pekerjaan,&penduduk.Kewarganegaraan)
                // and then print out the tag's Name attribute
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		 fmt.Println("Endpoint Hit: returnAllPendudukIslam")
    
   		 json.NewEncoder(w).Encode(penduduk)

	}
   
}
func Hindu(w http.ResponseWriter, r *http.Request){
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/penduduk_bandung")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	

	results, err := db.Query("SELECT nik,nama,tanggal_lahir,jenis_kelamin,alamat,agama,status_perkawinan,pekerjaan,kewarganegaraan FROM penduduk where agama like 'Hindu%'")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
   for results.Next() {
		var penduduk Penduduk
		// for each row, scan the result into our tag composite object
		
		err = results.Scan(&penduduk.NIK,&penduduk.Nama,&penduduk.Tanggal_Lahir,&penduduk.Jenis_Kelamin,&penduduk.Alamat,&penduduk.Agama,&penduduk.Status_Perkawinan,&penduduk.Pekerjaan,&penduduk.Kewarganegaraan)
                // and then print out the tag's Name attribute
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		 fmt.Println("Endpoint Hit: returnAllPendudukHindu")
    
   		 json.NewEncoder(w).Encode(penduduk)

	}
   
}
func Buddha(w http.ResponseWriter, r *http.Request){
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/penduduk_bandung")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	

	results, err := db.Query("SELECT nik,nama,tanggal_lahir,jenis_kelamin,alamat,agama,status_perkawinan,pekerjaan,kewarganegaraan FROM penduduk where agama like '%Buddha%'")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
   for results.Next() {
		var penduduk Penduduk
		// for each row, scan the result into our tag composite object
		
		err = results.Scan(&penduduk.NIK,&penduduk.Nama,&penduduk.Tanggal_Lahir,&penduduk.Jenis_Kelamin,&penduduk.Alamat,&penduduk.Agama,&penduduk.Status_Perkawinan,&penduduk.Pekerjaan,&penduduk.Kewarganegaraan)
                // and then print out the tag's Name attribute
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		 fmt.Println("Endpoint Hit: returnAllPendudukBuddha")
    
   		 json.NewEncoder(w).Encode(penduduk)

	}
   
}
func Kristen(w http.ResponseWriter, r *http.Request){
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/penduduk_bandung")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	

	results, err := db.Query("SELECT nik,nama,tanggal_lahir,jenis_kelamin,alamat,agama,status_perkawinan,pekerjaan,kewarganegaraan FROM penduduk where agama like '%Kristen%'")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
   for results.Next() {
		var penduduk Penduduk
		// for each row, scan the result into our tag composite object
		
		err = results.Scan(&penduduk.NIK,&penduduk.Nama,&penduduk.Tanggal_Lahir,&penduduk.Jenis_Kelamin,&penduduk.Alamat,&penduduk.Agama,&penduduk.Status_Perkawinan,&penduduk.Pekerjaan,&penduduk.Kewarganegaraan)
                // and then print out the tag's Name attribute
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		 fmt.Println("Endpoint Hit: returnAllPendudukKristen")
    
   		 json.NewEncoder(w).Encode(penduduk)

	}
   
}
func Katholik(w http.ResponseWriter, r *http.Request){
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/penduduk_bandung")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	

	results, err := db.Query("SELECT nik,nama,tanggal_lahir,jenis_kelamin,alamat,agama,status_perkawinan,pekerjaan,kewarganegaraan FROM penduduk where agama like '%Katholik%'")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
   for results.Next() {
		var penduduk Penduduk
		// for each row, scan the result into our tag composite object
		
		err = results.Scan(&penduduk.NIK,&penduduk.Nama,&penduduk.Tanggal_Lahir,&penduduk.Jenis_Kelamin,&penduduk.Alamat,&penduduk.Agama,&penduduk.Status_Perkawinan,&penduduk.Pekerjaan,&penduduk.Kewarganegaraan)
                // and then print out the tag's Name attribute
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		 fmt.Println("Endpoint Hit: returnAllPendudukKatholik")
    
   		 json.NewEncoder(w).Encode(penduduk)

	}
   
}
func WorkingPlace(w http.ResponseWriter, r *http.Request){
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/penduduk_bandung")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	

	results, err := db.Query("SELECT perusahaan.NIK,nama_perusahaan,alamat_perusahaan FROM perusahaan")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
   for results.Next() {
		var perusahaan Perusahaan
		// for each row, scan the result into our tag composite object
		
		err = results.Scan(&perusahaan.NIK,&perusahaan.Nama_Perusahaan,&perusahaan.Alamat_Perusahaan)
                // and then print out the tag's Name attribute
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		 fmt.Println("Endpoint Hit: returnAllWorkingPlace")
    
   		 json.NewEncoder(w).Encode(perusahaan)

	}
   
}
func GetAllWorkingPlace(w http.ResponseWriter, r *http.Request) {
        db, err := sql.Open("mysql",
                "root:@tcp(127.0.0.1:3306)/penduduk_bandung")

        if err != nil {
                log.Fatal(err)
        }
        defer db.Close()

        perusahaan := Perusahaan{}

        rows, err:=db.Query("SELECT NIK,nama_perusahaan,alamat_perusahaan FROM perusahaan")
        if err != nil {
                log.Fatal(err)
        }

        defer rows.Close()
        for rows.Next() {
                err:= rows.Scan(&perusahaan.NIK,&perusahaan.Nama_Perusahaan,&perusahaan.Alamat_Perusahaan)
                if err != nil{
                        log.Fatal(err)
                }
                json.NewEncoder(w).Encode(&perusahaan)
        }
        err=rows.Err()
}

//GetUser ..
func GetWorkingPlace(w http.ResponseWriter, r *http.Request, id string) {
        myid :=(id)
        db, err := sql.Open("mysql",
                "root:@tcp(127.0.0.1:3306)/penduduk_bandung")

        if err != nil {
                log.Fatal(err)
        }
        defer db.Close()

        perusahaan := Perusahaan{}

        rows, err:=db.Query("SELECT perusahaan.NIK,nama_perusahaan,alamat_perusahaan FROM perusahaan INNER JOIN penduduk ON penduduk.nik=perusahaan.NIK where nama like ?", myid)
        if err != nil {
                log.Fatal(err)
        }

        defer rows.Close()
        for rows.Next() {
                err:= rows.Scan(&perusahaan.NIK,&perusahaan.Nama_Perusahaan,&perusahaan.Alamat_Perusahaan)
                if err != nil{
                        log.Fatal(err)
                }
                json.NewEncoder(w).Encode(&perusahaan)
        }
        err=rows.Err()
        if err != nil {
                log.Fatal(err)
        }
}

//GetPostData
func GetPostDataWorkingPlace(w http.ResponseWriter, r *http.Request) {
        var t Perusahaan
        decoder:=json.NewDecoder(r.Body)
        err:=decoder.Decode(&t)
        if err!=nil {
                panic(err)
        }
        defer r.Body.Close()
        fmt.Printf("%s %s %s %s %s",t.NIK, t.Nama_Perusahaan, t.Alamat_Perusahaan)
}

func GetAllWorkingPlaceNIK(w http.ResponseWriter, r *http.Request) {
        db, err := sql.Open("mysql",
                "root:@tcp(127.0.0.1:3306)/penduduk_bandung")

        if err != nil {
                log.Fatal(err)
        }
        defer db.Close()

        perusahaan := Perusahaan{}

        rows, err:=db.Query("SELECT NIK,nama_perusahaan,alamat_perusahaan FROM perusahaan")
        if err != nil {
                log.Fatal(err)
        }

        defer rows.Close()
        for rows.Next() {
                err:= rows.Scan(&perusahaan.NIK,&perusahaan.Nama_Perusahaan,&perusahaan.Alamat_Perusahaan)
                if err != nil{
                        log.Fatal(err)
                }
                json.NewEncoder(w).Encode(&perusahaan)
        }
        err=rows.Err()
}

//GetUser ..
func GetWorkingPlaceNIK(w http.ResponseWriter, r *http.Request, id string) {
        myid :=(id)
        db, err := sql.Open("mysql",
                "root:@tcp(127.0.0.1:3306)/penduduk_bandung")

        if err != nil {
                log.Fatal(err)
        }
        defer db.Close()

        perusahaan := Perusahaan{}

        rows, err:=db.Query("SELECT NIK,nama_perusahaan,alamat_perusahaan FROM perusahaan where NIK like ?", myid)
        if err != nil {
                log.Fatal(err)
        }

        defer rows.Close()
        for rows.Next() {
                err:= rows.Scan(&perusahaan.NIK,&perusahaan.Nama_Perusahaan,&perusahaan.Alamat_Perusahaan)
                if err != nil{
                        log.Fatal(err)
                }
                json.NewEncoder(w).Encode(&perusahaan)
        }
        err=rows.Err()
        if err != nil {
                log.Fatal(err)
        }
}

//GetPostData
func GetPostDataWorkingPlaceNIK(w http.ResponseWriter, r *http.Request) {
        var t Perusahaan
        decoder:=json.NewDecoder(r.Body)
        err:=decoder.Decode(&t)
        if err!=nil {
                panic(err)
        }
        defer r.Body.Close()
        fmt.Printf("%s %s %s %s %s",t.NIK, t.Nama_Perusahaan, t.Alamat_Perusahaan)
}


func GetAllPendudukNIK(w http.ResponseWriter, r *http.Request) {
        db, err := sql.Open("mysql",
                "root:@tcp(127.0.0.1:3306)/penduduk_bandung")

        if err != nil {
                log.Fatal(err)
        }
        defer db.Close()

        penduduk := Penduduk{}

        rows, err:=db.Query("SELECT nik,nama,tanggal_lahir,jenis_kelamin,alamat,agama,status_perkawinan,pekerjaan,kewarganegaraan FROM penduduk")
        if err != nil {
                log.Fatal(err)
        }

        defer rows.Close()
        for rows.Next() {
                err:= rows.Scan(&penduduk.NIK,&penduduk.Nama,&penduduk.Tanggal_Lahir,&penduduk.Jenis_Kelamin,&penduduk.Alamat,&penduduk.Agama,&penduduk.Status_Perkawinan,&penduduk.Pekerjaan,&penduduk.Kewarganegaraan)
                if err != nil{
                        log.Fatal(err)
                }
                json.NewEncoder(w).Encode(&penduduk)
        }
        err=rows.Err()
}

//GetUser ..
func GetPendudukNIK(w http.ResponseWriter, r *http.Request, id string) {
        myid :=(id)
        db, err := sql.Open("mysql",
                "root:@tcp(127.0.0.1:3306)/penduduk_bandung")

        if err != nil {
                log.Fatal(err)
        }
        defer db.Close()

        penduduk := Penduduk{}

        rows, err:=db.Query("SELECT nik,nama,tanggal_lahir,jenis_kelamin,alamat,agama,status_perkawinan,pekerjaan,kewarganegaraan FROM penduduk where nik like ?", myid)
        if err != nil {
                log.Fatal(err)
        }

        defer rows.Close()
        for rows.Next() {
                err:= rows.Scan(&penduduk.NIK,&penduduk.Nama,&penduduk.Tanggal_Lahir,&penduduk.Jenis_Kelamin,&penduduk.Alamat,&penduduk.Agama,&penduduk.Status_Perkawinan,&penduduk.Pekerjaan,&penduduk.Kewarganegaraan)
                if err != nil{
                        log.Fatal(err)
                }
                json.NewEncoder(w).Encode(&penduduk)
        }
        err=rows.Err()
        if err != nil {
                log.Fatal(err)
        }
}

//GetPostData
func GetPostDataPendudukNIK(w http.ResponseWriter, r *http.Request) {
        var penduduk Penduduk
        decoder:=json.NewDecoder(r.Body)
        err:=decoder.Decode(&penduduk)
        if err!=nil {
                panic(err)
        }
        defer r.Body.Close()
        fmt.Printf("%s %s %s %s %s",penduduk.NIK,penduduk.Nama,penduduk.Tanggal_Lahir,penduduk.Jenis_Kelamin,penduduk.Alamat,penduduk.Agama,penduduk.Status_Perkawinan,penduduk.Pekerjaan,penduduk.Kewarganegaraan)
}

func GetAllPenduduk(w http.ResponseWriter, r *http.Request) {
        db, err := sql.Open("mysql",
                "root:@tcp(127.0.0.1:3306)/penduduk_bandung")

        if err != nil {
                log.Fatal(err)
        }
        defer db.Close()

        penduduk := Penduduk{}

        rows, err:=db.Query("SELECT nik,nama,tanggal_lahir,jenis_kelamin,alamat,agama,status_perkawinan,pekerjaan,kewarganegaraan FROM penduduk")
        if err != nil {
                log.Fatal(err)
        }

        defer rows.Close()
        for rows.Next() {
                err:= rows.Scan(&penduduk.NIK,&penduduk.Nama,&penduduk.Tanggal_Lahir,&penduduk.Jenis_Kelamin,&penduduk.Alamat,&penduduk.Agama,&penduduk.Status_Perkawinan,&penduduk.Pekerjaan,&penduduk.Kewarganegaraan)
                if err != nil{
                        log.Fatal(err)
                }
                json.NewEncoder(w).Encode(&penduduk)
        }
        err=rows.Err()
}

//GetUser ..
func GetPenduduk(w http.ResponseWriter, r *http.Request, id string) {
        myid :=(id)
        db, err := sql.Open("mysql",
                "root:@tcp(127.0.0.1:3306)/penduduk_bandung")

        if err != nil {
                log.Fatal(err)
        }
        defer db.Close()

        penduduk := Penduduk{}

        rows, err:=db.Query("SELECT nik,nama,tanggal_lahir,jenis_kelamin,alamat,agama,status_perkawinan,pekerjaan,kewarganegaraan FROM penduduk where nama like ?", myid)
        if err != nil {
                log.Fatal(err)
        }

        defer rows.Close()
        for rows.Next() {
                err:= rows.Scan(&penduduk.NIK,&penduduk.Nama,&penduduk.Tanggal_Lahir,&penduduk.Jenis_Kelamin,&penduduk.Alamat,&penduduk.Agama,&penduduk.Status_Perkawinan,&penduduk.Pekerjaan,&penduduk.Kewarganegaraan)
                if err != nil{
                        log.Fatal(err)
                }
                json.NewEncoder(w).Encode(&penduduk)
        }
        err=rows.Err()
        if err != nil {
                log.Fatal(err)
        }
}

//GetPostData
func GetPostDataPenduduk(w http.ResponseWriter, r *http.Request) {
        var penduduk Penduduk
        decoder:=json.NewDecoder(r.Body)
        err:=decoder.Decode(&penduduk)
        if err!=nil {
                panic(err)
        }
        defer r.Body.Close()
        fmt.Printf("%s %s %s %s %s",penduduk.NIK,penduduk.Nama,penduduk.Tanggal_Lahir,penduduk.Jenis_Kelamin,penduduk.Alamat,penduduk.Agama,penduduk.Status_Perkawinan,penduduk.Pekerjaan,penduduk.Kewarganegaraan)
}


func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/Penduduk", AllPenduduk)
    http.HandleFunc("/WorkingPlace/", WorkingPlace)
    http.HandleFunc("/Penduduk/Perempuan/", Perempuan)
    http.HandleFunc("/Penduduk/Laki-laki/", LakiLaki)
    http.HandleFunc("/Penduduk/Islam/", Islam)
    http.HandleFunc("/Penduduk/Hindu/", Hindu)
    http.HandleFunc("/Penduduk/Buddha/", Buddha)
    http.HandleFunc("/Penduduk/Kristen/", Kristen)
    http.HandleFunc("/Penduduk/Katholik/", Katholik)
    http.HandleFunc("/Penduduk/Indonesia/", Indonesia)
    http.HandleFunc("/Penduduk/NIK/", func(w http.ResponseWriter, r *http.Request) {

                switch r.Method {
                case "GET":
                        s:=r.URL.Path[len("/Penduduk/NIK/"):]
                        if s!=""{
                                GetPendudukNIK(w, r, s)
                        } else {
                                GetAllPendudukNIK(w, r)
                        }

                case "POST":
                        GetPostDataPendudukNIK(w, r)
                case "DELETE":

                default:
                        http.Error(w,"invalid",405)
                }
        })
    http.HandleFunc("/Penduduk/nama/", func(w http.ResponseWriter, r *http.Request) {

                switch r.Method {
                case "GET":
                        s:=r.URL.Path[len("/Penduduk/nama/"):]
                        if s!=""{
                                GetPenduduk(w, r, s)
                        } else {
                                GetAllPenduduk(w, r)
                        }

                case "POST":
                        GetPostDataPenduduk(w, r)
                case "DELETE":

                default:
                        http.Error(w,"invalid",405)
                }
        })
  	http.HandleFunc("/WorkingPlace/nama/", func(w http.ResponseWriter, r *http.Request) {

                switch r.Method {
                case "GET":
                        s:=r.URL.Path[len("/WorkingPlace/nama/"):]
                        if s!=""{
                                GetWorkingPlace(w, r, s)
                        } else {
                                GetAllWorkingPlace(w, r)
                        }

                case "POST":
                        GetPostDataWorkingPlace(w, r)
                case "DELETE":

                default:
                        http.Error(w,"invalid",405)
                }
        })

  	  	http.HandleFunc("/WorkingPlace/NIK/", func(w http.ResponseWriter, r *http.Request) {

                switch r.Method {
                case "GET":
                        s:=r.URL.Path[len("/WorkingPlace/NIK/"):]
                        if s!=""{
                                GetWorkingPlaceNIK(w, r, s)
                        } else {
                                GetAllWorkingPlaceNIK(w, r)
                        }

                case "POST":
                        GetPostDataWorkingPlaceNIK(w, r)
                case "DELETE":

                default:
                        http.Error(w,"invalid",405)
                }
        })
  	    log.Fatal(http.ListenAndServe(":8081", nil))
}