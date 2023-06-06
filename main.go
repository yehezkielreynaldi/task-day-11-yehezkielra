package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// e = echo package
	// GET/POST = run the method
	// "/" = endpoint/routing (ex. localhost:5000'/' | ex. dumbways.id'/lms')
	// helloWorld = function that will run if the routes are opened

	// Serve a static files from "public" directory
	e.Static("/public", "public")

	// Routing
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/my-project", myProject)
	e.GET("/project-detail/:id", projectDetail)
	e.POST("/add-project", addProject)
	e.GET("/testimonials", testimonials)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil { // null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"messsage": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func myProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/my-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Id":      id,
		"Title":   "Dumbways Web App",
		"Content": `Lorem ipsum dolor sit amet consectetur adipisicing elit. Rerum nam excepturi reprehenderit molestias vero laboriosam, accusamus aliquam? Voluptatem, libero consectetur quaerat aspernatur porro quidem error facere reprehenderit omnis earum nisi quos aperiam soluta vel tempora dignissimos possimus facilis quas, animi eaque nostrum suscipit perferendis optio ullam? Praesentium excepturi animi eius illum autem voluptates labore. Libero excepturi nisi ipsam veritatis est voluptatibus voluptates recusandae sapiente dolore distinctio! Cumque asperiores corporis necessitatibus, quisquam neque adipisci. Itaque, natus harum sint eum nesciunt ea ipsa perferendis porro soluta magni, corporis asperiores accusamus sed minus? Laudantium aperiam rem beatae voluptatum ipsum ipsam at dignissimos nobis. <br /> Lorem ipsum dolor sit amet consectetur adipisicing elit. Magni nam animi officiis ipsum? Eligendi voluptas dolorem, ab consequatur, neque magni veniam, modi quaerat labore quo ut dolorum velit a voluptates sint dolores dolor! Similique dolores unde adipisci neque iure exercitationem, distinctio sed debitis officiis nulla vero! Deleniti veniam quae, veritatis accusantium vero dicta, nihil modi atque voluptatem dolor rem dolorum dolore eum sequi harum nesciunt repellat quidem vitae architecto! Molestiae ipsa reprehenderit nam deserunt assumenda a magni, harum pariatur at exercitationem rem officiis ipsum repellendus nulla in laudantium rerum delectus natus facilis. Itaque nesciunt eveniet debitis consectetur veniam repellat modi! <br /> Lorem ipsum dolor sit amet consectetur adipisicing elit. Soluta quia eaque quae voluptatem nisi odit obcaecati rem vero ullam in porro maiores aliquam alias enim consequatur vitae tempora, nesciunt modi beatae animi. Dignissimos fugiat corrupti amet mollitia, et esse rem voluptas fuga incidunt dolores iure consectetur repudiandae placeat, minima adipisci praesentium nobis debitis minus distinctio atque. Aliquid mollitia totam nemo natus. Sunt et culpa blanditiis, commodi tempore eligendi itaque eius aliquam? Sequi incidunt molestiae odio cupiditate dicta voluptates a et commodi, facere nihil ratione ipsa sit quibusdam, ipsum autem mollitia? Quae fugit laboriosam cum numquam perferendis aperiam laudantium et vitae. <br /> Lorem ipsum dolor sit amet consectetur, adipisicing elit. Ipsam animi doloremque iste, ab iure vero praesentium repellat harum quisquam amet id dolorum unde dolorem magnam laborum modi. Earum voluptatibus, tempora minima vel culpa minus perferendis sapiente nostrum harum, inventore quaerat voluptates, explicabo obcaecati repudiandae possimus unde ullam sequi suscipit accusamus ea! Mollitia odio harum omnis nam porro aut corporis nisi sit nobis nulla, tempore explicabo animi non corrupti ipsam, libero ratione cum minus esse reiciendis! Dolore laboriosam, ab provident fuga sapiente praesentium natus aspernatur impedit excepturi ea quaerat sunt quasi voluptates ipsum veritatis architecto. Porro quas quos ratione eligendi rerum.`,
	}

	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func testimonials(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func addProject(c echo.Context) error {
	projectName := c.FormValue("projectName")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	description := c.FormValue("descriptionProject")
	tech1 := c.FormValue("tech1")
	tech2 := c.FormValue("tech2")
	tech3 := c.FormValue("tech3")
	tech4 := c.FormValue("tech4")

	println("Project Name : " + projectName)
	println("Start Date : " + startDate)
	println("End Date : " + endDate)
	println("Description : " + description)
	println("Technologies : " + tech1)
	println("Technologies : " + tech2)
	println("Technologies : " + tech3)
	println("Technologies : " + tech4)

	return c.Redirect(http.StatusMovedPermanently, "/")
}
