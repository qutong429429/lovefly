package controllers

import (
    // "encoding/json"
    // "lovefly/app/models"
    // "fmt"
    "github.com/revel/revel"
    "github.com/revmgo"
    // "gopkg.in/mgo.v2/bson"
    // "strconv"
    // "time"
)

type Category struct {
    *revel.Controller
    revmgo.MongoController
}

func (c Category) Index() revel.Result {
    controllerName := "list"
    return c.Render(controllerName)
}
