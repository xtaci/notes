(define (sum-int a b) 
	(if (> a b)
		0
		(+ a (sum-int (+ a 1) b))
	)
)
