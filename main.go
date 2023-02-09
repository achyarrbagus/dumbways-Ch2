package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

var Data = map[string]interface{}{
	"title":   "Personal Web",
	"isLogin": "on",
}

type Blog struct {
	Id           int
	StartDate    string
	EndDate      string
	Title        string
	Author       string
	Content      string
	Deference    string
	StartDateStr string
	EndDateStr   string
	//
	TechOne bool
	TechTwo bool
	TechTre bool
	TechFor bool
}

var BlogData = []Blog{
	{
		Id:           0,
		Title:        "Dumbways Mobile app",
		Author:       "Uciha-Yanto",
		Content:      "Dumbwasdasdasd",
		StartDate:    "12 Mei 2022",
		EndDate:      "15 Mei 2023",
		TechOne:      true,
		TechTwo:      true,
		TechFor:      true,
		TechTre:      true,
		Deference:    "Deadline: 3Day ",
		StartDateStr: "2022-02-08",
		EndDateStr:   "2021-04-08",
	},
	{
		Id:           1,
		Title:        "Dumbways Web app",
		Author:       "Senju-Akbar",
		Content:      "Dumbwasdasdasasdasdsadasldmjadfbsd kamdkeqmroqd",
		StartDate:    "12 Mei 2023",
		EndDate:      "16 Mei 2025",
		TechFor:      true,
		TechTwo:      true,
		Deference:    "Deadline: 4Day ",
		StartDateStr: "2023-02-08",
		EndDateStr:   "2023-04-08",
	},
}

func main() {

	route := mux.NewRouter()
	// crated static files
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/helloworld", helloWorld).Methods("GET")
	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/add-project", addProject).Methods("GET")
	route.HandleFunc("/add-project/blog", projectBlog).Methods("POST")
	route.HandleFunc("/contact-me", contact).Methods("GET")
	route.HandleFunc("/blog-content/{id}", blogContent).Methods("GET")
	route.HandleFunc("/delete-content/{id}", deleteBlog).Methods("GET")
	route.HandleFunc("/edit-content/{id}", editBlog).Methods("GET")
	route.HandleFunc("/update-content/blog{id}", updButton).Methods("POST")

	fmt.Println("server running on port : 5000")
	http.ListenAndServe("localhost:5000", route)
}

// update Button func handler
func updButton(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// getting value form
	projectName := r.PostForm.Get("projectName")
	projectBlog := r.PostForm.Get("projectBlog")
	// getting Date input
	startDateInput := r.PostForm.Get("startDate")
	endDateInput := r.PostForm.Get("endDate")
	// parsing time from string value to int value
	timeTemplate := "2006-01-02"
	getStartDt, _ := time.Parse(timeTemplate, startDateInput)
	getEndDt, _ := time.Parse(timeTemplate, endDateInput)
	// parsing time format from year/moth/date to date/month/year
	// final value
	startDate := getStartDt.Format("2 jan 2006")
	endDate := getEndDt.Format("2 jan 2006")

	// looking defferen from 2 value date, endDate - startDate
	var finalDeference string
	deference := getEndDt.Sub(getStartDt) //using method sub from time.time strik to get deference value from 2 date input

	if deference.Hours()/24 < 30 {
		fd := strconv.FormatFloat(deference.Hours()/24, 'f', 0, 64) //format float to convert int value to string, 'f to setting number behind the comma ,64 is int type
		finalDeference = "Duration :" + fd + "Day"                  //final value for one day
	} else if deference.Hours()/24/30 < 12 { //to get value month
		fd := strconv.FormatFloat(deference.Hours()/24/30, 'f', 0, 64) //same like before
		finalDeference = "Duration :" + fd + "Month"                   //final value for one day
	} else {
		fd := strconv.FormatFloat(deference.Hours()/24/30/12, 'f', 0, 64) //same like before
		finalDeference = "Duration: " + fd + "Year"
	}

	//thecnology value
	//this function to return value from on to true value
	var techOneInput bool
	if r.PostForm.Get("techOne") == "on" {
		techOneInput = true
	}
	var techTwoInput bool
	if r.PostForm.Get("techTwo") == "on" {
		// same like before
		techTwoInput = true
	}
	var techTreInput bool
	if r.PostForm.Get("techTre") == "on" {
		techTwoInput = true
	}
	var techForInput bool
	if r.PostForm.Get("techFor") == "on" {
		techForInput = true
	}
	BlogData[id].Id = id
	BlogData[id].Title = projectName
	BlogData[id].Content = projectBlog
	BlogData[id].Author = projectBlog
	BlogData[id].StartDateStr = startDateInput
	BlogData[id].EndDateStr = endDateInput
	BlogData[id].Deference = finalDeference
	BlogData[id].TechOne = techOneInput
	BlogData[id].TechTwo = techTwoInput
	BlogData[id].TechTre = techTreInput
	BlogData[id].TechFor = techForInput
	BlogData[id].StartDate = startDate
	BlogData[id].EndDate = endDate

	w.WriteHeader(http.StatusOK)
	fmt.Println(startDateInput)
}

// editblog Handler
func editBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text-html;charset=utf-8")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// parsing template html
	tmpl, _ := template.ParseFiles("./views/edit-project.html")

	BlogEdit := Blog{}

	for index, data := range BlogData {
		if index == id {
			BlogEdit = Blog{
				Id:           data.Id,
				Title:        data.Title,
				Content:      data.Content,
				Author:       data.Author,
				StartDateStr: data.StartDate,
				EndDateStr:   data.EndDate,
				Deference:    data.Deference,
				TechOne:      data.TechOne,
				TechTwo:      data.TechTwo,
				TechTre:      data.TechTre,
				TechFor:      data.TechFor,
			}
		}
	}
	resp := map[string]interface{}{
		"Title": Data,
		"data":  BlogEdit,
	}
	tmpl.Execute(w, resp)

}

// delete func
func deleteBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text-html;charset=utf-8")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	BlogData = append(BlogData[:id], BlogData[id+1:]...)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	fmt.Println(id)

}

// project blog handler
func projectBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text-html;charset=utf-8")
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	projectName := r.PostForm.Get("projectName")
	projectBlog := r.PostForm.Get("projectBlog")
	// getting Date input
	startDateInput := r.PostForm.Get("startDate")
	endDateInput := r.PostForm.Get("endDate")
	//
	getStartDt, _ := time.Parse("2006-01-02", startDateInput)
	startDate := getStartDt.Format("2 jan 2006")
	//
	getEndDt, _ := time.Parse("2006-01-02", endDateInput)
	endDate := getEndDt.Format("2 Jan 2006")

	var finalDeference string
	deference := getEndDt.Sub(getStartDt)

	if deference.Hours()/24 < 30 {
		fd := strconv.FormatFloat(deference.Hours()/24, 'f', 0, 64)
		finalDeference = "Duration: " + fd + "Day"
	} else if deference.Hours()/24/30 < 12 {
		fd := strconv.FormatFloat(deference.Hours()/24/30, 'f', 0, 64)
		finalDeference = "Duration: " + fd + "Month"
	} else {
		fd := strconv.FormatFloat(deference.Hours()/24/30/12, 'f', 0, 64)
		finalDeference = "Duration: " + fd + "Year"
	}

	//thecnology value
	//this function to return value from on to true value
	var techOneInput bool
	if r.PostForm.Get("techOne") == "on" {
		techOneInput = true
	}
	var techTwoInput bool
	if r.PostForm.Get("techTwo") == "on" {
		// same like before
		techTwoInput = true
	}
	var techTreInput bool
	if r.PostForm.Get("techTre") == "on" {
		techTwoInput = true
	}
	var techForInput bool
	if r.PostForm.Get("techFor") == "on" {
		techForInput = true
	}

	var newData = Blog{
		Title:        projectName,
		Author:       projectBlog,
		Content:      projectBlog,
		TechOne:      techOneInput,
		TechTwo:      techTwoInput,
		TechTre:      techTreInput,
		TechFor:      techForInput,
		StartDateStr: startDateInput,
		EndDateStr:   endDateInput,
		//
		StartDate: startDate,
		EndDate:   endDate,
		Deference: finalDeference,
	}

	BlogData = append(BlogData, newData)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

// blog content handler
func blogContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text-html;charset=utf-8")

	// parsing template html
	var tmpl, err = template.ParseFiles("./views/blog-content.html")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}

	BlogDetail := Blog{}

	for index, data := range BlogData {
		if index == id {
			BlogDetail = Blog{
				Id:        id,
				Title:     data.Title,
				Author:    data.Author,
				Content:   data.Content,
				TechOne:   data.TechOne,
				TechTwo:   data.TechTwo,
				TechTre:   data.TechTre,
				TechFor:   data.TechFor,
				StartDate: data.StartDate,
				EndDate:   data.EndDate,
			}
		}
	}
	resp := map[string]interface{}{
		"Title": Data,
		"data":  BlogDetail,
	}

	fmt.Println(resp, id)
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resp)
}

// add project handler
func addProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text-html;charset=utf-8")

	var tmpl, err = template.ParseFiles("./views/add-project.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)
}

// contact me handler
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html;charset=utf-8")

	var tmpl, err = template.ParseFiles("./views/contact-me.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)

}

// home handler
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("./views/home.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}
	resp := map[string]interface{}{
		"Title": Data,
		"Data":  BlogData,
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resp)

}

// handler hello world
func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world"))

	fmt.Println(http.StatusOK)

}
