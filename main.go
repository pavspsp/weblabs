package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	// Создаем новый reader для чтения данных из соединения
	reader := bufio.NewReader(conn)

	for {
		// Читаем строку из соединения
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			return
		}

		// Обрабатываем полученное сообщение
		response := message + "__Ответ\n"

		// Отправляем ответ обратно клиенту
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Ошибка отправки ответа:", err)
			return
		}

		fmt.Printf("Получено сообщение: %sОтправлен ответ: %s", message, response)
	}
}

func main() {
	// Задаем адрес и порт для слушания
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка разрешения адреса:", err)
		return
	}

	// Создаем новый TCP-сервер
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println("Ошибка создания сервера:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Сервер запущен на порту 8080. Ожидаем подключения...")

	for {
		// Принимаем новое соединение
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка приема соединения:", err)
			continue
		}

		// Обрабатываем новое соединение в отдельной горутине
		go handleConnection(conn)
	}
}
