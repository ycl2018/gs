
.PHONY: gen install

install:
	@echo "install antlr4-tools"
	pip install antlr4-tools && \
	antlr4

gen:
	antlr4 -Dlanguage=Go -o gen -package gen -visitor -no-listener -Werror  Gs.g4 && \
	sed '4,8d' ./gen/gs_base_visitor.go > tmp && mv tmp ./gen/gs_base_visitor.go
