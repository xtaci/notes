(define (fib N)
	(if (< N 2)
		N
		(+ (fib (- N 1)) (fib (- N 2)))
	)
)
