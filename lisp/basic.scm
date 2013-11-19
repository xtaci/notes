(define (abs x)
	(cond ((< x 0) (- x))
		((> x 0) x)
		((= x 0) 0)
	)
)

(define (average x y) (/ (+ x y) 2))
