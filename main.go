package main

import (
	"bufio"
	"fmt"
	"html"
	"net/http"
	"os"
	"strings"
	"sync"
)

const filename = "bookmarks.txt"

var (
	bookmarksMu sync.RWMutex
	bookmarks   = make(map[string]string)
)

func loadBookmarks() error {
	f, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		name := strings.TrimSpace(parts[0])
		url := strings.TrimSpace(parts[1])
		if name != "" && url != "" {
			bookmarks[name] = url
		}
	}
	return scanner.Err()
}

func saveBookmarks() error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for name, url := range bookmarks {
		if _, err := fmt.Fprintf(w, "%s=%s\n", name, url); err != nil {
			return err
		}
	}
	return w.Flush()
}

func main() {
	if err := loadBookmarks(); err != nil {
		fmt.Printf("err %v\n", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var message string

		if r.Method == http.MethodPost {
			r.ParseForm()
			action := r.FormValue("action")
			name := strings.TrimSpace(r.FormValue("name"))
			url := strings.TrimSpace(r.FormValue("url"))

			switch action {
			case "add":
				if name == "" || url == "" {
					message = "err"
				} else {
					bookmarksMu.Lock()
					bookmarks[name] = url
					err := saveBookmarks()
					bookmarksMu.Unlock()

					if err != nil {
						message = "err"
					} else {
						message = "added"
					}
				}
			case "delete":
				if name == "" {
					message = "err"
				} else {
					bookmarksMu.Lock()
					if _, exists := bookmarks[name]; exists {
						delete(bookmarks, name)
						err := saveBookmarks()
						if err != nil {
							message = "err"
						} else {
							message = "deleted"
						}
					} else {
						message = "not found"
					}
					bookmarksMu.Unlock()
				}
			}
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		fmt.Fprintln(w, "<h1></h1>")
		if message != "" {
			fmt.Fprintln(w, "<p><b>"+html.EscapeString(message)+"</b></p>")
		}

		searchQuery := strings.TrimSpace(r.URL.Query().Get("search"))

		type item struct{ name, url string }
		var displayList []item

		bookmarksMu.RLock()
		queryLower := strings.ToLower(searchQuery)
		for name, url := range bookmarks {
			if searchQuery == "" || strings.Contains(strings.ToLower(name), queryLower) || strings.Contains(strings.ToLower(url), queryLower) {
				displayList = append(displayList, item{name: name, url: url})
			}
		}
		isEmpty := len(displayList) == 0
		bookmarksMu.RUnlock()

		fmt.Fprintf(w, `
			<form method="GET" action="/">
				<input type="text" name="search" placeholder="Поиск" value="%s">
				<button type="submit">Найти</button>
				<a href="/">Сбросить</a>
			</form>
			<hr>
		`, html.EscapeString(searchQuery))

		if isEmpty {
			if searchQuery != "" {
				fmt.Fprintln(w, "<p>Ничего не найдено.</p>")
			} else {
				fmt.Fprintln(w, "<p>Список закладок пуст.</p>")
			}
		} else {
			if searchQuery != "" {
				fmt.Fprintf(w, "<p>Найдено по запросу: <b>%s</b></p>", html.EscapeString(searchQuery))
			}
			for _, it := range displayList {
				safeName := html.EscapeString(it.name)
				safeURL := html.EscapeString(it.url)
				fmt.Fprintf(w, "<p>%s: <a href=\"%s\" target=\"_blank\">%s</a></p>", safeName, safeURL, safeURL)
			}
		}
		fmt.Fprintln(w, `
			<hr>
			<h3>Добавить закладку</h3>
			<form method="POST">
				<input type="hidden" name="action" value="add">
				<p><input name="name" placeholder="Название" required></p>
				<p><input name="url" placeholder="Ссылка" required></p>
				<table width="100%%"><tr>
					<td align="left"><button type="submit">Добавить</button></td>
				</tr></table>
			</form>

			<h3>Удалить закладку</h3>
			<form method="POST">
				<input type="hidden" name="action" value="delete">
				<p><input name="name" placeholder="Название" required></p>
				<table width="100%%"><tr>
					<td align="left"><button type="submit">Удалить</button></td>
				</tr></table>
			</form>
		`)
	})

	http.ListenAndServe(":8080", nil)
}