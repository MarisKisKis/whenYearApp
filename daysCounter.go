package main

import (
	"github.com/gin-gonic/gin"
	"log"
	_ "log"
	"net/http"
	"strconv"
	"time"
)

func countDays(c *gin.Context) {
	year := c.Params.ByName("year")
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		log.Println("Введено не число!")
		panic(err)
	}
	// преобразование полученного года в дату
	t1 := time.Date(yearInt, 1, 1, 0, 0, 0, 0, time.UTC)
	// преобразование текущей даты
	t2 := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(),
		time.Now().Second(), time.Now().Nanosecond(), time.UTC)
	var (
		daysLeft int
		daysGone int
	)
	if t1.Before(t2) {
		daysGone = int(t2.Sub(t1) / time.Hour / 24)
		if daysGone == 106751 {
			c.Abort()
			log.Println("Вы ввели слишком мальнькое значение!")
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"Days gone": daysGone})
		}
	} else if t2.Before(t1) {
		daysLeft = int(t1.Sub(t2) / time.Hour / 24)
		if daysLeft == 106751 {
			c.Abort()
			log.Println("Вы ввели слишком большое значение!")
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"Days left": daysLeft})
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("Not found")
	}
}
