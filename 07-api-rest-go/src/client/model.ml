open Types_t

let empty = {
  topics = [] ;
  field = "" ;
}

let new_topic text = {
  id = Int64.zero ;
  text = text ;
  like = Int64.zero ;
}

let set_topics l m = 
  { m with topics = l }