package mongoshake

type ReplLsn struct {
	Unix int64  `json:"unix"`
	Time string `json:"time"`
	Ts   string `json:"ts"`
}

type ReplLsnAck struct {
	Unix int64  `json:"unix"`
	Time string `json:"time"`
	Ts   string `json:"ts"`
}

type ReplLsnCkpt struct {
	Unix int64  `json:"unix"`
	Time string `json:"time"`
	Ts   string `json:"ts"`
}

type ReplNow struct {
	Unix int64  `json:"unix"`
	Time string `json:"time"`
}

type Repl struct {
	Who         string      `json:"who"`
	Tag         string      `json:"tag"`
	Replset     string      `json:"replset"`
	LogsGet     int         `json:"logs_get"`
	LogsRepl    int         `json:"logs_repl"`
	LogsSuccess int         `json:"logs_success"`
	Tps         int         `json:"tps"`
	Lsn         ReplLsn     `json:"lsn"`
	LsnAck      ReplLsnAck  `json:"lsn_ack"`
	LsnCkpt     ReplLsnCkpt `json:"lsn_ckpt"`
	Now         ReplNow     `json:"now"`
	LogSizeAvg  string      `json:"log_size_avg"`
	LogSizeMax  string      `json:"log_size_max"`
}
