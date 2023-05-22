/*Разработайте RPC-сервер, который умеет вычислять расстояние между двумя точками на плоскости.

В качестве аргумента функция должна принимать тип Points, объявленный следующим образом:

// Точка на плоскости.
type Point struct {
    X, Y float64
}

// Аргумент для функции Dist.
type Points struct {
    A, B Point
}

Сервер должен экспортировать функцию Dist, вычисляющую расстояние между двумя точками и записывающую результат вычислений во второй аргумент.
*/
package main

import (
	//"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	//"time"
	"math"
)

const addr = ":12345"
const network = "tcp4"

// Тип данных для RPC сервера.
// Может быть любым пользовательским типом.
// Все экспортируемые методы этого типа
// будут доступны для удаленного вызова.

// Точка на плоскости.
type Point struct {
    X, Y float64
}

// Аргумент для функции Dist.
type Points struct {
    A, B Point
}


 type Server int

// Метод Time доступен для удаленного вызова.
func (s *Server) Dist(p  Points , resp *float64) error {
	
sq :=(p.A.X -p.B.X)*(p.A.X -p.B.X) + (p.A.Y-p.B.Y)*(p.A.Y-p.B.Y)
	*resp = math.Sqrt(sq)
	return nil
}
 
func main() {
	// Создаем указатель на переменную типа Server.
	srv := new(Server)
	// Регистрируем методы типа Server в службе RPC.
	rpc.Register(srv)
	// Регистрируем HTTP-обработчик для службы RPC.
	rpc.HandleHTTP()
	// Создаём сетевую службу.
	listener, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	// Запускаем HTTP-сервер поверх созданной сетевой службы.
	http.Serve(listener, nil)
}