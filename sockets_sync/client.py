
#............................................................
import socket
import json
import logging
import sys
import time

def load_config(filename):
    try:
        with open(filename, 'r') as file:
            config = json.load(file)
            return config['server_address'], config['server_port']
    except FileNotFoundError:
        print(f"Файл не найден.")
        return None, None
    except json.JSONDecodeError:
        print(f"Ошибка : {filename}.")
        return None, None

def main():
    ip, port = load_config('config.json')
    log_filename = "client_" + sys.argv[1] + ".log"
    logging.basicConfig(filename=log_filename, level=logging.INFO, 
                        format='%(asctime)s - %(levelname)s - %(message)s') 
    
    try:
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
            s.connect((ip, port))
            logging.info(f'Connected to {ip}:{port}')

            time.sleep(2)
            message = (sys.argv[1]+"_СпиридоновП.К._207м")
            logging.info(f'Отправленное сообщение: {message}')
            s.sendall((message+"\n").encode('utf-8'))
            data = s.recv(1024)
            print(data.decode())
            logging.info( "Ответ сервера: "+ data.decode())
            #logging.info(f"Ответ сервера: {data.decode('utf-8', errors='ignore')}")
    except ConnectionRefusedError:
        print(f"Ошибка соединения с {ip}:{port}")
    except socket.gaierror:
        print(f"Ошибка. {ip}.")
    except Exception as e:
        print(f"Ошибка: {e}")
main()