package controllers

import (
	"fmt"
	"golang-ptp/go-todos/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetTodos godoc
// @Summary Get all todos
// @Description get all todos
// @ID get-todos
// @Produce  json
// @Success 200 {array} models.Todo
// @Router /todos [get]
func GetTodos(c *gin.Context) {
	db := models.GetDB()

	var Todos []models.Todo

	err := db.Debug().Find(&Todos).Error

	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Internal server error",
		})
		return
	}

	c.JSON(200, Todos)
}

// GetTodo godoc
// @Summary Get a todo
// @Description get a todo
// @ID get-todo
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 200 {object} models.Todo
// @Router /todos/{id} [get]
func GetTodo(c *gin.Context) {
	db := models.GetDB()

	id := c.Param("id")
	var Todo models.Todo

	err := db.Debug().Where("id =  ?", id).Find(&Todo).Error

	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Internal server error",
		})
		return
	}

	c.JSON(200, Todo)
}

// CreateTodo godoc
// @Summary Create a todo
// @Description create a new todo
// @ID create-todo
// @Accept  json
// @Produce  json
// @Param requestbody body models.Todo true "request body json"
// @Success 200 {object} models.Todo
// @Router /todos [post]
func CreateTodo(c *gin.Context) {
	db := models.GetDB()

	var Todo models.Todo
	var TodoJSON = map[string]string{}

	if err := c.ShouldBindJSON(&TodoJSON); err != nil {
		c.JSON(500, gin.H{
			"messages": "error bind json",
			"err":      err,
		})
		return
	}

	TodoDueDate, _ := time.Parse("2006-01-02", TodoJSON["dueDate"])

	Todo = models.Todo{
		Title:       TodoJSON["title"],
		Description: TodoJSON["description"],
		DueDate:     TodoDueDate,
	}

	errCreate := db.Debug().Create(&Todo).Error

	if errCreate != nil {
		c.JSON(500, gin.H{
			"messages": "error create",
		})
		return
	}

	fmt.Println(Todo)

	c.JSON(201, Todo)
}

// UpdateTodo godoc
// @Summary Update a todo
// @Description update a todo
// @ID update-todo
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Param requestbody body models.Todo true "request body json"
// @Success 200 {object} models.Todo
// @Router /todos/{id} [put]
func PutTodo(c *gin.Context) {
	db := models.GetDB()

	id := c.Param("id")
	var Todo models.Todo
	var TodoJSON = map[string]string{}

	if err := c.ShouldBindJSON(&TodoJSON); err != nil {
		c.JSON(500, gin.H{
			"messages": "error bind json",
			"err":      err,
		})
		return
	}

	TodoDueDate, _ := time.Parse("2006-01-02", TodoJSON["dueDate"])

	Todo = models.Todo{
		Title:       TodoJSON["title"],
		Description: TodoJSON["description"],
		DueDate:     TodoDueDate,
	}

	errUpdate := db.Debug().Model(&Todo).Where("id = ?", id).Updates(Todo).Error

	if errUpdate != nil {
		c.JSON(500, gin.H{
			"messages": "error update",
		})
		return
	}

	res, _ := strconv.Atoi(id)

	Todo.ID = uint(res)
	c.JSON(200, Todo)
}

// DeleteTodo godoc
// @Summary Delete a todo
// @Description delete a todo
// @ID delete-todo
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 200 {object} models.Todo
// @Router /todos/{id} [delete]
func DeleteTodo(c *gin.Context) {
	db := models.GetDB()

	id := c.Param("id")
	var Todo models.Todo

	errDelete := db.Debug().Delete(&Todo, id).Error
	if errDelete != nil {
		c.JSON(500, gin.H{
			"messages": "error delete",
		})
		return
	}

	c.JSON(200, Todo)
}
