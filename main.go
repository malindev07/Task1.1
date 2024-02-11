package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Number string `json:"number"`
}

// не знаю как выполнить задачу не объявляя структуру глобальной
var users *Users

// инициализация и проверка данных в json
func checkJson() {
	data, err := os.ReadFile("db.json")
	err = json.Unmarshal(data, &users)
	if err != nil {
		panic(err)
	}

}

// Поиск телефона по имени
func (u *Users) FindNumber(name string) string {

	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].Name == name {
			return users.Users[i].Number
		}
	}
	return "Указанное имя не найдено"
}

// Поиск имени по телефону
func (u *Users) FindName(number string) string {
	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].Number == number {
			return users.Users[i].Name
		}
	}
	return "Указанный телефон не найден"
}

// Удаление строки Имя : Телефон
func (u *Users) DeleteRow(keyOrValue string) *Users {
	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].Name == keyOrValue || users.Users[i].Number == keyOrValue {
			users.Users = append(users.Users[:i], users.Users[i+1:]...)
		}
	}
	data, _ := json.MarshalIndent(users, "", "\t")
	_ = os.WriteFile("db.json", data, 02)
	return users
}

// Вывод всего списка телефонной книги
func (u *Users) ShowDB() *Users {
	return users
}

// Добавление строки Имя : Телефон
func (u *Users) AddRow(name, number string) *Users {
	isTrue := false
	str := User{Name: name, Number: number}

	for _, val := range u.Users {
		if val.Name == str.Name {
			fmt.Println("Уже есть", str.Name)
			isTrue = true
			break
		}
	}

	if !isTrue {
		u.Users = append(u.Users, str)
		fmt.Println(u.Users)
		data, _ := json.MarshalIndent(u, "", "\t")
		_ = os.WriteFile("db.json", data, 02)
	}
	return u

}

// обновление номера по имени
func (u *Users) ChangeNumber(name, number string) *Users {
	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].Name == name {
			users.Users[i].Number = number
		}
	}
	data, _ := json.MarshalIndent(users, "", "\t")
	_ = os.WriteFile("db.json", data, 02)

	return users

}

// далее идут обработчики на созданные функции
func FindNumberHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, users.FindNumber("Lidia"))
}
func FindNameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, users.FindName("89035043155"))
}
func ShowDBHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, users.ShowDB())
}

func DeleteRowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, users.DeleteRow("Vladimir"))
}

func AddRowHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, users.AddRow("Vladimir", "89039460606"))

}

func ChangeHumberHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, users.ChangeNumber("Vladimir", "89039460605"))
}

func main() {
	checkJson()
	//fmt.Println(users)
	users.AddRow("Tes", "Test")
	//wg := sync.WaitGroup{}
	//wg.Add(2)
	// вызываем обработчик с некой функцией внутри

	//http.HandleFunc("/", AddRowHandler)
	////resp, err := http.Get("localhost:8000")
	////
	////fmt.Println(resp.Body, err)
	////wg.Wait()
	//// Слушаем наш локальный сервер
	//http.ListenAndServe("localhost:8000", nil)

}
