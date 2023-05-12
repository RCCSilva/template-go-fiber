package domain

import "rccsilva.com/template-go/database"

type App struct {
	memberRepository MemberRepository
}

func New(db *database.Database) *App {
	return &App{
		memberRepository: MemberRepository{db: db},
	}
}
