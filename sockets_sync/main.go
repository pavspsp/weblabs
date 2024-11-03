package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func handleConnection(conn net.Conn) {
	defer func() {
		log.Printf("Клиент %s отключился", conn.RemoteAddr().String())
		conn.Close()
	}()
	log.Printf("Клиент %s подключился", conn.RemoteAddr().String())
	reader := bufio.NewReader(conn)
	for {
		//Читаем входящее сообщение
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Ошибка чтения от клиента %s: %v", conn.RemoteAddr().String(), err)
			continue
		}
		log.Printf("Получено сообщение от %s: %s", conn.RemoteAddr().String(), message)
		// Обрабатываем полученное сообщение
		time.Sleep(5 * time.Second)
		response := MakeResponce(message[:len(message)-1])
		// Отправляем ответ обратно клиенту
		_, err = conn.Write([]byte(response))
		if err != nil {
			log.Printf("Ошибка отправки ответа %s: %v", conn.RemoteAddr().String(), err)
			continue
		}
		log.Printf("Отправлен ответ клиенту %s: %s", conn.RemoteAddr().String(), response)
		break
	}
	time.Sleep(10 * time.Second)
	return
}

// Обработка сообщения
func MakeResponce(mes string) string {
	resp := []rune(mes)
	for i, j := 0, len(resp)-1; i < j; i, j = i+1, j-1 {
		resp[i], resp[j] = resp[j], resp[i]
	}
	return string(resp) + "_Сервер написан Спиридоновым П.К.207м"
}
func Test1(i string) {
	fmt.Println(i)
}

func main() {
	//лог-файл
	logfilepath := "C:/Users/mi/Documents/GitHub/weblabs/sockets_sync/server.log"
	logFile, err := os.OpenFile(logfilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	defer logFile.Close()
	log.SetOutput(logFile)

	addr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка адреса:", err)
		return
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println("Ошибка создания сервера:", err)
		return
	}
	defer listener.Close()
	log.Println("Запуск асинхронного сервера")
	fmt.Println("Асинхронный сервер запущен")
	for {
		// Принимаем новое соединение
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка приема соединения:", err)
			continue
		}

		//обработка соединения
		go handleConnection(conn)
	}
}
