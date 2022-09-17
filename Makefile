run_wp:
	go test -bench=. workerpoolsexample/books/using_wp
run_wwp:
	go test -bench=. workerpoolsexample/books/without_wp  
run_ex:
	go test -bench=. workerpoolsexample/theorical