package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/bymi15/kayounmusic/db"
	"github.com/bymi15/kayounmusic/db/models"
	"github.com/joho/godotenv"
)

type PageData struct {
	FilmMusics []models.FilmMusic
	Works      []models.Work
	Scores     []models.Score
}

var client db.MongoDbClient

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Failed to load .env file")
	}
	client = db.InitMongoClient()

	http.HandleFunc("/", handleTemplate)

	http.HandleFunc("/delete/filmmusic/", handleDeleteFilmMusic)
	http.HandleFunc("/delete/work/", handleDeleteWork)
	http.HandleFunc("/delete/score/", handleDeleteScore)

	log.Println("Listening on port 3000...")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func loadPageData() PageData {
	start := time.Now()
	filmMusics, err := client.GetFilmMusics()
	if err != nil {
		log.Fatal(err)
	}
	works, err := client.GetWorks()
	if err != nil {
		log.Fatal(err)
	}
	scores, err := client.GetScores()
	if err != nil {
		log.Fatal(err)
	}
	data := PageData{
		FilmMusics: filmMusics,
		Works:      works,
		Scores:     scores,
	}
	elapsed := time.Since(start)
	log.Printf("Time elapsed to fetch data from MongoDB: %s", elapsed)
	return data
}

func parseYoutubeIdFromUrl(youtubeUrl string) string {
	u, err := url.Parse(youtubeUrl)
	if err != nil {
		panic(err)
	}
	q := u.Query()
	return q.Get("v")
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		editId := r.FormValue("editId")
		switch dataType := r.FormValue("type"); dataType {
		case "filmMusic":
			comingSoon := false
			if r.FormValue("comingSoon") == "on" {
				comingSoon = true
			}
			filmMusic := models.FilmMusic{
				Title:      r.FormValue("title"),
				YoutubeId:  parseYoutubeIdFromUrl(r.FormValue("youtubeUrl")),
				Img:        r.FormValue("img"),
				ComingSoon: comingSoon,
			}
			var err error
			if editId != "" {
				err = client.UpdateFilmMusic(editId, filmMusic)
			} else {
				err = client.CreateFilmMusic(filmMusic)
			}
			if err != nil {
				fmt.Fprintf(w, "Create / Update film music err: %v", err)
				return
			}
		case "works":
			work := models.Work{
				Title:    r.FormValue("title"),
				Url:      r.FormValue("url"),
				Category: r.FormValue("category"),
			}
			var err error
			if editId != "" {
				err = client.UpdateWork(editId, work)
			} else {
				err = client.CreateWork(work)
			}
			if err != nil {
				fmt.Fprintf(w, "Create / Update work err: %v", err)
				return
			}
		case "scores":
			rescore := false
			if r.FormValue("rescore") == "on" {
				rescore = true
			}
			score := models.Score{
				Title:         r.FormValue("title"),
				Description:   r.FormValue("description"),
				PreviewFile:   r.FormValue("previewFile"),
				FullFile:      r.FormValue("fullFile"),
				Category:      r.FormValue("category"),
				Date:          r.FormValue("date"),
				SoundCloudUrl: r.FormValue("soundCloudUrl"),
				Img:           r.FormValue("img"),
				YoutubeId:     parseYoutubeIdFromUrl(r.FormValue("youtubeUrl")),
				Rescore:       rescore,
			}
			var err error
			if editId != "" {
				err = client.UpdateScore(editId, score)
			} else {
				err = client.CreateScore(score)
			}
			if err != nil {
				fmt.Fprintf(w, "Create / Update score err: %v", err)
				return
			}
		}
	}
	data := loadPageData()
	t := template.Must(template.ParseGlob("templates/*.gohtml"))
	err := t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleDeleteFilmMusic(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/delete/filmmusic/")
	log.Printf("Deleting FilmMusic id: %s", id)
	err := client.DeleteFilmMusic(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleDeleteWork(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/delete/work/")
	log.Printf("Deleting Work id: %s", id)
	err := client.DeleteWork(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/#works", http.StatusSeeOther)
}

func handleDeleteScore(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/delete/score/")
	log.Printf("Deleting Score id: %s", id)
	err := client.DeleteScore(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/#scores", http.StatusSeeOther)
}
