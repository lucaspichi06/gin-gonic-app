package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lucaspichi06/gin-gonic-app/internal/domain"
	post2 "github.com/lucaspichi06/gin-gonic-app/internal/post"
	"net/http"
	"strconv"
	"strings"
)

type post struct {
	service post2.Post
}

func NewPost(service post2.Post) post {
	return post{
		service: service,
	}
}

func (p *post) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	resp, err := p.service.Get(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (p *post) GetAll(c *gin.Context) {
	resp, err := p.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, resp)
}

func (p *post) GetAllByUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := p.service.GetByUserID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (p *post) Create(c *gin.Context) {
	var post domain.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("invalid json"))
		return
	}
	err = p.service.Create(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, post)
}

func (p *post) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var post domain.Post
	err = c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("invalid json"))
		return
	}

	err = p.service.Update(id, post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (p *post) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = p.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"error": "post deleted"})
}
