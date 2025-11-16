package main

import "doomsc2/sg3"

type Doomsg3 struct {
	*sg3.Bot
}

func (b *Doomsg3) OnGameReady() {

}
func (b *Doomsg3) OnStep() {
	if b.GameLoop() == 300 {
		b.BuildAt(b.MyNatural(), sg3.UnitTypeProtossPylon)
	}
}
func (b *Doomsg3) OnInit() {

}
func (b *Doomsg3) Plan() {
	// pylonPoint := sg3.APIPoint{X: 54.5, Y: 107.5}
	// gatewayPoint := sg3.APIPoint{X: 52.5, Y: 104.5}

}
func (b *Doomsg3) OnGameFinished() {

}
func main() {
	genericBot := sg3.NewBot(sg3.APIRaceProtoss, "Doom")
	Doomsg3 := &Doomsg3{Bot: genericBot}
	// ThunderbirdAIE
	// AcropolisAIE
	// InterloperAIE
	// EphemeronAIE
	// AutomatonAIE
	// AbyssalReefAIE
	// RomanticideAIE
	// Doomsg3.VsDockerMap("ThunderbirdAIE")

	Doomsg3.VsComputer(sg3.APIRaceTerran, sg3.APIDifficultyCheatInsane, sg3.APIAIBuildMacro, "MagannathaAIE_v2")
	Doomsg3.Start(Doomsg3)
}

type Dummysg3 struct {
	*sg3.Bot
}

func (b *Dummysg3) OnGameReady() {

}
func (b *Dummysg3) OnStep() {

}
func (b *Dummysg3) OnInit() {
	b.Chat("dummy here")
}
func (b *Dummysg3) OnGameFinished() {

}
