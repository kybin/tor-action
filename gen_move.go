package action

func MoveLeft(t *Text, c *Cursor, sel *Selection) Action {
	if !sel.Empty() {
		return MoveAction{
			c.BPos(),
			sel.Min(),
		}
	}
	return MoveAction{
		c.BPos(),
		c.Left(t),
	}
}

func MoveRight(t *Text, c *Cursor, sel *Selection) Action {
	if !sel.Empty() {
		return MoveAction{
			c.BPos(),
			sel.Max(),
		}
	}
	return MoveAction{
		c.BPos(),
		c.Right(t),
	}
}
