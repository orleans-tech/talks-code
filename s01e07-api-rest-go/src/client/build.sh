#!/bin/sh

cd atd/
for f in *.atd ; \
do ( \
	atdgen -t -j-std $f; \
	atdgen -j -j-std $f; \
); \
done
cd ../

ocamlbuild -use-ocamlfind \
  -pkgs atdgen,lwt.syntax,js_of_ocaml,js_of_ocaml.syntax,js_of_ocaml.tyxml,tyxml,react,reactiveData \
  -syntax camlp4o \
  main.byte ;

js_of_ocaml +weak.js --opt 3 -o js/main.js main.byte
