package pkg

import mychannels "golangbot/pkg/channels"

func Channels() {

	//channels
	mychannels.Channel_basic()
	mychannels.Channel_in_goroutine()
	mychannels.Channel_in_goroutine1()
	mychannels.Channel_advance_calc()
	mychannels.Channel_advance_unidirectional()
	mychannels.Channel_pings_pongs()
	mychannels.Channel_close_by_for()
	mychannels.Channel_close_by_range()
	mychannels.Channel_advance_calc_using_range()

	mychannels.Channel_buffered_basic()
	mychannels.Channel_buffered_close()

	mychannels.Channel_deadlock()
	mychannels.Channel_length_capacity()
	mychannels.Channel_waitGroup()
	mychannels.Channel_worker_pool()
	mychannels.Channel_select()
	mychannels.Channel_select1()
	mychannels.Channel_select_deadlock()
	mychannels.Channel_select_deadlock_handling()
	mychannels.Channel_nil_select() /**/
}
