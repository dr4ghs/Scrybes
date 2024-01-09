package game

import (
	"errors"
	"fmt"
)

// =============================================================================
// Card
// =============================================================================
type Card struct {
	CardType
	CardCost
	id     string
	attack uint
	health uint
}

// =============================================================================
// Card Type
// =============================================================================
type CardType struct {
	rareType    bool
	terrainType bool
}

func (cardType CardType) String() string {
	if cardType.rareType && cardType.terrainType {
		return "rareTerrain"
	}

	if cardType.rareType {
		return "rare"
	}

	if cardType.terrainType {
		return "terrain"
	}

	return "common"
}

// =============================================================================
// Card Cost
// =============================================================================
type CardCost struct {
	sacrifices uint8
	bones      uint8
	energy     uint8
	MagikCost
}

func NewCardCost() CardCost {
	return CardCost{}
}

func (c *CardCost) Equals(o *CardCost) bool {
	return c.sacrifices == o.sacrifices &&
		c.bones == o.bones &&
		c.energy == o.energy &&
		c.orangeMagik == o.orangeMagik &&
		c.greenMagik == o.greenMagik &&
		c.blueMagik == o.blueMagik
}

func (c *CardCost) Sacrifices(cost uint8) (*CardCost, error) {
	if cost > 4 {
		return nil, errors.New(
			fmt.Sprintf("Sacrifices cannot be more than 4: %d", cost),
		)
	}

	c.sacrifices = cost

	return c, nil
}

func (c *CardCost) Bones(cost uint8) *CardCost {
	c.bones = cost

	return c
}

func (c *CardCost) Energy(cost uint8) (*CardCost, error) {
	if cost > 6 {
		return nil, errors.New(
			fmt.Sprintf("Energy cannot be more than 6: %d", cost),
		)
	}

	c.energy = cost

	return c, nil
}

func (c *CardCost) Magiks(orange, green, blue bool) *CardCost {
	c.orangeMagik = orange
	c.greenMagik = green
	c.blueMagik = blue

	return c
}

// =============================================================================
// Magik Cost
// =============================================================================
type MagikCost struct {
	orangeMagik bool
	greenMagik  bool
	blueMagik   bool
}
