package main

import (
	"math"
)

type Coord struct {
	x, y int
}

func (c Coord) GetDistance(c2 Coord) int {
	return int(math.Abs(float64(c.x-c2.x)) + math.Abs(float64(c.y-c2.y)))
}

type Creature struct {
	coordinate Coord
	race       string
	health     int
}

func InitElf(x, y int) Creature {
	return Creature{Coord{x, y}, "E", 200}
}

func InitGoblin(x, y int) Creature {
	return Creature{Coord{x, y}, "G", 200}
}

func (c *Creature) Attacked() {
	c.health -= 3
}

func (c *Creature) MoveUp() {
	c.coordinate.y--
}

func (c *Creature) MoveDown() {
	c.coordinate.y++
}

func (c *Creature) MoveLeft() {
	c.coordinate.x--
}

func (c *Creature) MoveRight() {
	c.coordinate.x++
}
