package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/gomatic/go-kit-phases/api/moody"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want moody.SelfServer
	}{
		{
			name: "happy",
			want: Self{},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := New()
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_Self_Create(t *testing.T) {
	type fields struct {
		UnimplementedSelfServer moody.UnimplementedSelfServer
	}
	type args struct {
		ctx     context.Context
		feeling *moody.Feeling
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *moody.Overall
		wantErr bool
	}{
		{
			name: "happy",
			fields: fields{
				UnimplementedSelfServer: moody.UnimplementedSelfServer{},
			},
			args: args{
				ctx: nil,
				feeling: &moody.Feeling{
					Enjoyment: &moody.Enjoyment{
						Feeling: 0,
						Level:   0,
					},
					Anger: &moody.Anger{
						Feeling: 0,
						Level:   0,
					},
					Fear: &moody.Fear{
						Feeling: 0,
						Level:   0,
					},
					Sadness: &moody.Sadness{
						Feeling: 0,
						Level:   0,
					},
					Disgust: &moody.Disgust{
						Feeling: 0,
						Level:   0,
					},
					Date: 0,
				},
			},
			want: &moody.Overall{
				Date:    5,
				Average: 5,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				s := Self{}
				got, err := s.Create(tt.args.ctx, tt.args.feeling)
				require.False(t, (err != nil) != tt.wantErr)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_Self_List(t *testing.T) {
	type fields struct {
		UnimplementedSelfServer moody.UnimplementedSelfServer
	}
	type args struct {
		ctx     context.Context
		request *moody.Query
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *moody.Complicated
		wantErr bool
	}{
		{
			name: "happy",
			fields: fields{
				UnimplementedSelfServer: moody.UnimplementedSelfServer{},
			},
			args: args{
				ctx:     nil,
				request: &moody.Query{},
			},
			want: &moody.Complicated{
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
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				s := Self{}
				got, err := s.List(tt.args.ctx, tt.args.request)
				require.False(t, (err != nil) != tt.wantErr)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_Self_Retrieve(t *testing.T) {
	type fields struct {
		UnimplementedSelfServer moody.UnimplementedSelfServer
	}
	type args struct {
		ctx     context.Context
		request *moody.Query
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *moody.Feeling
		wantErr bool
	}{
		{
			name: "happy",
			fields: fields{
				UnimplementedSelfServer: moody.UnimplementedSelfServer{},
			},
			args: args{
				ctx:     nil,
				request: &moody.Query{},
			},
			want: &moody.Feeling{
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
				Date: 5,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				s := Self{}
				got, err := s.Retrieve(tt.args.ctx, tt.args.request)
				require.False(t, (err != nil) != tt.wantErr)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_Self_Update(t *testing.T) {
	type fields struct {
		UnimplementedSelfServer moody.UnimplementedSelfServer
	}
	type args struct {
		ctx     context.Context
		feeling *moody.Feeling
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *moody.Overall
		wantErr bool
	}{
		{
			name: "happy",
			fields: fields{
				UnimplementedSelfServer: moody.UnimplementedSelfServer{},
			},
			args: args{
				ctx: nil,
				feeling: &moody.Feeling{
					Enjoyment: &moody.Enjoyment{
						Feeling: 0,
						Level:   0,
					},
					Anger: &moody.Anger{
						Feeling: 0,
						Level:   0,
					},
					Fear: &moody.Fear{
						Feeling: 0,
						Level:   0,
					},
					Sadness: &moody.Sadness{
						Feeling: 0,
						Level:   0,
					},
					Disgust: &moody.Disgust{
						Feeling: 0,
						Level:   0,
					},
					Date: 0,
				},
			},
			want: &moody.Overall{
				Date:    5,
				Average: 5,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				s := Self{}
				got, err := s.Update(tt.args.ctx, tt.args.feeling)
				require.False(t, (err != nil) != tt.wantErr)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func TestSelf_Delete(t *testing.T) {
	type fields struct {
		UnimplementedSelfServer moody.UnimplementedSelfServer
	}
	type args struct {
		in0 context.Context
		in1 *moody.Feeling
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *moody.Feeling
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Self{
				UnimplementedSelfServer: tt.fields.UnimplementedSelfServer,
			}
			got, err := s.Delete(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}
