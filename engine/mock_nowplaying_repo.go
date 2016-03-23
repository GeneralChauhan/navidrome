package engine

import (
	"errors"
	"time"
)

func CreateMockNowPlayingRepo() *MockNowPlaying {
	return &MockNowPlaying{}
}

type MockNowPlaying struct {
	NowPlayingRepository
	data []NowPlayingInfo
	err  bool
}

func (m *MockNowPlaying) SetError(err bool) {
	m.err = err
}

func (m *MockNowPlaying) Enqueue(playerId int, playerName string, trackId, username string) error {
	if m.err {
		return errors.New("Error!")
	}
	info := NowPlayingInfo{}
	info.TrackId = trackId
	info.Username = username
	info.Start = time.Now()
	info.PlayerId = playerId
	info.PlayerName = playerName

	m.data = append(m.data, NowPlayingInfo{})
	copy(m.data[1:], m.data[0:])
	m.data[0] = info

	return nil
}

func (m *MockNowPlaying) Dequeue(playerId int) (*NowPlayingInfo, error) {
	if len(m.data) == 0 {
		return nil, nil
	}
	l := len(m.data)
	info := m.data[l-1]
	m.data = m.data[:l-1]

	return &info, nil
}

func (m *MockNowPlaying) Count(playerId int) (int64, error) {
	return int64(len(m.data)), nil
}

func (m *MockNowPlaying) GetAll() ([]*NowPlayingInfo, error) {
	np, err := m.Head(1)
	if np == nil || err != nil {
		return nil, err
	}
	return []*NowPlayingInfo{np}, err
}

func (m *MockNowPlaying) Head(playerId int) (*NowPlayingInfo, error) {
	if len(m.data) == 0 {
		return nil, nil
	}
	info := m.data[0]
	return &info, nil
}

func (m *MockNowPlaying) Tail(playerId int) (*NowPlayingInfo, error) {
	if len(m.data) == 0 {
		return nil, nil
	}
	info := m.data[len(m.data)-1]
	return &info, nil
}

func (m *MockNowPlaying) ClearAll() {
	m.data = make([]NowPlayingInfo, 0)
	m.err = false
}
