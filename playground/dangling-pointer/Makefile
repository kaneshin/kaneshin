TARGET = c go
OBJ = a.out

.SUFFIXES: .c .go

all: $(TARGET)

c:
	gcc -o $(OBJ) main.c && ./$(OBJ)

go:
	go run main.go

clean:
	rm -f $(OBJ)
