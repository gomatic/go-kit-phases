package service

import (
	"context"
	"time"

	"github.com/gomatic/go-kit-phases/api/moody"
)

//
func (s Self) Create(_ context.Context, _ *moody.Feeling) (*moody.Overall, error) {
	return &moody.Overall{
		Date:    uint64(time.Now().UTC().Unix()),
		Average: 5,
	}, nil
}

//
func (s Self) Retrieve(_ context.Context, _ *moody.Query) (*moody.Feeling, error) {
	return &moody.Feeling{
		Enjoyment: &moody.Enjoyment{
			Feeling: 5,
			Level:   5,
		},
		Anger: &moody.Anger{
			Feeling: 5,
			Level:   5,
		},
		Fear: &moody.Fear{
			Feeling: 5,
			Level:   5,
		},
		Sadness: &moody.Sadness{
			Feeling: 5,
			Level:   5,
		},
		Disgust: &moody.Disgust{
			Feeling: 5,
			Level:   5,
		},
		Date: uint64(time.Now().UTC().Unix()),
	}, nil
}

//
func (s Self) Update(_ context.Context, _ *moody.Feeling) (*moody.Overall, error) {
	return &moody.Overall{
		Date:    uint64(time.Now().UTC().Unix()),
		Average: 5,
	}, nil
}

//
func (s Self) Delete(_ context.Context, _ *moody.Feeling) (*moody.Feeling, error) {
	return &moody.Feeling{
		Enjoyment: &moody.Enjoyment{
			Feeling: 1,
			Level:   1,
		},
		Anger: &moody.Anger{
			Feeling: 1,
			Level:   1,
		},
		Fear: &moody.Fear{
			Feeling: 1,
			Level:   1,
		},
		Sadness: &moody.Sadness{
			Feeling: 1,
			Level:   1,
		},
		Disgust: &moody.Disgust{
			Feeling: 1,
			Level:   1,
		},
		Date: 1,
	}, nil
}

//
func (s Self) List(_ context.Context, _ *moody.Query) (*moody.Complicated, error) {
	return &moody.Complicated{
		Feeling: []*moody.Feeling{
			{
				Enjoyment: &moody.Enjoyment{
					Feeling: 1,
					Level:   1,
				},
				Anger: &moody.Anger{
					Feeling: 1,
					Level:   1,
				},
				Fear: &moody.Fear{
					Feeling: 1,
					Level:   1,
				},
				Sadness: &moody.Sadness{
					Feeling: 1,
					Level:   1,
				},
				Disgust: &moody.Disgust{
					Feeling: 1,
					Level:   1,
				},
				Date: 1,
			},
			{
				Enjoyment: &moody.Enjoyment{
					Feeling: 2,
					Level:   2,
				},
				Anger: &moody.Anger{
					Feeling: 2,
					Level:   2,
				},
				Fear: &moody.Fear{
					Feeling: 2,
					Level:   2,
				},
				Sadness: &moody.Sadness{
					Feeling: 2,
					Level:   2,
				},
				Disgust: &moody.Disgust{
					Feeling: 2,
					Level:   2,
				},
				Date: 2,
			},
		},
	}, nil
}
