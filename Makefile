ignore_expression = Spacing,Primary,Op,Expression,Grouping,BooleanOp
PEGS = encoding/binary/expression/expression.go

all: $(PEGS)
clean:
	rm $(PEGS)

include ../parser/PegRules.make
