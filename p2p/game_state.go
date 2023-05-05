package p2p

type GameStatus uint32

func (g GameStatus) String() string {
	switch g {
	case GameStatusWaiting:
		return "WAITING"
	case GameStatusDealing:
		return "DEALING"
	case GameStatusPreFlop:
		return "PRE FLOP"
	case GameStatusFlop:
		return "FLOP"
	case GameStatusTurn:
		return "TURN"
	case GameStatusRiver:
		return "RIVER"
	default:
		return "unknown"

	}
}

const (
	GameStatusWaiting GameStatus = iota
	GameStatusDealing
	GameStatusPreFlop
	GameStatusFlop
	GameStatusTurn
	GameStatusRiver
)

type GameRound uint8
type Player struct {
	Status GameStatus
}

type GameState struct {
	isDealer   bool // atomic accesable
	gameStatus GameStatus
	Players    map[string]*Player
}

func NewGameState() *GameState {
	return &GameState{
		isDealer:   false,
		gameStatus: status,
	}
}
func (g *GameState) loop() {
	for {
		select {}
	}
}
