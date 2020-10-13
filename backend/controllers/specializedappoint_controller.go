package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/jxxshark/app/ent"
	"github.com/jxxshark/app/ent/specializeddiag"
	"github.com/jxxshark/app/ent/patient"
	"github.com/jxxshark/app/ent/user"
	"github.com/gin-gonic/gin"
)

// SpecializedappointController defines the struct for the specializedappoint controller
type SpecializedappointController struct {
	client *ent.Client
	router gin.IRouter
}

// Specializedappoint defines the struct for the specializedappoint
type Specializedappoint struct {
	PatientID int
	SpecializeddiagID int
	UserID     int
	Date  string
}

// CreateSpecializedappoint handles POST requests for adding specializedappoint entities
// @Summary Create specializedappoint
// @Description Create specializedappoint
// @ID create-specializedappoint
// @Accept   json
// @Produce  json
// @Param specializedappoint body Specializedappoint true "Specializedappoint entity"
// @Success 200 {object} ent.Specializedappoint
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /specializedappoints [post]
func (ctl *SpecializedappointController) CreateSpecializedappoint(c *gin.Context) {
	obj := Specializedappoint{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "specializedappoint binding failed",
		})
		return
	}

	p, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.PatientID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "patient not found",
		})
		return
	}

	sd, err := ctl.client.Specializeddiag.
		Query().
		Where(specializeddiag.IDEQ(int(obj.SpecializeddiagID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "specialized diagnostic  not found",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.UserID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}
	times, err := time.Parse(time.RFC3339, obj.Date)

	sa, err := ctl.client.Specializedappoint.
		Create().
		SetPatient(p).
		SetSpecializeddiag(sd).
		SetUser(u).
		SetDate(times).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, sa)
}

// ListSpecializedappoint handles request to get a list of specializedappoint entities
// @Summary List specializedappoint entities
// @Description list specializedappoint entities
// @ID list-specializedappoint
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Specializedappoint
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /specializedappoints [get]
func (ctl *SpecializedappointController) ListSpecializedappoint(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	specializedappoints, err := ctl.client.Specializedappoint.
		Query().
		WithPatient().
		WithSpecializeddiag().
		WithUser().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, specializedappoints)
}

// DeleteSpecializedappoint handles DELETE requests to delete a specializedappoint entity
// @Summary Delete a specializedappoint entity by ID
// @Description get specializedappoint by ID
// @ID delete-specializedappoint
// @Produce  json
// @Param id path int true "Specializedappoint ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /specializedappoints/{id} [delete]
func (ctl *SpecializedappointController) DeleteSpecializedappoint(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Specializedappoint.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// NewSpecializedappointController creates and registers handles for the specializedappoint controller
func NewSpecializedappointController(router gin.IRouter, client *ent.Client) *SpecializedappointController {
	drc := &SpecializedappointController{
		client: client,
		router: router,
	}
	drc.register()
	return drc
}

// InitUserController registers routes to the main engine
func (ctl *SpecializedappointController) register() {
	specializedappoints := ctl.router.Group("/specializedappoints")

	specializedappoints.GET("", ctl.ListSpecializedappoint)

	// CRUD
	specializedappoints.POST("", ctl.CreateSpecializedappoint)
	specializedappoints.DELETE(":id", ctl.DeleteSpecializedappoint)
}
