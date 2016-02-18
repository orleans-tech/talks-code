(* Auto-generated from "types.atd" *)


type topic = { id: Int64.t; text: string; like: Int64.t }

type topic_list = topic list

type t = { topics: topic list; field: string }
