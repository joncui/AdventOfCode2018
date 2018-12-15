package main

var nextDirectionMap = map[string]string{
	">\\": "v",
	">/":  "^",
	"<\\": "^",
	"</":  "v",
	"^\\": "<",
	"^/":  ">",
	"v\\": ">",
	"v/":  "<",
}

var intersectionMap = map[string][]string{
	"^": []string{"<", "^", ">"},
	"v": []string{">", "v", "<"},
	"<": []string{"v", "<", "^"},
	">": []string{"^", ">", "v"},
}

type Coord struct {
	x, y int
}

type ByCoord []Cart

func (a ByCoord) Len() int {
	return len(a)
}

func (a ByCoord) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByCoord) Less(i, j int) bool {
	if a[i].GetY() == a[j].GetY() {
		return a[i].GetX() < a[j].GetX()
	}

	return a[i].GetY() < a[j].GetY()
}

type Cart struct {
	coordinates  Coord
	direction    string
	intersection int
	alive        bool
}

func InitCart(x, y int, direction string) Cart {
	return Cart{Coord{x, y}, direction, 0, true}
}

func (c *Cart) UpdateCoordinates(tracks [][]string, cartMap *map[Coord]*Cart) bool {
	delete(*cartMap, c.coordinates)
	currentDirection := c.direction
	switch currentDirection {
	case "^":
		c.UpdateY(-1)
	case "v":
		c.UpdateY(1)
	case "<":
		c.UpdateX(-1)
	case ">":
		c.UpdateX(1)
	}

	nextTrack := tracks[c.GetY()][c.GetX()]
	if _, ok := (*cartMap)[c.coordinates]; ok {
		return true
	}

	(*cartMap)[c.coordinates] = c
	if nextTrack == "+" {
		c.direction = intersectionMap[currentDirection][c.intersection]
		c.incrementIntersection()
	} else if nextDir, ok := nextDirectionMap[currentDirection+nextTrack]; ok {
		c.direction = nextDir
	}

	return false
}

func (c *Cart) UpdateX(num int) {
	c.coordinates.x += num
}

func (c *Cart) UpdateY(num int) {
	c.coordinates.y += num
}

func (c *Cart) incrementIntersection() {
	c.intersection = (c.intersection + 1) % 3
}

func (c Cart) GetX() int {
	return c.coordinates.x
}

func (c Cart) GetY() int {
	return c.coordinates.y
}

func (c *Cart) SetCrashed() {
	c.alive = false
}
