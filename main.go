package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Definición de los modelos
type Team struct {
	gorm.Model
	Nombre        string
	Puntos        int
	GolesAFavor   int
	GolesEnContra int
}

type Match struct {
	gorm.Model
	LocalTeamID       uint
	VisitanteTeamID   uint
	GolesDelLocal     int
	GolesDelVisitante int
}

func main() {
	r := gin.Default()

	dsn := "host=localhost user=postgres password=postgres dbname=cuadrangular port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Team{}, &Match{})

	r.POST("/api/teams", func(c *gin.Context) {
		var teams []Team
		if err := c.BindJSON(&teams); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		db.Create(&teams)

		c.JSON(200, gin.H{"message": "Teams registered successfully"})
	})

	r.POST("/api/matches", func(c *gin.Context) {
		var matches []Match
		if err := c.BindJSON(&matches); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		for _, match := range matches {
			var equipoLocal, equipoVisitante Team
			db.First(&equipoLocal, match.LocalTeamID)
			db.First(&equipoVisitante, match.VisitanteTeamID)

			equipoLocal.GolesAFavor += match.GolesDelLocal
			equipoLocal.GolesEnContra += match.GolesDelVisitante
			equipoVisitante.GolesAFavor += match.GolesDelVisitante
			equipoVisitante.GolesEnContra += match.GolesDelLocal

			if match.GolesDelLocal > match.GolesDelVisitante {
				equipoLocal.Puntos += 3
			} else if match.GolesDelLocal < match.GolesDelVisitante {
				equipoVisitante.Puntos += 3
			} else {
				equipoLocal.Puntos++
				equipoVisitante.Puntos++
			}

			db.Save(&equipoLocal)
			db.Save(&equipoVisitante)
		}

		c.JSON(200, gin.H{"message": "Match results submitted successfully"})
	})

	r.GET("/api/standings", func(c *gin.Context) {
		var teams []Team
		db.Order("puntos desc, goles_a_favor desc").Find(&teams)
		c.JSON(200, gin.H{"standings": teams})
	})

	r.Run(":8080")
}
