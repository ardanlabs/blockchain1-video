// Package state is the core API for the blockchain and implements all the
// business rules and processing.
package state

import (
	"github.com/ardanlabs/blockchain/foundation/blockchain/database"
	"github.com/ardanlabs/blockchain/foundation/blockchain/genesis"
)

// Config represents the configuration required to start
// the blockchain node.
type Config struct {
	Genesis genesis.Genesis
}

// State manages the blockchain database.
type State struct {
	genesis genesis.Genesis
	db      *database.Database
}

// New constructs a new blockchain for data management.
func New(cfg Config) (*State, error) {

	// Access the storage for the blockchain.
	db, err := database.New(cfg.Genesis)
	if err != nil {
		return nil, err
	}

	// Create the State to provide support for managing the blockchain.
	state := State{
		genesis: cfg.Genesis,
		db:      db,
	}

	return &state, nil
}

// Genesis returns a copy of the genesis information.
func (s *State) Genesis() genesis.Genesis {
	return s.genesis
}

// Accounts returns a copy of the database accounts.
func (s *State) Accounts() map[database.AccountID]database.Account {
	return s.db.Copy()
}

// QueryAccount returns a copy of the account from the database.
func (s *State) QueryAccount(account database.AccountID) (database.Account, error) {
	return s.db.Query(account)
}
