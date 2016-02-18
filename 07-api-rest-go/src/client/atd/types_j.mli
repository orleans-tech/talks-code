(* Auto-generated from "types.atd" *)


type topic = Types_t.topic = { id: Int64.t; text: string; like: Int64.t }

type topic_list = Types_t.topic_list

type t = Types_t.t = { topics: topic list; field: string }

val write_topic :
  Bi_outbuf.t -> topic -> unit
  (** Output a JSON value of type {!topic}. *)

val string_of_topic :
  ?len:int -> topic -> string
  (** Serialize a value of type {!topic}
      into a JSON string.
      @param len specifies the initial length
                 of the buffer used internally.
                 Default: 1024. *)

val read_topic :
  Yojson.Safe.lexer_state -> Lexing.lexbuf -> topic
  (** Input JSON data of type {!topic}. *)

val topic_of_string :
  string -> topic
  (** Deserialize JSON data of type {!topic}. *)

val write_topic_list :
  Bi_outbuf.t -> topic_list -> unit
  (** Output a JSON value of type {!topic_list}. *)

val string_of_topic_list :
  ?len:int -> topic_list -> string
  (** Serialize a value of type {!topic_list}
      into a JSON string.
      @param len specifies the initial length
                 of the buffer used internally.
                 Default: 1024. *)

val read_topic_list :
  Yojson.Safe.lexer_state -> Lexing.lexbuf -> topic_list
  (** Input JSON data of type {!topic_list}. *)

val topic_list_of_string :
  string -> topic_list
  (** Deserialize JSON data of type {!topic_list}. *)

val write_t :
  Bi_outbuf.t -> t -> unit
  (** Output a JSON value of type {!t}. *)

val string_of_t :
  ?len:int -> t -> string
  (** Serialize a value of type {!t}
      into a JSON string.
      @param len specifies the initial length
                 of the buffer used internally.
                 Default: 1024. *)

val read_t :
  Yojson.Safe.lexer_state -> Lexing.lexbuf -> t
  (** Input JSON data of type {!t}. *)

val t_of_string :
  string -> t
  (** Deserialize JSON data of type {!t}. *)

