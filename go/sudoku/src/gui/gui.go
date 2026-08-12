package gui

import (
	"fmt"
	"image/color"
	"sudokusolver/src/sudoku"
	"sudokusolver/src/sudoku/solver"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

const (
	cellW = 60
	gridW = 540
	winW  = 540
	winH  = 680
)

var (
	clueDigitImages   [10]*ebiten.Image
	userDigitImages   [10]*ebiten.Image
	easyTextActive    *ebiten.Image
	easyTextInactive  *ebiten.Image
	mediumTextActive  *ebiten.Image
	mediumTextInactive *ebiten.Image
	hardTextActive    *ebiten.Image
	hardTextInactive  *ebiten.Image
)

type Game struct {
	sudoku            *sudoku.Sudoku
	activeBoard       [9][9]int
	selectedRow       int
	selectedCol       int
	solveButton       button
	resetButton       button
	easyButton        button
	mediumButton      button
	hardButton        button
	currentDifficulty sudoku.Difficulty
	message           string
	messageImg        *ebiten.Image
}

type button struct {
	x, y, w, h int
	bg         color.Color
	img        *ebiten.Image
}

func createScaledTextImage(str string, clr color.Color, scale float64) *ebiten.Image {
	// Face7x13 width is ~7px per char, height is ~13px
	w := len(str) * 7
	h := 13
	textImg := ebiten.NewImage(w, h)
	text.Draw(textImg, str, basicfont.Face7x13, 0, 11, clr)

	scaledW := int(float64(w) * scale)
	scaledH := int(float64(h) * scale)
	scaledImg := ebiten.NewImage(scaledW, scaledH)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	scaledImg.DrawImage(textImg, op)
	return scaledImg
}

func initDigits() {
	blueColor := color.RGBA{0, 64, 221, 255}
	blackColor := color.RGBA{28, 28, 30, 255}
	for i := 1; i <= 9; i++ {
		clueDigitImages[i] = createScaledTextImage(fmt.Sprintf("%d", i), blueColor, 2.0)
		userDigitImages[i] = createScaledTextImage(fmt.Sprintf("%d", i), blackColor, 2.0)
	}

	whiteColor := color.White
	easyTextActive = createScaledTextImage("EASY", whiteColor, 1.3)
	easyTextInactive = createScaledTextImage("EASY", blackColor, 1.3)
	mediumTextActive = createScaledTextImage("MEDIUM", whiteColor, 1.3)
	mediumTextInactive = createScaledTextImage("MEDIUM", blackColor, 1.3)
	hardTextActive = createScaledTextImage("HARD", whiteColor, 1.3)
	hardTextInactive = createScaledTextImage("HARD", blackColor, 1.3)
}

func (b *button) clicked(mx, my int) bool {
	return mx >= b.x && mx < b.x+b.w && my >= b.y && my < b.y+b.h
}

func NewGame(s *sudoku.Sudoku) *Game {
	initDigits()

	solveBg := color.RGBA{0, 122, 255, 255}
	resetBg := color.RGBA{142, 142, 147, 255}

	solveText := createScaledTextImage("SOLVE", color.White, 1.5)
	resetText := createScaledTextImage("RESET", color.White, 1.5)

	return &Game{
		sudoku:            s,
		activeBoard:       s.GetBoard(),
		selectedRow:       0,
		selectedCol:       0,
		currentDifficulty: s.GetDifficulty(),
		solveButton: button{
			x:   60,
			y:   615,
			w:   160,
			h:   40,
			bg:  solveBg,
			img: solveText,
		},
		resetButton: button{
			x:   320,
			y:   615,
			w:   160,
			h:   40,
			bg:  resetBg,
			img: resetText,
		},
		easyButton: button{
			x:  60,
			y:  555,
			w:  120,
			h:  35,
		},
		mediumButton: button{
			x:  210,
			y:  555,
			w:  120,
			h:  35,
		},
		hardButton: button{
			x:  360,
			y:  555,
			w:  120,
			h:  35,
		},
	}
}

func (g *Game) Update() error {
	// Mouse click
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		if my < gridW && mx >= 0 && mx < gridW {
			g.selectedCol = mx / cellW
			g.selectedRow = my / cellW
		} else {
			if g.solveButton.clicked(mx, my) {
				g.solve()
			} else if g.resetButton.clicked(mx, my) {
				g.reset()
			} else if g.easyButton.clicked(mx, my) {
				g.loadDifficulty(sudoku.Easy)
			} else if g.mediumButton.clicked(mx, my) {
				g.loadDifficulty(sudoku.Medium)
			} else if g.hardButton.clicked(mx, my) {
				g.loadDifficulty(sudoku.Hard)
			}
		}
	}

	// Keyboard input
	g.handleKeyboardInput()

	return nil
}

func (g *Game) loadDifficulty(diff sudoku.Difficulty) {
	s, err := sudoku.NewSudoku(diff)
	if err != nil {
		g.messageImg = createScaledTextImage(fmt.Sprintf("ERROR: %v", err), color.RGBA{255, 59, 48, 255}, 1.5)
		return
	}
	g.sudoku = s
	g.activeBoard = s.GetBoard()
	g.selectedRow = 0
	g.selectedCol = 0
	g.currentDifficulty = diff
	g.messageImg = createScaledTextImage("NEW GAME", color.RGBA{0, 122, 255, 255}, 1.5)
}

func (g *Game) handleKeyboardInput() {
	if g.selectedRow < 0 || g.selectedRow > 8 || g.selectedCol < 0 || g.selectedCol > 8 {
		return
	}
	originalBoard := g.sudoku.GetBoard()
	if originalBoard[g.selectedRow][g.selectedCol] != 0 {
		return
	}

	keys := []ebiten.Key{
		ebiten.Key1, ebiten.Key2, ebiten.Key3,
		ebiten.Key4, ebiten.Key5, ebiten.Key6,
		ebiten.Key7, ebiten.Key8, ebiten.Key9,
	}
	for i, key := range keys {
		if inpututil.IsKeyJustPressed(key) {
			g.activeBoard[g.selectedRow][g.selectedCol] = i + 1
			g.messageImg = nil
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) || inpututil.IsKeyJustPressed(ebiten.KeyDelete) || inpututil.IsKeyJustPressed(ebiten.Key0) {
		g.activeBoard[g.selectedRow][g.selectedCol] = 0
		g.messageImg = nil
	}
}

func (g *Game) solve() {
	solverInstance := solver.NewSolver(g.sudoku, solver.BruteForceSolverType)
	_, solvedBoard, err := solverInstance.Solve()
	if err != nil {
		g.messageImg = createScaledTextImage("UNSOLVABLE!", color.RGBA{255, 59, 48, 255}, 1.5)
		return
	}
	g.activeBoard = solvedBoard
	g.messageImg = createScaledTextImage("SOLVED!", color.RGBA{52, 199, 89, 255}, 1.5)
}

func (g *Game) reset() {
	g.activeBoard = g.sudoku.GetBoard()
	g.messageImg = createScaledTextImage("RESET", color.RGBA{142, 142, 147, 255}, 1.5)
}

func drawCenteredImage(dst *ebiten.Image, img *ebiten.Image, cx, cy int) {
	if img == nil {
		return
	}
	w, h := img.Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(cx-w/2), float64(cy-h/2))
	dst.DrawImage(img, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Fill background
	screen.Fill(color.RGBA{245, 245, 247, 255})

	// Draw selection highlight
	if g.selectedRow >= 0 && g.selectedRow < 9 && g.selectedCol >= 0 && g.selectedCol < 9 {
		vector.DrawFilledRect(
			screen,
			float32(g.selectedCol*cellW),
			float32(g.selectedRow*cellW),
			cellW,
			cellW,
			color.RGBA{0, 122, 255, 40},
			false,
		)
	}

	// Draw grid lines
	for i := 0; i <= 9; i++ {
		thickness := float32(1)
		clr := color.RGBA{174, 174, 178, 255}
		if i%3 == 0 {
			thickness = float32(4)
			clr = color.RGBA{28, 28, 30, 255}
		}

		// Vertical line
		vector.DrawFilledRect(
			screen,
			float32(i*cellW)-thickness/2,
			0,
			thickness,
			float32(gridW),
			clr,
			false,
		)

		// Horizontal line
		vector.DrawFilledRect(
			screen,
			0,
			float32(i*cellW)-thickness/2,
			float32(gridW),
			thickness,
			clr,
			false,
		)
	}

	// Draw digits
	originalBoard := g.sudoku.GetBoard()
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			val := g.activeBoard[r][c]
			if val != 0 {
				cx := c*cellW + cellW/2
				cy := r*cellW + cellW/2
				if originalBoard[r][c] != 0 {
					drawCenteredImage(screen, clueDigitImages[val], cx, cy)
				} else {
					drawCenteredImage(screen, userDigitImages[val], cx, cy)
				}
			}
		}
	}

	// Draw difficulty buttons
	// Easy Button
	var easyBg color.Color
	var easyImg *ebiten.Image
	if g.currentDifficulty == sudoku.Easy {
		easyBg = color.RGBA{0, 122, 255, 255}
		easyImg = easyTextActive
	} else {
		easyBg = color.RGBA{225, 225, 230, 255}
		easyImg = easyTextInactive
	}
	vector.DrawFilledRect(
		screen,
		float32(g.easyButton.x),
		float32(g.easyButton.y),
		float32(g.easyButton.w),
		float32(g.easyButton.h),
		easyBg,
		false,
	)
	drawCenteredImage(
		screen,
		easyImg,
		g.easyButton.x+g.easyButton.w/2,
		g.easyButton.y+g.easyButton.h/2,
	)

	// Medium Button
	var mediumBg color.Color
	var mediumImg *ebiten.Image
	if g.currentDifficulty == sudoku.Medium {
		mediumBg = color.RGBA{0, 122, 255, 255}
		mediumImg = mediumTextActive
	} else {
		mediumBg = color.RGBA{225, 225, 230, 255}
		mediumImg = mediumTextInactive
	}
	vector.DrawFilledRect(
		screen,
		float32(g.mediumButton.x),
		float32(g.mediumButton.y),
		float32(g.mediumButton.w),
		float32(g.mediumButton.h),
		mediumBg,
		false,
	)
	drawCenteredImage(
		screen,
		mediumImg,
		g.mediumButton.x+g.mediumButton.w/2,
		g.mediumButton.y+g.mediumButton.h/2,
	)

	// Hard Button
	var hardBg color.Color
	var hardImg *ebiten.Image
	if g.currentDifficulty == sudoku.Hard {
		hardBg = color.RGBA{0, 122, 255, 255}
		hardImg = hardTextActive
	} else {
		hardBg = color.RGBA{225, 225, 230, 255}
		hardImg = hardTextInactive
	}
	vector.DrawFilledRect(
		screen,
		float32(g.hardButton.x),
		float32(g.hardButton.y),
		float32(g.hardButton.w),
		float32(g.hardButton.h),
		hardBg,
		false,
	)
	drawCenteredImage(
		screen,
		hardImg,
		g.hardButton.x+g.hardButton.w/2,
		g.hardButton.y+g.hardButton.h/2,
	)

	// Draw control buttons
	// Solve button
	vector.DrawFilledRect(
		screen,
		float32(g.solveButton.x),
		float32(g.solveButton.y),
		float32(g.solveButton.w),
		float32(g.solveButton.h),
		g.solveButton.bg,
		false,
	)
	drawCenteredImage(
		screen,
		g.solveButton.img,
		g.solveButton.x+g.solveButton.w/2,
		g.solveButton.y+g.solveButton.h/2,
	)

	// Reset button
	vector.DrawFilledRect(
		screen,
		float32(g.resetButton.x),
		float32(g.resetButton.y),
		float32(g.resetButton.w),
		float32(g.resetButton.h),
		g.resetButton.bg,
		false,
	)
	drawCenteredImage(
		screen,
		g.resetButton.img,
		g.resetButton.x+g.resetButton.w/2,
		g.resetButton.y+g.resetButton.h/2,
	)

	// Draw status message if present
	if g.messageImg != nil {
		drawCenteredImage(screen, g.messageImg, winW/2, 668)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return winW, winH
}
