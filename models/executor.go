package models

type Executor struct {
	Id            int      `json:"id"`
	Insert        int      `json:"insert"`
	Update        int      `json:"update"`
	Delete        int      `json:"delete"`
	Ddl           int      `json:"ddl"`
	Unknown       int      `json:"unknown"`
	Error         int      `json:"error"`
	InsertNsTop3  []NsTop3 `json:"insert_ns_top_3"`
	UpdateNsTop3  []NsTop3 `json:"update_ns_top_3"`
	DeleteNsTop3  []NsTop3 `json:"delete_ns_top_3"`
	DdlNsTop3     []NsTop3 `json:"ddl_ns_top_3"`
	UnknownNsTop3 []NsTop3 `json:"unknown_ns_top_3"`
	ErrorNsTop3   []NsTop3 `json:"error_ns_top_3"`
}

type NsTop3 struct {
	Key string `json:"Key"`
	Val int    `json:"Val"`
}
