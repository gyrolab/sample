/*
 *  Gyro - A modern UI toolkit for Golang
 *  Copyright (C) 2018 Roland Singer <roland@desertbit.com>
 */

package main

import (
	"log"

	"github.com/gyrolab/gyro"
	"github.com/gyrolab/gyro-backend-nanovgo"
)

func init() {
	gyro.RegisterBackend(nanovgo.New())
}

func main() {
	err := gyro.Run(initApp)
	if err != nil {
		log.Fatalln(err)
	}
}

func initApp() (err error) {
	win, err := gyro.NewWindow("title", 400, 300)
	if err != nil {
		return
	}

	rect := gyro.NewRect()
	rect.SetX(20)
	rect.SetY(20)
	rect.SetWidth(100)
	rect.SetHeight(100)
	rect.SetRadius(10)
	rect.SetBorder(10)
	rect.SetColor(gyro.RGBA(255, 255, 0, 255))
	rect.SetBorderColor(gyro.RGBA(255, 0, 0, 255))
	win.AddWidget(rect)

	rect2 := gyro.NewRect()
	rect2.SetX(10)
	rect2.SetY(10)
	rect2.SetWidth(50)
	rect2.SetHeight(50)
	rect2.SetRadius(10)
	rect2.SetBorder(10)
	rect2.SetColor(gyro.RGBA(255, 255, 255, 255))
	rect2.SetBorderColor(gyro.RGBA(255, 0, 255, 255))
	rect.AddWidget(rect2)

	return
}
