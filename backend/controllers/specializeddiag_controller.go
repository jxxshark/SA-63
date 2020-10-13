package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jxxshark/app/ent"
	"github.com/jxxshark/app/ent/user"
	"github.com/jxxshark/app/ent/specializeddiag"
)

// SpecializeddiagController defines the struct for the specializeddiag controller
type SpecializeddiagController struct {
	client *ent.Client
	router gin.IRouter
}

// Specializeddiag defines the struct for the Specializeddiag
type Specializeddiag struct {
	SPECIALIZEDDIAGTYPE string
	SPECIALISTID        int
}

// CreateSpecializeddiag handles POST requests for adding specializeddiag entities
// @Summary Create specializeddiag
// @Description Create specializeddiag
// @ID create-specializeddiag
// @Accept   json
// @Produce  json
// @Param specializeddiag body ent.Specializeddiag true "Specializeddiag entity"
// @Success 200 {object} ent.Specializeddiag
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /specializeddiags [post]
func (ctl *SpecializeddiagController) CreateSpecializeddiag(c *gin.Context) {
	obj := Specializeddiag{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "specialized diagnostic binding failed",
		})
		return
	}

	sd, err := ctl.client.Specializeddiag.
		Create().
		SetSpecializeddiacnostictype(obj.SPECIALIZEDDIAGTYPE).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	us,err := ctl.client.User.
		UpdateOneID(int(obj.SPECIALISTID)).
		AddSpecializeddiag(sd).
		Save(context.Background())

	    if err != nil {
		c.JSON(400, gin.H{
			"error": "saving edge failed",
		})
		return
	}

	c.JSON(200, us)
}

// GetSpecializeddiag handles GET requests to retrieve a specializeddiag entity
// @Summary Get a specializeddiag entity by ID
// @Description get specializeddiag by ID
// @ID get-specializeddiag
// @Produce  json./main
// @Param id path int true "Specializeddiag ID"
// @Success 200 {object} ent.Specializeddiag
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /specializeddiags/{id} [get]
func (ctl *SpecializeddiagController) GetSpecializeddiag(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Specializeddiag.
		Query().
		Where(specializeddiag.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListSpecializeddiag handles request to get a list of specializeddiag entities
// @Summary List specializeddiag entities
// @Description list specializeddiag entities
// @ID list-specializeddiag
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Specializeddiag
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /specializeddiags [get]
func (ctl *SpecializeddiagController) ListSpecializeddiag(c *gin.Context) {
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

	specializeddiags, err := ctl.client.Specializeddiag.
		Query().
		WithUser().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, specializeddiags)
}

// DeleteSpecializeddiag handles DELETE requests to delete a specializeddiag entity
// @Summary Delete a specializeddiag entity by ID
// @Description get specializeddiag by ID
// @ID delete-specializeddiag
// @Produce  json
// @Param id path int true "Specializeddiag ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /specializeddiags/{id} [delete]
func (ctl *SpecializeddiagController) DeleteSpecializeddiag(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Specializeddiag.
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

// UpdateSpecializeddiag handles PUT requests to update a specializeddiag entity
// @Summary Update a specializeddiag entity by ID
// @Description update specializeddiag by ID
// @ID update-specializeddiag
// @Accept   json
// @Produce  json
// @Param id path int true "Specializeddiag ID"
// @Param specializeddiag body ent.Specializeddiag true "Specializeddiag entity"
// @Success 200 {object} ent.Specializeddiag
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /specializeddiags/{id} [put]
func (ctl *SpecializeddiagController) UpdateSpecializeddiag(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := Specializeddiag{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "specialized diagnostic binding failed",
		})
		return
	}
	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.SPECIALISTID))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	sd, err := ctl.client.Specializeddiag.
		UpdateOneID(int(id)).
		SetSpecializeddiacnostictype(obj.SPECIALIZEDDIAGTYPE).
		SetUser(u).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "update video failed",
		})
		return
	}

	c.JSON(200, sd)
}
//NewSpecializeddiagController creates and registers handles for the specializeddiag controller
func NewSpecializeddiagController(router gin.IRouter, client *ent.Client) *SpecializeddiagController {
	sd := &SpecializeddiagController{
		client: client,
		router: router,
	}
	sd.register()
	return sd
}

// InitSpecializeddiagController registers routes to the main engine
func (ctl *SpecializeddiagController) register() {
	specializeddiags := ctl.router.Group("/specializeddiags")

	specializeddiags.GET("", ctl.ListSpecializeddiag)

	// CRUD
	specializeddiags.POST("", ctl.CreateSpecializeddiag)
	specializeddiags.GET(":id", ctl.GetSpecializeddiag)
	specializeddiags.PUT(":id", ctl.UpdateSpecializeddiag)
	specializeddiags.DELETE(":id", ctl.DeleteSpecializeddiag)
}
