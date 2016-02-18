open Types_t
open Types
open Action
open Tyxml_js

let bind_event ev elem handler =
  let handler evt _ = handler evt in
  Lwt_js_events.(async @@ (fun () -> ev elem handler))

let topic_entry ((r, f) : rp) =
  let topic_input =
    Html5.(input ~a:[
        a_class ["field"] ;
        a_input_type `Text ;
        a_placeholder "Propose your incredible topic" ;
        a_autofocus `Autofocus ;
        R.Html5.a_value (React.S.map (fun m -> m.field) r) ;
        a_onkeypress (fun evt -> if evt##keyCode = 13 then (Controller.update Add (r, f)); true) ;
      ] ())
  in
  let topic_input_dom = To_dom.of_input topic_input in
  bind_event Lwt_js_events.inputs topic_input_dom (fun _ ->
      Lwt.return @@ (Controller.update (Update_field topic_input_dom##value) (r, f))) ;
  Html5.(header [
      div ~a:[a_class ["row title"]] [
        div ~a:[a_class ["twelve columns"]] [
          h1 [ pcdata "Topics" ]
        ]
      ] ;        
      div ~a:[a_class ["row"]] [
        div ~a:[a_class ["twelve columns"]] [
          topic_input
        ]
      ]
    ])

let topic_item ((r, f) : rp) acc topic =
  Html5.(tr [
      td [pcdata topic.text] ;
      td [
        button ~a:[
          a_class ["button-primary like"] ;
          a_onclick (
            fun evt -> (Controller.update (Like topic.id) (r, f)); true;
          )
        ] [
          span ~a:[a_class ["nblike"]] [pcdata (Int64.to_string topic.like)];
          pcdata " ";
          span ~a:[a_class ["inclike"]] [pcdata "+"]]
      ]
    ]) :: acc

let topic_list ((r, f) : rp) =
  let topics m =
    List.rev(List.fold_left (topic_item (r, f)) [] m.topics)
  in
  let rl = Rl.list (React.S.map topics r) in
  Html5.(section [
      R.Html5.table ~a:[a_class ["u-full-width"]] rl
    ])

let view (r, f) =
  Html5.(
    div ~a:[a_class ["container"]] [
      topic_entry (r, f) ;
      topic_list (r, f)
    ])