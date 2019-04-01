package neurons

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/geometry"
)


/*
	3D тело нейрона для отрисовки и позиционирования
	инициализировать на момент отрисовки, функция затухания, и удалять
	плавная градация света от белый - красный - жёлтый - зеленый - синий - чёрный
*/
type neuron3Dbody struct{
	geom geometry.Geometry
	mat material.Material
	mesh graphic.Graphic
}