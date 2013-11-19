(define (sqrt-guess g x)
	(if (< (abs (- g (/ x g))) 0.1) 
		 g
		 (sqrt-guess (average g (/ x g)) x))
)
