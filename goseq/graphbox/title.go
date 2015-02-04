package graphbox

type TitleStyle struct {
    Font            Font
    FontSize        int
    Padding         Point
}


// A title
type Title struct {
    TC              int
    style           TitleStyle
    textBox         *TextBox
    textBoxRect     Rect
}

func NewTitle(toCol int, text string, style TitleStyle) *Title {
    textBox := NewTextBox(style.Font, style.FontSize, MiddleTextAlign)
    textBox.AddText(text)

    brect := textBox.BoundingRect()
    return &Title{toCol, style, textBox, brect}
}

func (al *Title) Constraint(r, c int) Constraint {
    h := al.textBoxRect.H + al.style.Padding.Y
    w := al.textBoxRect.W

    _ = h
    _ = w
    return Constraints([]Constraint{ 
        AddSizeConstraint{r, c, 0, 0, h, al.style.Padding.Y},
        TotalSizeConstraint{r, c, r + 1, al.TC, w + al.style.Padding.X * 2, 0},
    })
}

func (al *Title) Draw(ctx DrawContext, point Point) {
    fx, fy := point.X, point.Y
    if point, isPoint := ctx.PointAt(ctx.R, al.TC) ; isPoint {
        tx, _ := point.X, point.Y

        textX := fx + (tx - fx) / 2
        textY := fy - al.style.Padding.Y
        al.renderMessage(ctx, textX, textY)
    }
}

func (al *Title) renderMessage(ctx DrawContext, tx, ty int) {
    rect := al.textBoxRect.PositionAt(tx, ty, SouthGravity)

    ctx.Canvas.Rect(rect.X, rect.Y, rect.W, rect.H, "fill:white;stroke:white;")
    al.textBox.Render(ctx.Canvas, tx, ty, SouthGravity)
}