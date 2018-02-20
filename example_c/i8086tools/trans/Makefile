all: hello

hello: hello-8086.s csv.s
	8086v6-cc -o $@ $^

hello-8086.s: hello.s pdp2i8086.py
	python pdp2i8086.py $< > $@

hello.s: hello.c
	v6cc -S $<

clean:
	rm -f hello.s hello-8086.s hello
