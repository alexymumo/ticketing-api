package controllers

import (
	"database/sql"
	"events/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ping"})
	}
}

func AvailableTickets(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		eventid := ctx.Param("eventid")
		query := "SELECT capacity FROM event WHERE eventid = ?"
		var capacity int
		err := db.QueryRow(query, eventid).Scan(&capacity)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "no event found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query db"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"eventid":  eventid,
			"capacity": capacity,
		})
	}
}

func CreateTicket(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		eventid := ctx.Param("eventid")
		userid, exists := ctx.Get("id")
		if !exists {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			return
		}
		tx, err := db.Begin()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to begin transaction"})
			return
		}
		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to rollback tx"})
			}
		}()
		var capacity int
		query := "SELECT capacity FROM event WHERE eventid = ? FOR UPDATE"
		err = tx.QueryRow(query, eventid).Scan(&capacity)
		if err != nil {
			if err == sql.ErrNoRows {
				tx.Rollback()
				ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			} else {
				tx.Rollback()
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get event"})
			}
			return
		}
		if capacity <= 0 {
			tx.Rollback()
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "tickets sold out"})
			return
		}
		updateQuery := "UPDATE event SET capacity = capacity - 1 WHERE eventid = ?"
		_, err = tx.Exec(updateQuery, eventid)
		if err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update capacity"})
			return
		}
		queryInsert := "INSERT into ticket (eventid,userid) VALUES (?,?)"
		result, err := tx.Exec(queryInsert, eventid, userid)
		if err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to to create a ticket"})
			return
		}
		ticketId, err := result.LastInsertId()
		if err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save user"})
			return
		}
		if err := tx.Commit(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
			return
		}
		var ticket models.Ticket
		retrieve := "SELECT eventid,userid FROM ticket WHERE ticketid = ?"
		err = db.QueryRow(retrieve, ticketId).Scan(&ticket.TicketID, &ticket.UserID, &ticket.EventId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query ticket"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ticket successfully created",
			"ticket":  ticket,
		})
	}
}
