package main

import (
	"log"
	"net/http"
	"strconv"
)

type StudentItem struct {
	Rank        int
	PhotoURL    string
	Composition []string
	Name        string
}

var studentItems []StudentItem

func main() {
	studentItems = []StudentItem{
		{
			Rank:        5,
			PhotoURL:    "https://slovnet.ru/wp-content/uploads/2019/04/9-22.jpg",
			Composition: []string{"Учаться в 5 классе"},
			Name:        "Льюис",
		},
		{
			Rank:        5,
			PhotoURL:    "https://slovnet.ru/wp-content/uploads/2019/04/22-2.jpeg",
			Composition: []string{"Учаться в 5 классе"},
			Name:        "Уилбер",
		},
		{
			Rank:        5,
			PhotoURL:    "https://slovnet.ru/wp-content/uploads/2019/04/3-23.jpg",
			Composition: []string{"Учаться в 5 классе"},
			Name:        "Майки",
		},
	}

	http.HandleFunc("/menu", menuListHandler)

	err := http.ListenAndServe("localhost:8888", http.DefaultServeMux)
	if err != nil {
		log.Fatal(err)
	}
}

func menuListHandler(w http.ResponseWriter, r *http.Request) {
	rankParam := r.FormValue("rank")
	rank, err := strconv.Atoi(rankParam)
	if err == nil && rank != 0 {
		var filteredMenu []StudentItem
		for _, menuItem := range studentItems {
			if menuItem.Rank <= rank {
				filteredMenu = append(filteredMenu, menuItem)
			}
		}
		studentItems = filteredMenu
	}

	menuItemsHtml := ``

	for _, menuItem := range studentItems {
		compositionHtml := ``
		for _, item := range menuItem.Composition {
			compositionHtml += `
			<li>
				` + item + `
			</li>	
		`
		}

		menuItemsHtml += `
			<tr>
				<td>` + menuItem.Name + `</td>
				<td>
					<img src="` + menuItem.PhotoURL + `" alt="Бургер"
						 class="menu-item-photo">
				</td>
				<td>
					<ul>
						` + compositionHtml + `
					</ul>
				</td>
				<td>` + strconv.Itoa(menuItem.Rank) + ` класс</td>
				<td>
					<button class="button button1">Редактировать</button>
				</td>
				<td>
					<button class="button button2">Удалить</button>
				</td>
			</tr>
		`
	}

	template := `
		<!DOCTYPE html>
		<html>
		<head>
			<style>
				table {
					font-family: arial, sans-serif;
					border-collapse: collapse;
					width: 100%;
				}
		
				td, th {
					border: 1px solid #dddddd;
					text-align: left;
					padding: 8px;
				}
		
				tr:nth-child(even) {
					background-color: #dddddd;
				}
		
				.menu-item-photo {
					height: 100px;
					width: auto;
				}
				.button {
					border: none;
					color: white;
					padding: 15px 32px;
					text-align: center;
					text-decoration: none;
					display: inline-block;
					font-size: 16px;
					margin: 4px 2px;
					cursor: pointer;
				}
		
				.button1 {background-color: #4ca0af;} /* Green */
				.button2 {background-color: #ba0016;} /* Blue */
			</style>
			<meta charset="utf-8">
		</head>
		<body>
		
		<h2>Управление меню</h2>

		<form action="http://localhost:8888/menu">
			<label for="price">Класс до:</label><br>
			<input type="text" id="rank" name="rank"><br>
			<input type="submit" value="Фильтр">
		</form>
		
		<br>
		
		<table>
			<tr>
				<th>Имя</th>
				<th>Фото</th>
				<th>Информация</th>
				<th>Класс</th>
				<th>Редактировать</th>
				<th>Удалить</th>
			</tr>
			` + menuItemsHtml + `
		</table>
		
		</body>
		</html>
		`

	w.Write([]byte(template))
}
