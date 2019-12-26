package main

type collisionBox interface {
	posX() float64
	posY() float64
	sizeX() float64
	sizeY() float64
}

func collision(c1, c2 collisionBox) bool {
	x1In := c2.posX() <= c1.posX() && c1.posX() < c2.posX()+c2.sizeX()
	x2In := c1.posX() <= c2.posX() && c2.posX() < c1.posX()+c1.sizeX()
	y1In := c2.posY() <= c1.posY() && c1.posY() < c2.posY()+c2.sizeY()
	y2In := c1.posY() <= c2.posY() && c2.posY() < c1.posY()+c1.sizeY()

	return (x1In || x2In) && (y1In || y2In)
}
