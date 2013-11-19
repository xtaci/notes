(define (sqrt x)
	(define (sqrt-guess g x)
		(if (< (abs (- g (/ x g))) 0.001) 
			 g
			 (sqrt-guess (average g (/ x g)) x))
	)
	(sqrt-guess 1 x)
)
