package main

import (
	"context"
	"database/sql"
	"fmt"
	pb "golangSocket/grpc-server"
	"log"
	"net"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

var ctx = context.Background()

// Cambia estos valores con tus propias credenciales de la base de datos
const (
	DBHost     = "34.122.19.115"
	DBPort     = 3306
	DBUser     = "root"
	DBPassword = "test123$"
	DBName     = "proyecto2"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

const port = ":3001"

type Data struct {
	Carnet   string
	Nombre   string
	Curso    string
	Nota     string
	Semestre string
	Anio     string
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	fmt.Println("Recibí de cliente: ", in.GetCarnet())
	data := Data{
		Carnet:   in.GetCarnet(),
		Nombre:   in.GetNombre(),
		Curso:    in.GetCurso(),
		Nota:     in.GetNota(),
		Semestre: in.GetSemestre(),
		Anio:     in.GetAnio(),
	}
	fmt.Println(data)
	insertMysql(data)
	return &pb.ReplyInfo{Info: "Hola cliente, recibí el comentario"}, nil
}

func insertMysql(data Data) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", DBUser, DBPassword, DBHost, DBPort, DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// insert ALUMNO
	_, err = db.ExecContext(ctx, "INSERT INTO proyecto2.ALUMNO (carnet, nombre) VALUES (?, ?)",
		data.Carnet, data.Nombre)

	if err != nil {
		log.Fatalf("Error al insertar en la segunda tabla: %v", err)
	}

	// INSERT CALIFICACION
	_, err = db.ExecContext(ctx, "INSERT INTO proyecto2.CALIFICACION (nota, anio, carnet, codigo, idSemestre) VALUES (?, ?, ?, ?, ?)",
		data.Nota, data.Anio, data.Carnet, data.Curso, data.Semestre)

	if err != nil {
		log.Fatalf("Error al insertar en la base de datos: %v", err)
	}
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
